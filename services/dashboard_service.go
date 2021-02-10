package services

import (
	"awacs_smart_api_service/graph/model"

	db "github.com/brkelkar/common_utils/databases"
	"github.com/brkelkar/common_utils/logger"
)

//BuyerDashboard get order details for dashboard
func BuyerDashboard(Buyerdashboard *[]*model.BuyerDashboard, buyerID string, fromDate string, toDate string) (err error) {
	err = db.DB["smartdb"].Raw("exec SP_DASHBOARD_ORDERSTATUS  ?, ?, ?", buyerID, fromDate, toDate).Scan(Buyerdashboard).Error
	if err != nil {
		logger.Error("Buyer Dashboard: ", err)
	}
	return
}
