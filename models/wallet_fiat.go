package models

import (
	"github.com/astaxie/beego/orm"
)

func (wf *WalletFiat) TableName() string {
    return "wallet_address"
}

func (wf *WalletFiat) Insert() error {
	if _, err := orm.NewOrm().Insert(wf); err != nil {
		return err
	}
	return nil
}

func (wf *WalletFiat) Read(fields ...string) error {
	if err := orm.NewOrm().Read(wf, fields...); err != nil {
		return err
	}
	return nil
}

func (wf *WalletFiat) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(wf, field, fields...)
}

func (wf *WalletFiat) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(wf, fields...); err != nil {
		return err
	}
	return nil
}

func (wf *WalletFiat) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(wf, fields ...); err != nil {
		return err
	}
	return nil
}

func WalletFiats(walletId int64) (int64, []WalletFiat, error) {
	var table WalletFiat
	var walletFiats []WalletFiat
	num, err := orm.NewOrm().QueryTable(table).Filter("wallet_id", walletId).OrderBy("-created_at").All(&walletFiats)
	return num, walletFiats, err
}

func init() {
    orm.RegisterModel(new(WalletFiat))
}