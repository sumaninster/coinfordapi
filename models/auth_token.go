package models

import (
	"github.com/astaxie/beego/orm"
)

func (a *AuthToken) TableName() string {
    return "auth_token"
}

func (a *AuthToken) Insert() error {
	if _, err := orm.NewOrm().Insert(a); err != nil {
		return err
	}
	return nil
}

func (a *AuthToken) Read(fields ...string) error {
	if err := orm.NewOrm().Read(a, fields...); err != nil {
		return err
	}
	return nil
}

func (a *AuthToken) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(a, field, fields...)
}

func (a *AuthToken) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}

func (a *AuthToken) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(a, fields ...); err != nil {
		return err
	}
	return nil
}

func init() {
    orm.RegisterModel(new(AuthToken))
}