package gocvent

import (
	"errors"
	"log"

	"github.com/matthewpoer/gocvent/gosoap"
)

const wsdlSandbox string = "https://sandbox-api.cvent.com/soap/V200611.ASMX?WSDL"
const wsdlProduction string = "https://api.cvent.com/soap/V200611.ASMX?WSDL"
const xmlns string = "http://schemas.cvent.com/api/2006-11"

// Auth Cvent API Login
func (c *CventAPI) Auth(accountNumber string, user string, pass string) (bool, error) {
	soap, err := gosoap.SoapClient(wsdlProduction)
	if err != nil {
		log.Printf("error not expected on soap invocation: %s", err)
		return false, errors.New("WSDL Load Failure")
	}
	c.soap = soap
	params := gosoap.Params{
		"AccountNumber": accountNumber,
		"UserName":      user,
		"Password":      pass,
	}
	err = c.soap.Call("Login", params)
	if err != nil {
		log.Printf("error not expected on cvent Login: %s", err)
		return false, errors.New("SOAP Call Failure")
	}

	var r LoginResponse
	c.soap.Unmarshal(&r)
	if r.LoginResult.LoginSuccess != "true" {
		log.Printf("login was not successful?: %s", r.LoginResult.LoginSuccess)
		return false, errors.New("Login Failure")
	}

	// store the retrieved Server URL and Header, and go ahead and set the soap
	// object up for re-use by re-setting the URL and including the header
	c.ServerURL = r.LoginResult.ServerURL + "?WSDL"
	c.CventSessionHeader = r.LoginResult.CventSessionHeader
	c.soap, err = gosoap.SoapClient(c.ServerURL)
	if err != nil {
		log.Printf("error not expected on soap re-invocation: %s", err)
		return false, errors.New("New WSDL Load Failure")
	}
	c.soap.HeaderName = "CventSessionHeader"
	c.soap.HeaderParams = make(map[string]string)
	c.soap.HeaderParams["CventSessionValue"] = c.CventSessionHeader
	return true, nil
}

// DescribeCvObject get information about one or more Cvent objects (e.g. Event, Contact)
func (c *CventAPI) DescribeCvObject(objectTypes []string) ([]DescribeCvObjectResult, error) {
	var r DescribeCvObjectResponse

	params := gosoap.Params{}
	ObjectTypes := make(map[string][]string)
	ObjectTypes["CvObjectType"] = objectTypes
	params["ObjectTypes"] = ObjectTypes
	err := c.soap.Call("DescribeCvObject", params)
	if err != nil {
		log.Printf("error not expected on cvent DescribeCvObject: %s", err)
		return r.DescribeCvObjectResults, errors.New("SOAP Call Failure")
	}

	c.soap.Unmarshal(&r)
	return r.DescribeCvObjectResults, nil
}

// DescribeGlobal get API settings for your account
func (c *CventAPI) DescribeGlobal() (DescribeGlobalResult, error) {
	var r DescribeGlobalResponse

	params := gosoap.Params{}
	err := c.soap.Call("DescribeGlobal", params)
	if err != nil {
		log.Printf("error not expected on cvent DescribeGlobal: %s", err)
		return r.DescribeGlobalResult, errors.New("SOAP Call Failure")
	}

	c.soap.Unmarshal(&r)
	return r.DescribeGlobalResult, nil
}

// Search is used to Search any Cvent Object using an optional set of filters
func (c *CventAPI) Search(ObjectType string, Filters []Filter) (SearchResult, error) {
	var r SearchResponse

	params := gosoap.Params{}
	params["ObjectType"] = ObjectType
	CvSearchObject := make(map[string][]Filter)
	CvSearchObject["Filter"] = Filters
	params["CvSearchObject"] = CvSearchObject

	err := c.soap.Call("Search", params)
	if err != nil {
		log.Printf("error not expected on cvent DescribeCvObject: %s", err)
		return r.SearchResult, errors.New("SOAP Call Failure")
	}

	c.soap.Unmarshal(&r)
	return r.SearchResult, nil
}
