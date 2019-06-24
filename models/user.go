package models

import (
	"github.com/astaxie/beego/orm"
)

func (u *User) TableName() string {
    return "user"
}

func (u *User) Insert() error {
	if _, err := orm.NewOrm().Insert(u); err != nil {
		return err
	}
	return nil
}

func (u *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(u, fields...); err != nil {
		return err
	}
	return nil
}

func (u *User) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(u, field, fields...)
}

func (u *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}
	return nil
}

func (u *User) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(u, fields ...); err != nil {
		return err
	}
	return nil
}

func init() {
	orm.RegisterModel(new(User))
}
