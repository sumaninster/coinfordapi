package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

func (ow *OrderWallet) TableName() string {
    return "order_wallet"
}

func (ow *OrderWallet) Insert() error {
	if _, err := orm.NewOrm().Insert(ow); err != nil {
		return err
	}
	return nil
}

func (ow *OrderWallet) Read(fields ...string) error {
	if err := orm.NewOrm().Read(ow, fields...); err != nil {
		return err
	}
	return nil
}

func (ow *OrderWallet) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(ow, field, fields...)
}

func (ow *OrderWallet) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(ow, fields...); err != nil {
		return err
	}
	return nil
}

func (ow *OrderWallet) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(ow, fields ...); err != nil {
		return err
	}
	return nil
}

func MyOrderWallet(user *User, currencyId int64, rateCurrencyId int64) *OrderWallet {
	orderWallet := OrderWallet{UserId: user.Id, CurrencyId: currencyId, RateCurrencyId: rateCurrencyId}
  err := orderWallet.Read("user_id", "currency_id", "rate_currency_id")
	if(err == nil) {
		return &orderWallet
	} else {
		return nil
	}
}

func UpdateOrderWallet(user *User, currencyId int64, rateCurrencyId int64, currencyWalletId int64, rateCurrencyWalletId int64) error {
	orderWallet := OrderWallet{UserId: user.Id,
		CurrencyId: currencyId, RateCurrencyId: rateCurrencyId,
		CurrencyWalletId: currencyWalletId, RateCurrencyWalletId: rateCurrencyWalletId}
  _, _, err := orderWallet.ReadOrCreate("user_id", "currency_id", "rate_currency_id")
	fmt.Println(err, orderWallet)
  if err == nil {
		orderWallet.CurrencyWalletId = currencyWalletId
		orderWallet.RateCurrencyWalletId = rateCurrencyWalletId
		err = orderWallet.Update()
    return err
  }
  return err
}

func init() {
    orm.RegisterModel(new(OrderWallet))
}
