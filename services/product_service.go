package services

import (
	"awacs_smart_api_service/graph/model"
	"reflect"

	db "github.com/brkelkar/common_utils/databases"
	"github.com/brkelkar/common_utils/logger"
)

var (
	productQuery = `ItemCode ProductCode ,
					SKU ProductName,
					BRAND Brand,
					COMPANY Company,
					DRUG_TYPE DrugType,
					STRENGTH Strength,
					PACK Pack,
					MRP Mrp,
					PTR Ptr,
					PTS Pts`
)

//Products get product details
func Products(Product *[]*model.Product, productName string) (err error) {
	var P []*model.ProductDetails
	err = db.DB["smartdb"].Limit(1000).Where("ProductName like '" + productName + "%'").Order("ProductName").Find(&P).Error

	if err != nil {
		logger.Error("Product details: ", err)
		return

	}
	productMap := make(map[string]model.Product, len(P)+1)
	for _, val := range P {
		key := val.ProductName + val.DistributorID
		if _, ok := productMap[key]; !ok {
			//fill product details
			tempProduct := BindProductDetails(*val)
			tempUser := BindUserDetails("Buyer", *val)
			tempProduct.Distributors = &tempUser
			productMap[key] = tempProduct
		}
		tempProduct := productMap[key]
		//fill buyer details
		productMap[key] = tempProduct

	}

	for _, val := range productMap {
		temp := val
		*Product = append(*Product, &temp)
	}
	return nil
}

//ProductsAwacs product awacs
func ProductsAwacs(AwacsProduct *[]*model.AWACSProduct, productName string) (err error) {
	err = db.DB["awacs"].Limit(5000).Select(productQuery).Where(
		"STATUS='Active' and SKU LIKE '" + productName + "%'").Find(
		AwacsProduct).Error
	if err != nil {
		logger.Error("Product Awacs details: ", err)
	}

	return
}

//ProductByItemCode get product details item code
func ProductByItemCode(Product *[]*model.AWACSProduct, itemCode string) (err error) {
	err = db.DB["awacs"].Select(productQuery).Where(
		"STATUS='Active' and ItemCode=" + itemCode).Find(Product).Error
	if err != nil {
		logger.Error("Product details by itemcode: ", err)
	}
	return
}

//BindProductDetails assign to product details
func BindProductDetails(val interface{}) (product model.Product) {
	value := reflect.ValueOf(val)
	product.ProductCode = value.FieldByName("ProductCode").String()
	product.ProductName = value.FieldByName("ProductName").String()
	product.Company = value.FieldByName("Company").String()
	product.Quantity = value.FieldByName("Quantity").Float()
	product.FreeQuantity = value.FieldByName("FreeQuantity").Float()
	product.SchemePercentage = value.FieldByName("SchemePercentage").Float()
	product.Closing = value.FieldByName("ProductClosing").Float()
	product.DrugType = value.FieldByName("DrugType").String()
	product.Strength = value.FieldByName("Strength").String()
	product.Pack = value.FieldByName("ProductPack").String()
	product.Mrp = value.FieldByName("Mrp").Float()
	product.Ptr = value.FieldByName("Ptr").Float()
	product.Pts = value.FieldByName("Pts").Float()
	return
}

//BindAdditionalProductDetails assign to product details
func BindAdditionalProductDetails(val interface{}, product *model.Product) {
	value := reflect.ValueOf(val)
	product.Rate = value.FieldByName("Rate").Float()
	product.Amount = value.FieldByName("Amount").Float()
	product.Discount = value.FieldByName("Discount").Float()
	product.DiscountAmount = value.FieldByName("DiscountAmount").Float()
	product.AddlScheme = value.FieldByName("AddlScheme").Float()
	product.AddlSchemeAmount = value.FieldByName("AddlSchemeAmount").Float()
	product.AddlDiscount = value.FieldByName("AddlDiscount").Float()
	product.AddlDiscountAmount = value.FieldByName("AddlDiscountAmount").Float()
	product.DeductableBeforeDiscount = value.FieldByName("DeductableBeforeDiscount").Float()
	product.MrpInclusiveTax = value.FieldByName("MrpInclusiveTax").Bool()
	product.VatApplication = value.FieldByName("VatApplication").String()
	product.Vat = value.FieldByName("Vat").Float()
	product.AddlTax = value.FieldByName("AddlTax").Float()
	product.Cst = value.FieldByName("Cst").Float()
	product.SGst = value.FieldByName("SGst").Float()
	product.CGst = value.FieldByName("CGst").Float()
	product.IGst = value.FieldByName("IGst").Float()
	product.BaseSchemeQuantity = value.FieldByName("BaseSchemeQuantity").Float()
	product.BaseSchemeFreeQuantity = value.FieldByName("BaseSchemeFreeQuantity").Float()
	product.Remarks = value.FieldByName("Remarks").String()
	product.BatchID = value.FieldByName("BatchID").String()
	product.SgstAmount = value.FieldByName("SgstAmount").Float()
	product.CgstAmount = value.FieldByName("CgstAmount").Float()
	product.IgstAmount = value.FieldByName("IgstAmount").Float()
	product.Cess = value.FieldByName("Cess").Float()
	product.CessAmount = value.FieldByName("CessAmount").Float()
	product.Hsn = value.FieldByName("Hsn").String()
	product.OrderNumber = value.FieldByName("OrderNumber").String()
	product.OrderDate = value.FieldByName("OrderDate").String()
	product.Barcode = value.FieldByName("Barcode").String()
	product.OrderSource = value.FieldByName("OrderSource").String()

	return
}
