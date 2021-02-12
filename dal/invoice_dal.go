package dal

import (
	"awacs_smart_api_service/graph/model"

	db "github.com/brkelkar/common_utils/databases"
)

//InvoiceDetailsDal get invoice details
func InvoiceDetails(InvoiceDetails *[]*model.InvoiceDetails, clouse string) (err error) {
	err = db.DB["smartdb"].Where(clouse).Find(&InvoiceDetails).Error
	return
}
