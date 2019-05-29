package gocvent

// UserRetrieveResult defines the result wrapper
type UserRetrieveResult struct {
	CvObject User `xml:"RetrieveResult>CvObject"`
}

// User defines the CvObject
type User struct {
	Active                bool   `xml:"Active,attr"`
	Address1              string `xml:"Address1,attr"`
	Address2              string `xml:"Address2,attr"`
	Address3              string `xml:"Address3,attr"`
	AllEventVisibility    bool   `xml:"AllEventVisibility,attr"`
	AllRFPVisibility      bool   `xml:"AllRFPVisibility,attr"`
	AllSurveyVisibility   bool   `xml:"AllSurveyVisibility,attr"`
	ChangePasswordOnLogin bool   `xml:"ChangePasswordOnLogin,attr"`
	City                  string `xml:"City,attr"`
	Company               string `xml:"Company,attr"`
	Country               string `xml:"Country,attr"`
	CountryCode           string `xml:"CountryCode,attr"`
	CreatedBy             string `xml:"CreatedBy,attr"`
	CreatedDate           string `xml:"CreatedDate,attr"`
	DefaultContactGroupId string `xml:"DefaultContactGroupId,attr"`
	Email                 string `xml:"Email,attr"`
	FederatedId           string `xml:"FederatedId,attr"`
	FirstName             string `xml:"FirstName,attr"`
	HomeFax               string `xml:"HomeFax,attr"`
	HomePhone             string `xml:"HomePhone,attr"`
	Id                    string `xml:"Id,attr"`
	LastLoginDate         string `xml:"LastLoginDate,attr"`
	LastModifiedBy        string `xml:"LastModifiedBy,attr"`
	LastModifiedDate      string `xml:"LastModifiedDate,attr"`
	LastName              string `xml:"LastName,attr"`
	MobilePhone           string `xml:"MobilePhone,attr"`
	Pager                 string `xml:"Pager,attr"`
	Password              string `xml:"Password,attr"`
	Passwordstrong        string `xml:"Password (strong),attr"`
	PasswordSalt          string `xml:"Password Salt,attr"`
	PostalCode            string `xml:"PostalCode,attr"`
	Prefix                string `xml:"Prefix,attr"`
	State                 string `xml:"State,attr"`
	StateCode             string `xml:"StateCode,attr"`
	Title                 string `xml:"Title,attr"`
	Username              string `xml:"Username,attr"`
	UserRole              string `xml:"UserRole,attr"`
	UserRoleId            string `xml:"UserRoleId,attr"`
	UserType              string `xml:"UserType,attr"`
	WorkFax               string `xml:"WorkFax,attr"`
	WorkPhone             string `xml:"WorkPhone,attr"`
}
