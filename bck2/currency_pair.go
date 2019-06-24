package models

import (
	"github.com/astaxie/beego/orm"
)

func (e *CurrencyPair) TableName() string {
    return "currency_pair"
}

func (e *CurrencyPair) Insert() error {
	if _, err := orm.NewOrm().Insert(e); err != nil {
		return err
	}
	return nil
}

func (e *CurrencyPair) Read(fields ...string) error {
	if err := orm.NewOrm().Read(e, fields...); err != nil {
		return err
	}
	return nil
}

func (e *CurrencyPair) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(e, field, fields...)
}

func (e *CurrencyPair) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(e, fields...); err != nil {
		return err
	}
	return nil
}

func (e *CurrencyPair) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(e, fields ...); err != nil {
		return err
	}
	return nil
}

func IsCurrencyPairSupported(from_currency_id, to_currency_id int64) (bool, string) {
	currencypair := CurrencyPair{FromCurrencyId: from_currency_id, ToCurrencyId: to_currency_id, Supported: "YES"}
	err := currencypair.Read("from_currency_id", "to_currency_id", "supported")
	if err == nil {
		if currencypair.Id > 0 {
			if currencypair.PairType == "FIAT_TO_FIAT" || currencypair.PairType == "FIAT_TO_CRYPTO" || currencypair.PairType == "CRYPTO_TO_FIAT" {
				return true, "Fiat"
			} else if currencypair.PairType == "CRYPTO_TO_CRYPTO" {
				return true, "Digital"
			}
		}
	}
	return false, "Unknown"
}

func ListCurrencyPair(currencypairtype string) (int64, []CurrencyPair, error) {
	var table CurrencyPair
	var currencypairs []CurrencyPair
	var num int64
	var err error
	if currencypairtype == "FIAT_TO_FIAT" || currencypairtype == "FIAT_TO_CRYPTO" || currencypairtype == "CRYPTO_TO_FIAT" || currencypairtype == "CRYPTO_TO_CRYPTO" {
		num, err = orm.NewOrm().QueryTable(table).Filter("pair_type", currencypairtype).Filter("deleted_at__isnull", true).OrderBy("code").All(&currencypairs)
	} else {
		num, err = orm.NewOrm().QueryTable(table).Filter("deleted_at__isnull", true).OrderBy("code").All(&currencypairs)
	}
	return num, currencypairs, err
}

func init() {
    orm.RegisterModel(new(CurrencyPair))
}