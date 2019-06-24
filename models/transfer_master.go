package models

import (
	"github.com/astaxie/beego/orm"
)

func (tm *TransferMaster) TableName() string {
    return "transfer_master"
}

func (tm *TransferMaster) Insert() error {
	if _, err := orm.NewOrm().Insert(tm); err != nil {
		return err
	}
	return nil
}

func (tm *TransferMaster) Read(fields ...string) error {
	if err := orm.NewOrm().Read(tm, fields...); err != nil {
		return err
	}
	return nil
}

func (tm *TransferMaster) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(tm, field, fields...)
}

func (tm *TransferMaster) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(tm, fields...); err != nil {
		return err
	}
	return nil
}

func (tm *TransferMaster) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(tm, fields ...); err != nil {
		return err
	}
	return nil
}

func TransferMasters(userId int64) (int64, []TransferMaster, error) {
	var table TransferMaster
	var transferMasters []TransferMaster
	num, err := orm.NewOrm().QueryTable(table).Filter("user_id", userId).OrderBy("-created_at").All(&transferMasters)
	return num, transferMasters, err
}

func init() {
	orm.RegisterModel(new(TransferMaster))
}