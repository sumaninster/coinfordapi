
func (dc *DataController) prepareDataList(userId int64, countryId int64, fieldType string, message string) DataList {
    _, fields, _ := models.Fields(countryId, fieldType)
    _, datas, _ := models.DataList(userId, countryId, fieldType)
    var datagroups []DataGroup
    var err error
    for _, data := range datas {
        var datadetails []DataDetail
        for _, field := range fields {
            datadetail := DataDetail{Field: &field}
            switch field.DataType {
            case "TEXT":
                dt := models.DataText{DataId: data.Id, FieldId: field.Id}
                err = dt.Read("data_id", "field_id")
                if err == nil {
                    datadetail.DataText = &dt
                }
            case "CATEGORY":
                dc := models.DataCategory{DataId: data.Id, FieldId: field.Id}
                err := dc.Read("data_id", "field_id")
                if err == nil {
                    fc := models.FieldCategory{Id: dc.FieldCategoryId}
                    err = fc.Read("id")
                    if err == nil {
                        datadetail.FieldCategory = &fc
                    }
                    datadetail.DataCategory = &dc
                }
            case "FILE":
                df := models.DataFile{DataId: data.Id, FieldId: field.Id}
                err := df.Read("data_id", "field_id")
                if err == nil {
                    datadetail.DataFile = &df
                }
            }
            datadetails = append(datadetails, datadetail)
        }
        datagroup := DataGroup{Data: data, DataDetail: datadetails}
        datagroups = append(datagroups, datagroup)
    }
    return DataList{DataGroup: datagroups, ResponseCode: 200, ResponseDescription: "List of Data"}
}

/*
type Bank struct {
    Id              int64
    UserId          int64 
    CountryId       int64
    Nickname        string
    Primary         string
    Active          string
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type BankDetail struct {
    Id              int64
    BankId          int64 
    FieldId         int64
    Value           string
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}*/

/*
type KycDetail struct {
    Id              int64
    UserId          int64
    FieldId         int64
    Value           string
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type KycDocument struct {
    Id              int64
    UserId          int64
    FieldId         int64
    Value           string
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}*/
//Response
/*
type DataDetail struct {    
    FieldId                 int64
    CountryId               int64 
    FieldName               string
    FieldDescription        string
    FieldType               string
    FieldOrder              int64
    HasInputText            bool
    HasCategory             bool
    HasFile                 bool
    DataId                  int64
    Text                    string
    FieldCategoryId         int64
    FieldCategoryName       string
    FileName                string
    FileExtension           string
}

type DataList struct {  
    DataDetail              []DataDetail
}

type DataAll struct {
    DataList                []DataList
    ResponseCode            int
    ResponseDescription     string
}*/
//user : data
/*
func (dc *DataController) prepareDataList(userId int64, countryId int64, fieldType string, message string) DataAll {
    _, fields, _ := models.Fields(countryId, fieldType)
    _, datas, _ := models.DataList(userId, countryId, fieldType)
    var datalists []DataList
    for _, data := range datas {
        var datadetails []DataDetail
        for _, field := range fields {
            datadetail := DataDetail{FieldId: field.Id, CountryId: countryId, FieldName: field.Name, FieldDescription: field.Description, 
                FieldType: field.FieldType, FieldOrder: field.Order, HasInputText: field.HasInputText, HasCategory: field.HasCategory, 
                HasFile: field.HasFile, DataId: data.Id}
            if field.HasInputText {
                dt := models.DataText{DataId: data.Id, FieldId: field.Id}
                err := dt.Read("data_id", "field_id")
                if err == nil {
                    datadetail.Text = dt.Text
                    datadetails = append(datadetails, datadetail)
                }
            } else if field.HasCategory {
                dc := models.DataCategory{DataId: data.Id, FieldId: field.Id}
                err := dc.Read("data_id", "field_id")
                if err == nil {
                    cn := models.FieldCategory{Id: dc.FieldCategoryId, FieldId: field.Id}
                    err = cn.Read("id","field_id")
                    if err == nil {
                        datadetail.FieldCategoryName = cn.Name
                    }
                    datadetail.FieldCategoryId = dc.FieldCategoryId
                    datadetails = append(datadetails, datadetail)
                }
            } else if field.HasFile {
                df := models.DataFile{DataId: data.Id, FieldId: field.Id}
                err := df.Read("data_id", "field_id")
                if err == nil {
                    datadetail.FileName = df.Name
                    datadetail.FileExtension = df.Extension
                    datadetails = append(datadetails, datadetail)
                }
            }
        }
        datalist := DataList{DataDetail: datadetails}
        datalists = append(datalists, datalist)
    }
    return DataAll{DataList: datalists, ResponseCode: 200, ResponseDescription: "List of Data"}
}*/



/*
func Fields(countryId int64, fieldType string) (int64, []Field, error) {
    var table Field
    var fields []Field
    var num int64
    var err error
    if fieldType == "ADDRESS" || fieldType == "FILE" || fieldType == "BANK" {
        num, err = orm.NewOrm().QueryTable(table).Filter("field_type", fieldType).Filter("country_id", countryId).Filter("deleted_at__isnull", true).OrderBy("order").All(&fields)
    } else if fieldType == "ALL" {
        num, err = orm.NewOrm().QueryTable(table).Filter("country_id", countryId).Filter("deleted_at__isnull", true).OrderBy("order").All(&fields)
    }
    return num, fields, err
}*/


func (uc *UserCountryController) Add() {
    var rqd UserCountryIdAdd
    json.Unmarshal(uc.Ctx.Input.RequestBody, &rqd)
    jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
    if isLogin {
        num, _, _ := models.UserCountries(user, "NO")
        if num < configs.MaxUnverifiedCountry {
            usercountry := models.UserCountry{UserId: user.Id, CountryId: rqd.CountryId}
            err := usercountry.Read("user_id", "country_id")
            if err == nil {
                //usercountry.UpdatedAt = time.Now()
                usercountry.DeletedAt = *configs.NullTime
                err = usercountry.Update()
                if err == nil {
                    uc.Data["json"] = uc.prepareCountryList(user, &rqd.RequestCountries, "User Country added successfully")
                } else {
                    uc.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Read and write error 1"}
                }
            } else {
                usercountry := models.UserCountry{UserId: user.Id, CountryId: rqd.CountryId, Eligible: "NO", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: *configs.NullTime}
                err := usercountry.Insert()
                if err == nil {
                    uc.Data["json"] = uc.prepareCountryList(user, &rqd.RequestCountries, "User Country added successfully")
                } else {
                    fmt.Println(err)
                    uc.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Read and write error 2"}
                }
            }
        } else {
            uc.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "You have exceeded the maximum number of unverified countries. Please vefify them first or delete them before adding a new country."}
        }
    } else {
        uc.Data["json"] = jres
    }
    uc.ServeJSON()
}


func (ft *TransferController) updateAdressStatus(transfer *models.Transfer, rqd *TransferDetails) {
    var num int64
    var err error
    if rqd.Primary {
        num, _, err = models.TransferAddressPrimaryFalse(transfer.Id)
    } else {
        num, _, err = models.TransferAddresses(transfer.Id)
    }

    if num == 0 && err == nil {
        if rqd.Nickname == *configs.NullString {
            rqd.Nickname = "Primary"
        }
        rqd.Primary = true
    } else {
        if rqd.Nickname == *configs.NullString {
            rqd.Nickname = "No Name"
        }
    }
    transfer.HasAddress = true
    transfer.Update()
}

func (ft *TransferController) addTransferAddress(transfer *models.Transfer, rqd *TransferDetails) error {
    switch transfer.CurrencyId {
        case models.GetCurrencyId("BTC"):
            return ft.addBTCTransferAddress(transfer, rqd)
        case models.GetCurrencyId("ETH"):
            return ft.addETHTransferAddress(transfer, rqd)
        default:
            return nil
    }
    return nil
}

func (ft *TransferController) addBTCTransferAddress(transfer *models.Transfer, rqd *TransferDetails) error {
    ft.updateAdressStatus(transfer, rqd)
    btcprocess := process.BTCProcess{}
    address, err := btcprocess.GetNewAddress()
    if err == nil {
        transferaddress := models.TransferAddress{TransferId: transfer.Id, Address: address, Nickname: rqd.Nickname, Primary: rqd.Primary}
        err = transferaddress.Insert()
        fmt.Println(err)
        return err
    } else {
        return err
    }
    return nil
}

func (ft *TransferController) addETHTransferAddress(transfer *models.Transfer, rqd *TransferDetails) error {
    ft.updateAdressStatus(transfer, rqd)
    expirationTime := time.Now().Add(time.Hour * time.Duration(configs.ETHPasswordExpiryHour))
    passphrase, _ := token(expirationTime.Unix())
    //passphrase := configs.RandString(100)
    //fmt.Println(passphrase)
    transferpassword := models.TransferPassword{TransferId: transfer.Id, Password: passphrase}
    err := transferpassword.Insert()
    if err == nil {
        ethprocess := process.ETHProcess{}
        address, err := ethprocess.GetNewAddress(passphrase)
        if err == nil {
            transferaddress := models.TransferAddress{TransferId: transfer.Id, Address: address, Nickname: rqd.Nickname, Primary: rqd.Primary}
            err = transferaddress.Insert()
            fmt.Println(err)
            return err
        } else {
            return err
        }
    }
    return nil
}
/*

// @Title AddCrypto
// @Description Add Transfer Crypto Address
// @Param   body        body    TransferCryptoAdd       true        "Add Transfer Crypto Address"
// @Success 200 {string} Transfer Crypto Address added successfully!
// @Failure 403 Add Transfer Crypto Address failed
// @router /addcrypto [post]
func (tc *TransferController) AddCrypto() {
    var rqd TransferCryptoAdd
    json.Unmarshal(tc.Ctx.Input.RequestBody, &rqd)
    jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
    if isLogin {
        transfercrypto := models.TransferCrypto{UserId: user.Id, Address: rqd.Address, Nickname: rqd.Nickname, CurrencyId: rqd.CurrencyId}
        err := transfercrypto.Insert()
        if err == nil {
            tc.Data["json"] = CommonResponse{ResponseCode: 200, ResponseDescription: "Account added successfully"}
        } else {
            tc.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Read and write error"}
        }
    } else {
        tc.Data["json"] = jres
    }
    tc.ServeJSON()
}*/
