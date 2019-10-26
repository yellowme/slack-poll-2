package core

import (
	"testing"

	"github.com/jerolan/slack-poll/core/entity"
	"github.com/jerolan/slack-poll/test"
	"github.com/stretchr/testify/assert"
)

func TestCreatePoll(t *testing.T) {
	question := "Test question?"
	options := []string{"Option 1", "Option 2"}
	pollRepository := test.TestUserRepository{}

	poll := CreatePoll(pollRepository, question, options)
	expectedPoll := entity.Poll{Question: question, Options: options}

	assert.Equal(t, poll, expectedPoll)
}
