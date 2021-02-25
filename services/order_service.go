package services

import (
	"awacs_smart_api_service/graph/model"
	"reflect"
	"strings"

	"awacs_smart_api_service/dal"

	"github.com/brkelkar/common_utils/logger"
)

//OrderByUserID get order details by user id
func OrderByUserID(Order *[]*model.Order, buyerID string) (err error) {
	var OrderDetails []*model.OrderDetails
	clouse := "buyerId='" + buyerID + "'"
	err = dal.OrderDetails(&OrderDetails, clouse)
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
	clouse := "WorkspaceId='" + workspaceID + "' and OrderDate  between '" + fromDate + "' and '" + toDate + "'"
	err = dal.OrderDetails(&OrderDetails, clouse)
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
	clouse := "OrderNumber='" + orderNumber + "'"
	err = dal.OrderDetails(&OrderDetails, clouse)
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
	clouse := "BuyerId='" + BuyerID + "'and Orderdate between '" + fromDate + "' and '" + toDate + "' order by 2 desc"
	err = dal.OrderstatusByBuyerID(&Orderdetails, clouse)
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
	clouse := "SupplierId='" + SupplierID + "'and Orderdate between '" + fromDate + "' and '" + toDate + "' order by 2 desc"
	err = dal.OrderstatusBySupplierID(&Orderdetails, clouse)
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
	clouse := "UserName = '" + retailerName + "' AND OrderDate between '" + fromDate + "' and '" + toDate + "'"
	err = dal.MyOrders(&SmartOrderDetails, clouse)
	if err != nil {
		logger.Error("MyOrders for Affregator and Retailer: ", err)
		return
	}
	if len(SmartOrderDetails) > 0 {
		aggregatororderMap := make(map[string]model.Order, len(SmartOrderDetails))
		retailerorderMap := make(map[string]model.Order, len(SmartOrderDetails))
		retailerMap := make(map[string]model.User)
		aggregatorMap := make(map[string]model.User)

		for _, val := range SmartOrderDetails {
			orderkey := val.RetailerCode + "_" + val.OrderNumber

			if val.OrderType == "Aggregator" {
				if _, isretailer := aggregatorMap[val.RetailerCode]; !isretailer {
					//fill user details
					retailer := BindUserDetails("Retailer", *val)
					aggregatorMap[val.RetailerCode] = retailer
				}

				if _, ok := aggregatororderMap[orderkey]; !ok {
					//fill order details
					tempOrder := BindOrderDetails(*val)
					//fill user details
					supplier := BindUserDetails("Supplier", *val)
					tempOrder.Supplier = &supplier

					aggregatororderMap[orderkey] = tempOrder
				}
				o, _ := aggregatororderMap[orderkey]

				//fill product details
				product := BindProductDetails(*val)
				o.Products = append(o.Products, &product)

				aggregatororderMap[orderkey] = o

			} else if val.OrderType == "Retailer" {
				if _, isretailer := retailerMap[val.RetailerCode]; !isretailer {
					//fill user details
					retailer := BindUserDetails("Retailer", *val)
					retailerMap[val.RetailerCode] = retailer
				}

				if _, ok := retailerorderMap[orderkey]; !ok {
					//fill order details
					tempOrder := BindOrderDetails(*val)
					//fill user details
					supplier := BindUserDetails("Supplier", *val)
					tempOrder.Supplier = &supplier

					retailerorderMap[orderkey] = tempOrder
				}
				o, _ := retailerorderMap[orderkey]
				//fill product details
				product := BindProductDetails(*val)
				o.Products = append(o.Products, &product)

				retailerorderMap[orderkey] = o
			}
		}

		for key, mapVal := range retailerorderMap {
			temp := mapVal
			rr, _ := retailerMap[strings.Split(key, "_")[0]]
			rr.Orders = append(rr.Orders, &temp)
			retailerMap[strings.Split(key, "_")[0]] = rr
		}

		for _, mapVal := range retailerMap {
			temp := mapVal
			Myorder.RetailerOrders = append(Myorder.RetailerOrders, &temp)
		}

		for key, mapVal := range aggregatororderMap {
			temp := mapVal
			rr, _ := aggregatorMap[strings.Split(key, "_")[0]]
			rr.Orders = append(rr.Orders, &temp)
			aggregatorMap[strings.Split(key, "_")[0]] = rr
		}

		for _, mapVal := range aggregatorMap {
			temp := mapVal
			Myorder.AggregatorOrders = append(Myorder.AggregatorOrders, &temp)
		}
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
