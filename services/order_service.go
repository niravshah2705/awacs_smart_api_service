package services

import (
	"awacs_smart_api_service/graph/model"
	"log"

	db "github.com/brkelkar/common_utils/databases"
)

func OrderByUserId(Order *[]*model.Order, buyerID string) (err error) {

	var OrderDetails []*model.OrderDetails
	err = db.DB["smartdb"].Where("buyerId='" + buyerID + "'").Find(&OrderDetails).Error
	if err != nil {
		log.Print(err)
		return

	}
	orderMap := make(map[string]model.Order, len(OrderDetails)+1)
	for _, val := range OrderDetails {

		if _, ok := orderMap[val.OrderNumber]; !ok {
			var tempOrder model.Order
			tempOrder.ID = val.ID
			tempOrder.Source = val.Source
			tempOrder.BuyerID = val.BuyerID
			tempOrder.OrderNumber = val.OrderNumber
			tempOrder.OrderDate = val.OrderDate
			tempOrder.SupplierID = val.SupplierID
			tempOrder.TakenOn = val.TakenOn
			tempOrder.DeliveryOption = val.DeliveryOption
			tempOrder.DeliveryMode = val.DeliveryMode
			tempOrder.Remarks = val.Remarks
			tempOrder.WorkspaceID = val.WorkspaceID
			tempOrder.Status = val.Status
			tempOrder.InvoiceNumber = val.InvoiceNumber
			tempOrder.IsBounced = val.IsBounced
			orderMap[val.OrderNumber] = tempOrder
		}
		o, _ := orderMap[val.OrderNumber]
		var product model.Product
		product.ProductCode = val.ProductCode
		product.ProductName = val.ProductName
		product.Company = val.Company
		product.Pack = val.Pack
		product.Quantity = val.Quantity
		product.FreeQuantity = val.FreeQuantity
		product.SchemePercentage = val.SchemePercentage
		product.Closing = val.Closing
		product.Mrp = val.Mrp
		product.Ptr = val.Ptr
		product.Pts = val.Pts
		o.Products = append(o.Products, &product)
		orderMap[val.OrderNumber] = o

	}

	for _, mapVal := range orderMap {
		temp := mapVal
		*Order = append(*Order, &temp)

	}
	return nil
}

//OrderstatusByBuyerID get status count by buyer
func OrderstatusByBuyerID(Orderstatus *model.OrderBuyerStatus, BuyerID string, fromDate string, toDate string) (err error) {
	var Orderdetails []*model.OrderBuyerStatusDetails
	err = db.DB["smartdb"].Where("BuyerId='" + BuyerID + "'and Orderdate between '" + fromDate + "' and '" + toDate + "' order by 2 desc").Find(&Orderdetails).Error
	if err != nil {
		log.Print(err)
		return
	}

	flag:=1
	for _, val := range Orderdetails {	
		if flag==1{
			var user model.User
			user.Name=val.BuyerName
			user.City=val.City
			user.State=val.State
			user.Email=val.Email
			user.Mobile=val.Mobile
			user.PhoneNo=val.PhoneNo
			user.Pincode=&val.Pincode
			user.Username=val.BuyerId

			Orderstatus.Buyer=&user
			flag=0
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
	var Orderdetails []*model.OrderStatusSupplierDetails
	err = db.DB["smartdb"].Where("SupplierId='" + SupplierID + "'and Orderdate between '" + fromDate + "' and '" + toDate + "' order by 2 desc").Find(&Orderdetails).Error
	if err != nil {
		log.Print(err)
		return
	}

	flag:=1
	for _, val := range Orderdetails {	
		if flag==1{
			var user model.User
			user.Name=val.SupplierName
			user.City=val.City
			user.State=val.State
			user.Email=val.Email
			user.Mobile=val.Mobile
			user.PhoneNo=val.PhoneNo
			user.Pincode=&val.Pincode
			user.Username=val.SupplierId

			Orderstatus.Supplier=&user
			flag=0
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