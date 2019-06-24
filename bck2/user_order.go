package user

import (
	"coinford_api/models"
	"coinford_api/configs"
	"time"
	"github.com/astaxie/beego"
	"encoding/json"
	//"fmt"
)

type OrderController struct {
	beego.Controller
}

// @Title RegisterWallet
// @Description Register New User
// @Param	body		body 	OrderDetails		true		"Order Details"
// @Success 200 {int} response
// @Failure 403 Authentication Failed
// @router /add [post]
func (o *OrderController) Add() {
	var rqd OrderDetails
	json.Unmarshal(o.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		order := models.Order{UserId: user.Id, CurrencyId: rqd.CurrencyId, Amount: rqd.Amount, Rate: rqd.Rate, RateCurrencyId: rqd.RateCurrencyId, OrderType: rqd.OrderType, ProcessedType: "NOT_PROCESSED", ProcessedAt: *configs.NullTime, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: *configs.NullTime}
		order.Insert()
		o.Data["json"] = o.prepareOrderList(user.Id, &rqd.RequestOrders)
	} else {
		o.Data["json"] = jres
	}
	o.ServeJSON()
}

// @Title GetAll
// @Description get all the wallets for the user
// @Param	body		body 	RequestOrderList		true		"Token for Authentication"
// @Success 200 {int} response
// @Failure 403 Authentication Failed
// @router /list [post]
func (o *OrderController) GetAll() {
	var rqd RequestOrderList
	json.Unmarshal(o.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		o.Data["json"] = o.prepareOrderList(user.Id, &rqd.RequestOrders)
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
// @router /delete [delete]
func (o *OrderController) Delete() {
	var rqd OrderId
	json.Unmarshal(o.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		order := models.Order{Id: rqd.OrderId, UserId: user.Id}
		err := order.Read("id", "user_id")
		if err == nil {
			order.DeletedAt = time.Now()
			err = order.Update()
			if err == nil {
				o.Data["json"] = o.prepareOrderList(user.Id, &rqd.RequestOrders)
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

func (o *OrderController) prepareOrderList(userId int64, rqd *RequestOrders) OrderList {
	var orders []models.Order
	if rqd.AllUser == "YES" {
		if rqd.ExcludeMine == "YES" {
			if rqd.IsProcessed == "YES" || rqd.IsProcessed == "NO" {
				_, orders, _ = models.OrdersIsProcessedExcludeMine(userId, rqd.OrderType, rqd.IsProcessed)
			} else {
				_, orders, _ = models.OrdersExcludeMine(userId, rqd.OrderType)
			}
		} else {
			if rqd.IsProcessed == "YES" || rqd.IsProcessed == "NO" {
				_, orders, _ = models.OrdersIsProcessed(userId, rqd.OrderType, rqd.IsProcessed)
			} else {
				_, orders, _ = models.Orders(userId, rqd.OrderType)
			}
		}
	} else {
		if rqd.IsProcessed == "YES" || rqd.IsProcessed == "NO" {
			_, orders, _ = models.UserOrdersIsProcessed(userId, rqd.OrderType, rqd.IsProcessed)
		} else {
			_, orders, _ = models.UserOrders(userId, rqd.OrderType)
		}
	}
	return OrderList{Orders: orders, ResponseCode: 200, ResponseDescription: "List of Orders"}
}
