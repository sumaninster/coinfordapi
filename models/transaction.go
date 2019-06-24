package models

import (
	"github.com/astaxie/beego/orm"
)

func (t *Transaction) TableName() string {
    return "transaction"
}

func (t *Transaction) Insert() error {
	if _, err := orm.NewOrm().Insert(t); err != nil {
		return err
	}
	return nil
}

func (t *Transaction) Read(fields ...string) error {
	if err := orm.NewOrm().Read(t, fields...); err != nil {
		return err
	}
	return nil
}

func (t *Transaction) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(t, field, fields...)
}

func (t *Transaction) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(t, fields...); err != nil {
		return err
	}
	return nil
}

func (t *Transaction) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(t, fields ...); err != nil {
		return err
	}
	return nil
}

func Transactions(user *User) (int64, []Transaction, error) {
	var table Transaction
	var transactions []Transaction
	num, err := orm.NewOrm().QueryTable(table).Filter("user_id", user.Id).Filter("deleted_at__isnull", true).OrderBy("-updated_at").All(&transactions)
	return num, transactions, err
}

func init() {
    orm.RegisterModel(new(Transaction))
}