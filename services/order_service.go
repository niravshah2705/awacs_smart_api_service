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
