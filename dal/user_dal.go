package dal

import (
	"awacs_smart_api_service/graph/model"

	db "github.com/brkelkar/common_utils/databases"
)

//UserByID get user data
func UserByID(User *model.User, clouse string) (err error) {
	err = db.DB["smartdb"].Where(clouse).First(User).Error
	return
}