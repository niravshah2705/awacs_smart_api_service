package model

//Outstanding model
type Outstanding struct {
	UserID         string `json:"userId"`
	CustomerCode   string `json:"customerCode"`
	DocumentNumber string `json:"documentNumber"`
	DocumentDate   string `json:"documentDate"`
}

//TableName retunrs source table name
func (Outstanding) TableName() string {
	return "dbo.Outstanding"
}
