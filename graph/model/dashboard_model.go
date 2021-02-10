package model

type BuyerDashboard struct {
	SupplierID         string `json:"SupplierId"`
	SupplierName       string `json:"SupplierName"`
	Pending            string `json:"Pending"`
	Bounced            string `json:"Bounced"`
	Billed             string `json:"Billed"`
	Short              string `json:"Short"`
	TotalOrder         string `json:"TotalOrder"`
	CurrentOutstanding string `json:"CurrentOutstanding"`
	ProductCount       string `json:"ProductCount"`
}
