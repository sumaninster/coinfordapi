package user

import (
	"coinford_api/models"
	"coinford_api/configs"
	"coinford_api/process"
	"time"
	"fmt"
	"encoding/json"
	"github.com/astaxie/beego"
)

type WalletController struct {
	beego.Controller
}

// @Title RegisterWallet
// @Description Register New User
// @Param	body		body 	WalletDetails		true		"Wallet Details"
// @Success 200 {int} response
// @Failure 403 Authentication Failed
// @router /add [post]
func (w *WalletController) Add() {
	var rqd WalletDetails
	json.Unmarshal(w.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		walletMaster := models.WalletMaster{UserId: user.Id, CurrencyId: rqd.CurrencyId, CurrencyType: "CRYPTO"}
		_, _, err := walletMaster.ReadOrCreate("user_id", "currency_id")
		if err == nil {
			notDeleted := true
			if walletMaster.DeletedAt != *configs.NullTime {
				walletMaster.DeletedAt = *configs.NullTime
				err = walletMaster.Update()
				notDeleted = false
			}
			if err == nil {
				if notDeleted {
					w.updateStatus(walletMaster.Id, &rqd)
					wallet := models.Wallet{WalletMasterId: walletMaster.Id, Amount: 0, AmountLocked: 0, Nickname: rqd.Nickname, Primary: rqd.Primary}
					err := wallet.Insert()
					if err == nil {
						err = w.addWalletCrypto(wallet.Id, walletMaster.CurrencyId)
					} else {
						fmt.Println(err)
						w.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Error Creating New Wallet 1"}
					}
				}
				if err == nil {
					w.Data["json"] = w.prepareWalletList(user.Id)
				} else {
					fmt.Println(err)
					w.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Error Creating New Wallet 2"}
				}
			} else {
				fmt.Println(err)
				w.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Error Creating New Wallet 3"}
			}
		} else {
			fmt.Println(err)
			w.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Error Creating New Wallet 4"}
		}
	} else {
		w.Data["json"] = jres
	}
	w.ServeJSON()
}

// @Title GetAll
// @Description get all the orders for the user
// @Param	body		body 	UserToken		true		"Token for Authentication"
// @Success 200 {int} response
// @Failure 403 Authentication Failed
// @router /list [post]
func (w *WalletController) GetAll() {
	var rqd UserToken
	json.Unmarshal(w.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		w.Data["json"] = w.prepareWalletList(user.Id)
	} else {
		w.Data["json"] = jres
	}
	w.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	body		body 	WalletMasterId		true		"Wallet Details"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /delete [delete]
func (w *WalletController) Delete() {
	var rqd WalletMasterId
	json.Unmarshal(w.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		wallet := models.WalletMaster{Id: rqd.Id, UserId: user.Id}
		err := wallet.Read("id", "user_id")
		if err == nil {
			wallet.DeletedAt = time.Now()
			err = wallet.Update()
			if err == nil {
				w.Data["json"] = w.prepareWalletList(user.Id)
			} else {
				debugMessage("Delete: ", err)
				w.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Delete Failed 1"}
			}
		} else {
			debugMessage("Delete: ", err)
			w.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Delete Failed 2"}
		}
	} else {
		w.Data["json"] = jres
	}
	w.ServeJSON()
}


func (w *WalletController) prepareWalletList(userId int64) WalletList {
	var walletMasterDatas []WalletMasterData
	_, walletMasters, _ := models.WalletMasters(userId)
	for _, walletMaster := range walletMasters {
		currency := models.Currency{Id: walletMaster.CurrencyId}
		err := currency.Read("Id")
		if err == nil {
			walletMasterData := WalletMasterData{WalletMaster: walletMaster, Currency: currency}
			_, wallets, _ := models.Wallets(walletMaster.Id)
			var walletDatas []WalletData
			for _, wallet := range wallets {
				walletData := WalletData{Wallet: wallet}
				switch walletMaster.CurrencyType {
				case "CRYPTO":
					walletCrypto := models.WalletCrypto{WalletId: wallet.Id}
					walletCrypto.Read("wallet_id")
					walletData.Crypto = &walletCrypto
				case "FIAT":
					walletFiat := models.WalletFiat{WalletId: wallet.Id}
					err := walletFiat.Read("wallet_id")
					if err == nil {
						field := models.Field{CountryId: currency.CountryId, FieldType: "BANK", Key: "PRIMARY"}
						err = field.Read("country_id", "field_type", "key")
						if err == nil {
							dataText := models.DataText{DataId: walletFiat.DataId, FieldId: field.Id}
							walletData.Fiat = &dataText
						}
					}
				}
				walletDatas = append(walletDatas, walletData)
			}
			walletMasterData.WalletData = walletDatas
			walletMasterDatas = append(walletMasterDatas, walletMasterData)
		}
	}
	return WalletList{WalletMasterData: walletMasterDatas, ResponseCode: 200, ResponseDescription: "List of Wallets"}
}

func (w *WalletController) updateStatus(walletMasterId int64, rqd *WalletDetails) {
	var num int64
	var err error
	if rqd.Primary {
		num, _, err = models.WalletPrimaryFalse(walletMasterId)
	} else {
		num, _, err = models.Wallets(walletMasterId)
	}

	if num == 0 && err == nil {
		if rqd.Nickname == *configs.NullString {
			rqd.Nickname = "Primary"
		}
		rqd.Primary = true
	} else {
		if rqd.Nickname == *configs.NullString {
			rqd.Nickname = "No Name"
		}
	}
}

func (w *WalletController) addWalletCrypto(walletId int64, CurrencyId int64) error {
	switch CurrencyId {
		case models.GetCurrencyId("BTC"):
			return w.addBTCWalletCrypto(walletId)
		case models.GetCurrencyId("ETH"):
			return w.addETHWalletCrypto(walletId)
		default:
			return nil
	}
	return nil
}

func (w *WalletController) addBTCWalletCrypto(walletId int64) error {
	btcprocess := process.BTCProcess{}
	address, err := btcprocess.GetNewAddress()
	fmt.Println(address)
	if err == nil {
		walletaddress := models.WalletCrypto{WalletId: walletId, Address: address}
		err = walletaddress.Insert()
		fmt.Println(err)
		return err
	} else {
		return err
	}
	return nil
}

func (w *WalletController) addETHWalletCrypto(walletId int64) error {
	expirationTime := time.Now().Add(time.Hour * time.Duration(configs.ETHPasswordExpiryHour))
	passphrase, _ := token(expirationTime.Unix())
	//passphrase := configs.RandString(100)
	//fmt.Println(passphrase)
	walletPassphrase := models.WalletPassphrase{WalletId: walletId, Passphrase: passphrase}
	err := walletPassphrase.Insert()
	if err == nil {
		ethprocess := process.ETHProcess{}
		address, err := ethprocess.GetNewAddress(passphrase)
		if err == nil {
			walletAddress := models.WalletCrypto{WalletId: walletId, Address: address}
			err = walletAddress.Insert()
			fmt.Println(err)
			return err
		} else {
			return err
		}
	}
	return nil
}