package user

import (
	"coinford_api/models"
	"github.com/astaxie/beego"
	"encoding/json"
	//"fmt"
)

type FieldController struct {
	beego.Controller
}

// @Title GetAll
// @Description get the list of fiels for specific type
// @Param	body		body 	FieldType		true		"Token for Authentication"
// @Success 200 {int} response
// @Failure 403 Authentication Failed
// @router /list [post]
func (a *FieldController) GetAll() {
	var rqd FieldType
	json.Unmarshal(a.Ctx.Input.RequestBody, &rqd)
	jres, _, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
		if rqd.RequestField.FieldType == "ALL" {
			a.Data["json"] = a.allFieldList(&rqd.RequestField)
		} else {
			a.Data["json"] = a.prepareFieldList(&rqd.RequestField)
		}	
	} else {
		a.Data["json"] = jres
	}
	a.ServeJSON()
}

func (a *FieldController) allFieldList(fieldType *RequestField) FieldListAll {
	_, kyc, _ := models.Fields(fieldType.CountryId, "KYC")
	_, address, _ := models.Fields(fieldType.CountryId, "ADDRESS")
	_, bank, _ := models.Fields(fieldType.CountryId, "BANK")

	var kycf []FieldDetails
	for _, field := range kyc {
		var kycdetails FieldDetails
		if field.HasCategory {
			_, kyc_cat, _ := models.FieldCategories(field.Id)
			kycdetails = FieldDetails{Field: field, Category: kyc_cat}
		} else {
			kycdetails = FieldDetails{Field: field, Category: []models.FieldCategory{}}
		}
		kycf = append(kycf, kycdetails)
	}

	var addressf []FieldDetails
	for _, field := range address {
		var addressdetails FieldDetails
		if field.HasCategory {
			_, address_cat, _ := models.FieldCategories(field.Id)
			addressdetails = FieldDetails{Field: field, Category: address_cat}
 		} else {
			addressdetails = FieldDetails{Field: field, Category: []models.FieldCategory{}}
		}
		addressf = append(addressf, addressdetails)
	}

	var bankf []FieldDetails
	for _, field := range bank {
		var bankdetails FieldDetails
		if field.HasCategory {
			_, bank_cat, _ := models.FieldCategories(field.Id)
			bankdetails = FieldDetails{Field: field, Category: bank_cat}
 		} else {
			bankdetails = FieldDetails{Field: field, Category: []models.FieldCategory{}}
		}
		bankf = append(bankf, bankdetails)
	}

	return FieldListAll{KYC: kycf, ADDRESS: addressf, BANK: bankf, ResponseCode: 200, ResponseDescription: "List of Fields"}
}

func (a *FieldController) prepareFieldList(fieldType *RequestField) FieldList {
	_, fields, _ := models.Fields(fieldType.CountryId, fieldType.FieldType)
	var fieldf []FieldDetails
	for _, field := range fields {
		var fielddetails FieldDetails
		if field.HasCategory {
			_, field_cat, _ := models.FieldCategories(field.Id)
			fielddetails = FieldDetails{Field: field, Category: field_cat}
 		} else {
			fielddetails = FieldDetails{Field: field, Category: []models.FieldCategory{}}
		}
		fieldf = append(fieldf, fielddetails)
	}
	return FieldList{Fields: fieldf, ResponseCode: 200, ResponseDescription: "List of Fields"}
}