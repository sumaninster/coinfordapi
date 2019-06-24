package user

import (
	"coinford_api/models"
	"coinford_api/configs"
	"time"
	"fmt"
	"encoding/json"
	"github.com/astaxie/beego"
)

type TransferController struct {
	beego.Controller
}

// @Title RegisterTransfer
// @Description Register New User
// @Param	body		body 	TransferAdd		true		"Transfer Details"
// @Success 200 {int} response
// @Failure 403 Authentication Failed
// @router /add [post]
func (ft *TransferController) Add() {
	var rqd TransferAdd
	json.Unmarshal(ft.Ctx.Input.RequestBody, &rqd)
	fmt.Println(rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		transferMaster := models.TransferMaster{UserId: user.Id, CurrencyId: rqd.TransferSearch.CurrencyId, CurrencyType: rqd.TransferSearch.CurrencyType}
		_, _, err := transferMaster.ReadOrCreate("user_id", "currency_id")
		if err == nil {
			notDeleted := true
			if transferMaster.DeletedAt != *configs.NullTime {
				transferMaster.DeletedAt = *configs.NullTime
				err = transferMaster.Update()
				notDeleted = false
			}
			if err == nil {
				if notDeleted {
					transfer := models.Transfer{TransferMasterId: transferMaster.Id, Amount: rqd.Amount, TransferType: rqd.TransferSearch.TransferType, FromUserId: user.Id, ToUserId: user.Id}
					err = transfer.Insert()
					if err == nil {
						switch rqd.TransferSearch.TransferType {
						case "WALLET":
							transferWallet := models.TransferWallet{TransferId: transfer.Id, FromWalletId: rqd.FromWalletId, ToWalletId: rqd.ToWalletId}
							err = transferWallet.Insert()
						case "BANK":
							transferBank := models.TransferBank{TransferId: transfer.Id, FromWalletId: rqd.FromWalletId, ToDataId: rqd.ToDataId}
							err = transferBank.Insert()
						case "PAYEE":
							transferPayee := models.TransferPayee{TransferId: transfer.Id, FromWalletId: rqd.FromWalletId, ToPayeeId: rqd.ToPayeeId}
							err = transferPayee.Insert()
						}
					} else {
						fmt.Println(err)
					}
				}
				if err == nil {
					ft.Data["json"] = ft.prepareTransferList(user.Id)
				} else {
					fmt.Println(err)
					ft.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Error Creating New Transfer 1"}
				}
			} else {
				fmt.Println(err)
				ft.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Error Creating New Transfer 2"}
			}
		} else {
			ft.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Error Creating New Transfer 3"}
		}
	} else {
		ft.Data["json"] = jres
	}
	ft.ServeJSON()
}

// @Title GetAll
// @Description get all the orders for the user
// @Param	body		body 	TransferSearchAll		true		"Token for Authentication"
// @Success 200 {int} response
// @Failure 403 Authentication Failed
// @router /list [post]
func (ft *TransferController) GetAll() {
	var rqd TransferSearchAll
	json.Unmarshal(ft.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		ft.Data["json"] = ft.prepareTransferList(user.Id)
	} else {
		ft.Data["json"] = jres
	}
	ft.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	body		body 	TransferId		true		"Transfer Details"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /delete [delete]
func (ft *TransferController) Delete() {
	var rqd TransferId
	json.Unmarshal(ft.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		transfer := models.TransferMaster{Id: rqd.Id, UserId: user.Id}
		fmt.Println(transfer)
		err := transfer.Read("id", "user_id")
		if err == nil {
			transfer.DeletedAt = time.Now()
			err = transfer.Update()
			if err == nil {
				ft.Data["json"] = ft.prepareTransferList(user.Id)
			} else {
				debugMessage("Delete: ", err)
				ft.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Delete Failed 1"}
			}
		} else {
			debugMessage("Delete: ", err)
			ft.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Delete Failed 2"}
		}
	} else {
		ft.Data["json"] = jres
	}
	ft.ServeJSON()
}

func (ft *TransferController) transferWalletDetail(transferId int64) TransferWalletDetail{
	var transferWalletDetail TransferWalletDetail
	t := models.TransferWallet{TransferId: transferId}
	err := t.Read("transfer_id")
	if err == nil {
		fromWallet := models.Wallet{Id: t.FromWalletId}
		err = fromWallet.Read("id")
		if err == nil {
			transferWalletDetail.FromWallet = &fromWallet
		}
		toWallet := models.Wallet{Id: t.ToWalletId}
		err = toWallet.Read("id")
		if err == nil {
			transferWalletDetail.ToWallet = &toWallet
		}
	}
	return transferWalletDetail
}

func (ft *TransferController) transferPayeeDetail(transferId int64) TransferPayeeDetail {
	var transferPayeeDetail TransferPayeeDetail
	t := models.TransferPayee{TransferId: transferId}
	err := t.Read("transfer_id")
	if err == nil {
		fromWallet := models.Wallet{Id: t.FromWalletId}
		err = fromWallet.Read("id")
		if err == nil {
			transferPayeeDetail.FromWallet = &fromWallet
		}
		toPayee := models.Payee{Id: t.ToPayeeId}
		err = toPayee.Read("id")
		if err == nil {
			transferPayeeDetail.ToPayee = &toPayee
		}
	}
	return transferPayeeDetail
}

func (ft *TransferController) transferBankDetail(transferId int64) TransferBankDetail {
	var transferBankDetail TransferBankDetail
	t := models.TransferBank{TransferId: transferId}
	err := t.Read("transfer_id")
	if err == nil {
		fromWallet := models.Wallet{Id: t.FromWalletId}
		err = fromWallet.Read("id")
		if err == nil {
			transferBankDetail.FromWallet = &fromWallet
		}
		toData := models.DataText{Id: t.ToDataId}
		err = toData.Read("id")
		if err == nil {
			transferBankDetail.ToData = &toData
		}
	}
	return transferBankDetail
}

func (ft *TransferController) prepareTransferList(userId int64) TransferList {
	var transferMasterDatas []TransferMasterData
	_, transferMasters, _ := models.TransferMasters(userId)
	for _, transferMaster := range transferMasters {
		currency := models.Currency{Id: transferMaster.CurrencyId}
		err := currency.Read("Id")
		if err == nil {
			transferMasterData := TransferMasterData{TransferMaster: transferMaster, Currency: currency}
			_, transfers, _ := models.Transfers(transferMaster.Id)
			var transferDatas []TransferData
			for _, transfer := range transfers {
				transferData := TransferData{Transfer: transfer}
				switch transfer.TransferType {
				case "WALLET":
					wallets := ft.transferWalletDetail(transfer.Id)
					transferData.Wallet = wallets
				case "PAYEE":
					payees := ft.transferPayeeDetail(transfer.Id)
					transferData.Payee = payees
				case "BANK":
					banks := ft.transferBankDetail(transfer.Id)
					transferData.Bank = banks
				}
				transferDatas = append(transferDatas, transferData)
			}
			transferMasterData.TransferData = transferDatas
			transferMasterDatas = append(transferMasterDatas, transferMasterData)
		}
	}
	return TransferList{TransferMasterData: transferMasterDatas, ResponseCode: 200, ResponseDescription: "List of Transfers"}
}