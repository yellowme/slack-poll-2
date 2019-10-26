package core

import (
	"testing"

	"github.com/jerolan/slack-poll/core/entity"
	"github.com/jerolan/slack-poll/test"
	"github.com/stretchr/testify/assert"
)

func TestUpsertPollAnswer(t *testing.T) {
	question := "Test question?"
	options := []string{"Option 1", "Option 2"}
	expectedPollAnswer := entity.PollAnswer{Option: "Option 1"}
	pollService := test.TestPollService{}

	poll := CreatePoll(pollService, question, options)
	expectedPollAnswer.Poll = poll

	pollAnswer := UpsertPollAnswer(pollService, poll.ID, "Option 1")

	assert.Equal(t, pollAnswer, expectedPollAnswer)
}
