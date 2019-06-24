// @Param	name			query 	string	true		"The name of the user"
// @Param	email			query 	string	true		"The email id which will be used for login"
// @Param	password		query 	string	true		"The password for login"
// @Param	token			query 	string	true		"The generated token"

func (u *UserController) Add() {
	var useradd UserAdd
	json.Unmarshal(u.Ctx.Input.RequestBody, &useradd)
	err := parseToken(u.GetString("token"))
	if err == nil {
		user := models.User{Name: u.GetString("name"), Email: u.GetString("email"), Password: configs.GetSha512(u.GetString("password")), EditNameTimes: 0, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: *configs.NullTime}
		err = u.validateUser(&user)
		if err == nil {
			err = user.Insert()
			if err == nil {
				var tokenString string
				var jres JSONResponse
				jres, tokenString, err = saveAuthToken(&user)
				if err == nil {
					u.Data["json"] = u.prepareLoginResponse(tokenString, &user)
				} else {
					debugMessage("RegisterUser Error 1: ", err)
					u.Data["json"] = jres
				}
			} else {
				debugMessage("RegisterUser Error 2: ", err)
				u.Data["json"] = JSONResponse{ResponseCode: 403, Description: "Please Register Again"}
			}
		} else {
			debugMessage("RegisterUser Error 3: ", err)
			u.Data["json"] = JSONResponse{ResponseCode: 403, Description: "Invalid Data"}	
		}
	} else {
		debugMessage("RegisterUser Error 4: ", err)
		u.Data["json"] = JSONResponse{ResponseCode: 403, Description: "Invalid Token"}
	}
	u.ServeJSON()
}

// @Param	token	query 	string	true		"Token for authentication"

// @Param	currrentpassword	query 	string	true		"Current Password"
// @Param	newpassword			query 	string	true		"New Password"
// @Param	token				query 	string	true		"Token for authentication"

// @Param	currrentpassword	query 	string	true		"Current Password"
// @Param	newemail			query 	string	true		"New Password"
// @Param	token				query 	string	true		"Token for authentication"

// @Param	currrentpassword	query 	string	true		"Current Password"
// @Param	newname				query 	string	true		"New Name"
// @Param	token				query 	string	true		"Token for authentication"

// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Param	token			query 	string	true		"The token for login"

// @Param	token			query 	string	true		"The token for login"

// @Param	token	query 	string	true		"Token for authentication"

// @Param	currency		query 	string	true		"The currency code"
// @Param	nickname		query 	string	false		"The nickname for the wallet, default is Primary"
// @Param	token			query 	string	true		"Token for authentication"

// @Param	token			query 	string	true		"Token for authentication"

// @Param	id		query 	string	true		"The wallet id you want to delete"
// @Param	token	query 	string	true		"Token for authentication"

// @Param	currency		query 	string	true		"Type of Digital Currency"
// @Param	amount			query 	string	true		"Amount of Digital Currency Required"
// @Param	rate			query 	string	true		"The rate at which you want to buy Digital Currency"
// @Param	ratecurrtype	query 	string	true		"The currency type of rate"
// @Param	ordertype		query 	string	true		"Buy or Sell order type"
// @Param	token			query 	string	true		"The generated token"

// @Param	token			query 	string	true		"The generated token"

// @Param	id		query 	string	true		"The order id you want to delete"
// @Param	token	query 	string	true		"Token for authentication"