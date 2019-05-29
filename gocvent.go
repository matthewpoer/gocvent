package gocvent

import (
	"errors"
	"io/ioutil"
	"os/exec"
	"strings"

	"github.com/matthewpoer/gocvent/gosoap"
)

const wsdlSandbox string = "https://sandbox-api.cvent.com/soap/V200611.ASMX?WSDL"
const wsdlProduction string = "https://api.cvent.com/soap/V200611.ASMX?WSDL"
const xmlns string = "http://schemas.cvent.com/api/2006-11"

// Auth Cvent API Login
func (c *CventAPI) Auth(accountNumber string, user string, pass string) (bool, error) {
	soap, err := gosoap.SoapClient(wsdlProduction)
	if err != nil {
		return false, errors.New("CventAPI.Auth error loading WSDL: " + err.Error())
	}
	c.soap = soap
	params := gosoap.Params{
		"AccountNumber": accountNumber,
		"UserName":      user,
		"Password":      pass,
	}
	err = c.soap.Call("Login", params)
	if err != nil {
		return false, errors.New("CventAPI.Auth Soap Login Failure: " + err.Error())
	}

	var r LoginResponse
	err = c.soap.Unmarshal(&r)
	if err != nil {
		return false, errors.New("CventAPI.Auth received SOAP Fault: " + err.Error())
	}

	if r.LoginResult.LoginSuccess != "true" {
		return false, errors.New("CventAPI.Auth Login Failure: " + r.LoginResult.ErrorMessage)
	}

	// store the retrieved Server URL and Header, and go ahead and set the soap
	// object up for re-use by re-setting the URL and including the header
	c.ServerURL = r.LoginResult.ServerURL + "?WSDL"
	c.CventSessionHeader = r.LoginResult.CventSessionHeader
	c.soap, err = gosoap.SoapClient(c.ServerURL)
	if err != nil {
		return false, errors.New("CventAPI.Auth error loading revised WSDL: " + err.Error())
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
		return r.DescribeCvObjectResults, errors.New("CventAPI.DescribeCvObject Soap DescribeCvObject Failure: " + err.Error())
	}

	err = c.soap.Unmarshal(&r)
	if err != nil {
		return r.DescribeCvObjectResults, errors.New("CventAPI.DescribeCvObject received SOAP Fault: " + err.Error())
	}
	return r.DescribeCvObjectResults, nil
}

// DescribeGlobal get API settings for your account
func (c *CventAPI) DescribeGlobal() (DescribeGlobalResult, error) {
	var r DescribeGlobalResponse

	params := gosoap.Params{}
	err := c.soap.Call("DescribeGlobal", params)
	if err != nil {
		return r.DescribeGlobalResult, errors.New("CventAPI.DescribeGlobal Soap DescribeGlobal Failure: " + err.Error())
	}

	err = c.soap.Unmarshal(&r)
	if err != nil {
		return r.DescribeGlobalResult, errors.New("CventAPI.DescribeGlobal received SOAP Fault: " + err.Error())
	}

	return r.DescribeGlobalResult, nil
}

// Retrieve is used to get a single Cvent Object
func (c *CventAPI) Retrieve(ObjectType string, ID string, objectDef interface{}) error {

	params := gosoap.Params{}
	params["ObjectType"] = ObjectType
	params["Ids"] = ID

	err := c.soap.Call("Retrieve", params)
	if err != nil {
		return errors.New("CventAPI.Retrieve Soap Retrieve Failure: " + err.Error())
	}

	err = c.soap.Unmarshal(&objectDef)
	if err != nil {
		return errors.New("CventAPI.Retrieve received SOAP Fault: " + err.Error())
	}
	return nil
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
		return r.SearchResult, errors.New("CventAPI.Search Soap Search Failure: " + err.Error())
	}

	err = c.soap.Unmarshal(&r)
	if err != nil {
		return r.SearchResult, errors.New("CventAPI.Search received SOAP Fault: " + err.Error())
	}
	return r.SearchResult, nil
}

// StructGen is used to generate a Golang struct based on the fields in a CvObject
func (c *CventAPI) StructGen(filePath string, objectType string) error {

	// get fields and info. about the CvObject
	var objectList = make([]string, 1)
	objectList[0] = objectType
	r, err := c.DescribeCvObject(objectList)
	if err != nil {
		return err
	}
	foundObjectType := false
	for _, CvObjectMetadata := range r {
		if CvObjectMetadata.Name == objectType {
			foundObjectType = true
		}
	}
	if !foundObjectType {
		return errors.New("DescribeCvObject failure")
	}

	output := "package gocvent\n"
	output += "// " + objectType + "RetrieveResult defines the result wrapper\n"
	output += "type " + objectType + "RetrieveResult struct {\n"
	output += "CvObject " + objectType + " `xml:\"RetrieveResult>CvObject\"`\n"
	output += "}\n"
	output += "// " + objectType + " defines the CvObject\n"
	output += "type " + objectType + " struct {\n"

	for _, CvObjectMetadata := range r {
		if CvObjectMetadata.Name == objectType {
			for _, objectField := range CvObjectMetadata.Fields {

				// ignore anything that isn't on the object we're tracking
				if objectField.ObjectLocation != "/"+objectType {
					continue
				}

				dataType := "string"
				if objectField.DataType == "Boolean" {
					dataType = "bool"
				}
				if objectField.DataType == "Date Time" {
					dataType = "string" // @todo should be whatever golang likes to call datetime
				}

				xmlFieldName := strings.TrimSpace(objectField.Name)             // drop leading/trailing spaces
				structFieldname := strings.Replace(xmlFieldName, " ", "", -1)   // don't allow spaces in field name
				structFieldname = strings.Replace(structFieldname, "(", "", -1) // don't allow parens in field name
				structFieldname = strings.Replace(structFieldname, ")", "", -1) // don't allow parens in field name

				output += "\t" + structFieldname + " " +
					dataType +
					" `xml:\"" + xmlFieldName + ",attr\"`  " +
					"\n"
			}
		}
	}

	output += "}\n"

	fileName := filePath + "/" + objectType + ".go"
	data := []byte(output)
	err = ioutil.WriteFile(fileName, data, 0644)
	if err != nil {
		return err
	}

	// run gofmt on the new definitions file
	cmd := exec.Command("gofmt", "-w", fileName)
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
