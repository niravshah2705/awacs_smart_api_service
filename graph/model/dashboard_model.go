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
	Buyer               *User                 `json:"Buyer"`
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
	BuyerName          string  `json:"buyerName"`
	BuyerEmail         string  `json:"buyerEmail"`
	BuyerMobile        string  `json:"buyerMobile"`
	BuyerPincode       string  `json:"buyerPincode"`
	BuyerCity          string  `json:"buyerCity"`
	BuyerState         string  `json:"buyerState"`
	BuyerAddress1      string  `json:"buyerAddress1"`
}
