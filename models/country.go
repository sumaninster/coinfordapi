package models

import (
	"github.com/astaxie/beego/orm"
	//"coinford_api/configs"
)

func (c *Country) TableName() string {
    return "country"
}

func (c *Country) Read(fields ...string) error {
	if err := orm.NewOrm().Read(c, fields...); err != nil {
		return err
	}
	return nil
}

func Countries(userId int64, onlyMine string, eligible string) (int64, []Country, error) {
	var table Country
	var countries []Country
	var num int64
	var err error
	o := orm.NewOrm()
	qs := o.QueryTable(table)
	cond := orm.NewCondition()
	var cond1 *orm.Condition
	icountry_ids := UserCountryIds(userId, eligible)
	if onlyMine == "YES" {
		cond1 = cond.And("id__in", icountry_ids...)
	} else {
		cond1 = cond.AndNot("id__in", icountry_ids...)//configs.GLOBAL_CODE)
	}
	num, err = qs.SetCond(cond1).Filter("deleted_at__isnull", true).OrderBy("name").All(&countries)

	return num, countries, err
}

func init() {
	orm.RegisterModel(new(Country))
}