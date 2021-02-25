package dal

import (
	"awacs_smart_api_service/graph/model"

	db "github.com/brkelkar/common_utils/databases"
)

//Products get product data
func Products(Productdetail *[]*model.ProductDetails, clouse string) (err error) {
	err = db.DB["smartdb"].Limit(100).Where(clouse).Order("ProductName").Find(&Productdetail).Error
	return
}

//ProductsAwacs get AWACS product data
func ProductsAwacs(Productdetail *[]*model.AWACSProduct, selectquery string, clouse string) (err error) {
	err = db.DB["awacs"].Limit(100).Select(selectquery).Where(clouse).Order("ProductName").Find(&Productdetail).Error
	return
}
