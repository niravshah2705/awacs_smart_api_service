package model

type NewProduct struct {
	ProductName string `json:"ProductName"`
}
type AWACSProduct struct {
	ProductCode string   `json:"productCode"`
	ProductName string   `json:"productName"`
	Brand       string   `json:"brand"`
	Company     string   `json:"company"`
	DrugType    *string  `json:"drugType"`
	Strength    *string  `json:"strength"`
	Pack        string   `json:"pack"`
	Mrp         float64  `json:"mrp"`
	Ptr         *float64 `json:"ptr"`
	Pts         *float64 `json:"pts"`
}
type Product struct {
	ProductCode              string   `json:"productCode"`
	ProductName              string   `json:"productName"`
	Company                  string   `json:"company"`
	Quantity                 float64  `json:"quantity"`
	FreeQuantity             *float64 `json:"freeQuantity"`
	SchemePercentage         *float64 `json:"schemePercentage"`
	Closing                  *float64 `json:"closing"`
	DrugType                 *string  `json:"drugType"`
	Strength                 *string  `json:"strength"`
	Pack                     string   `json:"pack"`
	Mrp                      float64  `json:"mrp"`
	Ptr                      *float64 `json:"ptr"`
	Pts                      *float64 `json:"pts"`
	BatchID                  *string  `json:"batchId"`
	Hsn                      *string  `json:"hsn"`
	TaxableAmount            *float64 `json:"TaxableAmount"`
	Remarks                  *string  `json:"remarks"`
	Batch                    *string  `json:"batch"`
	Expiry                   *string  `json:"expiry"`
	Rate                     *float64 `json:"rate"`
	Amount                   *float64 `json:"amount"`
	Discount                 *float64 `json:"discount"`
	DiscountAmount           *float64 `json:"discountAmount"`
	AddlScheme               *float64 `json:"addlScheme"`
	AddlSchemeAmount         *float64 `json:"addlSchemeAmount"`
	AddlDiscount             *float64 `json:"addlDiscount"`
	AddlDiscountAmount       *float64 `json:"addlDiscountAmount"`
	DeductableBeforeDiscount *float64 `json:"deductableBeforeDiscount"`
	MrpInclusiveTax          *bool    `json:"mrpInclusiveTax"`
	VatApplication           *string  `json:"vatApplication"`
	Vat                      *float64 `json:"vat"`
	AddlTax                  *float64 `json:"addlTax"`
	Cst                      *float64 `json:"cst"`
	SGst                     *float64 `json:"sGST"`
	CGst                     *float64 `json:"cGST"`
	IGst                     *float64 `json:"iGST"`
	OrderNumber              *string  `json:"orderNumber"`
	OrderDate                *string  `json:"orderDate"`
	Barcode                  *string  `json:"barcode"`
	OrderSource              *string  `json:"orderSource"`
	BaseSchemeQuantity       *float64 `json:"baseSchemeQuantity"`
	BaseSchemeFreeQuantity   *float64 `json:"baseSchemeFreeQuantity"`
	SgstAmount               *float64 `json:"sgstAmount"`
	CgstAmount               *float64 `json:"cgstAmount"`
	IgstAmount               *float64 `json:"igstAmount"`
	Cess                     *float64 `json:"cess"`
	CessAmount               *float64 `json:"cessAmount"`
	Distributors             []*User  `json:"distributors"`
}
type ProductDetails struct {
	ProductCode      string   `json:"ProductCode"`
	ProductName      string   `json:"ProductName"`
	DistributorID    string   `json:"DistributorId"`
	Brand            string   `json:"BRAND"`
	Company          string   `json:"Company"`
	Quantity         float64  `json:"Quantity"`
	FreeQuantity     *float64 `json:"FreeQuantity"`
	SchemePercentage *float64 `json:"SchemePercentage"`
	Closing          *float64 `json:"Closing"`
	DrugType         *string  `json:"DRUG_TYPE"`
	Strength         *string  `json:"STRENGTH"`
	Pack             string   `json:"PACK"`
	Mrp              float64  `json:"MRP"`
	Ptr              *float64 `json:"PTR"`
	Pts              *float64 `json:"PTS"`
	Name             string   `json:"name"`
	CityId           int      `json:"CityId"`
	StateId          int      `json:"StateId"`
	AwacsId          string   `json:"AwacsId"`
	Email            string   `json:"Email"`
	Pincode          string   `json:"Pincode"`
	Mobile           string   `json:"Mobile"`
	PhoneNo          string   `json:"PhonNo"`
	City             string   `json:"city"`
	State            string   `json:"state"`
}

//TableName retunrs source table name
func (AWACSProduct) TableName() string {
	return "dbo.AWACS_ItemMaster_Hdr"
}

//TableName retunrs source table name
func (ProductDetails) TableName() string {
	return "dbo.V_Product_Distributor_details"
}
