package gocvent

import (
	"os"
	"testing"

	objDefs "github.com/matthewpoer/gocvent/definitions"
	log "github.com/sirupsen/logrus"

	"github.com/stretchr/testify/assert"
)

func genericAuth() (CventAPI, bool, error) {
	var cvent CventAPI
	success, err := cvent.Auth(
		os.Getenv("CVENT_ACCOUNT_NUMBER"),
		os.Getenv("CVENT_USERNAME"),
		os.Getenv("CVENT_PASSWORD"),
	)
	return cvent, success, err
}

func TestAuth(t *testing.T) {
	cvent, success, err := genericAuth()
	if err != nil {
		log.Printf("TestAuth err from genericAuth(): %s", err)
	}
	assert.Nil(t, err)

	if success != true {
		log.Println("TestAuth success is false")
	}
	assert.True(t, success)

	assert.NotEmpty(t, cvent.ServerURL)
	assert.NotEmpty(t, cvent.CventSessionHeader)
}

func TestDescribeCvObjectMultiple(t *testing.T) {
	cvent, success, err := genericAuth()
	if err != nil || !success {
		t.Errorf("TestDescribeCvObjectMultiple fails because authorization or invocation failed")
		return
	}

	var objectList = make([]string, 3)
	objectList[0] = "Contact"
	objectList[1] = "Event"
	objectList[2] = "User"

	r, err := cvent.DescribeCvObject(objectList)
	if err != nil {
		log.Printf("TestDescribeCvObjectMultiple err from cvent.DescribeCvObject: %s", err)
	}
	assert.Nil(t, err)

	// make sure that we found a DescribeCvObjectResult to represent each of our elements.
	foundContact := false
	foundEvent := false
	foundUser := false
	for _, CvObjectMetadata := range r {
		if CvObjectMetadata.Name == "Contact" {
			foundContact = true
		} else if CvObjectMetadata.Name == "Event" {
			foundEvent = true
		} else if CvObjectMetadata.Name == "User" {
			foundUser = true
		}
	}
	assert.True(t, foundContact, "Could not find Contact Object information")
	assert.True(t, foundEvent, "Could not find Event Object information")
	assert.True(t, foundUser, "Could not find User Object information")

	// check that we have relevant information about the out-of-box Contact field "Company"
	foundCompanyField := false
	for _, CvObjectMetadata := range r {
		if CvObjectMetadata.Name == "Contact" {
			for _, ContactField := range CvObjectMetadata.Fields {
				if ContactField.Name == "Company" {
					foundCompanyField = true
				}
			}
		}
	}
	assert.True(t, foundCompanyField, "Could not find Company field on the Contact Object")

	// look for at least a single CustomField on any of the objects... this
	// will break if the Cvent system has zero custom fields
	foundAnyCustomField := false
	for _, CvObjectMetadata := range r {
		for _, CustomField := range CvObjectMetadata.CustomFields {
			if CustomField.Name != "" {
				foundAnyCustomField = true
			}
		}
	}
	assert.True(t, foundAnyCustomField, "Could not find any Custom Fields")
}

func TestDescribeCvObjectSingle(t *testing.T) {
	cvent, success, err := genericAuth()
	if err != nil || !success {
		t.Errorf("TestDescribeCvObjectSingle fails because authorization or invocation failed")
		return
	}

	var objectList = make([]string, 1)
	objectList[0] = "Contact"

	r, err := cvent.DescribeCvObject(objectList)
	if err != nil {
		log.Printf("TestDescribeCvObjectMultiple err from cvent.DescribeCvObject: %s", err)
	}
	assert.Nil(t, err)

	// make sure that we found a DescribeCvObjectResult to represent each of our elements.
	foundContact := false
	for _, CvObjectMetadata := range r {
		if CvObjectMetadata.Name == "Contact" {
			foundContact = true
		}
	}
	assert.True(t, foundContact, "Could not find Contact Object information")

	// check that we have relevant information about the out-of-box Contact field "Company"
	foundCompanyField := false
	for _, CvObjectMetadata := range r {
		if CvObjectMetadata.Name == "Contact" {
			for _, ContactField := range CvObjectMetadata.Fields {
				if ContactField.Name == "Company" {
					foundCompanyField = true
				}
			}
		}
	}
	assert.True(t, foundCompanyField, "Could not find Company field on the Contact Object")
}

func TestDescribeCvObjectSingleBadObject(t *testing.T) {
	cvent, success, err := genericAuth()
	if err != nil || !success {
		t.Errorf("TestDescribeCvObjectSingleBadObject fails because authorization or invocation failed")
		return
	}

	var objectList = make([]string, 1)
	objectList[0] = "SomeFakeObjectName"

	r, err := cvent.DescribeCvObject(objectList)
	// skip the conventional err != nil check. This test is expecting a failure.
	assert.NotEmpty(t, err)
	assert.Zero(t, len(r))
}
func TestDescribeGlobal(t *testing.T) {
	cvent, success, err := genericAuth()
	if err != nil || !success {
		t.Errorf("TestDescribeGlobal fails because authorization or invocation failed")
		return
	}

	r, err := cvent.DescribeGlobal()
	if err != nil {
		log.Printf("TestDescribeGlobal err from cvent.DescribeGlobal: %s", err)
	}
	assert.Nil(t, err)
	assert.NotEmpty(t, r.CurrentAPICalls)
	assert.NotEmpty(t, r.CvObjectTypes)
	assert.NotEmpty(t, r.MaxAPICalls)
	assert.NotEmpty(t, r.MaxBatchSize)
	assert.NotEmpty(t, r.MaxRecordSet)
}

func TestRetrieveEvent(t *testing.T) {
	cvent, success, err := genericAuth()
	if err != nil || !success {
		t.Errorf("TestRetrieveEvent fails because authorization or invocation failed")
		return
	}

	// get a list of Events, there should always be a few to work with
	searchRes, err := cvent.Search("Event", []Filter{})
	if err != nil {
		log.Printf("TestRetrieveEvent err from cvent.Search: %s", err)
	}
	assert.Nil(t, err)
	assert.Greater(t, len(searchRes.Ids), 0)

	// No need to test against all events, 3 should be plenty
	if len(searchRes.Ids) > 3 {
		searchRes.Ids = searchRes.Ids[:3]
	}

	// attempt to retrieve basic data about each User
	for _, v := range searchRes.Ids {
		var objectDef objDefs.EventRetrieveResult
		err := cvent.Retrieve("Event", v, &objectDef)
		if err != nil {
			log.Printf("TestRetrieveEvent err from cvent.Retrieve: %s", err)
		}

		assert.Nil(t, err)
		assert.NotNil(t, objectDef.CvObject.Capacity)
		assert.NotEmpty(t, objectDef.CvObject.Capacity)
		assert.NotNil(t, objectDef.CvObject.Currency)
		assert.NotEmpty(t, objectDef.CvObject.Currency)
		assert.NotNil(t, objectDef.CvObject.EventStatus)
		assert.NotEmpty(t, objectDef.CvObject.EventStatus)
		assert.NotNil(t, objectDef.CvObject.EventTitle)
		assert.NotEmpty(t, objectDef.CvObject.EventTitle)
		assert.NotNil(t, objectDef.CvObject.Id)
		assert.NotEmpty(t, objectDef.CvObject.Id)
	}

}

func TestRetrieveUser(t *testing.T) {
	cvent, success, err := genericAuth()
	if err != nil || !success {
		t.Errorf("TestRetrieveUser fails because authorization or invocation failed")
		return
	}

	// get a list of users, there should always be a few users to work with
	searchRes, err := cvent.Search("User", []Filter{})
	if err != nil {
		log.Printf("TestRetrieveUser err from cvent.Search: %s", err)
	}
	assert.Nil(t, err)
	assert.Greater(t, len(searchRes.Ids), 0)

	// No need to test against all users, 3 should be plenty
	if len(searchRes.Ids) > 3 {
		searchRes.Ids = searchRes.Ids[:3]
	}

	// attempt to retrieve basic data about each User
	for _, v := range searchRes.Ids {
		var objectDef objDefs.UserRetrieveResult
		err := cvent.Retrieve("User", v, &objectDef)
		if err != nil {
			log.Printf("TestRetrieveUser err from cvent.Retrieve: %s", err)
		}
		assert.Nil(t, err)
		assert.NotNil(t, objectDef.CvObject.Email)
		assert.NotEmpty(t, objectDef.CvObject.Email)
		assert.NotNil(t, objectDef.CvObject.Id)
		assert.NotEmpty(t, objectDef.CvObject.Id)
		assert.NotNil(t, objectDef.CvObject.LastName)
		assert.NotEmpty(t, objectDef.CvObject.LastName)
	}

}

func TestSearchNoFilter(t *testing.T) {
	cvent, success, err := genericAuth()
	if err != nil || !success {
		t.Errorf("TestSearchNoFilter fails because authorization or invocation failed")
		return
	}

	r, err := cvent.Search("Contact", []Filter{})
	if err != nil {
		log.Printf("TestSearchNoFilter err from cvent.Search: %s", err)
	}
	assert.Nil(t, err)
	assert.Greater(t, len(r.Ids), 0)
}

func TestSearchWithFilters(t *testing.T) {
	cvent, success, err := genericAuth()
	if err != nil || !success {
		t.Errorf("TestSearchWithFilters fails because authorization or invocation failed")
		return
	}

	Filters := make([]Filter, 2)
	Filters[0] = Filter{
		Field:    "Company",
		Operator: "Not Equal to",
		Value:    "Some Junk Value",
	}
	Filters[1] = Filter{
		Field:    "LastName",
		Operator: "Equals",
		Value:    "Smith",
	}

	r, err := cvent.Search("Contact", Filters)
	if err != nil {
		log.Printf("TestSearchWithFilters err from cvent.Search: %s", err)
	}
	assert.Nil(t, err)
	assert.Greater(t, len(r.Ids), 0)

	numberOfSmiths := len(r.Ids)

	Filters = make([]Filter, 1)
	Filters[0] = Filter{
		Field:    "LastName",
		Operator: "Includes",
	}
	Filters[0].ValueArray = append(Filters[0].ValueArray, "Smith")
	Filters[0].ValueArray = append(Filters[0].ValueArray, "Johnson")
	Filters[0].ValueArray = append(Filters[0].ValueArray, "Williams")
	Filters[0].ValueArray = append(Filters[0].ValueArray, "Jones")
	r, err = cvent.Search("Contact", Filters)
	if err != nil {
		log.Printf("TestSearchWithFilters err from cvent.Search: %s", err)
	}
	assert.Nil(t, err)
	assert.Greater(t, len(r.Ids), 0)
	assert.Greater(t, len(r.Ids), numberOfSmiths)
}

func TestStructGen(t *testing.T) {
	cvent, success, err := genericAuth()
	if err != nil || !success {
		t.Errorf("TestStructGen fails because authorization or invocation failed")
		return
	}
	assert.Nil(t, err)

	err = cvent.StructGen("definitions", "Contact")
	assert.Nil(t, err)

	err = cvent.StructGen("definitions", "Event")
	assert.Nil(t, err)

	err = cvent.StructGen("definitions", "Registration")
	assert.Nil(t, err)

	err = cvent.StructGen("definitions", "User")
	assert.Nil(t, err)
}
