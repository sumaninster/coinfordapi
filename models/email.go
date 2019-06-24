package models

import (
	"github.com/astaxie/beego/orm"
)

func (e *Email) TableName() string {
    return "email"
}

func (e *Email) Insert() error {
	if _, err := orm.NewOrm().Insert(e); err != nil {
		return err
	}
	return nil
}

func (e *Email) Read(fields ...string) error {
	if err := orm.NewOrm().Read(e, fields...); err != nil {
		return err
	}
	return nil
}

func (e *Email) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(e, field, fields...)
}

func (e *Email) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(e, fields...); err != nil {
		return err
	}
	return nil
}

func (e *Email) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(e, fields ...); err != nil {
		return err
	}
	return nil
}

func Emails(user *User) (int64, []Email, error) {
	var table Email
	var emails []Email
	num, err := orm.NewOrm().QueryTable(table).Filter("user_id", user.Id).Filter("deleted_at__isnull", true).OrderBy("-updated_at").All(&emails)
	return num, emails, err
}

func init() {
    orm.RegisterModel(new(Email))
}