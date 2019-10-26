package core

import (
	"github.com/jerolan/slack-poll/core/entity"
	"github.com/jerolan/slack-poll/core/service"
)

func GetPolls(pollService service.PollService, pollID string) entity.Poll {
	poll := pollService.Find(pollID)
	return poll
}
