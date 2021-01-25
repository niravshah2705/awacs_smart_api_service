package services

import (
	"awacs_smart_api_service/graph/model"
	"fmt"

	db "github.com/brkelkar/common_utils/databases"
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

func Products(Product *[]*model.Product, productName string) (err error) {
	var P []*model.ProductDetails
	err = db.DB["smartdb"].Limit(1000).Where("ProductName like '" + productName + "%'").Order("ProductName").Find(&P).Error

	if err != nil {
		fmt.Print(err)
		return

	}
	productMap := make(map[string]model.Product, len(P)+1)
	for _, pd := range P {
		key := pd.ProductName + pd.DistributorID
		if _, ok := productMap[key]; !ok {
			var tempProduct model.Product
			tempProduct.ProductCode = pd.ProductCode
			tempProduct.ProductName = pd.ProductName
			tempProduct.Company = pd.Company
			tempProduct.Quantity = pd.Quantity
			tempProduct.FreeQuantity = pd.FreeQuantity
			tempProduct.SchemePercentage = pd.SchemePercentage
			tempProduct.Closing = pd.Closing
			tempProduct.DrugType = pd.DrugType
			tempProduct.Strength = pd.Strength
			tempProduct.Pack = pd.Pack
			tempProduct.Mrp = pd.Mrp
			tempProduct.Ptr = pd.Ptr
			tempProduct.Pts = pd.Pts
			productMap[key] = tempProduct
		}
		tempProduct := productMap[key]
		var tempUser model.User
		tempUser.Name = pd.Name
		tempUser.City = pd.City
		tempUser.State = pd.State
		tempUser.AwacsID = &pd.AwacsId
		tempUser.Email = pd.Email
		tempUser.Pincode = &pd.Pincode
		tempUser.Mobile = pd.Mobile
		tempUser.PhoneNo = pd.PhoneNo
		tempProduct.Distributors = append(tempProduct.Distributors, &tempUser)
		productMap[key] = tempProduct

	}

	for _, val := range productMap {
		temp := val
		*Product = append(*Product, &temp)
	}
	return nil
}

func ProductsAwacs(AwacsProduct *[]*model.AWACSProduct, productName string) (err error) {
	err = db.DB["awacs"].Limit(5000).Select(productQuery).Where(
		"STATUS='Active' and SKU LIKE '" + productName + "%'").Find(
		AwacsProduct).Error
	return
}

func ProductByItemCode(Product *[]*model.AWACSProduct, itemCode string) (err error) {
	err = db.DB["awacs"].Select(productQuery).Where(
		"STATUS='Active' and ItemCode=" + itemCode).Find(Product).Error
	return
}
