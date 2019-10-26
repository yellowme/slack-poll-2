package core

import (
	"testing"

	"github.com/jerolan/slack-poll/test"
	"github.com/stretchr/testify/assert"
)

func TestGetPolls(t *testing.T) {
	question := "Test question?"
	options := []string{"Option 1", "Option 2"}
	pollService := test.TestPollService{}

	createdPoll := CreatePoll(pollService, question, options)
	poll := GetPolls(pollService, createdPoll.ID)

	assert.Equal(t, poll, createdPoll)
}
