package services

import (
	"awacs_smart_api_service/graph/model"

	"github.com/brkelkar/common_utils/logger"

	db "github.com/brkelkar/common_utils/databases"
)

//Getoutstandingbyuserid get outstanding data
func Getoutstandingbyuserid(Outstanding *[]*model.Outstanding, userid string) (err error) {
	err = db.DB["smartdb"].Where("UserId ='" + userid + "'").Find(Outstanding).Error
	if err != nil {
		logger.Error("Invoice details: ", err)
		return
	}
	return
}
