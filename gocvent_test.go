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
