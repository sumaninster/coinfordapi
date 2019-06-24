package models

import (
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    "time"
    "fmt"
    _ "github.com/lib/pq"
)

var (
    DB_PREFIX string
    DB_USER string
    DB_PASSWORD string
    DB_HOST string
    DB_NAME string
    DB_PORT string
    Runmode string
)

type Admin struct {
}

type AdminAccess struct {
}

type AuthToken struct {
    Id              int64
    UserId          int64
    Token           string      `orm:"unique"`
    ExpirationTime  time.Time
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type Country struct {
    Id              int64
    Name            string
    IsoCode         string
    DialCode        string
    Code            string
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type Currency struct {
    Id              int64
    Code            string
    Name            string
    Description     string
    Type            string
    CountryId       int64
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type Data struct {
    Id              int64
    UserId          int64
    CountryId       int64
    FieldType       string
    Nickname        string
    Primary         bool
    Active          bool
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type DataText struct {
    Id              int64
    DataId          int64
    FieldId         int64
    Text            string
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type DataCategory struct {
    Id              int64
    DataId          int64
    FieldId         int64
    FieldCategoryId int64
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type DataFile struct {
    Id              int64
    DataId          int64
    FieldId         int64
    Name            string
    Extension       string
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type Email struct {
    Id              int64
    UserId          int64
    Email           string      `valid:"Required;Email;MaxSize(250)"`
    Primary         string
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type Field struct {
    Id              int64
    CountryId       int64
    Name            string
    Description     string
    FieldType       string
    DataType        string
    Order           int64
    Key             string
    HasCategory     bool
    HasInputText    bool
    HasFile         bool
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type FieldCategory struct {
    Id              int64
    FieldId         int64
    Name            string
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type FieldSubcategory struct {
    Id              int64
    FieldId         int64
    SubFieldId      int64
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type GlobalSetting struct {
    Id              int64
    CountryId       int64
    Key             string
    Value           string
    AdminId         string
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type Mobile struct {
    Id              int64
    UserId          int64
    Number          string
    Primary         string
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type Order struct {
    Id              int64
    SellerUserId    int64
    BuyerUserId     int64
    CurrencyId      int64
    RateCurrencyId  int64
    Amount          float64
    Rate            float64
    TotalAmount     float64
    SellerCurrencyWalletId      int64
    SellerRateCurrencyWalletId  int64
    BuyerCurrencyWalletId       int64
    BuyerRateCurrencyWalletId   int64
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type OrderBuy struct {
    Id                    int64
    UserId                int64
    CurrencyId            int64
    RateCurrencyId        int64
    Amount                float64
    Rate                  float64
    TotalAmount           float64
    CurrencyWalletId      int64
    RateCurrencyWalletId  int64
    Lock                  bool
    CreatedAt             time.Time   `orm:"auto_now_add;type(datetime)"`
}

type OrderCurrency struct {
    Id              int64
    UserId          int64
    CurrencyId      int64
    RateCurrencyId  int64
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type OrderGraph struct {
    Id                    int64
    last_order_id         int64
    CurrencyId            int64
    RateCurrencyId        int64
    Open                  float64
    High                  float64
    Low                   float64
    Close                 float64
    Volume                float64
    Split                 float64
    Dividend              float64
    AbsoluteChange        float64
    PercentChange         float64
    Date                  time.Time   `orm:"auto_now_add;type(datetime)"`
}

type OrderSell struct {
    Id                    int64
    UserId                int64
    CurrencyId            int64
    RateCurrencyId        int64
    Amount                float64
    Rate                  float64
    TotalAmount           float64
    CurrencyWalletId      int64
    RateCurrencyWalletId  int64
    Lock                  bool
    CreatedAt             time.Time   `orm:"auto_now_add;type(datetime)"`
}

type OrderWallet struct {
    Id                    int64
    UserId                int64
    CurrencyId            int64
    RateCurrencyId        int64
    CurrencyWalletId      int64
    RateCurrencyWalletId  int64
    CreatedAt             time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt             time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt             time.Time
}

type Password struct {
    Id              int64
    UserId          int64
    Password        string      `valid:"Required;MinSize(8)"`
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}


type Payee struct {
    Id              int64
    PayeeMasterId   int64
    Name            string
    Nickname        string
    Email           string
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type PayeeCrypto struct {
    Id              int64
    PayeeId          int64
    Address         string
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type PayeeFiat struct {
    Id              int64
    PayeeId         int64
    DataId          int64
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type PayeeMaster struct {
    Id              int64
    UserId          int64
    CurrencyId      int64
    CurrencyType    string
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}
type Transaction struct {
    Id              int64
    OrderId         int64
    FromUser        int64
    ToUser          int64
    CurrencyId      int64
    Amount          float64
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type Transfer struct {
    Id              int64
    TransferMasterId int64
    Amount          float64
    TransferType    string
    FromUserId      int64
    ToUserId        int64
    ProcessedAt     time.Time
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type TransferBank struct {
    Id              int64
    TransferId      int64
    FromWalletId    int64
    ToDataId        int64
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type TransferMaster struct {
    Id              int64
    UserId          int64
    CurrencyId      int64
    CurrencyType    string
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type TransferPayee struct {
    Id              int64
    TransferId      int64
    FromWalletId    int64
    ToPayeeId       int64
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type TransferWallet struct {
    Id              int64
    TransferId      int64
    FromWalletId    int64
    ToWalletId      int64
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type User struct {
    Id              int64
    Name            string      `valid:"Required;MaxSize(250)"`
    Username        string      `valid:"Required;MaxSize(250)"`
    EditNameTimes   int
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type UserCountry struct {
    Id              int64
    UserId          int64
    CountryId       int64
    Eligible        string
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type UserSetting struct {
    Id              int64
    UserId          int64
    Key             string
    Value           string
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type Wallet struct {
    Id              int64
    WalletMasterId  int64
    Amount          float64
    AmountLocked    float64
    Nickname        string
    Primary         bool
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type WalletCrypto struct {
    Id              int64
    WalletId        int64
    Address         string
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type WalletFiat struct {
    Id              int64
    WalletId        int64
    DataId          int64
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type WalletMaster struct {
    Id              int64
    UserId          int64
    CurrencyId      int64
    CurrencyType    string
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type WalletPassphrase struct {
    Id              int64
    WalletId        int64
    Passphrase      string
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

func init() {
    DB_PREFIX = beego.AppConfig.String("DB_PREFIX")
    DB_USER = beego.AppConfig.String("DB_USER")
    DB_PASSWORD = beego.AppConfig.String("DB_PASSWORD")
    DB_HOST = beego.AppConfig.String("DB_HOST")
    DB_NAME = beego.AppConfig.String("DB_NAME")
    DB_PORT = beego.AppConfig.String("DB_PORT")
    Runmode = beego.AppConfig.String("runmode")

    switch Runmode {
    case "dev":
        orm.RegisterDriver("postgres", orm.DRPostgres)
        orm.RegisterDataBase("default", "postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?port=%i&sslmode=disable",
        DB_USER, DB_PASSWORD, DB_HOST, DB_NAME, DB_PORT), 30)
        orm.Debug = true
    }
}
