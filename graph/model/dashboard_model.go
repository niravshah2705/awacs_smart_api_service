package model

type BuyerDashboard struct {
	TotalOrders         int                   `json:"TotalOrders"`
	TotalPendingOrders  int                   `json:"TotalPendingOrders"`
	TotalBilledOrders   int                   `json:"TotalBilledOrders"`
	TotalShortOrders    int                   `json:"TotalShortOrders"`
	TotalBounceOrders   int                   `json:"TotalBounceOrders"`
	TotalOutstanding    float64               `json:"TotalOutstanding"`
	TotalProductOrdered int                   `json:"TotalProductOrdered"`
	Supper              []*SupperOrderDeatils `json:"Supper"`
}

type SupperOrderDeatils struct {
	SupplierID         string  `json:"SupplierId"`
	SupplierName       string  `json:"SupplierName"`
	Pending            int     `json:"Pending"`
	Bounced            int     `json:"Bounced"`
	Billed             int     `json:"Billed"`
	Short              int     `json:"Short"`
	TotalOrder         int     `json:"TotalOrder"`
	CurrentOutstanding float64 `json:"CurrentOutstanding"`
	ProductCount       int     `json:"ProductCount"`
}
