package model

type NewProduct struct {
	ProductName string `json:"ProductName"`
}
type AWACSProduct struct {
	ProductCode string  `json:"productCode"`
	ProductName string  `json:"productName"`
	Brand       string  `json:"brand"`
	Company     string  `json:"company"`
	DrugType    string  `json:"drugType"`
	Strength    string  `json:"strength"`
	Pack        string  `json:"pack"`
	Mrp         float64 `json:"mrp"`
	Ptr         float64 `json:"ptr"`
	Pts         float64 `json:"pts"`
}
type Product struct {
	ProductCode              string  `json:"productCode"`
	ProductName              string  `json:"productName"`
	Company                  string  `json:"company"`
	Quantity                 float64 `json:"quantity"`
	FreeQuantity             float64 `json:"freeQuantity"`
	SchemePercentage         float64 `json:"schemePercentage"`
	Closing                  float64 `json:"closing"`
	DrugType                 string  `json:"drugType"`
	Strength                 string  `json:"strength"`
	Pack                     string  `json:"pack"`
	Mrp                      float64 `json:"mrp"`
	Ptr                      float64 `json:"ptr"`
	Pts                      float64 `json:"pts"`
	BatchID                  string  `json:"batchId"`
	Hsn                      string  `json:"hsn"`
	TaxableAmount            float64 `json:"TaxableAmount"`
	Remarks                  string  `json:"remarks"`
	Batch                    string  `json:"batch"`
	Expiry                   string  `json:"expiry"`
	Rate                     float64 `json:"rate"`
	Amount                   float64 `json:"amount"`
	Discount                 float64 `json:"discount"`
	DiscountAmount           float64 `json:"discountAmount"`
	AddlScheme               float64 `json:"addlScheme"`
	AddlSchemeAmount         float64 `json:"addlSchemeAmount"`
	AddlDiscount             float64 `json:"addlDiscount"`
	AddlDiscountAmount       float64 `json:"addlDiscountAmount"`
	DeductableBeforeDiscount float64 `json:"deductableBeforeDiscount"`
	MrpInclusiveTax          bool    `json:"mrpInclusiveTax"`
	VatApplication           string  `json:"vatApplication"`
	Vat                      float64 `json:"vat"`
	AddlTax                  float64 `json:"addlTax"`
	Cst                      float64 `json:"cst"`
	SGst                     float64 `json:"sGST"`
	CGst                     float64 `json:"cGST"`
	IGst                     float64 `json:"iGST"`
	OrderNumber              string  `json:"orderNumber"`
	OrderDate                string  `json:"orderDate"`
	Barcode                  string  `json:"barcode"`
	OrderSource              string  `json:"orderSource"`
	BaseSchemeQuantity       float64 `json:"baseSchemeQuantity"`
	BaseSchemeFreeQuantity   float64 `json:"baseSchemeFreeQuantity"`
	SgstAmount               float64 `json:"sgstAmount"`
	CgstAmount               float64 `json:"cgstAmount"`
	IgstAmount               float64 `json:"igstAmount"`
	Cess                     float64 `json:"cess"`
	CessAmount               float64 `json:"cessAmount"`
	Distributors             *User   `json:"distributors"`
}

type ProductDetails struct {
	ProductCode         string  `json:"ProductCode"`
	ProductName         string  `json:"ProductName"`
	DistributorID       string  `json:"DistributorId"`
	DistributorName     string  `json:"DistributorName"`
	DistributorEmail    string  `json:"DistributorEmail"`
	DistributorPincode  string  `json:"DistributorPincode"`
	DistributorMobile   string  `json:"DistributorMobile"`
	DistributorPhoneNo  string  `json:"DistributorPhoneNo"`
	DistributorCity     string  `json:"DistributorCity"`
	DistributorState    string  `json:"DistributorState"`
	DistributorAddress1 string  `json:"DistributorAddress1"`
	Brand               string  `json:"BRAND"`
	Company             string  `json:"Company"`
	Quantity            float64 `json:"Quantity"`
	FreeQuantity        float64 `json:"FreeQuantity"`
	SchemePercentage    float64 `json:"SchemePercentage"`
	ProductClosing      float64 `json:"ProductClosing"`
	DrugType            string  `json:"DRUG_TYPE"`
	Strength            string  `json:"STRENGTH"`
	ProductPack         string  `json:"ProductPack"`
	Mrp                 float64 `json:"Mrp"`
	Ptr                 float64 `json:"Ptr"`
	Pts                 float64 `json:"Pts"`
	CityId              int     `json:"CityId"`
	StateId             int     `json:"StateId"`
	AwacsId             string  `json:"AwacsId"`
	BuyerName           string  `json:"BuyerName"`
	BuyerEmail          string  `json:"BuyerEmail"`
	BuyerPincode        string  `json:"BuyerPincode"`
	BuyerMobile         string  `json:"BuyerMobile"`
	BuyerPhoneNo        string  `json:"BuyerPhoneNo"`
	BuyerCity           string  `json:"BuyerCity"`
	BuyerState          string  `json:"BuyerState"`
	BuyerAddress1       string  `json:"BuyerAddress1"`
}

//TableName retunrs source table name
func (AWACSProduct) TableName() string {
	return "dbo.AWACS_ItemMaster_Hdr"
}

//TableName retunrs source table name
func (ProductDetails) TableName() string {
	return "dbo.V_Product_Distributor_details"
}
