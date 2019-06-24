// @APIVersion 1.0.0
// @Title Coinford API
// @Description This document describes all the API for coinford.com
// @Contact coinford@gmail.com
// @TermsOfServiceUrl http://coinford.com/
// @License @Coinford 2017
// @LicenseUrl http://coinford.com
package routers

import (
	"coinford_api/user"
	//"coinford_api/admin"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
	    AllowAllOrigins: true,
	    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	    AllowHeaders:     []string{"Origin", "Authorization", "X-Requested-With"},//, "Access-Control-Allow-Origin"
	    ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},//, "Access-Control-Allow-Headers"
	}))
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&user.UserController{},
			),
		),
		beego.NSNamespace("/password",
			beego.NSInclude(
				&user.PasswordController{},
			),
		),
		beego.NSNamespace("/email",
			beego.NSInclude(
				&user.EmailController{},
			),
		),
		beego.NSNamespace("/wallet",
			beego.NSInclude(
				&user.WalletController{},
			),
		),
		beego.NSNamespace("/order",
			beego.NSInclude(
				&user.OrderController{},
			),
		),
		beego.NSNamespace("/trade",
			beego.NSInclude(
				&user.TradeHistoryController{},
			),
		),
		beego.NSNamespace("/transfer",
			beego.NSInclude(
				&user.TransferController{},
			),
		),
		beego.NSNamespace("/payee",
			beego.NSInclude(
				&user.PayeeController{},
			),
		),
		beego.NSNamespace("/currency",
			beego.NSInclude(
				&user.CurrencyController{},
			),
		),
		beego.NSNamespace("/country",
			beego.NSInclude(
				&user.CountryController{},
			),
		),
		beego.NSNamespace("/usercountry",
			beego.NSInclude(
				&user.UserCountryController{},
			),
		),
		beego.NSNamespace("/field",
			beego.NSInclude(
				&user.FieldController{},
			),
		),
		beego.NSNamespace("/data",
			beego.NSInclude(
				&user.DataController{},
			),
		),
		beego.NSNamespace("/file",
			beego.NSInclude(
				&user.FileController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
