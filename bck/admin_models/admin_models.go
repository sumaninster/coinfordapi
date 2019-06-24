package admin_models

import (
    "time"
)

type Admin struct {
    Id              int64
    Name            string      `valid:"Required;MaxSize(250)"`
    Username        string      `valid:"Required;MaxSize(250)"`
    EditNameTimes   int
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type AdminAccess struct {

}

type AdminPassword struct {
    Id              int64
    UserId          int64 
    Password        string      `valid:"Required;MinSize(8)"`
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type AdminToken struct {
    Id              int64
    UserId          int64 
    Token           string      `orm:"unique"`
    ExpirationTime  time.Time
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type AdminCountry struct {
    Id              int64
    UserId          int64
    CountryId       int64
    Eligible        string
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type AdminSetting struct {
    Id              int64
    UserId          int64 
    Key             string
    Value           string
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}