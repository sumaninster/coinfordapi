package user

import (
	"coinford_api/models"
	"coinford_api/configs"
	"time"
	"github.com/astaxie/beego"
	"encoding/json"
	"errors"
	"fmt"
)

type DataController struct {
	beego.Controller
}

// @Title Add Data
// @Description add data
// @Param	body		body 	DataUpload		true		"Data Details"
// @Success 200 {int} response
// @Failure 403 Authentication Failed
// @router /add [post]
func (dc *DataController) Add() {
	var rqd DataUpload
	var errInsert error
	partialSuccess := false
	json.Unmarshal(dc.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		dc.updateStatus(user.Id, &rqd)
		data := models.Data{UserId: user.Id, CountryId: rqd.CountryId, FieldType: rqd.FieldType, Nickname: rqd.Nickname, Primary: rqd.Primary, Active: rqd.Active}
		err := data.Insert()
		//err := data.ReadOrCreate("user_id", "country_id", "field_type")
		if err == nil {
			if rqd.FieldType == "BANK" {
				dc.createFiatWallet(user.Id, data.Id, &rqd)
			}
			for _, datarq := range rqd.DataText {
				datatext := models.DataText{DataId: data.Id, FieldId: datarq.FieldId, Text: datarq.InputText}
				err := datatext.Insert()
				if err == nil {
					errInsert = nil
					partialSuccess = true
				} else {
					errInsert = errors.New("Unable to add data text")
				}
			}//End of for loop
			for _, datarq := range rqd.DataCategory {
				datacategory := models.DataCategory{DataId: data.Id, FieldId: datarq.FieldId, FieldCategoryId: datarq.FieldCategoryId}
				err := datacategory.Insert()
				if err == nil {
					errInsert = nil
					partialSuccess = true
				} else {
					errInsert = errors.New("Unable to add data category")
				}
			}//End of for loop
			for _, datarq := range rqd.DataFile {
				datafile := models.DataFile{DataId: data.Id, FieldId: datarq.FieldId, Name: datarq.Name, Extension: datarq.Extension}
				err := datafile.Insert()
				if err == nil {
					errInsert = nil
					partialSuccess = true
				} else {
					errInsert = errors.New("Unable to add data file")
				}
			}//End of for loop
			if errInsert == nil {
				dc.Data["json"] = dc.prepareDataList(user.Id, data.CountryId, data.FieldType, "Data added successfully")
			} else if partialSuccess {
				dc.Data["json"] = dc.prepareDataList(user.Id, data.CountryId, data.FieldType, "Some data could not be added because of internal error")
			} else {
				dc.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Data details insert failed"}
			}
		} else {
			dc.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Data insert failed"}
		}
	} else {
		dc.Data["json"] = jres
	}
	dc.ServeJSON()
}

// @Title GetAll
// @Description get all the data for the user and the type
// @Param	body		body 	FieldType		true		"Token for Authentication"
// @Success 200 {int} response
// @Failure 403 Authentication Failed
// @router /list [post]
func (dc *DataController) GetAll() {
	var rqd FieldType
	json.Unmarshal(dc.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		dc.Data["json"] = dc.prepareDataList(user.Id, rqd.RequestField.CountryId, rqd.RequestField.FieldType, "List Data")
	} else {
		dc.Data["json"] = jres
	}
	dc.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	body		body 	DataId		true		"Token for Authentication"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /delete [delete]
func (dc *DataController) Delete() {
	var rqd DataId
	json.Unmarshal(dc.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		data := models.Data{Id: rqd.Id, UserId: user.Id}
		err := data.Read("id", "user_id")
		if err == nil {
			data.DeletedAt = time.Now()
			err = data.Update()
			if err == nil {
				dc.Data["json"] = dc.prepareDataList(user.Id, data.CountryId, data.FieldType, "Delete success")
			} else {
				dc.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Delete Failed"}
			}
		} else {
			dc.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Delete Failed"}
		}
	} else {
		dc.Data["json"] = jres
	}
	dc.ServeJSON()
}

func (dc *DataController) createFiatWallet(userId int64, dataId int64, rqd *DataUpload) {
	currency := models.Currency{CountryId: rqd.CountryId}
	walletMaster := models.WalletMaster{UserId: userId, CurrencyId: currency.Id, CurrencyType: currency.Type}
	_, _, err := walletMaster.ReadOrCreate("user_id", "currency_id")
	if err == nil {
		eligible := models.IsEligible(userId, rqd.CountryId)
		notDeleted := true
		if eligible {
			if walletMaster.DeletedAt != *configs.NullTime {
				walletMaster.DeletedAt = *configs.NullTime
				err = walletMaster.Update()
				notDeleted = false
			}
		} else {
			walletMaster.DeletedAt = time.Now()
			err = walletMaster.Update()
		}

		if err == nil {
			if notDeleted {
				wallet := models.Wallet{WalletMasterId: walletMaster.Id, Amount: 0, AmountLocked: 0, Nickname: rqd.Nickname, Primary: rqd.Primary}
				err := wallet.Insert()
				if err == nil {
					walletFiat := models.WalletFiat{WalletId: wallet.Id, DataId: dataId}
					walletFiat.Insert()
				} else {
					fmt.Println(err)
				}
			}
		} else {
			fmt.Println(err)
		}
	}
}

func (dc *DataController) updateStatus(userId int64, rqd *DataUpload) {
	var num int64
	var err error
	if rqd.Primary {
		num, _, err = models.DataPrimaryFalse(userId, rqd.CountryId, rqd.FieldType)
	} else {
		num, _, err = models.DataList(userId, rqd.CountryId, rqd.FieldType)
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
}

func (dc *DataController) prepareDataList(userId int64, countryId int64, fieldType string, message string) DataList {
	_, fields, _ := models.Fields(countryId, fieldType)
	_, datas, _ := models.DataList(userId, countryId, fieldType)
	var datagroups []DataGroup
	var err error
	for _, data := range datas {
		var datadetails []DataDetail
		for _, field := range fields {
			datadetail := DataDetail{Field: field}
			if field.HasInputText {
				dt := models.DataText{DataId: data.Id, FieldId: field.Id}
				err = dt.Read("data_id", "field_id")
				if err == nil {
					datadetail.DataText = &dt
				}
			}
			if field.HasCategory {
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
			}
			if field.HasFile {
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