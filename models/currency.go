package models

import (
	"github.com/astaxie/beego/orm"
	"coinford_api/configs"
)

func (e *Currency) TableName() string {
    return "currency"
}
/*
func (e *Currency) Insert() error {
	if _, err := orm.NewOrm().Insert(e); err != nil {
		return err
	}
	return nil
}
*/
func (e *Currency) Read(fields ...string) error {
	if err := orm.NewOrm().Read(e, fields...); err != nil {
		return err
	}
	return nil
}
/*
func (e *Currency) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(e, field, fields...)
}

func (e *Currency) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(e, fields...); err != nil {
		return err
	}
	return nil
}

func (e *Currency) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(e, fields ...); err != nil {
		return err
	}
	return nil
}
*/
func GetCurrencyId(currency string) int64 {
	e := Currency{Code: currency}
	err := e.Read("code")
	if err == nil {
		return e.Id
	}
	return 0
}

func UserCurrencyIds(userId int64) []interface{} {
	num, user_currencies, err := Currencies(userId, "")
	currency_ids := []int64{}
	if num > 0 && err == nil {
		for _, v := range user_currencies {
		    currency_ids = append(currency_ids, v.Id)
		}
	}
	icurrency_ids := configs.Int64ToInterface(currency_ids)
	return icurrency_ids
}

func Currencies(userId int64, currencyType string) (int64, []Currency, error) {
	var table Currency
	var currencies []Currency
	var num int64
	var err error
	icountry_ids := UserCountryIds(userId, "YES")

	if currencyType == "FIAT" || currencyType == "CRYPTO" {
		num, err = orm.NewOrm().QueryTable(table).Filter("type", currencyType).Filter("deleted_at__isnull", true).Filter("country_id__in", icountry_ids...).OrderBy("code").All(&currencies)
	} else {
		num, err = orm.NewOrm().QueryTable(table).Filter("deleted_at__isnull", true).Filter("country_id__in", icountry_ids...).OrderBy("code").All(&currencies)	
	}
	return num, currencies, err
}

func init() {
    orm.RegisterModel(new(Currency))
}