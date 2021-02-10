package services

import (
	"awacs_smart_api_service/graph/model"
	"reflect"

	"github.com/brkelkar/common_utils/logger"

	db "github.com/brkelkar/common_utils/databases"
)

//OrderByUserID get order details by user id
func OrderByUserID(Order *[]*model.Order, buyerID string) (err error) {
	var OrderDetails []*model.OrderDetails
	err = db.DB["smartdb"].Where("buyerId='" + buyerID + "'").Find(&OrderDetails).Error
	if err != nil {
		logger.Error("Order details by userId: ", err)
		return
	}
	orderMap := make(map[string]model.Order, len(OrderDetails))
	for _, val := range OrderDetails {
		if _, ok := orderMap[val.OrderNumber]; !ok {
			//fill order details
			tempOrder := BindOrderDetails(*val)
			orderMap[val.OrderNumber] = tempOrder
		}
		o, _ := orderMap[val.OrderNumber]
		//fill product details
		product := BindProductDetails(*val)
		o.Products = append(o.Products, &product)
		orderMap[val.OrderNumber] = o
	}

	for _, mapVal := range orderMap {
		temp := mapVal
		*Order = append(*Order, &temp)
	}
	return nil
}

//OrderByWorkspaceID get order details by workspace id
func OrderByWorkspaceID(Order *model.OrderByWorkspaceID, workspaceID string, fromDate string, toDate string) (err error) {
	var OrderDetails []*model.OrderDetails
	err = db.DB["smartdb"].Where("WorkspaceId='" + workspaceID + "' and OrderDate  between '" + fromDate + "' and '" + toDate + "'").Find(&OrderDetails).Error
	if err != nil {
		logger.Error("Order details by workspaceId: ", err)
		return
	}
	orderMap := make(map[string]model.Order, len(OrderDetails)+1)
	for _, val := range OrderDetails {
		//fill buyer details
		buyer := BindUserDetails("Buyer", *val)
		Order.Buyer = &buyer

		//fill user details
		supplier := BindUserDetails("Supplier", *val)
		Order.Supplier = &supplier

		if _, ok := orderMap[val.OrderNumber]; !ok {
			//fill order details
			tempOrder := BindOrderDetails(*val)
			orderMap[val.OrderNumber] = tempOrder
		}
		o, _ := orderMap[val.OrderNumber]
		//fill product details
		product := BindProductDetails(*val)
		o.Products = append(o.Products, &product)
		orderMap[val.OrderNumber] = o
	}

	for _, mapVal := range orderMap {
		temp := mapVal
		Order.Order = append(Order.Order, &temp)
	}
	return nil
}

//OrderByOrderNumber get order details by order number
func OrderByOrderNumber(Order *model.Order, orderNumber string) (err error) {
	var OrderDetails []*model.OrderDetails
	err = db.DB["smartdb"].Where("OrderNumber='" + orderNumber + "'").Find(&OrderDetails).Error
	if err != nil {
		logger.Error("Order details by Ordernumber: ", err)
		return
	}

	//fill order details
	*Order = BindOrderDetails(*OrderDetails[0])

	//fill buyer details
	buyer := BindUserDetails("Buyer", *OrderDetails[0])
	Order.Buyer = &buyer

	//fill supplier details
	supplier := BindUserDetails("Supplier", *OrderDetails[0])
	Order.Supplier = &supplier

	for _, val := range OrderDetails {
		//fill product details
		product := BindProductDetails(*val)
		Order.Products = append(Order.Products, &product)
	}

	return nil
}

//OrderstatusByBuyerID get status count by buyer
func OrderstatusByBuyerID(Orderstatus *model.OrderBuyerStatus, BuyerID string, fromDate string, toDate string) (err error) {
	var Orderdetails []*model.OrderBuyerStatusDetails
	err = db.DB["smartdb"].Where("BuyerId='" + BuyerID + "'and Orderdate between '" + fromDate + "' and '" + toDate + "' order by 2 desc").Find(&Orderdetails).Error
	if err != nil {
		logger.Error("Order status by buyerId: ", err)
		return
	}

	flag := 1
	for _, val := range Orderdetails {
		if flag == 1 {
			//fill buyer details
			user := BindUserDetails("Buyer", *val)
			Orderstatus.Buyer = &user
			flag = 0
		}

		var status model.OrderStatus
		status.OrderDate = val.OrderDate
		status.Pending = val.Pending
		status.Billed = val.Billed
		status.Bounced = val.Bounced
		Orderstatus.OrderStatus = append(Orderstatus.OrderStatus, &status)
	}

	return nil
}

//OrderstatusBySupplierID get status and count by supplierId
func OrderstatusBySupplierID(Orderstatus *model.OrderSupplierStatus, SupplierID string, fromDate string, toDate string) (err error) {
	var Orderdetails []*model.OrderSupplierStatusDetails
	err = db.DB["smartdb"].Where("SupplierId='" + SupplierID + "'and Orderdate between '" + fromDate + "' and '" + toDate + "' order by 2 desc").Find(&Orderdetails).Error
	if err != nil {
		logger.Error("Order status by supplierId: ", err)
		return
	}

	flag := 1
	for _, val := range Orderdetails {
		if flag == 1 {
			//fill supplier details
			user := BindUserDetails("Supplier", *val)
			Orderstatus.Supplier = &user
			flag = 0
		}

		var status model.OrderStatus
		status.OrderDate = val.OrderDate
		status.Pending = val.Pending
		status.Billed = val.Billed
		status.Bounced = val.Bounced
		Orderstatus.OrderStatus = append(Orderstatus.OrderStatus, &status)
	}

	return nil
}

//MyOrders get aggregator my orders
func MyOrders(Myorder *model.SmartOrders, retailerName string, fromDate string, toDate string) (err error) {
	var SmartOrderDetails []*model.SmartOrdersDetails
	err = db.DB["awacs"].Where("UserName = '" + retailerName + "' AND OrderDate between '" + fromDate + "' and '" + toDate + "'").Find(&SmartOrderDetails).Error
	if err != nil {
		logger.Error("Order details by userId: ", err)
		return
	}
	aggregatorMap := make(map[string]model.Order, len(SmartOrderDetails))
	retailerMap := make(map[string]model.Order, len(SmartOrderDetails))

	retailer := BindUserDetails("Retailer", *SmartOrderDetails[0])
	Myorder.Retailer = &retailer

	for _, val := range SmartOrderDetails {
		if val.OrderType == "Aggregator" {
			if _, ok := aggregatorMap[val.OrderNumber]; !ok {
				//fill order details
				tempOrder := BindOrderDetails(*val)
				aggregatorMap[val.OrderNumber] = tempOrder
			}
			o, _ := aggregatorMap[val.OrderNumber]
			//fill product details
			product := BindProductDetails(*val)
			o.Products = append(o.Products, &product)
			aggregatorMap[val.OrderNumber] = o
		}
		if val.OrderType == "Aggregator" {
			if _, ok := retailerMap[val.OrderNumber]; !ok {
				//fill order details
				tempOrder := BindOrderDetails(*val)
				retailerMap[val.OrderNumber] = tempOrder
			}
			// o, _ := retailerMap[val.OrderNumber]
			// //fill product details
			// product := BindProductDetails(*val)
			// o.Products = append(o.Products, &product)
			// retailerMap[val.OrderNumber] = o
		}
	}

	for _, mapVal := range aggregatorMap {
		temp := mapVal
		Myorder.AggregatorOrders = append(Myorder.AggregatorOrders, &temp)
	}

	for _, mapVal := range retailerMap {
		temp := mapVal
		Myorder.RetailerOrders = append(Myorder.AggregatorOrders, &temp)
	}

	return
}

//BindOrderDetails bind order details
func BindOrderDetails(val interface{}) (order model.Order) {
	value := reflect.ValueOf(val)
	order.ID = value.FieldByName("ID").String()
	order.Source = value.FieldByName("Source").String()
	order.BuyerID = value.FieldByName("BuyerID").String()
	order.OrderNumber = value.FieldByName("OrderNumber").String()
	order.OrderDate = value.FieldByName("OrderDate").String()
	order.SupplierID = value.FieldByName("SupplierID").String()
	order.TakenOn = value.FieldByName("TakenOn").String()
	order.DeliveryOption = value.FieldByName("DeliveryOption").String()
	order.DeliveryMode = value.FieldByName("DeliveryMode").String()
	order.Remarks = value.FieldByName("Remarks").String()
	//order.WorkspaceID = int(value.FieldByName("WorkspaceID").Int())
	order.Status = value.FieldByName("Status").String()
	order.InvoiceNumber = value.FieldByName("InvoiceNumber").String()
	//order.IsBounced = value.FieldByName("IsBounced").Bool()

	return
}
