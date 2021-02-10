package model

//Workspaces model 
type Workspaces struct {
	ID                  string `json:"Id"`
	Name                string `json:"Name"`
	Address1            string `json:"Address1"`
	Ddress2             string `json:"ddress2"`
	Address3            string `json:"Address3"`
	Country             string `json:"Country"`
	State               string `json:"State"`
	City                string `json:"City"`
	Email               string `json:"Email"`
	Pincode             string `json:"Pincode"`
	PhoneNo             string `json:"PhoneNo"`
	Mobile              string `json:"Mobile"`
	WorkSpace           string `json:"WorkSpace"`
	OrderToEmail        string `json:"OrderToEmail"`
	InvoiceToEmail      string `json:"InvoiceToEmail"`
	LiveorderStartFrom  string `json:"LiveorderStartFrom"`
	LiveorderExpireOn   string `json:"LiveorderExpireOn"`
	InfoserverValidUpTo string `json:"InfoserverValidUpTo"`
	HasAdvt             string `json:"HasAdvt"`
	IsLiveorderBlocked  string `json:"IsLiveorderBlocked"`
	LiveorderCount      string `json:"LiveorderCount"`
	LastIPAddress       string `json:"LastIpAddress"`
	Password            string `json:"Password"`
	PictureID           string `json:"PictureID"`
	ProviderID          string `json:"ProviderId"`
	ProviderType        string `json:"ProviderType"`
	Remarks             string `json:"Remarks"`
	TimeZoneID          string `json:"TimeZoneId"`
}

//TableName retunrs source table name
func (Workspaces) TableName() string {
	return "dbo.V_GET_WORKSPCEDETAILS"
}
