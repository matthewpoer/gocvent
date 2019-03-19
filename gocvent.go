package gocvent

import (
	"errors"
	"log"

	"github.com/tiaguinho/gosoap"
)

const wsdlSandbox string = "https://sandbox-api.cvent.com/soap/V200611.ASMX?WSDL"
const wsdlProduction string = "https://api.cvent.com/soap/V200611.ASMX?WSDL"

type CventAPI struct {
	ServerURL          string
	CventSessionHeader string
	soap               *gosoap.Client
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

	c.ServerURL = r.LoginResult.ServerURL + "?WSDL"
	c.CventSessionHeader = r.LoginResult.CventSessionHeader
	c.soap, err = gosoap.SoapClient(c.ServerURL)
	if err != nil {
		log.Printf("error not expected on soap re-invocation: %s", err)
		return false, errors.New("New WSDL Load Failure")
	}
	return true, nil
}
