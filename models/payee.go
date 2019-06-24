package models

import (
	"github.com/astaxie/beego/orm"
)

func (p *Payee) TableName() string {
    return "payee"
}

func (p *Payee) Insert() error {
	if _, err := orm.NewOrm().Insert(p); err != nil {
		return err
	}
	return nil
}

func (p *Payee) Read(fields ...string) error {
	if err := orm.NewOrm().Read(p, fields...); err != nil {
		return err
	}
	return nil
}

func (p *Payee) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(p, field, fields...)
}

func (p *Payee) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(p, fields...); err != nil {
		return err
	}
	return nil
}

func (p *Payee) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(p, fields ...); err != nil {
		return err
	}
	return nil
}

func Payees(payeeMasterId int64) (int64, []Payee, error) {
	var table Payee
	var payees []Payee
	num, err := orm.NewOrm().QueryTable(table).Filter("payee_master_id", payeeMasterId).Filter("deleted_at__isnull", true).OrderBy("-updated_at").All(&payees)
	return num, payees, err
}

func init() {
	orm.RegisterModel(new(Payee))
}