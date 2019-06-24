package models

import (
	"github.com/astaxie/beego/orm"
)

func (op *OrderPartial) TableName() string {
    return "order_partial"
}

func (op *OrderPartial) Insert() error {
	if _, err := orm.NewOrm().Insert(op); err != nil {
		return err
	}
	return nil
}

func (op *OrderPartial) Read(fields ...string) error {
	if err := orm.NewOrm().Read(op, fields...); err != nil {
		return err
	}
	return nil
}

func (op *OrderPartial) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(op, field, fields...)
}

func (op *OrderPartial) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(op, fields...); err != nil {
		return err
	}
	return nil
}

func (op *OrderPartial) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(op, fields ...); err != nil {
		return err
	}
	return nil
}

func OrderPartials(user *User) (int64, []OrderPartial, error) {
	var table OrderPartial
	var orderpartials []OrderPartial
	num, err := orm.NewOrm().QueryTable(table).Filter("user_id", user.Id).Filter("deleted_at__isnull", true).OrderBy("-updated_at").All(&orderpartials)
	return num, orderpartials, err
}

func init() {
    orm.RegisterModel(new(OrderPartial))
}