package model

type Order struct {
	ID             string     `json:"id"`
	Source         *string    `json:"source"`
	BuyerID        *string    `json:"buyerId"`
	OrderNumber    string     `json:"orderNumber"`
	OrderDate      *string    `json:"orderDate"`
	SupplierID     string     `json:"supplierId"`
	TakenOn        *string    `json:"takenOn"`
	DeliveryOption *string    `json:"deliveryOption"`
	DeliveryMode   *string    `json:"deliveryMode"`
	Remarks        *string    `json:"remarks"`
	WorkspaceID    *int       `json:"workspaceId"`
	Status         *string    `json:"status"`
	InvoiceNumber  *string    `json:"invoiceNumber"`
	IsBounced      *bool      `json:"isBounced"`
	Products       []*Product `json:"products"`
	Buyer          *User      `json:"buyer"`
	Supplier       *User      `json:"supplier"`
}

type OrderDetails struct {
	ID               string   `json:"id"`
	Source           *string  `json:"Source"`
	BuyerID          *string  `json:"buyerId"`
	OrderNumber      string   `json:"orderNumber"`
	OrderDate        *string  `json:"orderDate"`
	SupplierID       string   `json:"supplierId"`
	TakenOn          *string  `json:"takenOn"`
	DeliveryOption   *string  `json:"deliveryOption"`
	DeliveryMode     *string  `json:"deliveryMode"`
	Remarks          *string  `json:"remarks"`
	WorkspaceID      *int     `json:"workspaceId"`
	Status           *string  `json:"status"`
	InvoiceNumber    *string  `json:"invoiceNumber"`
	IsBounced        *bool    `json:"isBounced"`
	ProductCode      string   `json:"ProductCode"`
	ProductName      string   `json:"ProductName"`
	Company          string   `json:"Company"`
	Pack             string   `json:"Pack"`
	Quantity         float64  `json:"Quantity"`
	FreeQuantity     *float64 `json:"FreeQuantity"`
	SchemePercentage *float64 `json:"SchemePercentage"`
	Closing          *float64 `json:"Closing"`
	Mrp              float64  `json:"MRP"`
	Ptr              *float64 `json:"PTR"`
	Pts              *float64 `json:"PTS"`
	SupplierName     string   `json:"supplierName"`
	SupplierEmail    string   `json:"supplierEmail"`
	SupplierMobile   string   `json:"supplierMobile"`
	SupplierPincode  string   `json:"supplierPincode"`
	SupplierCity     string   `json:"supplierCity"`
	SupplierState    string   `json:"supplierState"`
	BuyerName        string   `json:"buyerName"`
	BuyerEmail       string   `json:"buyerEmail"`
	BuyerMobile      string   `json:"buyerMobile"`
	BuyerPincode     string   `json:"buyerPincode"`
	BuyerCity        string   `json:"buyerCity"`
	BuyerState       string   `json:"buyerState"`
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
	BuyerId   string `json:"buyerId"`
	OrderDate string `json:"orderDate"`
	Pending   string `json:"pending"`
	Bounced   string `json:"bounced"`
	Billed    string `json:"billed"`
	BuyerName string `json:"buyerName"`
	City      string `json:"City"`
	Email     string `json:"Email"`
	Mobile    string `json:"Mobile"`
	PhoneNo   string `json:"PhoneNo"`
	Pincode   string `json:"Pincode"`
	State     string `json:"State"`
}

type OrderStatusSupplierDetails struct {
	SupplierId   string `json:"supplierID"`
	OrderDate    string `json:"orderDate"`
	Pending      string `json:"pending"`
	Bounced      string `json:"bounced"`
	Billed       string `json:"billed"`
	SupplierName string `json:"supplierName"`
	City         string `json:"city"`
	Email        string `json:"email"`
	Mobile       string `json:"mobile"`
	PhoneNo      string `json:"phoneNo"`
	Pincode      string `json:"pincode"`
	State        string `json:"state"`
}

type OrderByWorkspaceID struct {
	Buyer    *User    `json:"buyer"`
	Supplier *User    `json:"supplier"`
	Order    []*Order `json:"Order"`
}

//TableName retunrs source table name
func (OrderDetails) TableName() string {
	return "dbo.V_ORDERS_BY_ORDER_NUMBER"
}

//TableName retunrs source table name
func (OrderStatusSupplierDetails) TableName() string {
	return "dbo.V_SUPPLIERID_WISE_ORDERSTATUS_COUNT"
}

//TableName retunrs source table name
func (OrderBuyerStatusDetails) TableName() string {
	return "dbo.V_BUYERID_WISE_ORDERSTATUS_COUNT"
}
