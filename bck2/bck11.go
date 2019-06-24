//num, err := orm.NewOrm().QueryTable(table).Filter("from_user_id", user.Id).Filter("currency_id", currencyId).Filter("rate_currency_id", rateCurrencyId).OrderBy("-created_at").Limit(limit).All(&orders)

type CurrencyPair struct {
    Id              int64
    FromCurrencyId  int64
    ToCurrencyId    int64
    PairType        string
    Supported       string
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}
