package core

import (
	"github.com/jerolan/slack-poll/core/entity"
)

func DeletePoll(pollRepository PollRepository, pollId string) entity.Poll {
	poll := pollRepository.Find(pollId)
	pollRepository.Delete(pollId)
	return poll
}
