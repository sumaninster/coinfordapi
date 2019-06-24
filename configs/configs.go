package configs

import (
	"github.com/astaxie/beego"
	"io/ioutil"
	"crypto/sha512"
	"encoding/hex"
	"log"
	"time"
	"math/rand"
	)

var (
	SignBytes []byte
	VerifyBytes []byte
	PreLoginTokenTime int
	PostLoginTokenTime int
	EditNameMaximumTimes int
	err error

	FILE_PATH string
	GLOBAL_CODE int64
	MaxUnverifiedCountry int64

	NullTime	*time.Time
	NullString	*string

	BTC_IN_ACCOUNT string
	BTC_OUT_ACCOUNT string
	BTC_VAULT_ACCOUNT string
	BTC_OUT_ADDRESS string
	BTC_MIN_CONF int

	ETHAuthUrl string
	ETHPasswordExpiryHour int64
	ETH_OUT_ADDRESS string
	ETH_MIN_CONF int
	ETH_OUT_PASSPHRASE string
)

func Init() {
	SignBytes, err = ioutil.ReadFile("conf/app.rsa")
	fatal("SignBytes: ", err)

	VerifyBytes, err = ioutil.ReadFile("conf/app.rsa.pub")
	fatal("VerifyBytes: ", err)

	PreLoginTokenTime, _ = beego.AppConfig.Int("preLoginTokenTime")
	PostLoginTokenTime, _ = beego.AppConfig.Int("postLoginTokenTime")
	EditNameMaximumTimes, _ = beego.AppConfig.Int("editNameMaximumTimes")

	GLOBAL_CODE, _ = beego.AppConfig.Int64("GLOBAL_CODE")
	MaxUnverifiedCountry, _ = beego.AppConfig.Int64("MaxUnverifiedCountry")

	FILE_PATH = beego.AppConfig.String("FILE_PATH")

	BTC_IN_ACCOUNT = beego.AppConfig.String("BTC_IN_ACCOUNT")
	BTC_OUT_ACCOUNT = beego.AppConfig.String("BTC_OUT_ACCOUNT")
	BTC_VAULT_ACCOUNT = beego.AppConfig.String("BTC_VAULT_ACCOUNT")
	BTC_OUT_ADDRESS = beego.AppConfig.String("BTC_OUT_ADDRESS")
	BTC_MIN_CONF, _ = beego.AppConfig.Int("BTC_MIN_CONF")

	NullTime = new(time.Time)
	NullString = new(string)

	ETHAuthUrl = "http://localhost:8545"
	ETHPasswordExpiryHour = 1000000
	ETH_OUT_ADDRESS = beego.AppConfig.String("ETH_OUT_ADDRESS")
	ETH_OUT_PASSPHRASE = beego.AppConfig.String("ETH_OUT_PASSPHRASE")
	ETH_MIN_CONF, _ = beego.AppConfig.Int("ETH_MIN_CONF")
}

func GetSha512(s string) string {
    h := sha512.New()
    h.Write([]byte(s))
    return hex.EncodeToString(h.Sum([]byte(s)))
}

func RandString(str_size int) string {
    alphanum := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!@#$%^&*()_-+=<>?:;"
    var bytes = make([]byte, str_size)
    rand.Read(bytes)
    for i, b := range bytes {
        bytes[i] = alphanum[b%byte(len(alphanum))]
    }
    return string(bytes)
}

func Int64ToInterface(t []int64) []interface{} {
	s := make([]interface{}, len(t))
	for i, v := range t {
	    s[i] = v
	}
	return s
}

func fatal(tag string, err error) {
	if err != nil {
		log.Fatal(tag, err)
	}
}