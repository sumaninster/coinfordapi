package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["coinford_api/controllers:AddressController"] = append(beego.GlobalControllerRouter["coinford_api/controllers:AddressController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/controllers:AddressController"] = append(beego.GlobalControllerRouter["coinford_api/controllers:AddressController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/list`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/controllers:AddressController"] = append(beego.GlobalControllerRouter["coinford_api/controllers:AddressController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/controllers:CountryController"] = append(beego.GlobalControllerRouter["coinford_api/controllers:CountryController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/controllers:CountryController"] = append(beego.GlobalControllerRouter["coinford_api/controllers:CountryController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/list`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/controllers:CountryController"] = append(beego.GlobalControllerRouter["coinford_api/controllers:CountryController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/controllers:CurrencyController"] = append(beego.GlobalControllerRouter["coinford_api/controllers:CurrencyController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/controllers:CurrencyController"] = append(beego.GlobalControllerRouter["coinford_api/controllers:CurrencyController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/list`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/controllers:CurrencyController"] = append(beego.GlobalControllerRouter["coinford_api/controllers:CurrencyController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/controllers:EmailController"] = append(beego.GlobalControllerRouter["coinford_api/controllers:EmailController"],
		beego.ControllerComments{
			Method: "ChangeEmail",
			Router: `/change`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/controllers:EmailController"] = append(beego.GlobalControllerRouter["coinford_api/controllers:EmailController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/list`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/controllers:KycDocumentController"] = append(beego.GlobalControllerRouter["coinford_api/controllers:KycDocumentController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/change`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/controllers:OrderController"] = append(beego.GlobalControllerRouter["coinford_api/controllers:OrderController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/controllers:OrderController"] = append(beego.GlobalControllerRouter["coinford_api/controllers:OrderController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/list`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/controllers:OrderController"] = append(beego.GlobalControllerRouter["coinford_api/controllers:OrderController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/controllers:PasswordController"] = append(beego.GlobalControllerRouter["coinford_api/controllers:PasswordController"],
		beego.ControllerComments{
			Method: "ChangePassword",
			Router: `/change`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/controllers:TransferController"] = append(beego.GlobalControllerRouter["coinford_api/controllers:TransferController"],
		beego.ControllerComments{
			Method: "AddCrypto",
			Router: `/addcrypto`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/controllers:UserController"] = append(beego.GlobalControllerRouter["coinford_api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/register`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/controllers:UserController"] = append(beego.GlobalControllerRouter["coinford_api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/controllers:UserController"] = append(beego.GlobalControllerRouter["coinford_api/controllers:UserController"],
		beego.ControllerComments{
			Method: "ChangeUsername",
			Router: `/changeusername`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/controllers:UserController"] = append(beego.GlobalControllerRouter["coinford_api/controllers:UserController"],
		beego.ControllerComments{
			Method: "ChangeName",
			Router: `/changename`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/controllers:UserController"] = append(beego.GlobalControllerRouter["coinford_api/controllers:UserController"],
		beego.ControllerComments{
			Method: "IssueToken",
			Router: `/token`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/controllers:UserController"] = append(beego.GlobalControllerRouter["coinford_api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/controllers:UserController"] = append(beego.GlobalControllerRouter["coinford_api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Authenticate",
			Router: `/auth`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/controllers:UserController"] = append(beego.GlobalControllerRouter["coinford_api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/controllers:UserController"] = append(beego.GlobalControllerRouter["coinford_api/controllers:UserController"],
		beego.ControllerComments{
			Method: "IsUniqueUsername",
			Router: `/isuniqueusername`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/controllers:UserCountryController"] = append(beego.GlobalControllerRouter["coinford_api/controllers:UserCountryController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/controllers:UserCountryController"] = append(beego.GlobalControllerRouter["coinford_api/controllers:UserCountryController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/list`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/controllers:UserCountryController"] = append(beego.GlobalControllerRouter["coinford_api/controllers:UserCountryController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/controllers:WalletController"] = append(beego.GlobalControllerRouter["coinford_api/controllers:WalletController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/controllers:WalletController"] = append(beego.GlobalControllerRouter["coinford_api/controllers:WalletController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/list`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_api/controllers:WalletController"] = append(beego.GlobalControllerRouter["coinford_api/controllers:WalletController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
