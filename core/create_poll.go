package core

import (
	"github.com/jerolan/slack-poll/core/entity"
	"github.com/jerolan/slack-poll/core/service"
)

func CreatePoll(pollService service.PollService, poll *entity.Poll) error {
	if poll.Mode == 0 {
		poll.Mode = entity.PollModeSingle
	}

	err := pollService.CreatePoll(poll)

	if err != nil {
		return err
	}

	return nil
}
