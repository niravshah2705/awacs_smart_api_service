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
	var dummySupper model.BuyerDashboard
	dummySupper.SupplierID = "Total"
	dummySupper.SupplierName = "Total"

	for _, val := range *Buyerdashboard {
		dummySupper.Billed = dummySupper.Billed + val.Billed
		dummySupper.Pending = dummySupper.Pending + val.Pending
		dummySupper.Bounced = dummySupper.Bounced + val.Bounced
		dummySupper.Short = dummySupper.Short + val.Short
		dummySupper.CurrentOutstanding = dummySupper.CurrentOutstanding + val.CurrentOutstanding
		dummySupper.ProductCount = dummySupper.ProductCount + val.ProductCount

	}
	*Buyerdashboard = append(*Buyerdashboard, &dummySupper)
	return
}
