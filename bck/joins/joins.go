package joins

import (
	//"coinford_api/configs"
	"coinford_api/models"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    "time"
    "fmt"
    _ "github.com/lib/pq"
)

var (
    DB_PREFIX string
    DB_USER string
    DB_PASSWORD string
    DB_HOST string
    DB_NAME string
    DB_PORT string
    Runmode string
)

type Country1 struct {
    Id              int64
    Name            string 
    IsoCode         string
    DialCode        string
    Code            string
    CreatedAt       time.Time
    UpdatedAt       time.Time
    DeletedAt       time.Time
}

type UserCountry1 struct {
    Id              int64
    UserId          int64
    Eligible        string
    CreatedAt       time.Time
    UpdatedAt       time.Time
    DeletedAt       time.Time
    Country         *Country1  `orm:"rel(fk)"`
}

func UserDetailCountries(user *models.User, eligible string) (int64, []UserCountry1, error) {
	//var table UserCountry
	//var usercountries []UserCountryDetail
	//orm.RegisterModel(new(UserCountry))
	var usercountries []UserCountry1
	var num int64
	var err error
	if eligible == "YES" || eligible == "NO" {
		num, err = orm.NewOrm().QueryTable("user_country").Filter("user_id", 9).Filter("eligible", eligible).Filter("deleted_at__isnull", true).OrderBy("country_id").RelatedSel().All(&usercountries)
		fmt.Println(usercountries)
	} else {
		num, err = orm.NewOrm().QueryTable("user_country").Filter("user_id", 9).Filter("deleted_at__isnull", true).OrderBy("country_id").RelatedSel().All(&usercountries)
		fmt.Println(usercountries)
	}
	return num, usercountries, err
}

func init() {
    DB_PREFIX = beego.AppConfig.String("DB_PREFIX")
    DB_USER = beego.AppConfig.String("DB_USER")
    DB_PASSWORD = beego.AppConfig.String("DB_PASSWORD")
    DB_HOST = beego.AppConfig.String("DB_HOST")
    DB_NAME = beego.AppConfig.String("DB_NAME")
    DB_PORT = beego.AppConfig.String("DB_PORT")
    Runmode = beego.AppConfig.String("runmode")
	orm.RegisterModel(new(UserCountry1))
	orm.RegisterModel(new(Country1))
    switch Runmode {
    case "dev":
        orm.RegisterDriver("postgres", orm.DRPostgres)
        orm.RegisterDataBase("default", "postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?port=%i&sslmode=disable",
        DB_USER, DB_PASSWORD, DB_HOST, DB_NAME, DB_PORT), 30)
        //orm.Debug = true
    }
}