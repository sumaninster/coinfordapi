package user

import (
	"coinford_api/models"
	//"coinford_api/configs"
	//"github.com/astaxie/beego/validation"
	//"time"
	"encoding/json"
	"github.com/astaxie/beego"
)

type EmailController struct {
	beego.Controller
}


// @Title ChangeEmail
// @Description change email for the user
// @Param	body		body 	UserChangeEmail		true		"Change Email"
// @Success 200 {string} email change success!
// @Failure 403 email change failed
// @router /change [post]
func (e *EmailController) ChangeEmail() {
	var rqd UserChangeEmail
	json.Unmarshal(e.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		email := models.Email{Id: rqd.Id, UserId: user.Id}
		err := email.Read()
		email.Email = rqd.NewEmail
		err = email.Update()
		if err == nil {
			e.Data["json"] = CommonResponse{ResponseCode: 200, ResponseDescription: "Email Changed Successfully"}
		} else {
			e.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Read and write error"}
		}
	} else {
		e.Data["json"] = jres
	}
	e.ServeJSON()
}

// @Title GetAll
// @Description get all the email ids for the user
// @Param	body		body 	UserToken		true		"Token for Authentication"
// @Success 200 {int} response
// @Failure 403 Authentication Failed
// @router /list [post]
func (e *EmailController) GetAll() {
	var rqd UserToken
	json.Unmarshal(e.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		e.Data["json"] = e.prepareEmailList(user)
	} else {
		e.Data["json"] = jres
	}
	e.ServeJSON()
}

func (e *EmailController) prepareEmailList(user *models.User) EmailList {
	_, emails, _ := models.Emails(user)
	return EmailList{Emails: emails, ResponseCode: 200, ResponseDescription: "List of Email Ids"}
}