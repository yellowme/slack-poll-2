package core

import (
	"github.com/jerolan/slack-poll/core/entity"
)

func GetPolls(pollRepository PollRepository, pollID string) entity.Poll {
	poll := pollRepository.Find(pollID)
	return poll
}
