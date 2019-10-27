package core

import (
	"testing"

	"github.com/jerolan/slack-poll/core/entity"
	"github.com/jerolan/slack-poll/test"
	"github.com/stretchr/testify/assert"
)

func TestUpsertPollAnswer(t *testing.T) {
	test.InitializeTestDatabase()

	t.Run("ok", func(t *testing.T) {
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

		pollAnswer := entity.PollAnswer{
			ID:     id,
			Option: options[0],
			Owner:  owner,
			PollID: poll.ID,
		}

		CreatePoll(pollService, &poll)
		UpsertPollAnswer(pollService, &pollAnswer)

		assert.Equal(t, poll.ID, pollAnswer.PollID)
	})

	/*
		t.Run("toggle response if already checked", func(t *testing.T) {
			test.InitializeTestDatabase()
			pollService := test.TestPollService{}
			uuid := test.UUID{}

			id := uuid.Generate()
			owner := "Test Owner"
			question := "Test question?"
			options := []string{"Option 1", "Option 2"}

			poll := CreatePoll(pollService, entity.Poll{
				ID:       id,
				Question: question,
				Owner:    owner,
				Options:  options,
			})

			firstPollAnswer := UpsertPollAnswer(pollService, poll.ID, owner, "Option 1")
			assert.Equal(t, firstPollAnswer.Poll, poll)

			secondPollAnswer := UpsertPollAnswer(pollService, poll.ID, owner, "Option 1")
			assert.Equal(t, secondPollAnswer.Poll, nil)
		})

		t.Run("change selected response", func(t *testing.T) {
			test.InitializeTestDatabase()
		})
	*/
}
