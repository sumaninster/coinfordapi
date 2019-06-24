package models

import (
	"github.com/astaxie/beego/orm"
)

func (pm *PayeeMaster) TableName() string {
    return "payee_master"
}

func (pm *PayeeMaster) Insert() error {
	if _, err := orm.NewOrm().Insert(pm); err != nil {
		return err
	}
	return nil
}

func (pm *PayeeMaster) Read(fields ...string) error {
	if err := orm.NewOrm().Read(pm, fields...); err != nil {
		return err
	}
	return nil
}

func (pm *PayeeMaster) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(pm, field, fields...)
}

func (pm *PayeeMaster) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(pm, fields...); err != nil {
		return err
	}
	return nil
}

func (pm *PayeeMaster) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(pm, fields ...); err != nil {
		return err
	}
	return nil
}

func PayeeMasters(userId int64) (int64, []PayeeMaster, error) {
	var table PayeeMaster
	var payeeMasters []PayeeMaster
	num, err := orm.NewOrm().QueryTable(table).Filter("user_id", userId).OrderBy("-created_at").All(&payeeMasters)
	return num, payeeMasters, err
}

func init() {
	orm.RegisterModel(new(PayeeMaster))
}