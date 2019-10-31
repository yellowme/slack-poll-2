package controller

import (
	"github.com/jerolan/slack-poll/core"
	"github.com/jerolan/slack-poll/core/entity"
	"github.com/jerolan/slack-poll/core/service"
	"github.com/jerolan/slack-poll/infra/webserver"
)

func CreatePoll(pollService service.PollService, command string, poll *entity.Poll) {
	pollFromCommand := webserver.ConvertCommandToStruct(command)
	poll.Question = pollFromCommand.Question
	poll.Mode = pollFromCommand.Mode
	poll.Options = pollFromCommand.Options
	core.CreatePoll(pollService, poll)
}
