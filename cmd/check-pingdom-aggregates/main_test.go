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
	plugin.Critical = 0
	plugin.Warning = 0
	_, err := checkArgs(event)
	assert.Error(err)

	plugin.Critical = 1
	_, err = checkArgs(event)
	assert.Error(err)
	plugin.Warning = 2
	_, err = checkArgs(event)
	assert.Error(err)
	plugin.Critical = 4
	_, err = checkArgs(event)
	assert.NoError(err)
}
