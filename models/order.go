package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	//"fmt"
)

func (o *Order) TableName() string {
    return "order"
}

func (o *Order) Insert() error {
	if _, err := orm.NewOrm().Insert(o); err != nil {
		return err
	}
	return nil
}

func (o *Order) Read(fields ...string) error {
	if err := orm.NewOrm().Read(o, fields...); err != nil {
		return err
	}
	return nil
}

func (o *Order) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(o, field, fields...)
}

func (o *Order) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(o, fields...); err != nil {
		return err
	}
	return nil
}

func (o *Order) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(o, fields ...); err != nil {
		return err
	}
	return nil
}

func Orders(limit int64, currencyId int64, rateCurrencyId int64) (int64, []Order, error) {
	var table Order
	var orders []Order
	num, err := orm.NewOrm().QueryTable(table).Filter("currency_id", currencyId).Filter("rate_currency_id", rateCurrencyId).OrderBy("-created_at").Limit(limit).All(&orders)
	return num, orders, err
}

func MyOrders(user *User, limit int64, currencyId int64, rateCurrencyId int64) (int64, []Order, error) {
	var table Order
	var orders []Order
	o := orm.NewOrm()
	qs := o.QueryTable(table)
	cond := orm.NewCondition()
	var cond1, cond2, cond3 *orm.Condition

	cond1 = cond.And("currency_id", currencyId).And("rate_currency_id", rateCurrencyId).And("seller_user_id", user.Id)
	cond2 = cond.And("currency_id", currencyId).And("rate_currency_id", rateCurrencyId).And("buyer_user_id", user.Id)
	cond3 = cond.AndCond(cond1).OrCond(cond2)

	num, err := qs.SetCond(cond3).Filter("deleted_at__isnull", true).OrderBy("-created_at").Limit(limit).All(&orders)
	return num, orders, err
}

func AllTradeHistory(duration string, currencyId int64, rateCurrencyId int64) (int64, error, float64, float64, float64, []Order) {
	var table Order
	var orders []Order
	o := orm.NewOrm()
	qs := o.QueryTable(table)
	cond := orm.NewCondition()
	var cond1 *orm.Condition

	cond1 = cond.And("currency_id", currencyId).And("rate_currency_id", rateCurrencyId)

	switch duration {
	case "all":
		//cond1 = cond1
	case "1y":
		start := time.Now()
		start = start.AddDate(-1, 0, 0)
		cond1 = cond.And("created_at__gte", start).AndCond(cond1)
	case "1m":
		start := time.Now()
		start = start.AddDate(0, -1, 0)
		cond1 = cond.And("created_at__gte", start).AndCond(cond1)
	case "1w":
		start := time.Now()
		start = start.AddDate(0, 0, -7)
		cond1 = cond.And("created_at__gte", start).AndCond(cond1)
	case "1d":
		start := time.Now()
		start = start.AddDate(0, 0, -1)
		cond1 = cond.And("created_at__gte", start).AndCond(cond1)
	case "12h":
		start := time.Now()
		start = start.Add(-12 * time.Hour)
		cond1 = cond.And("created_at__gte", start).AndCond(cond1)
	case "1h":
		start := time.Now()
		start = start.Add(-1 * time.Hour)
		cond1 = cond.And("created_at__gte", start).AndCond(cond1)
	}

	low := TradeHistoryLow(cond1)
	high := TradeHistoryHigh(cond1)
	last := TradeHistoryLast(cond1)

	num, err := qs.SetCond(cond1).Filter("deleted_at__isnull", true).OrderBy("created_at").All(&orders)

	return num, err, low, high, last, orders
}

func TradeHistoryLow(cond1 *orm.Condition) float64 {
	var table Order
	var orders []Order
	o := orm.NewOrm()
	qs := o.QueryTable(table)
	err := qs.SetCond(cond1).Filter("deleted_at__isnull", true).OrderBy("Rate").One(&orders)
	if err == nil {
		return orders[0].Rate
	}
	return 0
}

func TradeHistoryHigh(cond1 *orm.Condition) float64 {
	var table Order
	var orders []Order
	o := orm.NewOrm()
	qs := o.QueryTable(table)
	err := qs.SetCond(cond1).Filter("deleted_at__isnull", true).OrderBy("-Rate").One(&orders)
	if err == nil {
		return orders[0].Rate
	}
	return 0
}

func TradeHistoryLast(cond1 *orm.Condition) float64 {
	var table Order
	var orders []Order
	o := orm.NewOrm()
	qs := o.QueryTable(table)
	err := qs.SetCond(cond1).Filter("deleted_at__isnull", true).OrderBy("-created_at").One(&orders)
	if err == nil {
		return orders[0].Rate
	}
	return 0
}

func init() {
	orm.RegisterModel(new(Order))
}
