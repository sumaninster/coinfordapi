package user

import (
	"coinford_api/models"
	"github.com/astaxie/beego"
	"encoding/json"
)

type CurrencyController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all the wallets for the user
// @Param	body		body 	UserToken		true		"Token for Authentication"
// @Success 200 {int} response
// @Failure 403 Authentication Failed
// @router /list [post]
func (o *CurrencyController) GetAll() {
	var rqd UserToken
	json.Unmarshal(o.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		o.Data["json"] = o.allCurrencyList(user.Id)
	} else {
		o.Data["json"] = jres
	}
	o.ServeJSON()
}

func (o *CurrencyController) allCurrencyList(userId int64) CurrencyListAll {
	_, allcurrencies, _ := models.Currencies(userId, "")
	_, fiatcurrencies, _ := models.Currencies(userId, "FIAT")
	_, cryptocurrencies, _ := models.Currencies(userId, "CRYPTO")
	return CurrencyListAll{AllCurrencies: allcurrencies, FiatCurrencies: fiatcurrencies, CryptoCurrencies: cryptocurrencies, ResponseCode: 200, ResponseDescription: "List of Currencies"}
}

func (o *CurrencyController) prepareCurrencyList(userId int64, currencyType string) CurrencyList {
	_, currencies, _ := models.Currencies(userId, currencyType)
	return CurrencyList{Currencies: currencies, ResponseCode: 200, ResponseDescription: "List of Currencies"}
}