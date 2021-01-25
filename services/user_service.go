package services

import (
	"awacs_smart_api_service/graph/model"

	db "github.com/brkelkar/common_utils/databases"
)

func UserByID(User *model.User, id string) (err error) {
	err = db.DB["smartdb"].First(User, id).Error
	return
}

func UserByMobile(User *model.User, mobile string) (err error) {
	err = db.DB["smartdb"].Where("Mobile='" + mobile + "'").First(User).Error
	return
}

func UserByUserName(User *model.User, username string) (err error) {
	err = db.DB["smartdb"].Where("Username='" + username + "'").First(User).Error
	return
}
