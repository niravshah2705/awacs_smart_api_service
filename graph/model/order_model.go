package model

//Order model
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

//OrderDetails model
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

//OrderBuyerStatus model
type OrderBuyerStatus struct {
	Buyer       *User          `json:"buyer"`
	OrderStatus []*OrderStatus `json:"orderStatus"`
}

//OrderSupplierStatus model
type OrderSupplierStatus struct {
	Supplier    *User          `json:"supplier"`
	OrderStatus []*OrderStatus `json:"orderStatus"`
}

//OrderStatus order status model
type OrderStatus struct {
	OrderDate string `json:"orderDate"`
	Billed    string `json:"billed"`
	Bounced   string `json:"bounced"`
	Pending   string `json:"pending"`
}

//OrderBuyerStatusDetails buyer order status 
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

//OrderSupplierStatusDetails supplier order status model
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

//OrderByWorkspaceID order worksapce model
type OrderByWorkspaceID struct {
	Buyer    *User    `json:"buyer"`
	Supplier *User    `json:"supplier"`
	Order    []*Order `json:"Order"`
}

//SmartOrders samrt orders
type SmartOrders struct {
	AggregatorOrders []*User `json:"AggregatorOrders"`
	RetailerOrders   []*User `json:"RetailerOrders"`
}

//SmartOrdersDetails model
type SmartOrdersDetails struct {
	OrderType        string  `json:"OrderType"`
	ID               int64   `json:"ID"`
	Source           string  `json:"Source"`
	BuyerID          string  `json:"BuyerID"`
	OrderNumber      string  `json:"OrderNumber"`
	OrderDate        string  `json:"OrderDate"`
	SupplierID       string  `json:"SupplierID"`
	TakenOn          string  `json:"TakenOn"`
	ProductCode      string  `json:"ProductCode"`
	ProductName      string  `json:"ProductName"`
	Company          string  `json:"Company"`
	ProductClosing   float64 `json:"ProductClosing"`
	DrugType         string  `json:"DrugType"`
	Strength         string  `json:"Strength"`
	Pack             string  `json:"Pack"`
	Mrp              float64 `json:"Mrp"`
	Ptr              float64 `json:"Ptr"`
	Pts              float64 `json:"Pts"`
	SchemePercentage float64 `json:"SchemePercentage"`
	Quantity         float64 `json:"Quantity"`
	FreeQuantity     float64 `json:"FreeQuantity"`
	ReceivedOn       string  `json:"ReceivedOn"`
	DeliveryOption   string  `json:"DeliveryOption"`
	DeliveryMode     string  `json:"DeliveryMode"`
	Remarks          string  `json:"Remarks"`
	WorkspaceID      int64   `json:"WorkspaceID"`
	Status           string  `json:"Status"`
	SuppliedQty      float64 `json:"SuppliedQty"`
	InvoiceNumber    string  `json:"InvoiceNumber"`
	IsBounced        int64   `json:"IsBounced"`
	Supplied_FreeQty float64 `json:"Supplied_FreeQty"`
	RetailerCode     string  `json:"RetailerCode"`
	RetailerName     string  `json:"RetailerName"`
	RetailerEmail    string  `json:"RetailerEmail"`
	RetailerMobile   string  `json:"RetailerMobile"`
	RetailerPincode  string  `json:"RetailerPincode"`
	RetailerAddress1 string  `json:"RetailerAddress1"`
	RetailerAddress2 string  `json:"RetailerAddress2"`
	RetailerCity     string  `json:"RetailerCity"`
	RetailerState    string  `json:"RetailerState"`
	SupplierName     string  `json:"SupplierName"`
	SupplierEmail    string  `json:"SupplierEmail"`
	SupplierMobile   string  `json:"SupplierMobile"`
	SupplierPincode  string  `json:"SupplierPincode"`
	SupplierAddress1 string  `json:"SupplierAddress1"`
	SupplierAddress2 string  `json:"SupplierAddress2"`
	SupplierCity     string  `json:"SupplierCity"`
	SupplierState    string  `json:"SupplierState"`
	UserName         string  `json:"UserName"`
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
