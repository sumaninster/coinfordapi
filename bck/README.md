bee run -downdoc=true -gendoc=true
openssl genrsa -out app.rsa 2048
openssl rsa -in app.rsa -pubout > app.rsa.pub

sudo -i -u postgres
psql


1) Install postgresql

2) Import the attached database (coinford) with the below command.
createdb coinford
psql coinford < path_of_database_dump

3) Create beego api project with the following command:
bee api coinford_api
   //"github.com/astaxie/beego/validation"
	userid, err := u.GetInt64("userid")
	// @Param	userid			query 	string	true		"The username for login"
	

https://godoc.org/github.com/btcsuite/btcrpcclient
https://godoc.org/github.com/btcsuite/btcutil
https://golang.org/pkg/sort/#Sort

	amount,_ := btcp.client.GetBalance(account)
	
		//f := client.ListAccountsAsync()
		//acc, _ := f.Receive()
		//acc, _ := client.ListAccounts()
		//fmt.Println(acc)
		//var add []btcjson.ListReceivedByAddressResult
		

type TransferRequestCrypto struct {
    Id              int64
    UserId          int64
    Amount          float64
    CurrencyId      int64
    CreatedAt       time.Time
    UpdatedAt       time.Time
    ProcessedAt     time.Time
    DeletedAt       time.Time
    TransferCryptoObj     *TransferCrypto
    TransferCryptoMapObj  *TransferCryptoMap
}


func isPairSupported(sellCurrency int64, buyCurrency int64) (bool, string){
	if (sellCurrency == "BTC" && buyCurrency == "ETH") {
		return true, "Digital"
	}
	if (sellCurrency == "ETH" && buyCurrency == "BTC") {
		return true, "Digital"
	}
	return false, "Not_Supported"
}

func updateDigitalOrder(sellOrder *models.Order, buyOrders *[]models.Order) {
    for _, buyOrder := range *buyOrders {
		if sellOrder.Rate <= buyOrder.Rate {
			if walletBalanceSufficient(sellOrder, &buyOrder) {
				fmt.Println("walletBalanceSufficient")
				if sellOrder.Amount <= buyOrder.Amount {
					fmt.Println("sellOrder.Amount <= buyOrder.Amount")
					creditWallet(buyOrder.UserId, sellOrder.Amount, sellOrder.CurrencyType)
					debitWallet(sellOrder.UserId, sellOrder.Amount, sellOrder.CurrencyType)
					recordTransaction(sellOrder, sellOrder.UserId, buyOrder.UserId, sellOrder.Amount, sellOrder.CurrencyType)

					creditWallet(sellOrder.UserId, sellOrder.Amount * sellOrder.Rate, sellOrder.RateCurrencyType)
					debitWallet(buyOrder.UserId, sellOrder.Amount * sellOrder.Rate, sellOrder.RateCurrencyType)
					recordTransaction(&buyOrder, buyOrder.UserId, sellOrder.UserId, sellOrder.Amount * sellOrder.Rate, sellOrder.RateCurrencyType)
					
					if sellOrder.Amount < buyOrder.Amount {
						postNewOrder(&buyOrder, buyOrder.Amount - sellOrder.Amount)
					}
				} else if sellOrder.Amount > buyOrder.Amount{
					fmt.Println("sellOrder.Amount > buyOrder.Amount")
					creditWallet(buyOrder.UserId, buyOrder.Amount, sellOrder.CurrencyType)
					debitWallet(sellOrder.UserId, buyOrder.Amount, sellOrder.CurrencyType)
					recordTransaction(sellOrder, sellOrder.UserId, buyOrder.UserId, buyOrder.Amount, sellOrder.CurrencyType)

					creditWallet(sellOrder.UserId, buyOrder.Amount * sellOrder.Rate, sellOrder.RateCurrencyType)
					debitWallet(buyOrder.UserId, buyOrder.Amount * sellOrder.Rate, sellOrder.RateCurrencyType)
					recordTransaction(&buyOrder, buyOrder.UserId, sellOrder.UserId, buyOrder.Amount * sellOrder.Rate, sellOrder.RateCurrencyType)

					postNewOrder(sellOrder, sellOrder.Amount - buyOrder.Amount)
				}
			}
			fmt.Printf("Sell: %f , Amount: %f || Buy: %f , Amount: %f \n", sellOrder.Rate, sellOrder.Amount, buyOrder.Rate, buyOrder.Amount)
		}
	}
	fmt.Println("\n END \n")
}




func processDigital(sellOrder *models.Order, buyOrders *[]models.Order) {
	var mBuyOrders []models.Order
	var buyOrder models.Order
	fmt.Printf("Sell: %f , Amount: %f \n", sellOrder.Rate, sellOrder.Amount)
	for _, buyOrder = range *buyOrders {
		if sellOrder.Rate <= buyOrder.Rate && sellOrder.UserId != buyOrder.UserId {
			mBuyOrders = append(mBuyOrders, buyOrder)
		}
	}
	for _, mBuyOrder := range mBuyOrders {
		updateDigitalOrder(sellOrder, &mBuyOrder)
		fmt.Printf("Buy: %f , Amount: %f \n", mBuyOrder.Rate, mBuyOrder.Amount)
	}
}


------------------
-------------------
	j, ld, hd int64
	j = 0
	ld = diffArr[0]
	hd = diffArr[0]
	for i, diff := range diffArr {
		if diff == 0{
			j = i
		} else if diff < 0 {
			ld = 
		} else if diff > 0 {

		}
	}

			/*sellOrder := postNewOrder(sellOrder, sellOrder.Amount - buyOrder.Amount)
			buyNum, err := models.GetBuyOrders(&sellOrder, &buyOrders)
			if err == nil && buyNum > 0 {
				buyOrder = buyOrders[0]
				processDigital(sellOrder, buyOrder)
			}*/

		fmt.Printf("\nProcessing Digital Orders: Buy- %d , Sell- %d\n", buyNum, sellNum)
		fmt.Printf("\nProcessing Fiat Orders: Buy- %d , Sell- %d\n", buyNum, sellNum)

func processDigital1(sellOrder *models.Order, buyOrders *[]models.Order) {
	fmt.Printf("Sell: %f , Amount: %f \n", sellOrder.Rate, sellOrder.Amount)
	amount := sellOrder.Amount
	for _, buyOrder := range *buyOrders {
		updateDigitalOrder(sellOrder, &buyOrder, amount)
		fmt.Printf("Buy: %f , Amount: %f \n", mBuyOrder.Rate, mBuyOrder.Amount)
		amount -= buyOrder.Amount
	}
}

func updateDigitalOrder(sellOrder *models.Order, buyOrder *models.Order, amount float64) {
	if sellOrder.Rate <= buyOrder.Rate {
		if walletBalanceSufficient(sellOrder, &buyOrder) {
			fmt.Println("walletBalanceSufficient")
			if sellOrder.Amount <= buyOrder.Amount {
				fmt.Println("sellOrder.Amount <= buyOrder.Amount")
				creditWallet(buyOrder.UserId, sellOrder.Amount, sellOrder.CurrencyType)
				debitWallet(sellOrder.UserId, sellOrder.Amount, sellOrder.CurrencyType)
				recordTransaction(sellOrder, sellOrder.UserId, buyOrder.UserId, sellOrder.Amount, sellOrder.CurrencyType)

				creditWallet(sellOrder.UserId, sellOrder.Amount * sellOrder.Rate, sellOrder.RateCurrencyType)
				debitWallet(buyOrder.UserId, sellOrder.Amount * sellOrder.Rate, sellOrder.RateCurrencyType)
				recordTransaction(&buyOrder, buyOrder.UserId, sellOrder.UserId, sellOrder.Amount * sellOrder.Rate, sellOrder.RateCurrencyType)
				
				if sellOrder.Amount < buyOrder.Amount {
					postNewOrder(&buyOrder, buyOrder.Amount - sellOrder.Amount)
				}
			} else if sellOrder.Amount > buyOrder.Amount{
				fmt.Println("sellOrder.Amount > buyOrder.Amount")
				creditWallet(buyOrder.UserId, buyOrder.Amount, sellOrder.CurrencyType)
				debitWallet(sellOrder.UserId, buyOrder.Amount, sellOrder.CurrencyType)
				recordTransaction(sellOrder, sellOrder.UserId, buyOrder.UserId, buyOrder.Amount, sellOrder.CurrencyType)

				creditWallet(sellOrder.UserId, buyOrder.Amount * sellOrder.Rate, sellOrder.RateCurrencyType)
				debitWallet(buyOrder.UserId, buyOrder.Amount * sellOrder.Rate, sellOrder.RateCurrencyType)
				recordTransaction(&buyOrder, buyOrder.UserId, sellOrder.UserId, buyOrder.Amount * sellOrder.Rate, sellOrder.RateCurrencyType)

				postNewOrder(sellOrder, sellOrder.Amount - buyOrder.Amount)
			}
		}
		fmt.Printf("Sell: %f , Amount: %f || Buy: %f , Amount: %f \n", sellOrder.Rate, sellOrder.Amount, buyOrder.Rate, buyOrder.Amount)
	}
}

	"encoding/json"
// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	uid, err := u.GetInt64(":uid")
	if uid > 0 && err == nil {
		user := models.User{Id: uid}
		err := user.Read("id")
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Update() {
	uid := u.GetString(":uid")
	if uid != "" {
		user := new(models.User)
		json.Unmarshal(u.Ctx.Input.RequestBody, user)
		err := user.Update(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

DB_DATASOURCE = "postgres:Coin@123@tcp(127.0.0.1:5432)/auth?charset=utf8"

`orm:"rel(fk)"`

func (m *Container) GetID() int64 {
	return m.Id
}

func Users() orm.QuerySeter {
	var table User
	return orm.NewOrm().QueryTable(table).OrderBy("id")
}

/*
type CustomerInfo struct {
	Name string
	Kind string
}

type CustomClaimsExample struct {
	*jwt.StandardClaims
	TokenType string
	CustomerInfo
}*/

/*
token := jwt.New(jwt.SigningMethodRS512)
token.Claims = &CustomClaimsExample{
	&jwt.StandardClaims{
		// set the expire time
		// see http://tools.ietf.org/html/draft-ietf-oauth-json-web-token-20#section-4.1.4
		ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
	},
	"level1",
	CustomerInfo{"user", "human"},
}*/

		//log.Fatal(tag, err)
		        //sql.Open("postgres", "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full")
        //https://github.com/lib/pq/blob/master/doc.go

                //var w io.Writer
        //orm.DebugLog = orm.NewLog(w)
            //db, err := orm.GetDB()
    //fmt.Println(db, err)

        //"github.com/go-pg/pg"
    //"github.com/go-pg/pg/orm"
        //"io"
    //"database/sql"

    	/*orm.RegisterModelWithPrefix(
	DB_PREFIX,
	new(User))*/

		//"github.com/astaxie/beego/validation"
		/*
func (u *User) Valid(v *validation.Validation) {
	if u.Password != u.Repassword {
		v.SetError("Repassword", "Does not matched password, repassword")
	}
}*/

	user := new(models.User)
	json.Unmarshal(u.Ctx.Input.RequestBody, user)
	user.Insert()
	u.Data["json"] = map[string]string{"uid": fmt.Sprintf("%d",user.GetID())}



// @Title ChangePassword
// @Description change password for the user
// @Param	body		body 	UserChangePassword		true		"Change Password"
// @Success 200 {string} password change success!
// @Failure 403 password change failed
// @router /changepassword [post]
func (u *UserController) ChangePassword() {
	var rqd UserChangePassword
	json.Unmarshal(u.Ctx.Input.RequestBody, &rqd)
	jres, user, tokenString, err := apiAuthenticate(rqd.Token)
	if err == nil {
		password := models.Password{UserId: user.Id, Password: configs.GetSha512(rqd.CurrentPassword), DeletedAt: *configs.NullTime}
		err := password.Read("user_id", "password")
		if err == nil {
			password.DeletedAt = time.Now()
			password.UpdatedAt = time.Now()
			password.Update()
			newpassword := models.Password{UserId: user.Id, Password: configs.GetSha512(rqd.NewPassword), CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: *configs.NullTime}
			err = newpassword.Insert()
			if err == nil {
				jres := JSONResponse{ResponseCode: 200, Description: "Password Changed Successfully"}
				u.Data["json"] = CommonResponse{Token: tokenString, Response: jres}
			} else {
				u.Data["json"] = JSONResponse{ResponseCode: 403, Description: "Invalid User"}
			}
		} else {
			u.Data["json"] = JSONResponse{ResponseCode: 403, Description: "Read and write error"}
		}
	} else {
		u.Data["json"] = jres
	}
	u.ServeJSON()
}

				//} else {
				//	debugMessage("Login Error 2: ", err)
				//	u.Data["json"] = JSONResponse{ResponseCode: 403, Description: "You are trying to use old password"}
				//}