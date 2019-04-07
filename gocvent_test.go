package gocvent

import (
	"os"
	"testing"

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
	assert.Nil(t, err)
	assert.True(t, success)
	assert.NotEmpty(t, cvent.ServerURL)
	assert.NotEmpty(t, cvent.CventSessionHeader)
}

func TestDescribeCvObjectMultiple(t *testing.T) {
	var objectList = make([]string, 3)
	objectList[0] = "Contact"
	objectList[1] = "Event"
	objectList[2] = "User"
	cvent, _, _ := genericAuth()
	r, err := cvent.DescribeCvObject(objectList)
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
	var objectList = make([]string, 1)
	objectList[0] = "Contact"
	cvent, _, _ := genericAuth()
	r, err := cvent.DescribeCvObject(objectList)
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

func TestDescribeGlobal(t *testing.T) {
	cvent, _, _ := genericAuth()
	r, err := cvent.DescribeGlobal()
	assert.Nil(t, err)
	assert.NotEmpty(t, r.CurrentAPICalls)
	assert.NotEmpty(t, r.CvObjectTypes)
	assert.NotEmpty(t, r.MaxAPICalls)
	assert.NotEmpty(t, r.MaxBatchSize)
	assert.NotEmpty(t, r.MaxRecordSet)
}
