package controller

import (
	"github.com/jerolan/slack-poll/domain/entity"
	"github.com/jerolan/slack-poll/usecase"
	"github.com/jerolan/slack-poll/utils"
)

type PollController struct {
	useCase *usecase.UseCase
}

func NewPollController(uc *usecase.UseCase) *PollController {
	return &PollController{
		useCase: uc,
	}
}

func (pcl *PollController) CreatePoll(owner, command string) (poll entity.Poll) {
	pollFromCommand := utils.ConvertCommandToPoll(command)

	poll.Owner = owner
	poll.Mode = pollFromCommand.Mode
	poll.Question = pollFromCommand.Question
	poll.Options = pollFromCommand.Options

	pcl.useCase.CreatePoll(&poll)
	return poll
}
