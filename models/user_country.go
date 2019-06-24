package models

import (
	"github.com/astaxie/beego/orm"
	"coinford_api/configs"
)

func (uc *UserCountry) TableName() string {
    return "user_country"
}

func (uc *UserCountry) Insert() error {
	if _, err := orm.NewOrm().Insert(uc); err != nil {
		return err
	}
	return nil
}

func (uc *UserCountry) Read(fields ...string) error {
	if err := orm.NewOrm().Read(uc, fields...); err != nil {
		return err
	}
	return nil
}

func (uc *UserCountry) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(uc, field, fields...)
}

func (uc *UserCountry) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(uc, fields...); err != nil {
		return err
	}
	return nil
}

func (uc *UserCountry) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(uc, fields ...); err != nil {
		return err
	}
	return nil
}

func UserCountryIds(userId int64, eligible string) []interface{} {
	num, user_countries, err := UserCountries(userId, eligible)
	country_ids := []int64{configs.GLOBAL_CODE}
	if num > 0 && err == nil {
		for _, v := range user_countries {
		    country_ids = append(country_ids, v.CountryId)
		}
	}
	icountry_ids := configs.Int64ToInterface(country_ids)
	return icountry_ids
}

func IsEligible(userId int64, countryId int64) bool {
	var table UserCountry
	var usercountries []UserCountry
	num, err := orm.NewOrm().QueryTable(table).Filter("user_id", userId).Filter("country_id", countryId).Filter("eligible", "YES").Filter("deleted_at__isnull", true).All(&usercountries)
	if num == 1 && err == nil{
		return true
	}
	return false
}

func UserCountries(userId int64, eligible string) (int64, []UserCountry, error) {
	var table UserCountry
	var usercountries []UserCountry
	var num int64
	var err error
	if eligible == "YES" || eligible == "NO" {
		num, err = orm.NewOrm().QueryTable(table).Filter("user_id", userId).Filter("eligible", eligible).Filter("deleted_at__isnull", true).OrderBy("country_id").All(&usercountries)
	} else {
		num, err = orm.NewOrm().QueryTable(table).Filter("user_id", userId).Filter("deleted_at__isnull", true).OrderBy("country_id").All(&usercountries)
	}
	return num, usercountries, err
}

func init() {
	orm.RegisterModel(new(UserCountry))
	
}
