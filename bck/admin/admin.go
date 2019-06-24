package admin

import (
	//"coinford_api/user"
	//"coinford_api/configs"
	//"github.com/astaxie/beego/validation"
	//"time"
	//"encoding/json"
	"github.com/astaxie/beego"
	//"fmt"
	//"github.com/astaxie/beego/orm"
)

type AdminController struct {
	beego.Controller
}
/*
// @Title RegisterAdmin
// @Description Register New User
// @Param	body		body 	UserAdd		true		"New User Registration Data"
// @Success 200 {int} response
// @Failure 403 body is empty
// @router /register [post]
func (u *AdminController) Add() {
	var rqd user.UserAdd
	json.Unmarshal(u.Ctx.Input.RequestBody, &rqd)
	fmt.Println(rqd)
	u.Data["json"] = controllers.CommonResponse{ResponseCode: 403, ResponseDescription: "Invalid Token"}
	u.ServeJSON()
}

// @Title Add Country
// @Description Add new country
// @Param	body		body 	CountryDetails		true		"Country Details"
// @Success 200 {int} response
// @Failure 403 Authentication Failed
// @router /add [post]
func (c *AdminController) Add() {
	var rqd user.CountryDetails
	json.Unmarshal(c.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := user.apiAuthenticate(rqd.Token)
	if isLogin {
		//country := models.Country{Name: rqd.Name, IsoCode: rqd.IsoCode, DialCode: rqd.DialCode, Code: rqd.Code, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: *configs.NullTime}
		//country.Insert()
		c.Data["json"] = c.prepareCountryList(user, &rqd.RequestCountries)
	} else {
		c.Data["json"] = jres
	}
	c.ServeJSON()
}


// @Title Delete
// @Description delete the Country
// @Param	body		body 	CountryId		true		"Token for Authentication"
// @Success 200 {string} delete success!
// @Failure 403 Id is empty
// @router /delete [delete]
func (c *AdminController) Delete() {
	var rqd user.CountryId
	json.Unmarshal(c.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := user.apiAuthenticate(rqd.Token)
	if isLogin {
		country := models.Country{Id: rqd.Id}
		err := country.Read("id", "user_id")
		if err == nil {
			country.DeletedAt = time.Now()
			err = country.Update()
			if err == nil {
				c.Data["json"] = c.prepareCountryList(user, &rqd.RequestCountries)
			} else {
				c.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Delete Failed"}
			}
		} else {
			c.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Delete Failed"}
		}
	} else {
		c.Data["json"] = jres
	}
	c.ServeJSON()
}


// @Title RegisterWallet
// @Description Register New User
// @Param	body		body 	CurrencyDetails		true		"Currency Details"
// @Success 200 {int} response
// @Failure 403 Authentication Failed
// @router /add [post]
func (o *AdminController) Add() {
	var rqd user.CurrencyDetails
	json.Unmarshal(o.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := user.apiAuthenticate(rqd.Token)//*configs.NullString
	if isLogin {
		currency := models.Currency{CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: *configs.NullTime}
		currency.Insert()
		o.Data["json"] = o.prepareCurrencyList(user, "")
	} else {
		o.Data["json"] = jres
	}
	o.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	body		body 	CurrencyId		true		"Token for Authentication"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /delete [delete]
func (o *CurrencyController) Delete() {
	var rqd user.CurrencyId
	json.Unmarshal(o.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := user.apiAuthenticate(rqd.Token)
	if isLogin {
		currency := models.Currency{Id: rqd.CurrencyId}
		err := currency.Read("id", "user_id")
		if err == nil {
			currency.DeletedAt = time.Now()
			err = currency.Update()
			if err == nil {
				o.Data["json"] = o.prepareCurrencyList(user, "")
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
}*/