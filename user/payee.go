package user

import (
	"coinford_api/models"
	"coinford_api/configs"
	"time"
	"fmt"
	"encoding/json"
	"github.com/astaxie/beego"
)

type PayeeController struct {
	beego.Controller
}

// @Title RegisterWallet
// @Description Register New User
// @Param	body		body 	PayeeDetails		true		"Wallet Details"
// @Success 200 {int} response
// @Failure 403 Authentication Failed
// @router /add [post]
func (p *PayeeController) Add() {
	var rqd PayeeDetails
	json.Unmarshal(p.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		isCrypto := p.isCrypto(rqd.CurrencyId)
		if isCrypto {
			payeeMaster := models.PayeeMaster{UserId: user.Id, CurrencyId: rqd.CurrencyId, CurrencyType: "CRYPTO"}
			_, _, err := payeeMaster.ReadOrCreate("user_id", "currency_id")
			if err == nil {
				notDeleted := true
				if payeeMaster.DeletedAt != *configs.NullTime {
					payeeMaster.DeletedAt = *configs.NullTime
					err = payeeMaster.Update()
					notDeleted = false
				}
				if err == nil {
					if notDeleted {
						payee := models.Payee{PayeeMasterId: payeeMaster.Id, Name: rqd.Name, Email: rqd.Email, Nickname: rqd.Nickname}
						err := payee.Insert()
						if err == nil {
							payeeCrypto := models.PayeeCrypto{PayeeId: payee.Id, Address: rqd.Address}
							err = payeeCrypto.Insert()
						} else {
							fmt.Println(err)
							p.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Error Creating New Wallet 1"}
						}
					}
					if err == nil {
						p.Data["json"] = p.preparePayeeList(user.Id)
					} else {
						fmt.Println(err)
						p.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Error Creating New Wallet 2"}
					}
				} else {
					fmt.Println(err)
					p.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Error Creating New Wallet 3"}
				}
			} else {
				fmt.Println(err)
				p.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Error Creating New Wallet 4"}
			}
		} else {
			p.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Not a crypto currency"}
		}
	} else {
		p.Data["json"] = jres
	}
	p.ServeJSON()
}

// @Title GetAll
// @Description get all the orders for the user
// @Param	body		body 	UserToken		true		"Token for Authentication"
// @Success 200 {int} response
// @Failure 403 Authentication Failed
// @router /list [post]
func (p *PayeeController) GetAll() {
	var rqd UserToken
	json.Unmarshal(p.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		p.Data["json"] = p.preparePayeeList(user.Id)
	} else {
		p.Data["json"] = jres
	}
	p.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	body		body 	PayeeMasterId		true		"Wallet Details"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /delete [delete]
func (p *PayeeController) Delete() {
	var rqd WalletMasterId
	json.Unmarshal(p.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		wallet := models.WalletMaster{Id: rqd.Id, UserId: user.Id}
		err := wallet.Read("id", "user_id")
		if err == nil {
			wallet.DeletedAt = time.Now()
			err = wallet.Update()
			if err == nil {
				p.Data["json"] = p.preparePayeeList(user.Id)
			} else {
				debugMessage("Delete: ", err)
				p.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Delete Failed 1"}
			}
		} else {
			debugMessage("Delete: ", err)
			p.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Delete Failed 2"}
		}
	} else {
		p.Data["json"] = jres
	}
	p.ServeJSON()
}

// @Title RegisterTransfer
// @Description Register New User
// @Param	body		body 	TransferSearchAll		true		"Transfer Details"
// @Success 200 {int} response
// @Failure 403 Authentication Failed
// @router /listall [post]
func (p *PayeeController) ListAll() {
	var rqd TransferSearchAll
	json.Unmarshal(p.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
	    var walletCrypto 	[]WalletData
	    var walletFiat 		[]WalletData
	    var payeeCrypto   	[]PayeeData
	    var err 			error
		walletMaster := models.WalletMaster{UserId: user.Id, CurrencyId: rqd.TransferSearch.CurrencyId}
		err = walletMaster.Read("user_id", "currency_id")
		if err == nil {
			var wallets []models.Wallet
			_, wallets, err = models.Wallets(walletMaster.Id)
			if err == nil {
				for _, wallet := range wallets {
					switch walletMaster.CurrencyType {
					case "CRYPTO":
						wc := models.WalletCrypto{WalletId: wallet.Id}
						err = wc.Read("wallet_id")
						if err == nil {
							wcd := WalletData{Wallet: wallet, Crypto: &wc, Fiat: nil}
							walletCrypto = append(walletCrypto, wcd)
						}
					case "FIAT":
						wf := models.WalletFiat{WalletId: wallet.Id}
						err = wf.Read("wallet_id")
						if err == nil {
							dt := models.DataText{DataId: wf.DataId}
							err = dt.Read("data_id")
							if err == nil {
								wfd := WalletData{Wallet: wallet, Crypto: nil, Fiat: &dt}
								walletFiat = append(walletFiat, wfd)
							} else {
								fmt.Println(err)
							}
						}
					}
				}
			}
		}

		payeeMaster := models.PayeeMaster{UserId: user.Id, CurrencyId: rqd.TransferSearch.CurrencyId}
		err = payeeMaster.Read("user_id", "currency_id")
		if err == nil {
			var payees []models.Payee
			_, payees, err = models.Payees(payeeMaster.Id)
			if err == nil {
				for _, payee := range payees {
					switch payeeMaster.CurrencyType {
					case "CRYPTO":
						pc := models.PayeeCrypto{PayeeId: payee.Id}
						err = pc.Read("payee_id")
						if err == nil {
							pcd := PayeeData{Payee: payee, Crypto: &pc}
							payeeCrypto = append(payeeCrypto, pcd)
						} else {
							fmt.Println(err)
						}
					}
				}
			}
		}
		payee := Payee{WalletCrypto: walletCrypto, WalletFiat: walletFiat, PayeeCrypto: payeeCrypto}
		p.Data["json"] = PayeeListAll{Payee: payee, ResponseCode: 200, ResponseDescription: "Payee List"}
	} else {
		p.Data["json"] = jres
	}
	p.ServeJSON()
}

func (p *PayeeController) isCrypto(CurrencyId int64) bool {
	currency := models.Currency{Id: CurrencyId}
	err := currency.Read("id")
	if err == nil {
		if currency.Type == "CRYPTO" {
			return true
		}
	}
	return false
}

func (p *PayeeController) preparePayeeList(userId int64) PayeeList {
	var payeeMasterDatas []PayeeMasterData
	_, payeeMasters, _ := models.PayeeMasters(userId)
	for _, payeeMaster := range payeeMasters {
		payeeMasterData := PayeeMasterData{PayeeMaster: payeeMaster}
		_, payees, _ := models.Payees(payeeMaster.Id)
		var payeeDatas []PayeeData
		for _, payee := range payees {
			payeeData := PayeeData{Payee: payee}
			switch payeeMaster.CurrencyType {
			case "CRYPTO":
				payeeCrypto := models.PayeeCrypto{PayeeId: payee.Id}
				payeeCrypto.Read("payee_id")
				payeeData.Crypto = &payeeCrypto
			case "FIAT":
			}
			payeeDatas = append(payeeDatas, payeeData)
		}
		payeeMasterData.PayeeData = payeeDatas
		payeeMasterDatas = append(payeeMasterDatas, payeeMasterData)
	}
	return PayeeList{PayeeMasterData: payeeMasterDatas, ResponseCode: 200, ResponseDescription: "List of Wallets"}
}