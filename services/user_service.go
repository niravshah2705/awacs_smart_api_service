package services

import (
	"awacs_smart_api_service/dal"
	"awacs_smart_api_service/graph/model"
	"reflect"

	"github.com/brkelkar/common_utils/logger"
)

//UserByID get user details by user id
func UserByID(User *model.User, id string) (err error) {
	clouse := "Id='" + id + "'"
	err = dal.UserByID(User, clouse)
	if err != nil {
		logger.Error("User details by Id: ", err)
	}
	return
}

//UserByMobile get user details by mobile number
func UserByMobile(User *model.User, mobile string) (err error) {
	clouse := "Mobile='" + mobile + "'"
	err = dal.UserByID(User, clouse)
	if err != nil {
		logger.Error("User details by Mobile: ", err)
	}
	return
}

//UserByUserName get user details
func UserByUserName(User *model.User, username string) (err error) {
	clouse := "Username='" + username + "'"
	err = dal.UserByID(User, clouse)
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
	user.Address1 = value.FieldByName(role + "Address1").String()
	return
}
