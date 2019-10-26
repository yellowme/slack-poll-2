package core

import (
	"testing"

	"github.com/jerolan/slack-poll/test"
	"github.com/stretchr/testify/assert"
)

func TestDeletePoll(t *testing.T) {
	question := "Test question?"
	options := []string{"Option 1", "Option 2"}
	pollService := test.TestPollService{}

	poll := CreatePoll(pollService, question, options)
	expectedPoll := DeletePoll(pollService, poll.ID)
	emptyPolls := GetPolls(pollService, poll.ID)

	assert.Equal(t, poll, expectedPoll)
	assert.Empty(t, emptyPolls)
}
