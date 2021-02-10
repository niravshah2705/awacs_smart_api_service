package services

import (
	"awacs_smart_api_service/graph/model"
	"reflect"

	db "github.com/brkelkar/common_utils/databases"
	"github.com/brkelkar/common_utils/logger"
)

//UserByID get user details by user id
func UserByID(User *model.User, id string) (err error) {
	err = db.DB["smartdb"].First(User, id).Error
	if err != nil {
		logger.Error("User details by Id: ", err)
	}
	return
}

//UserByMobile get user details by mobile number
func UserByMobile(User *model.User, mobile string) (err error) {
	err = db.DB["smartdb"].Where("Mobile='" + mobile + "'").First(User).Error
	if err != nil {
		logger.Error("User details by Mobile: ", err)
	}
	return
}

//UserByUserName get user details
func UserByUserName(User *model.User, username string) (err error) {
	err = db.DB["smartdb"].Where("Username='" + username + "'").First(User).Error
	if err != nil {
		logger.Error("User details by Uesrname: ", err)
	}
	return
}

//BindUserDetails assign to user details
func BindUserDetails(role string, val interface{}) (user model.User) {
	value := reflect.ValueOf(val)
	user.Name = value.FieldByName(role + "Name").String()
	user.Email = value.FieldByName(role + "Email").String()
	user.City = value.FieldByName(role + "City").String()
	user.State = value.FieldByName(role + "State").String()
	user.Mobile = value.FieldByName(role + "Mobile").String()
	user.Pincode = value.FieldByName(role + "Pincode").String()
	return
}
