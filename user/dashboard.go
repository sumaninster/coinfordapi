package user
/*
import (
	"coinford_api/models"
	"coinford_api/configs"
	"time"
	"github.com/astaxie/beego"
	"encoding/json"
	"errors"
	"fmt"
)

type DashboardController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all the orders for the user
// @Param	body		body 	UserToken		true		"Token for Authentication"
// @Success 200 {int} response
// @Failure 403 Authentication Failed
// @router /listwallet [post]
func (d *DashboardController) GetAll() {
	var rqd UserToken
	json.Unmarshal(d.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		d.Data["json"] = d.prepareWalletList(user.Id)
	} else {
		d.Data["json"] = jres
	}
	d.ServeJSON()
}

func (d *DashboardController) prepareWalletList() {
	
}		*/