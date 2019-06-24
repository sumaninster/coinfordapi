package user

import (
	"coinford_api/models"
	"coinford_api/configs"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"errors"
)

func apiAuthenticate(tokenString string) (CommonResponse, *models.User, bool, error) {
	err := parseToken(tokenString)
	if err == nil {
		authToken := models.AuthToken{Token: tokenString}
		err = authToken.Read("Token")
		if err == nil {
			if time.Now().Before(authToken.ExpirationTime) {
				user := models.User{Id: authToken.UserId}
				err = user.Read("id")
				if err == nil {
					if user.DeletedAt == *configs.NullTime {
						return CommonResponse{}, &user, true, nil
					} else {
						debugMessage("Authenticate Error 2: ", errors.New("User deleted"))
						return CommonResponse{ResponseCode: 403, ResponseDescription: "User deleted."}, nil, false, errors.New("User deleted.")
					}
				} else {
					debugMessage("Authenticate Error 3: ", err)
					return CommonResponse{ResponseCode: 403, ResponseDescription: "Unable to read or write data. Please relogin."}, nil, false, err
				}
			} else {
				debugMessage("Authenticate Error 4: ", errors.New("Token expired"))
				return CommonResponse{ResponseCode: 403, ResponseDescription: "Token expired. Please relogin."}, nil , false, errors.New("Token expired.")
			} 
		} else {
			debugMessage("Authenticate Error 5: ", err)
			return CommonResponse{ResponseCode: 403, ResponseDescription: "Authentication failed. Please relogin."}, nil , false, err
		} 
	} else {
		debugMessage("Authenticate Error 6: ", err)
		return CommonResponse{ResponseCode: 403, ResponseDescription: "Invalid token. Please relogin."}, nil , false, err
	}
	return CommonResponse{ResponseCode: 403, ResponseDescription: "Unrecognized error."}, nil, false, errors.New("Unrecognized error.")
}
/*
func apiAdminAuthenticate(tokenString string) (CommonResponse, *models.User, bool, bool, error) {
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {

	}
}*/

func saveAuthToken(user *models.User) (CommonResponse, string, error){
	expirationTime := time.Now().Add(time.Hour * time.Duration(configs.PostLoginTokenTime))
	tokenString, _ := token(expirationTime.Unix())
	authToken := models.AuthToken{UserId: user.Id, Token: tokenString, ExpirationTime: expirationTime}
	_, _, err := authToken.ReadOrCreate("user_id")
	if err == nil {
		authToken.Token = tokenString
		authToken.ExpirationTime = expirationTime
		authToken.UpdatedAt = time.Now()
		authToken.DeletedAt = *configs.NullTime
		err = authToken.Update()
		if err == nil {
			return CommonResponse{}, tokenString, nil		
		} else {
			debugMessage("saveAuthToken Error 1: ", err)
			return CommonResponse{ResponseCode: 403, ResponseDescription: "Unable to read or write data. Please relogin."}, *configs.NullString, err
		}
	} else {
		debugMessage("saveAuthToken Error 2: ", err)
		return CommonResponse{ResponseCode: 403, ResponseDescription: "Unable to read or write data. Please relogin."}, *configs.NullString, err
	}
	return CommonResponse{ResponseCode: 403, ResponseDescription: "Unrecognized error"}, *configs.NullString, errors.New("Unrecognized error")
}

func token(expirationTime int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, 
    jwt.MapClaims{
        "exp": expirationTime,
        "iat": time.Now().Unix() })

    privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(configs.SignBytes)
	debugMessage("ParseRSAPrivateKeyFromPEM Error: ", err)
	if err != nil {return "Invalid Token", err}

    tokenString, err := token.SignedString(privateKey)
	debugMessage("SignedString Error: ", err)
	if err != nil {return "Invalid Token", err}

	return tokenString, nil
}

func parseToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        publicKey, err := jwt.ParseRSAPublicKeyFromPEM(configs.VerifyBytes)
        return publicKey, err
    })

    if err == nil && token.Valid {
        debugMessage("ParseToken Success: Your token is valid.", nil)
        return nil
    }
    debugMessage("ParseToken Error: ", err)
    return err
}

func debugMessage(tag string, err error) {
	if err != nil && models.Runmode == "dev" {
		fmt.Println(tag, err)
	}
}