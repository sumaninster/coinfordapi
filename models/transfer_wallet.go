package models

import (
	"github.com/astaxie/beego/orm"
)

func (tw *TransferWallet) TableName() string {
    return "transfer_wallet"
}

func (tw *TransferWallet) Insert() error {
	if _, err := orm.NewOrm().Insert(tw); err != nil {
		return err
	}
	return nil
}

func (tw *TransferWallet) Read(fields ...string) error {
	if err := orm.NewOrm().Read(tw, fields...); err != nil {
		return err
	}
	return nil
}

func (tw *TransferWallet) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(tw, field, fields...)
}

func (tw *TransferWallet) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(tw, fields...); err != nil {
		return err
	}
	return nil
}

func (tw *TransferWallet) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(tw, fields ...); err != nil {
		return err
	}
	return nil
}

func TransferWallets(transferId int64, transferType string) (int64, []TransferWallet, error) {
	var table TransferWallet
	var transferWallets []TransferWallet
	num, err := orm.NewOrm().QueryTable(table).Filter("transfer_id", transferId).Filter("transfer_type", transferType).Filter("deleted_at__isnull", true).OrderBy("-updated_at").All(&transferWallets)
	return num, transferWallets, err
}

func init() {
	orm.RegisterModel(new(TransferWallet))
}