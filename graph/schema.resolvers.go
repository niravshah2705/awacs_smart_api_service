package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"awacs_smart_api_service/graph/generated"
	"awacs_smart_api_service/graph/model"
	"awacs_smart_api_service/services"
	"context"
	"fmt"
	"log"

	db "github.com/brkelkar/common_utils/databases"
)

func (r *mutationResolver) CreateProduct(ctx context.Context, input model.NewProduct) (*model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Products(ctx context.Context, productName string) ([]*model.Product, error) {
	var Product []*model.Product
	err := services.Products(&Product, productName)
	if err != nil {
		fmt.Print(err)
		return nil, err

	}
	return Product, nil
}

func (r *queryResolver) ProductsAwacs(ctx context.Context, productName string) ([]*model.AWACSProduct, error) {
	var AwacsProduct []*model.AWACSProduct
	err := services.ProductsAwacs(&AwacsProduct, productName)
	if err != nil {
		fmt.Print(err)
		return nil, err

	}
	return AwacsProduct, nil
}

func (r *queryResolver) ProductByItemCode(ctx context.Context, itemCode string) ([]*model.AWACSProduct, error) {
	var Product []*model.AWACSProduct
	err := services.ProductByItemCode(&Product, itemCode)
	if err != nil {
		fmt.Print(err)
		return nil, err

	}
	return Product, nil
}

func (r *queryResolver) UserByID(ctx context.Context, id string) (*model.User, error) {
	var User model.User
	err := services.UserByID(&User, id)
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
	err := services.UserByMobile(&User, mobile)
	if err != nil {
		log.Print(err)
		return nil, err

	}
	return &User, nil
}

func (r *queryResolver) OrderByUserID(ctx context.Context, buyerID string) ([]*model.Order, error) {
	var Order []*model.Order
	err := services.OrderByUserId(&Order, buyerID)
	if err != nil {
		log.Print(err)
		return nil, err

	}
	return Order, nil
}

func (r *queryResolver) InvoiceDetails(ctx context.Context, buyerID string, supplierID string, fromDate string, toDate string) ([]*model.Invoice, error) {
	var Invoice []*model.Invoice
	err := services.InvoiceDetails(&Invoice, buyerID, supplierID, fromDate, toDate)
	if err != nil {
		log.Print(err)
		return nil, err

	}
	return Invoice, nil
}

func (r *queryResolver) InvoiceByBuyer(ctx context.Context, buyerID string, fromDate string, toDate string) (*model.InvoiceBuyer, error) {
	var Invoice model.InvoiceBuyer
	err := services.InvoiceByBuyer(&Invoice, buyerID, fromDate, toDate)
	if err != nil {
		log.Print(err)
		return nil, err

	}
	return &Invoice, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *invoiceResolver) Suppliers(ctx context.Context, obj *model.Invoice) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *invoiceResolver) Buyers(ctx context.Context, obj *model.Invoice) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

type invoiceResolver struct{ *Resolver }

func (r *invoiceResolver) Products(ctx context.Context, obj *model.Invoice) ([]*model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *invoiceResolver) Supplier(ctx context.Context, obj *model.Invoice) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *invoiceResolver) Buyer(ctx context.Context, obj *model.Invoice) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) UserByPhoneNumber(ctx context.Context, phoneNo string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}
