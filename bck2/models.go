type Order struct {
    Id              int64
    UserId          int64
    CurrencyId      int64
    Amount          float64
    Rate            float64
    RateCurrencyId  int64
    OrderType       string
    ProcessedType   string
    ProcessedAt     time.Time
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type OrderPartial struct {
    Id              int64
    ParentOrderId   int64
    ChildOrderId    int64
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}
