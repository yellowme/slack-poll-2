package core

import (
	"testing"

	"github.com/jerolan/slack-poll/test"
	"github.com/stretchr/testify/assert"
)

func TestGetPolls(t *testing.T) {
	question := "Test question?"
	options := []string{"Option 1", "Option 2"}
	pollRepository := test.TestUserRepository{}

	createdPoll := CreatePoll(pollRepository, question, options)
	poll := GetPolls(pollRepository, createdPoll.ID)

	assert.Equal(t, poll, createdPoll)
}
