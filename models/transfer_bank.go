package models

import (
	"github.com/astaxie/beego/orm"
)

func (tb *TransferBank) TableName() string {
    return "transfer_bank"
}

func (tb *TransferBank) Insert() error {
	if _, err := orm.NewOrm().Insert(tb); err != nil {
		return err
	}
	return nil
}

func (tb *TransferBank) Read(fields ...string) error {
	if err := orm.NewOrm().Read(tb, fields...); err != nil {
		return err
	}
	return nil
}

func (tb *TransferBank) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(tb, field, fields...)
}

func (tb *TransferBank) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(tb, fields...); err != nil {
		return err
	}
	return nil
}

func (tb *TransferBank) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(tb, fields ...); err != nil {
		return err
	}
	return nil
}

func init() {
	orm.RegisterModel(new(TransferBank))
}