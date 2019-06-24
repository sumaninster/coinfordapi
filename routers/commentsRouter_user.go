package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["coinford_api/user:CountryController"] = append(beego.GlobalControllerRouter["coinford_api/user:CountryController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/list`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:CurrencyController"] = append(beego.GlobalControllerRouter["coinford_api/user:CurrencyController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/list`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:DataController"] = append(beego.GlobalControllerRouter["coinford_api/user:DataController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:DataController"] = append(beego.GlobalControllerRouter["coinford_api/user:DataController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/list`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:DataController"] = append(beego.GlobalControllerRouter["coinford_api/user:DataController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:EmailController"] = append(beego.GlobalControllerRouter["coinford_api/user:EmailController"],
		beego.ControllerComments{
			Method: "ChangeEmail",
			Router: `/change`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:EmailController"] = append(beego.GlobalControllerRouter["coinford_api/user:EmailController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/list`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:FieldController"] = append(beego.GlobalControllerRouter["coinford_api/user:FieldController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/list`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:FileController"] = append(beego.GlobalControllerRouter["coinford_api/user:FileController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/upload`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:OrderController"] = append(beego.GlobalControllerRouter["coinford_api/user:OrderController"],
		beego.ControllerComments{
			Method: "LimitBuy",
			Router: `/limitbuy`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:OrderController"] = append(beego.GlobalControllerRouter["coinford_api/user:OrderController"],
		beego.ControllerComments{
			Method: "LimitSell",
			Router: `/limitsell`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:OrderController"] = append(beego.GlobalControllerRouter["coinford_api/user:OrderController"],
		beego.ControllerComments{
			Method: "DeleteBuy",
			Router: `/deletebuy`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:OrderController"] = append(beego.GlobalControllerRouter["coinford_api/user:OrderController"],
		beego.ControllerComments{
			Method: "DeleteSell",
			Router: `/deletesell`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:OrderController"] = append(beego.GlobalControllerRouter["coinford_api/user:OrderController"],
		beego.ControllerComments{
			Method: "AllList",
			Router: `/alllist`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:OrderController"] = append(beego.GlobalControllerRouter["coinford_api/user:OrderController"],
		beego.ControllerComments{
			Method: "MyCurrencyPair",
			Router: `/mycurrencypair`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:OrderController"] = append(beego.GlobalControllerRouter["coinford_api/user:OrderController"],
		beego.ControllerComments{
			Method: "UpdateMyCurrencyPair",
			Router: `/updatemycurrencypair`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:OrderController"] = append(beego.GlobalControllerRouter["coinford_api/user:OrderController"],
		beego.ControllerComments{
			Method: "UpdateMyOrderWallet",
			Router: `/updatemyorderwallet`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:OrderController"] = append(beego.GlobalControllerRouter["coinford_api/user:OrderController"],
		beego.ControllerComments{
			Method: "MyListBuy",
			Router: `/mylistbuy`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:OrderController"] = append(beego.GlobalControllerRouter["coinford_api/user:OrderController"],
		beego.ControllerComments{
			Method: "MyListSell",
			Router: `/mylistsell`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:OrderController"] = append(beego.GlobalControllerRouter["coinford_api/user:OrderController"],
		beego.ControllerComments{
			Method: "ListBuy",
			Router: `/listbuy`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:OrderController"] = append(beego.GlobalControllerRouter["coinford_api/user:OrderController"],
		beego.ControllerComments{
			Method: "ListSell",
			Router: `/listsell`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:PasswordController"] = append(beego.GlobalControllerRouter["coinford_api/user:PasswordController"],
		beego.ControllerComments{
			Method: "ChangePassword",
			Router: `/change`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:PayeeController"] = append(beego.GlobalControllerRouter["coinford_api/user:PayeeController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:PayeeController"] = append(beego.GlobalControllerRouter["coinford_api/user:PayeeController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/list`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:PayeeController"] = append(beego.GlobalControllerRouter["coinford_api/user:PayeeController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:PayeeController"] = append(beego.GlobalControllerRouter["coinford_api/user:PayeeController"],
		beego.ControllerComments{
			Method: "ListAll",
			Router: `/listall`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:TradeHistoryController"] = append(beego.GlobalControllerRouter["coinford_api/user:TradeHistoryController"],
		beego.ControllerComments{
			Method: "TradeHistory",
			Router: `/history`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:TransferController"] = append(beego.GlobalControllerRouter["coinford_api/user:TransferController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:TransferController"] = append(beego.GlobalControllerRouter["coinford_api/user:TransferController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/list`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:TransferController"] = append(beego.GlobalControllerRouter["coinford_api/user:TransferController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:UserController"] = append(beego.GlobalControllerRouter["coinford_api/user:UserController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/register`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:UserController"] = append(beego.GlobalControllerRouter["coinford_api/user:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:UserController"] = append(beego.GlobalControllerRouter["coinford_api/user:UserController"],
		beego.ControllerComments{
			Method: "ChangeUsername",
			Router: `/changeusername`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:UserController"] = append(beego.GlobalControllerRouter["coinford_api/user:UserController"],
		beego.ControllerComments{
			Method: "ChangeName",
			Router: `/changename`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:UserController"] = append(beego.GlobalControllerRouter["coinford_api/user:UserController"],
		beego.ControllerComments{
			Method: "IssueToken",
			Router: `/token`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:UserController"] = append(beego.GlobalControllerRouter["coinford_api/user:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:UserController"] = append(beego.GlobalControllerRouter["coinford_api/user:UserController"],
		beego.ControllerComments{
			Method: "Authenticate",
			Router: `/auth`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:UserController"] = append(beego.GlobalControllerRouter["coinford_api/user:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:UserController"] = append(beego.GlobalControllerRouter["coinford_api/user:UserController"],
		beego.ControllerComments{
			Method: "IsUniqueUsername",
			Router: `/isuniqueusername`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:UserCountryController"] = append(beego.GlobalControllerRouter["coinford_api/user:UserCountryController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:UserCountryController"] = append(beego.GlobalControllerRouter["coinford_api/user:UserCountryController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/list`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:UserCountryController"] = append(beego.GlobalControllerRouter["coinford_api/user:UserCountryController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:WalletController"] = append(beego.GlobalControllerRouter["coinford_api/user:WalletController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:WalletController"] = append(beego.GlobalControllerRouter["coinford_api/user:WalletController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/list`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/user:WalletController"] = append(beego.GlobalControllerRouter["coinford_api/user:WalletController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
