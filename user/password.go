package user

import (
	"coinford_api/models"
	"coinford_api/configs"
	//"github.com/astaxie/beego/validation"
	"time"
	"encoding/json"
	"github.com/astaxie/beego"
)

type PasswordController struct {
	beego.Controller
}

// @Title ChangePassword
// @Description change password for the user
// @Param	body		body 	UserChangePassword		true		"Change Password"
// @Success 200 {string} password change success!
// @Failure 403 password change failed
// @router /change [post]
func (p *PasswordController) ChangePassword() {
	var rqd UserChangePassword
	json.Unmarshal(p.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		var isvalid bool
		password := models.Password{UserId: user.Id, Password: configs.GetSha512(rqd.CurrentPassword)}
		isvalid, password = password.IsValid(user)
		if isvalid {
			password.DeletedAt = time.Now()
			password.UpdatedAt = time.Now()
			password.Update()
			newpassword := models.Password{UserId: user.Id, Password: configs.GetSha512(rqd.NewPassword), CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: *configs.NullTime}
			err := newpassword.Insert()
			if err == nil {
				p.Data["json"] = CommonResponse{ResponseCode: 200, ResponseDescription: "Password Changed Successfully"}
			} else {
				p.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Read and write error"}
			}
		} else {
			p.Data["json"] = CommonResponse{ResponseCode: 403, ResponseDescription: "Invalid User"}
		}
	} else {
		p.Data["json"] = jres
	}
	p.ServeJSON()
}