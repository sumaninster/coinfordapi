package models

import (
	"github.com/astaxie/beego/orm"
	"coinford_api/configs"
)

func (t *Transfer) TableName() string {
    return "transfer"
}

func (t *Transfer) Insert() error {
	if _, err := orm.NewOrm().Insert(t); err != nil {
		return err
	}
	return nil
}

func (t *Transfer) Read(fields ...string) error {
	if err := orm.NewOrm().Read(t, fields...); err != nil {
		return err
	}
	return nil
}

func (t *Transfer) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(t, field, fields...)
}

func (t *Transfer) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(t, fields...); err != nil {
		return err
	}
	return nil
}

func (t *Transfer) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(t, fields ...); err != nil {
		return err
	}
	return nil
}

func Transfers(transferMasterId int64) (int64, []Transfer, error) {
	var table Transfer
	var transfers []Transfer
	num, err := orm.NewOrm().QueryTable(table).Filter("transfer_master_id", transferMasterId).Filter("deleted_at__isnull", true).OrderBy("-updated_at").All(&transfers)
	return num, transfers, err
}

func (t *Transfer) CryptoAddress() string {
	transferPayee := TransferPayee{TransferId: t.Id}
	transferPayee.Read()
	payeeCrypto := PayeeCrypto{Id: transferPayee.ToPayeeId}
	payeeCrypto.Read()
	return payeeCrypto.Address
}

func UserTransferIds(transferMasterId int64) []interface{} {
	num, transfers, err := Transfers(transferMasterId)
	transferIds := []int64{}
	if num > 0 && err == nil {
		for _, v := range transfers {
		    transferIds = append(transferIds, v.Id)
		}
	}
	iTransferIds := configs.Int64ToInterface(transferIds)
	return iTransferIds
}
/*
func UserTransfers(userId int64, currencyId int64) (int64, []Transfer, error) {
	var table Transfer
	var transfers []Transfer
	num, err := orm.NewOrm().QueryTable(table).Filter("user_id", userId).Filter("currency_id", currencyId).Filter("deleted_at__isnull", true).OrderBy("-updated_at").All(&transfers)
	return num, transfers, err
}*/

func TransferProcess(currencyId int64) (int64, []Transfer, error) {
	var table Transfer
	var transfers []Transfer
	num, err := orm.NewOrm().QueryTable(table).Filter("currency_id", currencyId).Filter("deleted_at__isnull", true).OrderBy("-updated_at").All(&transfers)
	return num, transfers, err
}

func init() {
	orm.RegisterModel(new(Transfer))
}