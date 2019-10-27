package core

import (
	"github.com/jerolan/slack-poll/core/entity"
	"github.com/jerolan/slack-poll/core/service"
)

func UpsertPollAnswer(
	pollService service.PollService,
	pollAnswer *entity.PollAnswer,
) error {
	_, err := pollService.GetPollByID(pollAnswer.PollID)

	if err != nil {
		return err
	}

	err = pollService.CreatePollAnswer(pollAnswer)

	if err != nil {
		return err
	}

	return nil
}
