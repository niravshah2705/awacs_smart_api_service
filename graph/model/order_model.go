package model

type Order struct {
	ID             string     `json:"id"`
	Source         string     `json:"source"`
	BuyerID        string     `json:"buyerId"`
	OrderNumber    string     `json:"orderNumber"`
	OrderDate      string     `json:"orderDate"`
	SupplierID     string     `json:"supplierId"`
	TakenOn        string     `json:"takenOn"`
	DeliveryOption string     `json:"deliveryOption"`
	DeliveryMode   string     `json:"deliveryMode"`
	Remarks        string     `json:"remarks"`
	WorkspaceID    int        `json:"workspaceId"`
	Status         string     `json:"status"`
	InvoiceNumber  string     `json:"invoiceNumber"`
	IsBounced      bool       `json:"isBounced"`
	Products       []*Product `json:"products"`
	Buyer          *User      `json:"buyer"`
	Supplier       *User      `json:"supplier"`
}

type OrderDetails struct {
	ID               string  `json:"id"`
	Source           string  `json:"Source"`
	BuyerID          string  `json:"buyerId"`
	OrderNumber      string  `json:"orderNumber"`
	OrderDate        string  `json:"orderDate"`
	SupplierID       string  `json:"supplierId"`
	TakenOn          string  `json:"takenOn"`
	DeliveryOption   string  `json:"deliveryOption"`
	DeliveryMode     string  `json:"deliveryMode"`
	Remarks          string  `json:"remarks"`
	WorkspaceID      int     `json:"workspaceId"`
	Status           string  `json:"status"`
	InvoiceNumber    string  `json:"invoiceNumber"`
	IsBounced        bool    `json:"isBounced"`
	ProductCode      string  `json:"ProductCode"`
	ProductName      string  `json:"ProductName"`
	Company          string  `json:"Company"`
	ProductPack      string  `json:"ProductPack"`
	Quantity         float64 `json:"Quantity"`
	FreeQuantity     float64 `json:"FreeQuantity"`
	SchemePercentage float64 `json:"SchemePercentage"`
	ProductClosing   float64 `json:"ProductClosing"`
	Mrp              float64 `json:"MRP"`
	Ptr              float64 `json:"PTR"`
	Pts              float64 `json:"PTS"`
	SupplierName     string  `json:"supplierName"`
	SupplierEmail    string  `json:"supplierEmail"`
	SupplierMobile   string  `json:"supplierMobile"`
	SupplierPincode  string  `json:"supplierPincode"`
	SupplierCity     string  `json:"supplierCity"`
	SupplierState    string  `json:"supplierState"`
	BuyerName        string  `json:"buyerName"`
	BuyerEmail       string  `json:"buyerEmail"`
	BuyerMobile      string  `json:"buyerMobile"`
	BuyerPincode     string  `json:"buyerPincode"`
	BuyerCity        string  `json:"buyerCity"`
	BuyerState       string  `json:"buyerState"`
}

type OrderBuyerStatus struct {
	Buyer       *User          `json:"buyer"`
	OrderStatus []*OrderStatus `json:"orderStatus"`
}

type OrderSupplierStatus struct {
	Supplier    *User          `json:"supplier"`
	OrderStatus []*OrderStatus `json:"orderStatus"`
}

type OrderStatus struct {
	OrderDate string `json:"orderDate"`
	Billed    string `json:"billed"`
	Bounced   string `json:"bounced"`
	Pending   string `json:"pending"`
}

type OrderBuyerStatusDetails struct {
	BuyerId      string `json:"buyerId"`
	OrderDate    string `json:"orderDate"`
	Pending      string `json:"pending"`
	Bounced      string `json:"bounced"`
	Billed       string `json:"billed"`
	BuyerName    string `json:"buyerName"`
	BuyerCity    string `json:"City"`
	BuyerEmail   string `json:"Email"`
	BuyerMobile  string `json:"Mobile"`
	BuyerPhoneNo string `json:"PhoneNo"`
	BuyerPincode string `json:"Pincode"`
	BuyerState   string `json:"State"`
}

type OrderSupplierStatusDetails struct {
	SupplierId      string `json:"supplierID"`
	OrderDate       string `json:"orderDate"`
	Pending         string `json:"pending"`
	Bounced         string `json:"bounced"`
	Billed          string `json:"billed"`
	SupplierName    string `json:"supplierName"`
	SupplierCity    string `json:"city"`
	SupplierEmail   string `json:"email"`
	SupplierMobile  string `json:"mobile"`
	SupplierPhoneNo string `json:"phoneNo"`
	SupplierPincode string `json:"pincode"`
	SupplierState   string `json:"state"`
}

type OrderByWorkspaceID struct {
	Buyer    *User    `json:"buyer"`
	Supplier *User    `json:"supplier"`
	Order    []*Order `json:"Order"`
}

type SmartOrders struct {
	Retailer         *User    `json:"Retailer"`
	AggregatorOrders []*Order `json:"AggregatorOrders"`
	RetailerOrders   []*Order `json:"RetailerOrders"`
}

//SmartOrdersDetails model
type SmartOrdersDetails struct {
	OrderType           string `json:"OrderType"`
	ID                  string `json:"ID"`
	Source              string `json:"Source"`
	BuyerID             string `json:"BuyerID"`
	OrderNumber         string `json:"OrderNumber"`
	OrderDate           string `json:"OrderDate"`
	SupplierID          string `json:"SupplierID"`
	TakenOn             string `json:"TakenOn"`
	SupplierProductCode string `json:"SupplierProductCode"`
	BuyerProductCode    string `json:"BuyerProductCode"`
	BuyerProductName    string `json:"BuyerProductName"`
	BuyerProductPack    string `json:"BuyerProductPack"`
	Quantity            string `json:"Quantity"`
	FreeQuantity        string `json:"FreeQuantity"`
	ReceivedOn          string `json:"ReceivedOn"`
	DeliveryOption      string `json:"DeliveryOption"`
	DeliveryMode        string `json:"DeliveryMode"`
	Remarks             string `json:"Remarks"`
	WorkspaceID         string `json:"WorkspaceID"`
	CreatedBy           string `json:"CreatedBy"`
	CreatedDate         string `json:"CreatedDate"`
	Status              string `json:"Status"`
	SuppliedQty         string `json:"SuppliedQty"`
	InvoiceNumber       string `json:"InvoiceNumber"`
	IsBounced           string `json:"IsBounced"`
	IsHalfScheme        string `json:"IsHalfScheme"`
	Supplied_FreeQty    string `json:"Supplied_FreeQty"`
	RetailerCode        string `json:"RetailerCode"`
	RetailerName        string `json:"RetailerName"`
	RetailerEmail       string `json:"RetailerEmail"`
	RetailerMobile      string `json:"RetailerMobile"`
	RetailerPincode     string `json:"RetailerPincode"`
	RetailerAddress1    string `json:"RetailerAddress1"`
	RetailerAddress2    string `json:"RetailerAddress2"`
	UserName            string `json:"UserName"`
	RetailerCity        string `json:"RetailerCity"`
	RetailerState       string `json:"RetailerState"`
}

//TableName retunrs source table name
func (OrderDetails) TableName() string {
	return "dbo.V_ORDERS_BY_ORDER_NUMBER"
}

//TableName retunrs source table name
func (OrderSupplierStatusDetails) TableName() string {
	return "dbo.V_SUPPLIERID_WISE_ORDERSTATUS_COUNT"
}

//TableName retunrs source table name
func (OrderBuyerStatusDetails) TableName() string {
	return "dbo.V_BUYERID_WISE_ORDERSTATUS_COUNT"
}

//TableName retunrs source table name
func (SmartOrdersDetails) TableName() string {
	return "dbo.V_SMARTAGGREGATOR_MYORDERS"
}
