package gocvent

import "github.com/matthewpoer/gocvent/gosoap"

// CventAPI the primary receiver, referencing a gosoap handler and
// authenticated session information
type CventAPI struct {
	ServerURL          string
	CventSessionHeader string
	soap               *gosoap.Client
}

// CvObject is a generic result of a Retrieve call, so this is mostly just metadata fields
type CvObject struct {
	Type             string `xml:"type,attr"`
	ID               string `xml:"Id,attr"`
	CreatedBy        string `xml:"CreatedBy,attr"`
	CreatedDate      string `xml:"CreatedDate,attr"`
	LastModifiedDate string `xml:"LastModifiedDate,attr"`
}

// AnswerDetail looks and smells a lot like LookUpDetail IMO
type AnswerDetail struct {
	AnswerText string `xml:"AnswerText,attr"`
}

// CustomField is a lot like Field, but it's custom, so it has some additional
// attributes and may not exist at all
type CustomField struct {
	ID            string         `xml:"Id,attr"`
	Name          string         `xml:"Name,attr"`
	Category      string         `xml:"Category,attr"`
	FieldType     string         `xml:"FieldType,attr"`
	Format        string         `xml:"Format,attr"`
	SortOrder     int            `xml:"SortOrder,attr"`
	AnswerDetails []AnswerDetail `xml:"AnswerDetail"`
}

// DescribeCvObjectResponse part of the DescribeCvObject API Calls
type DescribeCvObjectResponse struct {
	DescribeCvObjectResults []DescribeCvObjectResult `xml:"DescribeCvObjectResult>DescribeCvObjectResult"`
}

// DescribeCvObjectResult part of the DescribeCvObject API Calls
type DescribeCvObjectResult struct {
	Name          string        `xml:"Name,attr"`
	Creatable     bool          `xml:"Creatable,attr"`
	Updateable    bool          `xml:"Updateable,attr"`
	Deletable     bool          `xml:"Deletable,attr"`
	Replicateable bool          `xml:"Replicateable,attr"`
	Retrieveable  bool          `xml:"Retrieveable,attr"`
	Searchable    bool          `xml:"Searchable,attr"`
	Fields        []Field       `xml:"Field"`
	CustomFields  []CustomField `xml:"CustomField"`
}

// DescribeGlobalResponse part of the DescribeGlobal API calls
type DescribeGlobalResponse struct {
	DescribeGlobalResult DescribeGlobalResult
}

// DescribeGlobalResult part of the DescribeGlobal API calls
type DescribeGlobalResult struct {
	CurrentAPICalls int      `xml:"CurrentAPICalls,attr"`
	CvObjectTypes   []string `xml:"CvObjectTypes"`
	MaxAPICalls     int      `xml:"MaxAPICalls,attr"`
	MaxBatchSize    int      `xml:"MaxBatchSize,attr"`
	MaxRecordSet    int      `xml:"MaxRecordSet,attr"`
	/*
		Lookups not yet supported, e.g.
		```
		<LookUps>
			<LookUp Type="ContactCustomField">
				<Id>12345678-1EBA-4936-9CB2-E7E13C21F4E6</Id>
				<Name>Emergency Contact Name</Name>
				<Code></Code>
			</LookUp>
			<LookUp Type="ContactCustomField">
				<Id>12345678-674C-4E50-BF52-914E5B439E49</Id>
				<Name>Emergency Contact Phone</Name>
				<Code></Code>
			</LookUp>
		</LookUps>
		```
	*/
}

// Field information about a single Field on an object
type Field struct {
	Name               string         `xml:"Name,attr"`
	ObjectLocation     string         `xml:"ObjectLocation,attr"`
	DataType           string         `xml:"DataType,attr"`
	Searchable         bool           `xml:"Searchable,attr"`
	Required           bool           `xml:"Required,attr"`
	ReadOnly           bool           `xml:"ReadOnly,attr"`
	DefaultValue       string         `xml:"DefaultValue,attr"`
	DefaultSearchValue string         `xml:"DefaultSearchValue,attr"`
	LookUpDetails      []LookUpDetail `xml:"LookUpDetail"`
}

// Filter is used by Search API calls
type Filter struct {
	Field      string   `xml:"Field"`
	Operator   string   `xml:"Operator"`
	Value      string   `xml:"Value"`
	ValueArray []string `xml:"ValueArray"`
}

// LoginResponse wrapper for LoginResult
type LoginResponse struct {
	LoginResult LoginResult
}

// LoginResult contains information about the Login attempt
type LoginResult struct {
	LoginSuccess       string `xml:"LoginSuccess,attr"`
	CventSessionHeader string `xml:"CventSessionHeader,attr"`
	ServerURL          string `xml:"ServerURL,attr"`
	ErrorMessage       string `xml:"ErrorMessage,attr"`
}

// LookUpDetail holds values available in a list (i.e. a dropdown or picklist)
// as part of a Field definition
type LookUpDetail struct {
	Value string `xml:"Value,attr"`
}

// RetrieveResponse wrapper for RetrieveResult
type RetrieveResponse struct {
	RetrieveResult RetrieveResult
}

// RetrieveResult contains information about the records returned by a Retrieve API Call
type RetrieveResult struct {
	CvObject CvObject `xml:"CvObject"`
}

// SearchResponse wrapper for SearchResult
type SearchResponse struct {
	SearchResult SearchResult
}

// SearchResult contains information about the IDs returned by a Search API Call
type SearchResult struct {
	Ids []string `xml:"Id"`
}
