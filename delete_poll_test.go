package core

import (
	"testing"

	"github.com/jerolan/slack-poll/test"
	"github.com/stretchr/testify/assert"
)

func TestDeletePoll(t *testing.T) {
	question := "Test question?"
	options := []string{"Option 1", "Option 2"}
	pollRepository := test.TestUserRepository{}

	poll := CreatePoll(pollRepository, question, options)
	expectedPoll := DeletePoll(pollRepository, poll.ID)
	emptyPolls := GetPolls(pollRepository, poll.ID)

	assert.Equal(t, poll, expectedPoll)
	assert.Empty(t, emptyPolls)
}
