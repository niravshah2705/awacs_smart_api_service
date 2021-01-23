package model
type Order struct {
	ID             string     `json:"id"`
	Source         *string    `json:"Source"`
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
}


type OrderDetails struct{
	ID             string          `json:"id"`
	Source         *string         `json:"Source"`
	BuyerID        *string         `json:"buyerId"`
	OrderNumber    string          `json:"orderNumber"`
	OrderDate      *string         `json:"orderDate"`
	SupplierID     string          `json:"supplierId"`
	TakenOn        *string         `json:"takenOn"`
	DeliveryOption *string         `json:"deliveryOption"`
	DeliveryMode   *string         `json:"deliveryMode"`
	Remarks        *string         `json:"remarks"`
	WorkspaceID    *int            `json:"workspaceId"`
	Status         *string         `json:"status"`
	InvoiceNumber  *string         `json:"invoiceNumber"`
	IsBounced      *bool           `json:"isBounced"`
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
}


//TableName retunrs source table name
func (OrderDetails) TableName() string {
	return "dbo.V_ORDERS_BY_ORDER_NUMBER"
}