package model
type NewProduct struct {
	ProductName string `json:"ProductName"`
}

type Product struct {
	ProductCode      string   `json:"ProductCode"`
	ProductName      string   `json:"ProductName"`
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
	Distributors     []*User  `json:"Distributors"`
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
	Name              string  `json:"name"` 
	CityId			 int 	   `json:"CityId"`	
	StateId			 int 	   `json:"StateId"`	
	AwacsId		 string 	   `json:"AwacsId"`	
	Email		 string 	   `json:"Email"`	
	Pincode      string 		`json:"Pincode"`
	Mobile      string 		`json:"Mobile"`	
	PhoneNo      string 		`json:"PhonNo"`	
}



//TableName retunrs source table name
func (Product) TableName() string {
	return "dbo.AWACS_ItemMaster_Hdr"
}

//TableName retunrs source table name
func (ProductDetails) TableName() string {
	return "dbo.V_Product_Distributor_details"
}
