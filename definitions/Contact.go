package gocvent

// ContactRetrieveResult defines the result wrapper
type ContactRetrieveResult struct {
	CvObject Contact `xml:"RetrieveResult>CvObject"`
}

// Contact defines the CvObject
type Contact struct {
	Active                       bool   `xml:"Active,attr"`
	ActivityId                   string `xml:"ActivityId,attr"`
	CCEmailAddress               string `xml:"CCEmailAddress,attr"`
	Company                      string `xml:"Company,attr"`
	ContactType                  string `xml:"ContactType,attr"`
	ContactTypeCode              string `xml:"ContactTypeCode,attr"`
	CreatedBy                    string `xml:"CreatedBy,attr"`
	CreatedDate                  string `xml:"CreatedDate,attr"`
	DateOfBirth                  string `xml:"DateOfBirth,attr"`
	Designation                  string `xml:"Designation,attr"`
	EmailAddress                 string `xml:"EmailAddress,attr"`
	EmailAddressStatus           string `xml:"EmailAddressStatus,attr"`
	ExcludedFromEmail            bool   `xml:"ExcludedFromEmail,attr"`
	ExpirationDate               string `xml:"ExpirationDate,attr"`
	FacebookURL                  string `xml:"FacebookURL,attr"`
	FirstName                    string `xml:"FirstName,attr"`
	Gender                       string `xml:"Gender,attr"`
	HomeAddress1                 string `xml:"HomeAddress1,attr"`
	HomeAddress2                 string `xml:"HomeAddress2,attr"`
	HomeAddress3                 string `xml:"HomeAddress3,attr"`
	HomeCity                     string `xml:"HomeCity,attr"`
	HomeCountry                  string `xml:"HomeCountry,attr"`
	HomeCountryCode              string `xml:"HomeCountryCode,attr"`
	HomeFax                      string `xml:"HomeFax,attr"`
	HomePhone                    string `xml:"HomePhone,attr"`
	HomePostalCode               string `xml:"HomePostalCode,attr"`
	HomeState                    string `xml:"HomeState,attr"`
	HomeStateCode                string `xml:"HomeStateCode,attr"`
	Id                           string `xml:"Id,attr"`
	ImageURL                     string `xml:"ImageURL,attr"`
	IsCreatedViaTestReg          bool   `xml:"IsCreatedViaTestReg,attr"`
	IsObfuscated                 bool   `xml:"IsObfuscated,attr"`
	JoinDate                     string `xml:"JoinDate,attr"`
	LastModifiedBy               string `xml:"LastModifiedBy,attr"`
	LastModifiedDate             string `xml:"LastModifiedDate,attr"`
	LastName                     string `xml:"LastName,attr"`
	LastOptOutBy                 string `xml:"LastOptOutBy,attr"`
	LastOptOutDate               string `xml:"LastOptOutDate,attr"`
	LastRenewalDate              string `xml:"LastRenewalDate,attr"`
	LinkedInURL                  string `xml:"LinkedInURL,attr"`
	LogDate                      string `xml:"LogDate,attr"`
	LogReason                    string `xml:"LogReason,attr"`
	LogResponse                  string `xml:"LogResponse,attr"`
	MembershipCode               string `xml:"MembershipCode,attr"`
	MiddleName                   string `xml:"MiddleName,attr"`
	MobilePhone                  string `xml:"MobilePhone,attr"`
	NationalIdentificationNumber string `xml:"NationalIdentificationNumber,attr"`
	Nickname                     string `xml:"Nickname,attr"`
	OptedIn                      bool   `xml:"OptedIn,attr"`
	Pager                        string `xml:"Pager,attr"`
	ParentContactId              string `xml:"ParentContactId,attr"`
	PassportCountry              string `xml:"PassportCountry,attr"`
	PassportCountryCode          string `xml:"PassportCountryCode,attr"`
	PassportNumber               string `xml:"PassportNumber,attr"`
	PrimaryAddress               string `xml:"PrimaryAddress,attr"`
	Salutation                   string `xml:"Salutation,attr"`
	SMTPCode                     string `xml:"SMTPCode,attr"`
	SocialSecurityNumber         string `xml:"SocialSecurityNumber,attr"`
	SourceId                     string `xml:"SourceId,attr"`
	Title                        string `xml:"Title,attr"`
	TwitterURL                   string `xml:"TwitterURL,attr"`
	WorkAddress1                 string `xml:"WorkAddress1,attr"`
	WorkAddress2                 string `xml:"WorkAddress2,attr"`
	WorkAddress3                 string `xml:"WorkAddress3,attr"`
	WorkCity                     string `xml:"WorkCity,attr"`
	WorkCountry                  string `xml:"WorkCountry,attr"`
	WorkCountryCode              string `xml:"WorkCountryCode,attr"`
	WorkFax                      string `xml:"WorkFax,attr"`
	WorkPhone                    string `xml:"WorkPhone,attr"`
	WorkPostalCode               string `xml:"WorkPostalCode,attr"`
	WorkState                    string `xml:"WorkState,attr"`
	WorkStateCode                string `xml:"WorkStateCode,attr"`
}
