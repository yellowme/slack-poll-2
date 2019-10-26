package core

import (
	"github.com/jerolan/slack-poll/core/entity"
	"github.com/jerolan/slack-poll/core/service"
)

func CreatePoll(pollService service.PollService, question string, options []string) entity.Poll {
	poll := pollService.Create(entity.Poll{Question: question, Options: options})
	return poll
}
