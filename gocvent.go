package gocvent

import (
	"errors"
	"log"

	"github.com/tiaguinho/gosoap"
)

const wsdlSandbox string = "https://sandbox-api.cvent.com/soap/V200611.ASMX?WSDL"
const wsdlProduction string = "https://api.cvent.com/soap/V200611.ASMX?WSDL"

type LoginResponse struct {
	LoginResult LoginResult
}

type LoginResult struct {
	LoginSuccess       string `xml:"LoginSuccess,attr"`
	CventSessionHeader string `xml:"CventSessionHeader,attr"`
	ServerURL          string `xml:"ServerURL,attr"`
	ErrorMessage       string `xml:"ErrorMessage,attr"`
}

var (
	r LoginResponse
)

// Auth Cvent API Login
func Auth(accountNumber string, user string, pass string) (string, string, error) {
	soap, err := gosoap.SoapClient(wsdlProduction)
	if err != nil {
		log.Printf("error not expected on soap invocation: %s", err)
		return "", "", errors.New("WSDL Load Failure")
	}
	params := gosoap.Params{
		"AccountNumber": accountNumber,
		"UserName":      user,
		"Password":      pass,
	}
	err = soap.Call("Login", params)
	if err != nil {
		log.Printf("error not expected on cvent call: %s", err)
		return "", "", errors.New("SOAP Call Failure")
	}

	soap.Unmarshal(&r)
	if r.LoginResult.LoginSuccess != "true" {
		log.Printf("login was not successful?: %s", r.LoginResult.LoginSuccess)
		return "", "", errors.New("Login Failure")
	}

	return r.LoginResult.ServerURL, r.LoginResult.CventSessionHeader, nil
}
