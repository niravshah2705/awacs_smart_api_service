package graph


				// This file will be automatically regenerated based on the schema, any resolver implementations
				// will be copied through when generating and any unknown code will be moved to the end.

import (
"context"
"fmt"

"errors"
"awacs_smart_api_service/graph/generated"
"awacs_smart_api_service/graph/model"
"log"
"strings"
db "github.com/brkelkar/common_utils/databases"
)


















func (r *mutationResolver) CreateProduct(ctx context.Context, input model.NewProduct) (*model.Product, error) {
		panic(fmt.Errorf("not implemented"))
	}

func (r *queryResolver) Products(ctx context.Context, productName string, typeArg string) ([]*model.Product, error) {
		typeArg= strings.ToLower(typeArg)
	var Product []*model.Product
	var err error
	switch typeArg {
	case "awacs": 
	err = db.DB["awacs"].Select(`ItemCode ProductCode ,
	SKU ProductName,
	BRAND Brand,
	COMPANY,
	DRUG_TYPE DrugType,
	STRENGTH Strength,
	PACK Pack,
	MRP Mrp,
	PTR Ptr,
	PTS Pts`).Where(
	"STATUS='Active' and SKU LIKE '" + productName + "%'").Find(&Product).Error

	case "distributor": 
	var P []*model.ProductDetails
	err = db.DB["smartdb"].Where("ProductName like '"+productName+"%'").Find(&P).Error
	if err != nil {
		fmt.Print(err)
		return nil, err

	}
	productMap := make(map[string]model.Product, len(P)+1)
	for _, pd := range P{
		key := pd.ProductName+pd.DistributorID
		if _, ok := productMap[key]; !ok {
			var tempProduct model.Product
			tempProduct.ProductCode=pd.ProductCode
			tempProduct.ProductName=pd.ProductName
			tempProduct.Brand=pd.Brand
			tempProduct.Company=pd.Company
			tempProduct.Quantity=pd.Quantity
			tempProduct.FreeQuantity=pd.FreeQuantity
			tempProduct.SchemePercentage=pd.SchemePercentage
			tempProduct.Closing=pd.Closing
			tempProduct.DrugType=pd.DrugType
			tempProduct.Strength=pd.Strength
			tempProduct.Pack=pd.Pack
			tempProduct.Mrp=pd.Mrp
			tempProduct.Ptr=pd.Ptr
			tempProduct.Pts=pd.Pts
			productMap[key] = tempProduct
		}
		tempProduct := productMap[key]
		var tempUser model.User
		tempUser.Name=pd.Name
		tempUser.CityID=pd.CityId
		tempUser.StateID=pd.StateId
		tempUser.AwacsID=&pd.AwacsId
		tempUser.Email=pd.Email
		tempUser.Pincode=&pd.Pincode
		tempUser.Mobile=pd.Mobile
		tempUser.PhoneNo=pd.PhoneNo
		tempProduct.Distributors = append(tempProduct.Distributors, &tempUser)
		productMap[key] = tempProduct

	}
	
	for _, val := range productMap{
		temp:= val
		Product = append(Product, &temp) 
	}
	return Product, nil
	default: return nil, errors.New("Incorrent product type expected [distributor|awacs]")
	}
	
	
	if err != nil {
		fmt.Print(err)
		return nil, err

	}
	return Product, nil
	}

func (r *queryResolver) ProductByItemCode(ctx context.Context, itemCode string) ([]*model.Product, error) {
		var Product []*model.Product
	err := db.DB["awacs"].Select(`ItemCode,
		SKU ProductName,
		BRAND Brand,
		COMPANY,
		DRUG_TYPE DrugType,
		STRENGTH Strength,
		PACK Pack,
		MRP Mrp,
		PTR Ptr,
		PTS Pts`).Where(
		"STATUS='Active' and ItemCode=" + itemCode).Find(&Product).Error
	if err != nil {
		fmt.Print(err)
		return nil, err

	}
	return Product, nil
	}

func (r *queryResolver) UserByID(ctx context.Context, id string) (*model.User, error) {
		var User model.User
	err := db.DB["smartdb"].First(&User, id).Error
	if err != nil {
		log.Print(err)
		return nil, err

	}
	return &User, nil
	}

func (r *queryResolver) UserByUserName(ctx context.Context, username string) (*model.User, error) {
		var User model.User
	err := db.DB["smartdb"].Where("Username='" + username + "'").First(&User).Error
	if err != nil {
		log.Print(err)
		return nil, err

	}
	return &User, nil
	}

func (r *queryResolver) UserByMobile(ctx context.Context, mobile string) (*model.User, error) {
		var User model.User
	err := db.DB["smartdb"].Where("Mobile='" + mobile + "'").First(&User).Error
	if err != nil {
		log.Print(err)
		return nil, err

	}
	return &User, nil
	}

func (r *queryResolver) OrderByUserID(ctx context.Context, buyerID string) ([]*model.Order, error) {
		var Order []*model.Order
	var OrderDetails []*model.OrderDetails
	err := db.DB["smartdb"].Where("buyerId='" + buyerID + "'").Find(&OrderDetails).Error
	if err != nil {
		log.Print(err)
		return nil, err

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
		Order = append(Order, &temp)

	}
	return Order, nil
	}



// Mutation returns generated.MutationResolver implementation.
	func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }
// Query returns generated.QueryResolver implementation.
	func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }


type mutationResolver struct { *Resolver }
type queryResolver struct { *Resolver }



    // !!! WARNING !!!
    // The code below was going to be deleted when updating resolvers. It has been copied here so you have
    // one last chance to move it out of harms way if you want. There are two reasons this happens:
	//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
	//    it when you're done.
	//  - You have helper methods in this file. Move them out to keep these resolver files clean.
	func (r *queryResolver) UserByPhoneNumber(ctx context.Context, phoneNo string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

