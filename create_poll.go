package core

import (
	"github.com/jerolan/slack-poll/core/entity"
)

func CreatePoll(pollRepository PollRepository, question string, options []string) entity.Poll {
	poll := pollRepository.Create(entity.Poll{Question: question, Options: options})
	return poll
}
