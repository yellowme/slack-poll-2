package core

import (
	"github.com/jerolan/slack-poll/core/entity"
	"github.com/jerolan/slack-poll/core/service"
)

func DeletePoll(pollRepository service.PollService, pollId string) entity.Poll {
	poll := pollRepository.Find(pollId)
	pollRepository.Delete(pollId)
	return poll
}
