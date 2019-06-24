	/*
	if onlyMine == "YES" {
		icountry_ids := UserCountryIds(userId, eligible)
		num, err = orm.NewOrm().QueryTable(table).Filter("id__in", icountry_ids...).Filter("deleted_at__isnull", true).OrderBy("name").All(&countries)
	} else {

		num, err = orm.NewOrm().QueryTable(table).Filter("id__not_in", configs.GLOBAL_CODE).Filter("deleted_at__isnull", true).OrderBy("name").All(&countries)	
	}*/

	/*
func (c *Country) Insert() error {
	if _, err := orm.NewOrm().Insert(c); err != nil {
		return err
	}
	return nil
}
*/

/*
func (c *Country) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(c, field, fields...)
}

func (c *Country) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(c, fields...); err != nil {
		return err
	}
	return nil
}

func (c *Country) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(c, fields ...); err != nil {
		return err
	}
	return nil
}*/