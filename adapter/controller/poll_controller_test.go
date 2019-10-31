package controller

import (
	"testing"

	"github.com/jerolan/slack-poll/test"
	"github.com/stretchr/testify/assert"
)

func TestPollController(t *testing.T) {
	test.InitializeTestDatabase()

	pollService := &test.TestPollService{}
	uuid := &test.UUID{}

	command := "/poll “Question” “Option 1” “Option 2”"
	owner := "Test User"

	poll := CreatePoll(pollService, uuid, owner, command)

	assert.Equal(t, "Question", poll.Question)
}
