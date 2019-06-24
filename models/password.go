package models

import (
	"github.com/astaxie/beego/orm"
)

func (p *Password) TableName() string {
    return "password"
}

func (p *Password) Insert() error {
	if _, err := orm.NewOrm().Insert(p); err != nil {
		return err
	}
	return nil
}

func (p *Password) Read(fields ...string) error {
	if err := orm.NewOrm().Read(p, fields...); err != nil {
		return err
	}
	return nil
}

func (p *Password) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(p, field, fields...)
}

func (p *Password) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(p, fields...); err != nil {
		return err
	}
	return nil
}

func (p *Password) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(p, fields ...); err != nil {
		return err
	}
	return nil
}

func (p *Password) IsValid(u *User) (bool, Password){
	var table Password
	var passwords []Password
	num, err := orm.NewOrm().QueryTable(table).Filter("user_id", u.Id).Filter("password", p.Password).Filter("deleted_at__isnull", true).OrderBy("-updated_at").All(&passwords)
	if err == nil {
		if num == 1 {
			return true, passwords[0]
		}
		return false, Password{}
	}
	return false, Password{}
}

func init() {
	orm.RegisterModel(new(Password))
}
