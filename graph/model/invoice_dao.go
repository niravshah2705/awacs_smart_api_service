package model

type Invoice struct {
	ID                  string     `json:"id"`
	DeveloperID         *string    `json:"developerId"`
	SupplierID          string     `json:"supplierId"`
	BillNumber          string     `json:"billNumber"`
	BillDate            string     `json:"billDate"`
	ChallanNumber       *string    `json:"challanNumber"`
	ChallanDate         *string    `json:"challanDate"`
	BuyerID             string     `json:"buyerId"`
	Reason              *string    `json:"reason"`
	Upc                 *string    `json:"upc"`
	PaymentDueDate      *string    `json:"paymentDueDate"`
	NetInvoiceAmount    *float64   `json:"netInvoiceAmount"`
	NetPurchaseAmount   *float64   `json:"netPurchaseAmount"`
	LastTransactionDate *string    `json:"lastTransactionDate"`
	ReceivedOn          *string    `json:"receivedOn"`
	EmailPrepared       *bool      `json:"emailPrepared"`
	EmailPreparedOn     *string    `json:"emailPreparedOn"`
	Taken               *string    `json:"taken"`
	TakenOn             *string    `json:"takenOn"`
	LastHostAddress     *string    `json:"lastHostAddress"`
	WorkspaceID         *int       `json:"workspaceId"`
	Ede2Taken           *bool      `json:"ede2_Taken"`
	Ede2TakenDate       *string    `json:"ede2_TakenDate"`
	Ede2TakenBatchID    *string    `json:"ede2_TakenBatchID"`
	TaxableAmount       *float64   `json:"taxableAmount"`
	Products            []*Product `json:"products"`
	Suppliers           []*User    `json:"suppliers"`
	Buyers              []*User    `json:"buyers"`
}
type InvoiceDetails struct {
	ID                       string   `json:"id"`
	DeveloperID              *string  `gorm:"column:DeveloperId";json:"developerId"`
	SupplierID               string   `gorm:"column:SupplierId";json:"supplierId"`
	BillNumber               string   `json:"billNumber"`
	BillDate                 string   `json:"billDate"`
	ChallanNumber            *string  `json:"challanNumber"`
	ChallanDate              *string  `json:"challanDate"`
	BuyerID                  string   `gorm:"column:BuyerId";json:"buyerId"`
	Reason                   *string  `json:"reason"`
	Upc                      *string  `json:"upc"`
	SupplierProductCode      string   `json:"supplierProductCode"`
	SupplierProductName      string   `json:"supplierProductName"`
	SupplierProductPack      string   `json:"supplierProductPack"`
	Mrp                      float64  `json:"mrp"`
	Batch                    string   `json:"batch"`
	Expiry                   string   `json:"expiry"`
	Quantity                 float64  `json:"quantity"`
	FreeQuantity             float64  `json:"freeQuantity"`
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
	BaseSchemeQuantity       *float64 `json:"baseSchemeQuantity"`
	BaseSchemeFreeQuantity   *float64 `json:"baseSchemeFreeQuantity"`
	PaymentDueDate           *string  `json:"paymentDueDate"`
	Remarks                  *string  `json:"remarks"`
	CompanyName              *string  `json:"companyName"`
	NetInvoiceAmount         *float64 `json:"netInvoiceAmount"`
	NetPurchaseAmount        *float64 `json:"netPurchaseAmount"`
	LastTransactionDate      *string  `json:"lastTransactionDate"`
	ReceivedOn               *string  `json:"receivedOn"`
	BatchID                  *string  `json:"batchId"`
	EmailPrepared            *bool    `json:"emailPrepared"`
	EmailPreparedOn          *string  `json:"emailPreparedOn"`
	Taken                    *string  `json:"taken"`
	TakenOn                  *string  `json:"takenOn"`
	LastHostAddress          *string  `json:"lastHostAddress"`
	WorkspaceID              *int     `json:"workspaceId"`
	Ede2Taken                *bool    `json:"ede2_Taken"`
	Ede2TakenDate            *string  `json:"ede2_TakenDate"`
	Ede2TakenBatchID         *string  `json:"ede2_TakenBatchID"`
	SgstAmount               *float64 `json:"sgstAmount"`
	CgstAmount               *float64 `json:"cgstAmount"`
	IgstAmount               *float64 `json:"igstAmount"`
	Cess                     *float64 `json:"cess"`
	CessAmount               *float64 `json:"cessAmount"`
	TaxableAmount            *float64 `json:"taxableAmount"`
	Hsn                      *string  `json:"hsn"`
	OrderNumber              *string  `json:"orderNumber"`
	OrderDate                *string  `json:"orderDate"`
	Barcode                  *string  `json:"barcode"`
	OrderSource              *string  `json:"orderSource"`
	ProductCode              string   `json:"ProductCode"`
	ProductName              string   `json:"ProductName"`
	ProductCompany           string   `json:"Company"`
	ProductPack              string   `json:"Pack"`
	ProductQuantity          float64  `json:"Quantity"`
	ProductFreeQuantity      *float64 `json:"FreeQuantity"`
	ProductSchemePercentage  *float64 `json:"SchemePercentage"`
	ProductClosing           *float64 `json:"Closing"`
	ProductMrp               float64  `json:"MRP"`
	ProductPtr               *float64 `json:"PTR"`
	ProductPts               *float64 `json:"PTS"`
	SupplierName             string   `json:"name"`
	SupplierAwacsId          string   `json:"AwacsId"`
	SupplierEmail            string   `json:"Email"`
	SupplierPincode          string   `json:"Pincode"`
	SupplierMobile           string   `json:"Mobile"`
	SupplierPhoneNo          string   `json:"PhonNo"`
	SupplierCity             string   `json:"city"`
	SupplierState            string   `json:"state"`
	BuyerName                string   `json:"name"`
	BuyerAwacsId             string   `json:"AwacsId"`
	BuyerEmail               string   `json:"Email"`
	BuyerPincode             string   `json:"Pincode"`
	BuyerMobile              string   `json:"Mobile"`
	BuyerPhoneNo             string   `json:"PhonNo"`
	BuyerCity                string   `json:"city"`
	BuyerState               string   `json:"state"`
}

//TableName retunrs source table name
func (Invoice) TableName() string {
	return "dbo.Invoices"
}

//TableName retunrs source table name
func (InvoiceDetails) TableName() string {
	return "dbo.V_INVOICE_DETAILS"
}
