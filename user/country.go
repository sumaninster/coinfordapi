package user

import (
	"coinford_api/models"
	"github.com/astaxie/beego"
	"encoding/json"
)

type CountryController struct {
	beego.Controller
}

// @Title GetAll
// @Description get the list of countries
// @Param	body		body 	RequestCountryList		true		"Token for Authentication"
// @Success 200 {int} response
// @Failure 403 Authentication Failed
// @router /list [post]
func (c *CountryController) GetAll() {
	var rqd RequestCountryList
	json.Unmarshal(c.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		c.Data["json"] = c.prepareCountryList(user.Id, &rqd.RequestCountries)
	} else {
		c.Data["json"] = jres
	}
	c.ServeJSON()
}

func (c *CountryController) prepareCountryList(userId int64, rqd *RequestCountries) CountryList {
	var countries []models.Country
	_, countries, _ = models.Countries(userId, rqd.OnlyMine, "")
	return CountryList{Countries: countries, ResponseCode: 200, ResponseDescription: "List of Countries"}
}