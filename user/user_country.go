package user

import (
	"coinford_api/models"
	//"coinford_api/models/joins"
	"coinford_api/configs"
	//"github.com/astaxie/beego/validation"
	"time"
	"encoding/json"
	"github.com/astaxie/beego"
	"fmt"
)

type UserCountryController struct {
	beego.Controller
}

// @Title Add
// @Description Add user country for the user
// @Param	body		body 	UserCountryIdAdd		true		"Add User Country"
// @Success 200 {string} user country added successfully!
// @Failure 403 user country add failed
// @router /add [post]
func (uc *UserCountryController) Add() {
	var rqd UserCountryIdAdd
	json.Unmarshal(uc.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		num, _, _ := models.UserCountries(user.Id, "NO")
		if num < configs.MaxUnverifiedCountry {
			usercountry := models.UserCountry{UserId: user.Id, CountryId: rqd.CountryId, Eligible: "NO"}
			_, _, err := usercountry.ReadOrCreate("user_id", "country_id")
			if err == nil {
				if usercountry.DeletedAt != *configs.NullTime {
					usercountry.DeletedAt = *configs.NullTime
					err = usercountry.Update()
				}
				if err == nil {
					uc.Data["json"] = uc.prepareCountryList(user.Id, &rqd.RequestCountries, "User Country added successfully")
				} else {
					uc.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Read and write error 1"}
				}
			} else {
				uc.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Read and write error 2"}
			}
		} else {
			uc.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "You have exceeded the maximum number of unverified countries. Please vefify them first or delete them before adding a new country."}
		}
	} else {
		uc.Data["json"] = jres
	}
	uc.ServeJSON()
}

// @Title GetAll
// @Description get the list of countries
// @Param	body		body 	RequestCountryList		true		"Token for Authentication"
// @Success 200 {int} response
// @Failure 403 Authentication Failed
// @router /list [post]
func (uc *UserCountryController) GetAll() {
	var rqd RequestCountryList
	json.Unmarshal(uc.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		uc.Data["json"] = uc.prepareCountryList(user.Id, &rqd.RequestCountries, "List of Countries")
	} else {
		uc.Data["json"] = jres
	}
	uc.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	body		body 	UserCountryId		true		"Token for Authentication"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /delete [delete]
func (uc *UserCountryController) Delete() {
	var rqd UserCountryId
	json.Unmarshal(uc.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		usercountry := models.UserCountry{Id: rqd.Id, UserId: user.Id}
		err := usercountry.Read("id", "user_id")
		if err == nil {
			usercountry.UpdatedAt = time.Now()
			usercountry.DeletedAt = time.Now()
			err = usercountry.Update()
			if err == nil {
				uc.Data["json"] = uc.prepareCountryList(user.Id, &rqd.RequestCountries, "Country deleted successfully");
			} else {
				uc.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Delete failed 1"}
				fmt.Println(err)
			}
		} else {
			uc.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Delete Failed 2"}
			fmt.Println(err)
		}
	} else {
		uc.Data["json"] = jres
	}
	uc.ServeJSON()
}

func (uc *UserCountryController) prepareCountryList(userId int64, rqd *RequestCountries, message string) UserCountryList {
	var userCountries []models.UserCountry
	var countries []models.Country
	_, userCountries, _ = models.UserCountries(userId, rqd.Eligible)
	_, countries, _ = models.Countries(userId, rqd.OnlyMine, rqd.Eligible)

	var userCountryDetails []UserCountryDetail
	for _, v := range userCountries {
		var userCountryDetail UserCountryDetail
		userCountryDetail.UserCountry = v
		for _, v1 := range countries {
			if v.CountryId == v1.Id {
				userCountryDetail.Country = v1
			} 
		}
		userCountryDetails = append(userCountryDetails, userCountryDetail)
	}

	return UserCountryList{Countries: userCountryDetails, ResponseCode: 200, ResponseDescription: message}
}