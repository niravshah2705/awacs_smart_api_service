package dal

import (
	"awacs_smart_api_service/graph/model"

	db "github.com/brkelkar/common_utils/databases"
)

//MyOrdersDal get query data
func MyOrders(SmartOrderDetails *[]*model.SmartOrdersDetails, clouse string) (err error) {
	err = db.DB["awacs"].Where(clouse).Find(&SmartOrderDetails).Error
	return
}

//OrderDetails get order details
func OrderDetails(OrderDetails *[]*model.OrderDetails, clouse string) (err error) {
	err = db.DB["smartdb"].Where(clouse).Find(&OrderDetails).Error
	return
}

//OrderstatusByBuyerIDDal get buyer order status
func OrderstatusByBuyerID(OrderDetails *[]*model.OrderBuyerStatusDetails, clouse string) (err error) {
	err = db.DB["smartdb"].Where(clouse).Find(&OrderDetails).Error
	return
}

//OrderstatusBySupplierIDDal get supplier order status
func OrderstatusBySupplierID(OrderDetails *[]*model.OrderSupplierStatusDetails, clouse string) (err error) {
	err = db.DB["smartdb"].Where(clouse).Find(&OrderDetails).Error
	return
}
