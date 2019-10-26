package core

import (
	"github.com/jerolan/slack-poll/core/entity"
	"github.com/jerolan/slack-poll/core/service"
)

func UpsertPollAnswer(
	pollService service.PollService,
	pollID string,
	option string,
) entity.PollAnswer {
	pollAnswer := pollService.AddOrUpdateAnswer(pollID, option)
	return pollAnswer
}
