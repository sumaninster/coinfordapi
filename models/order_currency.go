package models

import (
	"github.com/astaxie/beego/orm"
)

func (oc *OrderCurrency) TableName() string {
    return "order_currency"
}

func (oc *OrderCurrency) Insert() error {
	if _, err := orm.NewOrm().Insert(oc); err != nil {
		return err
	}
	return nil
}

func (oc *OrderCurrency) Read(fields ...string) error {
	if err := orm.NewOrm().Read(oc, fields...); err != nil {
		return err
	}
	return nil
}

func (oc *OrderCurrency) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(oc, field, fields...)
}

func (oc *OrderCurrency) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(oc, fields...); err != nil {
		return err
	}
	return nil
}

func (oc *OrderCurrency) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(oc, fields ...); err != nil {
		return err
	}
	return nil
}

func MyOrderCurrency(user *User) OrderCurrency {
	orderCurrency := OrderCurrency{UserId: user.Id}
  orderCurrency.Read("user_id")
  return orderCurrency
}

func UpdateOrderCurrency(user *User, currencyId int64, rateCurrencyId int64) error {
	orderCurrency := OrderCurrency{UserId: user.Id, CurrencyId: currencyId, RateCurrencyId: rateCurrencyId}
  _, _, err := orderCurrency.ReadOrCreate("user_id")
  if err == nil {
		orderCurrency.CurrencyId = currencyId
		orderCurrency.RateCurrencyId = rateCurrencyId
		err = orderCurrency.Update()
    return err
  }
  return err
}

func init() {
	orm.RegisterModel(new(OrderCurrency))
}
