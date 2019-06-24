package user

import (
	"coinford_api/models"
	"coinford_api/configs"
	//"github.com/astaxie/beego/validation"
	//"time"
	//"encoding/json"
	"github.com/astaxie/beego"
	"fmt"
)

type FileController struct {
	beego.Controller
}

// @Title RegisterUser
// @Description Register New User
// @Param	DataId			query 		int64	true		"The Id for the data group"
// @Param	FieldId			query 		int64	true		"The Id for the Field"
// @Param   File        	formData    file    true        "File"
// @Param	Token			query 		string	true		"The generated token"
// @Success 200 {int} response
// @Failure 403 body is empty
// @router /upload [post]
func (u *FileController) Add() {
	fmt.Println(u)
	jres, _, isLogin, _ := apiAuthenticate(u.GetString("Token"))
	isLogin = true
	if isLogin {
		DataId, err1 := u.GetInt64("DataId")
		FieldId, err2 := u.GetInt64("FieldId")
		fmt.Println(DataId, FieldId)
		if err1 == nil && err2 == nil {
			datafile := models.DataFile{DataId: DataId, FieldId: FieldId, Name: "tmp", Extension: "png"}
			err := datafile.Insert()
			if err == nil {
				filename := fmt.Sprintf("%s/%d.%s", configs.FILE_PATH , datafile.Id, datafile.Extension)
				err = u.SaveToFile("File", filename)
				if err == nil {
					datafile.Name = fmt.Sprintf("%d",datafile.Id)
					err = datafile.Update()
					if err == nil {
						u.Data["json"] = DataFileId{Id: datafile.Id, ResponseCode: 200, ResponseDescription: "Files saved successfully"}
					} else {
						u.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Unable to update name and extension"}
					}
				} else {
					u.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Unable to save file"}
				}
			} else {
				fmt.Println(err)
				u.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Unable to insert new file record"}
			}
		} else {
			u.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Unable to read DataId and FieldId from request data"}
		}
	} else {
		u.Data["json"] = jres
	}
	u.ServeJSON()
}