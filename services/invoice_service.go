package services

import (
	"awacs_smart_api_service/graph/model"
	"log"

	db "github.com/brkelkar/common_utils/databases"
)

func InvoiceByBillNumber(Invoice *[]*model.Invoice, billNumber string) (err error) {
	var InvoiceDetails []*model.InvoiceDetails
	err = db.DB["smartdb"].Where("TRIM(billNumber)='" + billNumber + "'").Find(&InvoiceDetails).Error
	if err != nil {
		log.Print(err)
		return

	}
	InvoiceMap := make(map[string]model.Invoice, len(InvoiceDetails)+1)
	for _, val := range InvoiceDetails {
		if _, ok := InvoiceMap[val.BillNumber]; !ok {
			var tempInvoice model.Invoice
			tempInvoice.ID = val.ID
			tempInvoice.DeveloperID = val.DeveloperID
			tempInvoice.SupplierID = val.SupplierID
			tempInvoice.BillNumber = val.BillNumber
			tempInvoice.BillDate = val.BillDate
			tempInvoice.ChallanNumber = val.ChallanNumber
			tempInvoice.ChallanDate = val.ChallanDate
			tempInvoice.BuyerID = val.BuyerID
			tempInvoice.Reason = val.Reason
			tempInvoice.Upc = val.Upc
			tempInvoice.PaymentDueDate = val.PaymentDueDate
			tempInvoice.NetInvoiceAmount = val.NetInvoiceAmount
			tempInvoice.NetPurchaseAmount = val.NetPurchaseAmount
			tempInvoice.LastTransactionDate = val.LastTransactionDate
			tempInvoice.ReceivedOn = val.ReceivedOn

			tempInvoice.EmailPrepared = val.EmailPrepared
			tempInvoice.EmailPreparedOn = val.EmailPreparedOn
			tempInvoice.Taken = val.Taken
			tempInvoice.TakenOn = val.TakenOn
			tempInvoice.LastHostAddress = val.LastHostAddress
			tempInvoice.WorkspaceID = val.WorkspaceID
			tempInvoice.Ede2Taken = val.Ede2Taken
			tempInvoice.Ede2TakenDate = val.Ede2TakenDate
			tempInvoice.Ede2TakenBatchID = val.Ede2TakenBatchID

			var tempSupplier model.User
			tempSupplier.Name = val.SupplierName
			tempSupplier.City = val.SupplierCity
			tempSupplier.State = val.SupplierState
			tempSupplier.AwacsID = &val.SupplierAwacsId
			tempSupplier.Email = val.SupplierEmail
			tempSupplier.Pincode = &val.SupplierPincode
			tempSupplier.Mobile = val.SupplierMobile
			tempSupplier.PhoneNo = val.SupplierPhoneNo
			tempInvoice.Suppliers = append(tempInvoice.Suppliers, &tempSupplier)

			var tempBuyer model.User
			tempBuyer.Name = val.BuyerName
			tempBuyer.City = val.BuyerCity
			tempBuyer.State = val.BuyerState
			tempBuyer.AwacsID = &val.BuyerAwacsId
			tempBuyer.Email = val.BuyerEmail
			tempBuyer.Pincode = &val.BuyerPincode
			tempBuyer.Mobile = val.BuyerMobile
			tempBuyer.PhoneNo = val.BuyerPhoneNo
			tempInvoice.Buyers = append(tempInvoice.Buyers, &tempBuyer)
			InvoiceMap[val.BillNumber] = tempInvoice
		}
		ti := InvoiceMap[val.BillNumber]
		var product model.Product
		product.ProductCode = val.SupplierProductCode
		product.ProductName = val.SupplierProductName
		product.Pack = val.SupplierProductPack
		product.Mrp = val.Mrp
		product.Batch = &val.Batch
		product.Expiry = &val.Expiry
		product.Quantity = val.Quantity
		product.FreeQuantity = &val.FreeQuantity
		product.Rate = val.Rate
		product.Amount = val.Amount
		product.Discount = val.Discount
		product.DiscountAmount = val.DiscountAmount
		product.AddlScheme = val.AddlScheme
		product.AddlSchemeAmount = val.AddlSchemeAmount
		product.AddlDiscount = val.AddlDiscount
		product.AddlDiscountAmount = val.AddlDiscountAmount
		product.DeductableBeforeDiscount = val.DeductableBeforeDiscount
		product.MrpInclusiveTax = val.MrpInclusiveTax
		product.VatApplication = val.VatApplication
		product.Vat = val.Vat
		product.AddlTax = val.AddlTax
		product.Cst = val.Cst
		product.SGst = val.SGst
		product.CGst = val.CGst
		product.IGst = val.IGst
		product.BaseSchemeQuantity = val.BaseSchemeQuantity
		product.BaseSchemeFreeQuantity = val.BaseSchemeFreeQuantity
		product.Remarks = val.Remarks
		product.Company = val.ProductCompany
		product.BatchID = val.BatchID
		product.SgstAmount = val.SgstAmount
		product.CgstAmount = val.CgstAmount
		product.IgstAmount = val.IgstAmount
		product.Cess = val.Cess
		product.CessAmount = val.CessAmount
		product.Hsn = val.Hsn
		product.OrderNumber = val.OrderNumber
		product.OrderDate = val.OrderDate
		product.Barcode = val.Barcode
		product.OrderSource = val.OrderSource
		product.Ptr = val.ProductPtr
		product.Pts = val.ProductPts
		ti.Products = append(ti.Products, &product)
		InvoiceMap[val.BillNumber] = ti

	}

	for _, mapVal := range InvoiceMap {
		temp := mapVal
		*Invoice = append(*Invoice, &temp)

	}
	return nil

}
