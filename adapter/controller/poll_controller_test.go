package controller

import (
	"testing"

	"github.com/jerolan/slack-poll/test"
	"github.com/jerolan/slack-poll/usecase"
	"github.com/stretchr/testify/assert"
)

func TestPollController(t *testing.T) {
	db := test.InitializeTestDatabase()

	uuid := test.NewTestUUIDPort()
	pollService := test.NewTestPollService(db)
	useCase := usecase.NewUseCase(pollService, uuid)
	pollController := NewPollController(useCase)

	command := "/poll “Question” “Option 1” “Option 2”"
	owner := "Test User"

	poll := pollController.CreatePoll(owner, command)

	assert.Equal(t, "Question", poll.Question)
}
