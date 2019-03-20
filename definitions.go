package gocvent

import "github.com/tiaguinho/gosoap"

type CventAPI struct {
	ServerURL          string
	CventSessionHeader string
	soap               *gosoap.Client
}

type DescribeGlobalResponse struct {
	DescribeGlobalResult DescribeGlobalResult
}

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

type LoginResponse struct {
	LoginResult LoginResult
}

type LoginResult struct {
	LoginSuccess       string `xml:"LoginSuccess,attr"`
	CventSessionHeader string `xml:"CventSessionHeader,attr"`
	ServerURL          string `xml:"ServerURL,attr"`
	ErrorMessage       string `xml:"ErrorMessage,attr"`
}
