package controller

import (
	"testing"

	"github.com/jerolan/slack-poll/core/entity"
	"github.com/jerolan/slack-poll/test"
	"github.com/stretchr/testify/assert"
)

func TestPollController(t *testing.T) {
	test.InitializeTestDatabase()

	pollService := &test.TestPollService{}

	command := "/poll “Question” “Option 1” “Option 2”"
	var poll entity.Poll
	CreatePoll(pollService, command, &poll)

	assert.Equal(t, "Question", poll.Question)
}
