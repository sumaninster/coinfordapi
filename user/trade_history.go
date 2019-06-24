package user

import (
	"coinford_api/models"
	"github.com/astaxie/beego"
	"encoding/json"
	"fmt"
)

type TradeHistoryController struct {
	beego.Controller
}

// @Title Trade History
// @Description Place new buy order
// @Param	body		body 	TradeHistory		true		"Token for Authentication"
// @Success 200 {int} response
// @Failure 403 Authentication Failed
// @router /history [post]
func (o *TradeHistoryController) TradeHistory() {
  var rqd TradeHistory
	json.Unmarshal(o.Ctx.Input.RequestBody, &rqd)
  fmt.Println(rqd)
	jres, _, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		if rqd.CurrencyId != rqd.RateCurrencyId {
			o.Data["json"] = o.prepareTradeHistory(rqd.Duration, rqd.CurrencyId, rqd.RateCurrencyId)
		} else {
			o.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Same currency pair not allowed."}
		}
	} else {
		o.Data["json"] = jres
	}
	o.ServeJSON()
}

func (o *TradeHistoryController) prepareTradeHistory(duration string, currencyId int64, rateCurrencyId int64) AllTradeHistory {
	_, _, low, high, last, orders := models.AllTradeHistory(duration, currencyId, rateCurrencyId)
	return AllTradeHistory{
    Low: low,
    High: high,
    Last: last,
		Orders: orders,
		ResponseCode: 200,
		ResponseDescription: "List of Orders"}
}
