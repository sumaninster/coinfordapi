package models

import (
	"github.com/astaxie/beego/orm"
)

func (m *Mobile) TableName() string {
    return "mobile"
}

func (m *Mobile) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Mobile) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Mobile) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(m, field, fields...)
}

func (m *Mobile) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Mobile) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(m, fields ...); err != nil {
		return err
	}
	return nil
}

func Mobiles(user *User) (int64, []Mobile, error) {
	var table Mobile
	var mobiles []Mobile
	num, err := orm.NewOrm().QueryTable(table).Filter("user_id", user.Id).Filter("deleted_at__isnull", true).OrderBy("-updated_at").All(&mobiles)
	return num, mobiles, err
}

func init() {
    orm.RegisterModel(new(Mobile))
}