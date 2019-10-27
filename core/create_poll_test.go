package core

import (
	"testing"

	"github.com/jerolan/slack-poll/core/entity"
	"github.com/jerolan/slack-poll/test"
	"github.com/stretchr/testify/assert"
)

func TestCreatePoll(t *testing.T) {
	test.InitializeTestDatabase()

	pollService := &test.TestPollService{}
	uuid := &test.UUID{}

	id := uuid.Generate()
	owner := "Test Owner"
	question := "Test question?"
	options := []string{"Option 1", "Option 2"}

	poll := entity.Poll{
		ID:       id,
		Question: question,
		Owner:    owner,
		Options:  options,
	}

	CreatePoll(pollService, &poll)

	assert.Equal(t, poll.Question, question)
	assert.Equal(t, poll.ID, id)
	assert.Equal(t, poll.Mode, entity.PollModeSingle)
}
