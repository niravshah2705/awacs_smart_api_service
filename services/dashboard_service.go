package services

import (
	"awacs_smart_api_service/dal"
	"awacs_smart_api_service/graph/model"

	"github.com/brkelkar/common_utils/logger"
)

//BuyerDashboard get order details for dashboard
func BuyerDashboard(Buyerdashboard *model.BuyerDashboard, buyerID string, fromDate string, toDate string) (err error) {
	var supplier []*model.SupperOrderDeatils
	err = dal.BuyerDashboard(&supplier, buyerID, fromDate, toDate)
	if err != nil {
		logger.Error("Buyer Dashboard: ", err)
	}

	for _, val := range supplier {
		Buyerdashboard.TotalBilledOrders = Buyerdashboard.TotalBilledOrders + val.Billed
		Buyerdashboard.TotalPendingOrders = Buyerdashboard.TotalPendingOrders + val.Pending
		Buyerdashboard.TotalBounceOrders = Buyerdashboard.TotalBounceOrders + val.Bounced
		Buyerdashboard.TotalShortOrders = Buyerdashboard.TotalShortOrders + val.Short
		Buyerdashboard.TotalOutstanding = Buyerdashboard.TotalOutstanding + val.CurrentOutstanding
		Buyerdashboard.TotalProductOrdered = Buyerdashboard.TotalProductOrdered + val.ProductCount
		Buyerdashboard.TotalOrders = Buyerdashboard.TotalOrders + val.TotalOrder
	}
	Buyerdashboard.Supper = supplier
	return
}
