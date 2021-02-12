package dal

import (
	"awacs_smart_api_service/graph/model"

	db "github.com/brkelkar/common_utils/databases"
)

//BuyerDashboardDal get dashboard data
func BuyerDashboard(supplier *[]*model.SupperOrderDeatils, buyerID string, fromDate string, toDate string) (err error) {
	err = db.DB["smartdb"].Raw("exec SP_DASHBOARD_ORDERSTATUS  ?, ?, ?", buyerID, fromDate, toDate).Scan(&supplier).Error

	return
}
