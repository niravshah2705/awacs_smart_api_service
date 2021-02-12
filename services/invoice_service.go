package services

import (
	"awacs_smart_api_service/dal"
	"awacs_smart_api_service/graph/model"
	"reflect"

	"github.com/brkelkar/common_utils/logger"
)

//InvoiceDetails get invoice details
func InvoiceDetails(Invoice *[]*model.Invoice, buyerID string, supplierID string, fromDate string, toDate string) (err error) {
	var InvoiceDetails []*model.InvoiceDetails
	clouse := " SupplierID='" + supplierID + "' and BuyerID='" + buyerID + "' and BillDate between '" + fromDate + "' and '" + toDate + "'"
	err = dal.InvoiceDetails(&InvoiceDetails, clouse)
	if err != nil {
		logger.Error("Invoice details: ", err)
		return
	}
	InvoiceMap := make(map[string]model.Invoice, len(InvoiceDetails)+1)
	for _, val := range InvoiceDetails {
		key := val.BillNumber + val.SupplierID + val.BuyerID
		if _, ok := InvoiceMap[key]; !ok {

			var tempInvoice model.Invoice
			//fill invoice details
			tempInvoice = BindInvoiceDetails(*val)

			//fill supplier details
			tempSupplier := BindUserDetails("Supplier", *val)
			tempInvoice.Suppliers = &tempSupplier

			//fill user details
			tempBuyer := BindUserDetails("Buyer", *val)
			tempInvoice.Buyers = &tempBuyer
			InvoiceMap[key] = tempInvoice
		}
		ti := InvoiceMap[key]

		//fill product details
		product := BindProductDetails(*val)
		//fill additional product details
		BindAdditionalProductDetails(*val, &product)
		ti.Products = append(ti.Products, &product)
		InvoiceMap[key] = ti
	}

	for _, mapVal := range InvoiceMap {
		temp := mapVal
		*Invoice = append(*Invoice, &temp)
	}
	return nil

}

//InvoiceByBuyer get invoice details
func InvoiceByBuyer(InvoiceBuyer *model.InvoiceBuyer, buyerID string, fromDate string, toDate string) (err error) {
	var InvoiceDetails []*model.InvoiceDetails
	clouse := "  BuyerID='" + buyerID + "' and BillDate between '" + fromDate + "' and '" + toDate + "'"
	err = dal.InvoiceDetails(&InvoiceDetails, clouse)
	if err != nil {
		logger.Error("Invoice details by buyer: ", err)
		return
	}
	iMap := make(map[string]model.Invoice, len(InvoiceDetails))
	for _, val := range InvoiceDetails {
		flag := 1
		key := val.BillNumber + "_" + val.SupplierID + "_" + val.BuyerID
		if flag == 1 {
			//fill buyer details
			tempBuyer := BindUserDetails("Buyer", *val)
			InvoiceBuyer.Buyers = &tempBuyer
			flag = 0
		}
		if _, ok := iMap[key]; !ok {
			var tempInvoice model.Invoice
			//fill invoice details
			tempInvoice = BindInvoiceDetails(*val)

			//fill supplier details
			tempSupplier := BindUserDetails("Supplier", *val)
			tempInvoice.Suppliers = &tempSupplier
			iMap[key] = tempInvoice

		}

		tP := iMap[key]
		//fill product details
		product := BindProductDetails(*val)
		//fill additional product details
		BindAdditionalProductDetails(*val, &product)
		tP.Products = append(tP.Products, &product)
		iMap[key] = tP
	}

	for _, val := range iMap {
		temp := val
		InvoiceBuyer.Invoices = append(InvoiceBuyer.Invoices, &temp)
	}

	return nil
}

//BindInvoiceDetails get invoice data
func BindInvoiceDetails(val interface{}) (invoice model.Invoice) {
	value := reflect.ValueOf(val)
	invoice.ID = value.FieldByName("ID").String()
	invoice.DeveloperID = value.FieldByName("DeveloperID").String()
	invoice.SupplierID = value.FieldByName("SupplierID").String()
	invoice.BillNumber = value.FieldByName("BillNumber").String()
	invoice.BillDate = value.FieldByName("BillDate").String()
	invoice.ChallanNumber = value.FieldByName("ChallanNumber").String()
	invoice.ChallanDate = value.FieldByName("ChallanDate").String()
	invoice.BuyerID = value.FieldByName("BuyerID").String()
	invoice.Reason = value.FieldByName("Reason").String()
	invoice.Upc = value.FieldByName("Upc").String()
	invoice.PaymentDueDate = value.FieldByName("PaymentDueDate").String()
	invoice.NetInvoiceAmount = value.FieldByName("NetInvoiceAmount").Float()
	invoice.NetPurchaseAmount = value.FieldByName("NetPurchaseAmount").Float()
	invoice.LastTransactionDate = value.FieldByName("LastTransactionDate").String()
	invoice.ReceivedOn = value.FieldByName("ReceivedOn").String()
	return
}
