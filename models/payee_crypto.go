package models

import (
	"github.com/astaxie/beego/orm"
)

func (pc *PayeeCrypto) TableName() string {
    return "payee_crypto"
}

func (pc *PayeeCrypto) Insert() error {
	if _, err := orm.NewOrm().Insert(pc); err != nil {
		return err
	}
	return nil
}

func (pc *PayeeCrypto) Read(fields ...string) error {
	if err := orm.NewOrm().Read(pc, fields...); err != nil {
		return err
	}
	return nil
}

func (pc *PayeeCrypto) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(pc, field, fields...)
}

func (pc *PayeeCrypto) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(pc, fields...); err != nil {
		return err
	}
	return nil
}

func (pc *PayeeCrypto) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(pc, fields ...); err != nil {
		return err
	}
	return nil
}

func ListPayeeCryptos(userId int64, currencyId int64) (int64, []PayeeCrypto, error) {
	var table PayeeCrypto
	var payeeCryptos []PayeeCrypto
	num, err := orm.NewOrm().QueryTable(table).Filter("user_id", userId).Filter("currency_id", currencyId).Filter("deleted_at__isnull", true).OrderBy("updated_at").All(&payeeCryptos)
	return num, payeeCryptos, err
}

func init() {
	orm.RegisterModel(new(PayeeCrypto))
}