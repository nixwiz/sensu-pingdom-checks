package main

import (
	"testing"

	corev2 "github.com/sensu/sensu-go/api/core/v2"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
}

func TestCharArgs(t *testing.T) {
	assert := assert.New(t)
	event := corev2.FixtureEvent("entity1", "check1")
	// Defaults
	plugin.CriticalAvailableSMS = -1
	plugin.WarningAvailableSMS = -1
	plugin.CriticalAvailableChecks = -1
	plugin.WarningAvailableChecks = -1
	_, err := checkArgs(event)
	assert.Error(err)

	plugin.CriticalAvailableSMS = 5
	_, err = checkArgs(event)
	assert.Error(err)
	plugin.WarningAvailableSMS = 2
	_, err = checkArgs(event)
	assert.Error(err)
	plugin.WarningAvailableSMS = 10
	_, err = checkArgs(event)
	assert.Error(err)
	plugin.CriticalAvailableChecks = 5
	_, err = checkArgs(event)
	assert.Error(err)
	plugin.WarningAvailableChecks = 2
	_, err = checkArgs(event)
	assert.Error(err)
	plugin.WarningAvailableChecks = 10
	_, err = checkArgs(event)
	assert.NoError(err)
}
