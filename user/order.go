package user

import (
	"coinford_api/models"
	//"coinford_api/configs"
	//"time"
	"github.com/astaxie/beego"
	"encoding/json"
	"fmt"
)

type OrderController struct {
	beego.Controller
}

// @Title OrderBuy
// @Description Place new buy order
// @Param	body		body 	OrderBuy		true		"Buy Order Details"
// @Success 200 {int} response
// @Failure 403 Authentication Failed
// @router /limitbuy [post]
func (o *OrderController) LimitBuy() {
	var rqd OrderBuy
	json.Unmarshal(o.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		if rqd.CurrencyId != rqd.RateCurrencyId {
			totalAmount := rqd.Amount * rqd.Rate
			err, err1 := models.CheckWalletLockAmount(user.Id,
				rqd.RateCurrencyWalletId, rqd.RateCurrencyId,
				rqd.CurrencyWalletId, rqd.CurrencyId,  totalAmount)
			if err == nil && err1 == nil {
				orderBuy := models.OrderBuy{UserId: user.Id,
					CurrencyId: rqd.CurrencyId, RateCurrencyId: rqd.RateCurrencyId,
					Amount: rqd.Amount, Rate: rqd.Rate, TotalAmount: totalAmount,
					CurrencyWalletId: rqd.CurrencyWalletId, RateCurrencyWalletId: rqd.RateCurrencyWalletId,
					Lock: false}
				err = orderBuy.Insert()
				if err == nil {
					o.Data["json"] = o.prepareMyBuyOrderList(user, orderBuy.CurrencyId, orderBuy.RateCurrencyId)
				} else {
					fmt.Println(err)
					o.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Unable to add buy order."}
				}
			} else {
				fmt.Println(err, err1)
				o.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Insufficient balance."}
			}
		} else {
			o.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Same currency pair not allowed."}
		}
	} else {
		o.Data["json"] = jres
	}
	o.ServeJSON()
}

// @Title OrderSell
// @Description Place new sell order
// @Param	body		body 	OrderSell		true		"Sell Order Details"
// @Success 200 {int} response
// @Failure 403 Authentication Failed
// @router /limitsell [post]
func (o *OrderController) LimitSell() {
	var rqd OrderSell
	json.Unmarshal(o.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		if rqd.CurrencyId != rqd.RateCurrencyId {
			totalAmount := rqd.Amount * rqd.Rate
			err, err1 := models.CheckWalletLockAmount(user.Id,
				rqd.CurrencyWalletId, rqd.CurrencyId,
				rqd.RateCurrencyWalletId, rqd.RateCurrencyId, rqd.Amount)
			if err == nil && err1 == nil {
				orderSell := models.OrderSell{UserId: user.Id,
					CurrencyId: rqd.CurrencyId, RateCurrencyId: rqd.RateCurrencyId,
					Amount: rqd.Amount, Rate: rqd.Rate, TotalAmount: totalAmount,
					CurrencyWalletId: rqd.CurrencyWalletId, RateCurrencyWalletId: rqd.RateCurrencyWalletId,
					Lock: false}
				err = orderSell.Insert()
				if err == nil {
					o.Data["json"] = o.prepareMySellOrderList(user, orderSell.CurrencyId, orderSell.RateCurrencyId)
				} else {
					fmt.Println(err)
					o.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Unable to add sell order."}
				}
			} else {
				fmt.Println(err, err1)
				o.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Insufficient balance."}
			}
		} else {
			o.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Same currency pair not allowed."}
		}
	} else {
		o.Data["json"] = jres
	}
	o.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	body		body 	OrderId		true		"Token for Authentication"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /deletebuy [delete]
func (o *OrderController) DeleteBuy() {
	var rqd OrderId
	json.Unmarshal(o.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		orderBuy := models.OrderBuy{Id: rqd.OrderId, UserId: user.Id, CurrencyId: rqd.CurrencyId, RateCurrencyId: rqd.RateCurrencyId}
		err := orderBuy.Read()
		if err == nil && orderBuy.Lock == false {
			err, err1 := models.CheckWalletUnlockAmount(user.Id,
				orderBuy.RateCurrencyWalletId, rqd.RateCurrencyId,
				orderBuy.CurrencyWalletId, rqd.CurrencyId, orderBuy.TotalAmount)
			if err == nil && err1 == nil {
				err = orderBuy.Delete("id", "user_id", "currency_id", "rate_currency_id")
				if err == nil {
					o.Data["json"] = o.prepareMyBuyOrderList(user, rqd.CurrencyId, rqd.RateCurrencyId)
				} else {
					o.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Delete Failed"}
				}
			} else {
				o.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Delete Failed"}
			}
		} else {
			o.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Delete Failed"}
		}
	} else {
		o.Data["json"] = jres
	}
	o.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	body		body 	OrderId		true		"Token for Authentication"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /deletesell [delete]
func (o *OrderController) DeleteSell() {
	var rqd OrderId
	json.Unmarshal(o.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		orderSell := models.OrderSell{Id: rqd.OrderId, UserId: user.Id, CurrencyId: rqd.CurrencyId, RateCurrencyId: rqd.RateCurrencyId}
		err := orderSell.Read()
		if err == nil && orderSell.Lock == false {
			err, err1 := models.CheckWalletUnlockAmount(user.Id,
				orderSell.CurrencyWalletId, rqd.CurrencyId,
				orderSell.RateCurrencyWalletId, rqd.RateCurrencyId, orderSell.Amount)
			if err == nil && err1 == nil {
				err = orderSell.Delete("id", "user_id", "currency_id", "rate_currency_id")
				if err == nil {
					o.Data["json"] = o.prepareMySellOrderList(user, rqd.CurrencyId, rqd.RateCurrencyId)
				} else {
					fmt.Println(err)
					o.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Delete Failed"}
				}
			} else {
				fmt.Println(err, err1)
				o.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Delete Failed"}
			}
		} else {
			o.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Delete Failed"}
		}
	} else {
		o.Data["json"] = jres
	}
	o.ServeJSON()
}

// @Title GetAll
// @Description get all the buy orders for the user
// @Param	body		body 	RequestOrderList		true		"Token for Authentication"
// @Success 200 {int} response
// @Failure 403 Authentication Failed
// @router /alllist [post]
func (o *OrderController) AllList() {
	var rqd RequestOrderList
	json.Unmarshal(o.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		o.Data["json"] = o.prepareAllOrderList(user, rqd.CurrencyId, rqd.RateCurrencyId)
	} else {
		o.Data["json"] = jres
	}
	o.ServeJSON()
}

// @Title GetAll
// @Description get all the buy orders for the user
// @Param	body		body 	UserToken		true		"Token for Authentication"
// @Success 200 {int} response
// @Failure 403 Authentication Failed
// @router /mycurrencypair [post]
func (o *OrderController) MyCurrencyPair() {
	var rqd UserToken
	json.Unmarshal(o.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		o.Data["json"] = o.myOrderCurrency(user)
	} else {
		o.Data["json"] = jres
	}
	o.ServeJSON()
}

// @Title GetAll
// @Description get all the buy orders for the user
// @Param	body		body 	RequestOrderList		true		"Token for Authentication"
// @Success 200 {int} response
// @Failure 403 Authentication Failed
// @router /updatemycurrencypair [post]
func (o *OrderController) UpdateMyCurrencyPair() {
	var rqd RequestOrderList
	json.Unmarshal(o.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		o.Data["json"] = o.updateOrderCurrency(user, rqd.CurrencyId, rqd.RateCurrencyId)
	} else {
		o.Data["json"] = jres
	}
	o.ServeJSON()
}

// @Title GetAll
// @Description get all the buy orders for the user
// @Param	body		body 	OrderWalletUpdate		true		"Token for Authentication"
// @Success 200 {int} response
// @Failure 403 Authentication Failed
// @router /updatemyorderwallet [post]
func (o *OrderController) UpdateMyOrderWallet() {
	var rqd OrderWalletUpdate
	json.Unmarshal(o.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		o.Data["json"] = o.updateOrderWallet(user, rqd.CurrencyId, rqd.RateCurrencyId, rqd.CurrencyWalletId, rqd.RateCurrencyWalletId)
	} else {
		o.Data["json"] = jres
	}
	o.ServeJSON()
}

// @Title GetAll
// @Description get all the buy orders for the user
// @Param	body		body 	RequestOrderList		true		"Token for Authentication"
// @Success 200 {int} response
// @Failure 403 Authentication Failed
// @router /mylistbuy [post]
func (o *OrderController) MyListBuy() {
	var rqd RequestOrderList
	json.Unmarshal(o.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		o.Data["json"] = o.prepareMyBuyOrderList(user, rqd.CurrencyId, rqd.RateCurrencyId)
	} else {
		o.Data["json"] = jres
	}
	o.ServeJSON()
}

// @Title GetAll
// @Description get all the sell orders for the user
// @Param	body		body 	RequestOrderList		true		"Token for Authentication"
// @Success 200 {int} response
// @Failure 403 Authentication Failed
// @router /mylistsell [post]
func (o *OrderController) MyListSell() {
	var rqd RequestOrderList
	json.Unmarshal(o.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		o.Data["json"] = o.prepareMySellOrderList(user, rqd.CurrencyId, rqd.RateCurrencyId)
	} else {
		o.Data["json"] = jres
	}
	o.ServeJSON()
}

// @Title GetAll
// @Description get all the buy orders for the user
// @Param	body		body 	RequestOrderList		true		"Token for Authentication"
// @Success 200 {int} response
// @Failure 403 Authentication Failed
// @router /listbuy [post]
func (o *OrderController) ListBuy() {
	var rqd RequestOrderList
	json.Unmarshal(o.Ctx.Input.RequestBody, &rqd)
	jres, _, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		o.Data["json"] = o.prepareBuyOrderList(rqd.CurrencyId, rqd.RateCurrencyId)
	} else {
		o.Data["json"] = jres
	}
	o.ServeJSON()
}

// @Title GetAll
// @Description get all the sell orders for the user
// @Param	body		body 	RequestOrderList		true		"Token for Authentication"
// @Success 200 {int} response
// @Failure 403 Authentication Failed
// @router /listsell [post]
func (o *OrderController) ListSell() {
	var rqd RequestOrderList
	json.Unmarshal(o.Ctx.Input.RequestBody, &rqd)
	jres, _, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		o.Data["json"] = o.prepareSellOrderList(rqd.CurrencyId, rqd.RateCurrencyId)
	} else {
		o.Data["json"] = jres
	}
	o.ServeJSON()
}

func (o *OrderController) prepareAllOrderList(user *models.User, currencyId int64, rateCurrencyId int64) AllOrderList {
	_, orderBuys, _ := models.BuyOrders(10, currencyId, rateCurrencyId)
	_, orderSells, _ := models.SellOrders(10, currencyId, rateCurrencyId)
	_, myOrderBuys, _ := models.MyBuyOrders(user, 10, currencyId, rateCurrencyId)
	_, myOrderSells, _ := models.MySellOrders(user, 10, currencyId, rateCurrencyId)

	_, orders, _ := models.Orders(10, currencyId, rateCurrencyId)
	_, myOrders, _ := models.MyOrders(user, 10, currencyId, rateCurrencyId)

	var myOrderDetails []MyOrderDetail
	for _, order := range myOrders {
		var myOrderDetail MyOrderDetail
		var orderType string
		if order.SellerUserId == user.Id {
			orderType = "SELL"
		} else if order.BuyerUserId == user.Id {
			orderType = "BUY"
		}
		myOrderDetail.Order = order
		myOrderDetail.OrderType = orderType

		myOrderDetails = append(myOrderDetails, myOrderDetail)
	}

	//currencyWallet, _ := models.PrimaryWallet(user.Id, currencyId)
	//rateCurrencyWallet, _ := models.PrimaryWallet(user.Id, rateCurrencyId)

	currencyWallets, rateCurrencyWallets := o.prepareMyWallets(user, currencyId, rateCurrencyId)
	orderWallet := models.MyOrderWallet(user, currencyId, rateCurrencyId)
	return AllOrderList{
		OrderBuys: orderBuys,
		OrderSells: orderSells,
		Orders: orders,
		MyOrderBuys: myOrderBuys,
		MyOrderSells: myOrderSells,
		MyOrders: myOrderDetails,
		CurrencyWallets: currencyWallets,
		RateCurrencyWallets: rateCurrencyWallets,
		OrderWallet: orderWallet,
		ResponseCode: 200, ResponseDescription: "List of Orders"}
}

func (o *OrderController) myOrderCurrency(user *models.User) DefaultCurrencyPair {
	defaultCurrencyPair := models.MyOrderCurrency(user)
	return DefaultCurrencyPair{DefaultCurrencyPair: defaultCurrencyPair, ResponseCode: 200, ResponseDescription: "Default currency pair."}
}

func (o *OrderController) updateOrderCurrency(user *models.User, currencyId int64, rateCurrencyId int64) CommonResponse {
	err := models.UpdateOrderCurrency(user, currencyId, rateCurrencyId)
	if err == nil {
		return CommonResponse{ResponseCode: 200, ResponseDescription: "Default currency pair succesfully added."}
	}
	return CommonResponse{ResponseCode: 404, ResponseDescription: "Failed to update deatult currency pair."}
}

func (o *OrderController) updateOrderWallet(user *models.User, currencyId int64, rateCurrencyId int64, currencyWalletId int64, rateCurrencyWalletId int64) OrderWalletList {
	err := models.UpdateOrderWallet(user, currencyId, rateCurrencyId, currencyWalletId, rateCurrencyWalletId)
	var orderWallet *models.OrderWallet
	var currencyWallets, rateCurrencyWallets []models.Wallet
	var ResponseDescription = ""
	var ResponseCode = 200
	if err == nil {
		currencyWallets, rateCurrencyWallets = o.prepareMyWallets(user, currencyId, rateCurrencyId)
		orderWallet = models.MyOrderWallet(user, currencyId, rateCurrencyId)
		ResponseDescription = "List of wallets."
		ResponseCode = 200
	} else {
		ResponseDescription = "Failed to update order wallet."
		ResponseCode = 404
	}
	return OrderWalletList{
		CurrencyWallets: currencyWallets,
		RateCurrencyWallets: rateCurrencyWallets,
		OrderWallet: orderWallet,
		ResponseCode: ResponseCode,
		ResponseDescription: ResponseDescription}
}

func (o *OrderController) prepareMyBuyOrderList(user *models.User, currencyId int64, rateCurrencyId int64) MyBuyOrderList {
	_, buyOrders, _ := models.MyBuyOrders(user, 10, currencyId, rateCurrencyId)
	currencyWallets, rateCurrencyWallets := o.prepareMyWallets(user, currencyId, rateCurrencyId)
	orderWallet := models.MyOrderWallet(user, currencyId, rateCurrencyId)
	return MyBuyOrderList{
		Orders: buyOrders,
		CurrencyWallets: currencyWallets,
		RateCurrencyWallets: rateCurrencyWallets,
		OrderWallet: orderWallet,
		ResponseCode: 200,
		ResponseDescription: "List of Orders"}
}

func (o *OrderController) prepareMySellOrderList(user *models.User, currencyId int64, rateCurrencyId int64) MySellOrderList {
	_, sellOrders, _ := models.MySellOrders(user, 10, currencyId, rateCurrencyId)
	currencyWallets, rateCurrencyWallets := o.prepareMyWallets(user, currencyId, rateCurrencyId)
	orderWallet := models.MyOrderWallet(user, currencyId, rateCurrencyId)
	return MySellOrderList{
		Orders: sellOrders,
		CurrencyWallets: currencyWallets,
		RateCurrencyWallets: rateCurrencyWallets,
		OrderWallet: orderWallet,
		ResponseCode: 200,
		ResponseDescription: "List of Orders"}
}

func (o *OrderController) prepareBuyOrderList(currencyId int64, rateCurrencyId int64) BuyOrderList {
	_, buyOrders, _ := models.BuyOrders(10, currencyId, rateCurrencyId)
	return BuyOrderList{Orders: buyOrders, ResponseCode: 200, ResponseDescription: "List of Orders"}
}

func (o *OrderController) prepareSellOrderList(currencyId int64, rateCurrencyId int64) SellOrderList {
	_, sellOrders, _ := models.SellOrders(10, currencyId, rateCurrencyId)
	return SellOrderList{Orders: sellOrders, ResponseCode: 200, ResponseDescription: "List of Orders"}
}

func (o *OrderController) prepareMyWallets(user *models.User, currencyId int64, rateCurrencyId int64) ([]models.Wallet, []models.Wallet) {
	_, currencyWallets, _ := models.MyWallets(user, currencyId)
	_, rateCurrencyWallets, _ := models.MyWallets(user, rateCurrencyId)
	return currencyWallets, rateCurrencyWallets
}
