package gocvent

// RegistrationRetrieveResult defines the result wrapper
type RegistrationRetrieveResult struct {
	CvObject Registration `xml:"RetrieveResult>CvObject"`
}

// Registration defines the CvObject
type Registration struct {
	CancelledDate        string `xml:"CancelledDate,attr"`
	CCEmailAddress       string `xml:"CCEmailAddress,attr"`
	Company              string `xml:"Company,attr"`
	ConfirmationNumber   string `xml:"ConfirmationNumber,attr"`
	ContactId            string `xml:"ContactId,attr"`
	Credit               string `xml:"Credit,attr"`
	EmailAddress         string `xml:"EmailAddress,attr"`
	EventCode            string `xml:"EventCode,attr"`
	EventId              string `xml:"EventId,attr"`
	EventStartDate       string `xml:"EventStartDate,attr"`
	EventTitle           string `xml:"EventTitle,attr"`
	FirstName            string `xml:"FirstName,attr"`
	GroupId              string `xml:"GroupId,attr"`
	GroupLeader          bool   `xml:"GroupLeader,attr"`
	Id                   string `xml:"Id,attr"`
	InternalNote         string `xml:"InternalNote,attr"`
	InvitedBy            string `xml:"InvitedBy,attr"`
	InviteeId            string `xml:"InviteeId,attr"`
	IsTestRegistrant     bool   `xml:"IsTestRegistrant,attr"`
	LastModifiedDate     string `xml:"LastModifiedDate,attr"`
	LastName             string `xml:"LastName,attr"`
	ModifiedBy           string `xml:"ModifiedBy,attr"`
	OriginalResponseDate string `xml:"OriginalResponseDate,attr"`
	Participant          bool   `xml:"Participant,attr"`
	ReferenceId          string `xml:"ReferenceId,attr"`
	RegistrationDate     string `xml:"RegistrationDate,attr"`
	RegistrationType     string `xml:"RegistrationType,attr"`
	RegistrationTypeCode string `xml:"RegistrationTypeCode,attr"`
	ResponseMethod       string `xml:"ResponseMethod,attr"`
	SourceId             string `xml:"SourceId,attr"`
	Status               string `xml:"Status,attr"`
	TargetedListId       string `xml:"TargetedListId,attr"`
	TargetedListName     string `xml:"TargetedListName,attr"`
	Title                string `xml:"Title,attr"`
	WorkPhone            string `xml:"WorkPhone,attr"`
}
