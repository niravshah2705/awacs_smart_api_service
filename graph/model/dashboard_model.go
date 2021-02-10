package model

type BuyerDashboard struct {
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