package model

type User struct {
	ID                    string   `json:"id"`
	ChangePasswordOnLogon bool     `json:"changePasswordOnLogon"`
	PictureID             int      `json:"pictureId"`
	WorkspaceID           int      `json:"workspaceId"`
	Username              string   `json:"username"`
	Role                  string   `json:"role"`
	Name                  string   `json:"name"`
	Address1              string   `json:"address1"`
	Address2              string   `json:"address2"`
	Address3              string   `json:"address3"`
	CityID                int      `json:"cityId"`
	City                  string   `json:"city"`
	State                 string   `json:"state"`
	StateID               int      `json:"stateId"`
	Country               string   `json:"country"`
	PhoneNo               string   `json:"phoneNo"`
	Mobile                string   `json:"mobile"`
	Email                 string   `json:"email"`
	IsLiveorderBlocked    bool     `json:"isLiveorderBlocked"`
	IsVerify              bool     `json:"isVerify"`
	IsAdmin               bool     `json:"isAdmin"`
	IsAdminDelegate       string   `json:"isAdminDelegate"`
	IsAdhoc               string   `json:"isAdhoc"`
	DeviceID              string   `json:"deviceId"`
	GCMRegisterKey        string   `json:"gCMRegisterKey"`
	Type                  string   `json:"type"`
	LoginType             string   `json:"loginType"`
	ProviderID            int      `json:"providerId"`
	ProviderType          string   `json:"providerType"`
	AwacsID               string   `json:"awacsId"`
	PObotp                string   `json:"pOBOTP"`
	GStin                 string   `json:"gSTIN"`
	Pincode               string   `json:"pincode"`
	RegistrationID        string   `json:"registrationId"`
	Source                string   `json:"source"`
	SourceType            string   `json:"sourceType"`
	SignSource            string   `json:"signSource"`
	SignDetail            string   `json:"signDetail"`
	Deleted               string   `json:"deleted"`
	DeletedBy             int      `json:"deletedBy"`
	DeletedOn             string   `json:"deletedOn"`
	LastActivity          string   `json:"lastActivity"`
	AppVersion            string   `json:"appVersion"`
	OSVersion             string   `json:"oSVersion"`
	Platform              string   `json:"platform"`
	UserKind              string   `json:"userKind"`
	Bank                  string   `json:"bank"`
	IFSCCode              string   `json:"iFSCCode"`
	UserID                string   `json:"userId"`
	Orders                []*Order `json:"orders"`
}

//TableName retunrs source table name
func (User) TableName() string {
	return "dbo.Users"
}
