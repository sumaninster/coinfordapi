package user

import (
	"coinford_api/models"
	"coinford_api/configs"
	"github.com/astaxie/beego/validation"
	"time"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UserController struct {
	beego.Controller
}

// @Title RegisterUser
// @Description Register New User
// @Param	body		body 	UserAdd		true		"New User Registration Data"
// @Success 200 {int} response
// @Failure 403 body is empty
// @router /register [post]
func (u *UserController) Add() {
	var rqd UserAdd
	json.Unmarshal(u.Ctx.Input.RequestBody, &rqd)
	err := parseToken(rqd.Token)
	if err == nil {
		user := models.User{Name: rqd.Name, Username: rqd.Username, EditNameTimes: 0, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: *configs.NullTime}
		err = u.validateUser(&user)
		if err == nil {
			if u.isUniqueUsername(user.Username) {
				err = user.Insert()
				if err == nil {
					password := models.Password{UserId: user.Id, Password: configs.GetSha512(rqd.Password), CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: *configs.NullTime}
					err1 := password.Insert()
					email := models.Email{UserId: user.Id, Email: rqd.Email, Primary: "YES", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: *configs.NullTime}
					err2 := email.Insert()
					debugMessage("RegisterUser Error (Password): ", err1)
					debugMessage("RegisterUser Error (Email): ", err2)
					var tokenString string
					var jres CommonResponse
					jres, tokenString, err = saveAuthToken(&user)
					if err == nil {
						u.Data["json"] = LoginResponse{Token: tokenString, Name: user.Name, ResponseCode: 200, ResponseDescription: "Registration successful"}
					} else {
						debugMessage("RegisterUser Error 1: ", err)
						u.Data["json"] = jres
					}
				} else {
					debugMessage("RegisterUser Error 2: ", err)
					u.Data["json"] = CommonResponse{ResponseCode: 403, ResponseDescription: "Unable to read and write data. Please try later."}
				}
			} else {
				debugMessage("RegisterUser Error 3: ", err)
				u.Data["json"] = CommonResponse{ResponseCode: 403, ResponseDescription: "Username exists"}
			}
		} else {
			debugMessage("RegisterUser Error 4: ", err)
			u.Data["json"] = CommonResponse{ResponseCode: 403, ResponseDescription: "Invalid data"}	
		}
	} else {
		debugMessage("RegisterUser Error 5: ", err)
		u.Data["json"] = CommonResponse{ResponseCode: 403, ResponseDescription: "Invalid token"}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	body		body 	UserToken		true		"Token for Authentication"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /delete [delete]
func (u *UserController) Delete() {
	var rqd UserToken
	json.Unmarshal(u.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		user.UpdatedAt = time.Now()
		user.DeletedAt = time.Now()
		err := user.Update()
		if err == nil {
			u.Data["json"] = CommonResponse{ResponseCode: 200, ResponseDescription: "User deleted successfully"}
		} else {
			u.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Delete failed"}
		}
	} else {
		u.Data["json"] = jres
	}
	u.ServeJSON()
}

// @Title ChangeUsername
// @Description change Username for the user
// @Param	body		body 	UserChangeUsername		true		"Change Username"
// @Success 200 {string} password change success!
// @Failure 403 password change failed
// @router /changeusername [post]
func (u *UserController) ChangeUsername() {
	var rqd UserChangeUsername
	json.Unmarshal(u.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		var isvalid bool
		password := models.Password{UserId: user.Id, Password: configs.GetSha512(rqd.CurrentPassword)}
		isvalid, _ = password.IsValid(user)
		if isvalid {
			user.Username = rqd.NewUsername
			if u.isUniqueUsername(user.Username) {
				u.updateUser(user, "Username changed successfully")
			} else {
				u.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Username exists"}				
			}
		} else {
			u.Data["json"] = CommonResponse{ResponseCode: 403, ResponseDescription: "Invalid user"}
		}
	} else {
		u.Data["json"] = jres
	}
	u.ServeJSON()
}

// @Title ChangeName
// @Description change name for the user (allowed only once). Please make sure this matches your bank account. You will not be able to change the name a second time.
// @Param	body		body 	UserChangeName		true		"Change Name"
// @Success 200 {string} password change success!
// @Failure 403 password change failed
// @router /changename [post]
func (u *UserController) ChangeName() {
	var rqd UserChangeName
	json.Unmarshal(u.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		var isvalid bool
		password := models.Password{UserId: user.Id, Password: configs.GetSha512(rqd.CurrentPassword)}
		isvalid, _ = password.IsValid(user)
		if isvalid {
			if user.EditNameTimes < configs.EditNameMaximumTimes {
				user.Name = rqd.NewName
				user.EditNameTimes += 1
				u.updateUser(user, "Name Changed Successfully")
			} else {
				u.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Edit name not allowed. You have exceeded the number of attemts"}
			}
		} else {
			u.Data["json"] = CommonResponse{ResponseCode: 403, ResponseDescription: "Invalid user"}
		}
	} else {
		u.Data["json"] = jres
	}
	u.ServeJSON()
}

// @Title IssueToken
// @Description Issue a new token
// @Success 200 {user} models.User
// @Failure 403 :uid is empty
// @router /token [get]
func (u *UserController) IssueToken() {
	expirationTime := time.Now().Add(time.Hour * time.Duration(configs.PreLoginTokenTime)).Unix()
	tokenString, _ := token(expirationTime)
    u.Data["json"] = map[string]string{"Token": tokenString}
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	body		body 	UserLogin		true		"Login Details"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [post]
func (u *UserController) Login() {
	var rqd UserLogin
	json.Unmarshal(u.Ctx.Input.RequestBody, &rqd)
	err := parseToken(rqd.Token)
	if err == nil {
		user := models.User{Username: rqd.Username}
		err := user.Read("username")
		if err == nil && user.DeletedAt == *configs.NullTime {
			var isvalid bool
			password := models.Password{UserId: user.Id, Password: configs.GetSha512(rqd.Password)}
			isvalid, password = password.IsValid(&user)
			if isvalid {
				var tokenString string
				var jres CommonResponse
				jres, tokenString, err = saveAuthToken(&user)
				if err == nil {
					u.Data["json"] = LoginResponse{Token: tokenString, Name: user.Name, ResponseCode: 200, ResponseDescription: "Login successful"}
				} else {
					debugMessage("Login Error 1: ", err)
					u.Data["json"] = jres
				}
			} else {
				debugMessage("Login Error 3: ", err)
				u.Data["json"] = CommonResponse{ResponseCode: 403, ResponseDescription: "Invalid password"}
			}	
		} else {
			debugMessage("Login Error 4: ", err)
			u.Data["json"] = CommonResponse{ResponseCode: 403, ResponseDescription: "User does not exists"}
		}
	} else {
		debugMessage("Login Error 5: ", err)
		u.Data["json"] = CommonResponse{ResponseCode: 403, ResponseDescription: "Invalid token"}
	}
	u.ServeJSON()
}

// @Title Authenticate
// @Description Authenticates the user into the system
// @Param	body		body 	UserToken		true		"Token for Authentication"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /auth [post]
func (u *UserController) Authenticate() {
	var rqd UserToken
	json.Unmarshal(u.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		//u.Data["json"] = CommonResponse{ResponseCode: 200, ResponseDescription: "Login successful"}
		jres, tokenString, err := saveAuthToken(user)
		if err == nil {
			u.Data["json"] = LoginResponse{Token: tokenString, Name: user.Name, ResponseCode: 200, ResponseDescription: "Login successful"}
		} else {
			debugMessage("Login Error 1: ", err)
			u.Data["json"] = jres
		}
	} else {
		u.Data["json"] = jres
	}
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Param	body		body 	UserToken		true		"Token for Authentication"
// @Success 200 {string} logout success
// @router /logout [post]
func (u *UserController) Logout() {
	var rqd UserToken
	json.Unmarshal(u.Ctx.Input.RequestBody, &rqd)
	jres, _, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		authToken := models.AuthToken{Token: rqd.Token}
		authToken.Read()
		authToken.ExpirationTime = time.Now()
		authToken.UpdatedAt = time.Now()
		authToken.DeletedAt = time.Now()
		err := authToken.Update()
		if err == nil {
			u.Data["json"] = CommonResponse{ResponseCode: 200, ResponseDescription: "Logout successful"}
		} else {
			u.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Unable to read or write data"}
		}
	} else {
		u.Data["json"] = jres
	}
	u.ServeJSON()
}

// @Title IsUniqueUsername
// @Description Logs out current logged in user session
// @Param	body		body 	UserUsername		true		"Username for uniqueness"
// @Success 200 {string} unique username
// @router /isuniqueusername [post]
func (u *UserController) IsUniqueUsername() {
	var rqd UserUsername
	json.Unmarshal(u.Ctx.Input.RequestBody, &rqd)
	err := parseToken(rqd.Token)
	if err == nil {
		if u.isUniqueUsername(rqd.Username) {
			u.Data["json"] = CommonResponse{ResponseCode: 200, ResponseDescription: "Username is unique"}
		} else {
			u.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Username exists"}
		}
	} else {
		u.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Invalid token"}
	}
	u.ServeJSON()
}

func (u *UserController) validateUser(user *models.User) error {
	valid := validation.Validation{}
	isvalid, err := valid.Valid(user)
	if isvalid {
		return nil
	}
	return err
}

func (u *UserController) isUniqueUsername(username string) bool {
	user := models.User{Username: username}
	err := user.Read("username")
	if err == orm.ErrNoRows {
		return true
	}
	return false
}

func (u *UserController) updateUser(user *models.User, message string) {
	err := u.validateUser(user)
	if err == nil {
		err = user.Update()
		if err == nil {
			u.Data["json"] = CommonResponse{ResponseCode: 200, ResponseDescription: message}
		} else {
			u.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Read and write error"}
		}	
	} else {
		u.Data["json"] = CommonResponse{ResponseCode: 403, ResponseDescription: "Invalid data"}
	}
}