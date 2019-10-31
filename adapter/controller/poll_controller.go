package controller

import (
	"github.com/jerolan/slack-poll/core"
	"github.com/jerolan/slack-poll/core/entity"
	"github.com/jerolan/slack-poll/core/port"
	"github.com/jerolan/slack-poll/core/service"
	"github.com/jerolan/slack-poll/utils"
)

func CreatePoll(
	pollService service.PollService,
	uuid port.UUID,
	owner,
	command string,
) (poll entity.Poll) {
	pollFromCommand := utils.ConvertCommandToPoll(command)

	poll.ID = uuid.Generate()
	poll.Owner = owner
	poll.Mode = pollFromCommand.Mode
	poll.Question = pollFromCommand.Question
	poll.Options = pollFromCommand.Options

	core.CreatePoll(pollService, &poll)
	return poll
}
