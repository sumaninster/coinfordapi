package models

import (
	"github.com/astaxie/beego/orm"
)

func (tp *TransferPayee) TableName() string {
    return "transfer_payee"
}

func (tp *TransferPayee) Insert() error {
	if _, err := orm.NewOrm().Insert(tp); err != nil {
		return err
	}
	return nil
}

func (tp *TransferPayee) Read(fields ...string) error {
	if err := orm.NewOrm().Read(tp, fields...); err != nil {
		return err
	}
	return nil
}

func (tp *TransferPayee) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(tp, field, fields...)
}

func (tp *TransferPayee) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(tp, fields...); err != nil {
		return err
	}
	return nil
}

func (tp *TransferPayee) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(tp, fields ...); err != nil {
		return err
	}
	return nil
}

func init() {
	orm.RegisterModel(new(TransferPayee))
}