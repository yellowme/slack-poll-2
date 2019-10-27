package core

import (
	"testing"

	"github.com/jerolan/slack-poll/core/entity"
	"github.com/jerolan/slack-poll/test"
	"github.com/stretchr/testify/assert"
)

func TestDeletePoll(t *testing.T) {
	test.InitializeTestDatabase()

	t.Run("works ok", func(t *testing.T) {
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

		err := CreatePoll(pollService, &poll)

		if err != nil {
			t.Error("not expected to fail")
		}

		err = DeletePoll(pollService, owner, poll.ID)

		if err != nil {
			t.Error("not expected to fail")
		}
	})

	t.Run("fail if not owner", func(t *testing.T) {
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

		err := CreatePoll(pollService, &poll)

		if err != nil {
			t.Error("not expected to fail")
		}

		err = DeletePoll(pollService, owner, poll.ID)

		assert.EqualError(t, err, "oh no")
	})
}
