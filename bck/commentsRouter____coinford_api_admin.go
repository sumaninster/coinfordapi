package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["coinford_api/admin:AdminController"] = append(beego.GlobalControllerRouter["coinford_api/admin:AdminController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/register`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
