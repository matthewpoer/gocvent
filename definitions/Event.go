package gocvent

// EventRetrieveResult defines the result wrapper
type EventRetrieveResult struct {
	CvObject Event `xml:"RetrieveResult>CvObject"`
}

// Event defines the CvObject
type Event struct {
	ArchiveDate               string `xml:"ArchiveDate,attr"`
	Capacity                  string `xml:"Capacity,attr"`
	Category                  string `xml:"Category,attr"`
	City                      string `xml:"City,attr"`
	ClosedBy                  string `xml:"ClosedBy,attr"`
	Country                   string `xml:"Country,attr"`
	CountryCode               string `xml:"CountryCode,attr"`
	CreatedBy                 string `xml:"CreatedBy,attr"`
	Currency                  string `xml:"Currency,attr"`
	EventCalendarAlternateURL string `xml:"EventCalendarAlternateURL,attr"`
	EventCalendarCompletedURL string `xml:"EventCalendarCompletedURL,attr"`
	EventCalendarLinkText     string `xml:"EventCalendarLinkText,attr"`
	EventCode                 string `xml:"EventCode,attr"`
	EventDescription          string `xml:"EventDescription,attr"`
	EventEndDate              string `xml:"EventEndDate,attr"`
	EventLaunchDate           string `xml:"EventLaunchDate,attr"`
	EventStartDate            string `xml:"EventStartDate,attr"`
	EventStatus               string `xml:"EventStatus,attr"`
	EventTitle                string `xml:"EventTitle,attr"`
	ExternalAuthentication    bool   `xml:"ExternalAuthentication,attr"`
	Hidden                    bool   `xml:"Hidden,attr"`
	Id                        string `xml:"Id,attr"`
	InternalNote              string `xml:"InternalNote,attr"`
	LastModifiedDate          string `xml:"LastModifiedDate,attr"`
	Location                  string `xml:"Location,attr"`
	MeetingRequestId          string `xml:"MeetingRequestId,attr"`
	MerchantAccount           string `xml:"MerchantAccount,attr"`
	MerchantAccountId         string `xml:"MerchantAccountId,attr"`
	PhoneNumber               string `xml:"PhoneNumber,attr"`
	PlannerCompany            string `xml:"PlannerCompany,attr"`
	PlannerEmailAddress       string `xml:"PlannerEmailAddress,attr"`
	PlannerFirstName          string `xml:"PlannerFirstName,attr"`
	PlannerLastName           string `xml:"PlannerLastName,attr"`
	PlannerPrefix             string `xml:"PlannerPrefix,attr"`
	PlannerTitle              string `xml:"PlannerTitle,attr"`
	PlanningStatus            string `xml:"PlanningStatus,attr"`
	PostalCode                string `xml:"PostalCode,attr"`
	RSVPbyDate                string `xml:"RSVPbyDate,attr"`
	StakeholderAddress1       string `xml:"StakeholderAddress1,attr"`
	StakeholderAddress2       string `xml:"StakeholderAddress2,attr"`
	StakeholderAddress3       string `xml:"StakeholderAddress3,attr"`
	StakeholderCity           string `xml:"StakeholderCity,attr"`
	StakeholderCompany        string `xml:"StakeholderCompany,attr"`
	StakeholderCountryCode    string `xml:"StakeholderCountryCode,attr"`
	StakeholderEmailAddress   string `xml:"StakeholderEmailAddress,attr"`
	StakeholderFirstName      string `xml:"StakeholderFirstName,attr"`
	StakeholderHomePhone      string `xml:"StakeholderHomePhone,attr"`
	StakeholderLastName       string `xml:"StakeholderLastName,attr"`
	StakeholderMobilePhone    string `xml:"StakeholderMobilePhone,attr"`
	StakeholderPostalCode     string `xml:"StakeholderPostalCode,attr"`
	StakeholderStateCode      string `xml:"StakeholderStateCode,attr"`
	StakeholderTitle          string `xml:"StakeholderTitle,attr"`
	StakeholderWorkFax        string `xml:"StakeholderWorkFax,attr"`
	StakeholderWorkPhone      string `xml:"StakeholderWorkPhone,attr"`
	State                     string `xml:"State,attr"`
	StateCode                 string `xml:"StateCode,attr"`
	StreetAddress1            string `xml:"StreetAddress1,attr"`
	StreetAddress2            string `xml:"StreetAddress2,attr"`
	StreetAddress3            string `xml:"StreetAddress3,attr"`
	Timezone                  string `xml:"Timezone,attr"`
}
