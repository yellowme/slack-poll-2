package usecase

import "github.com/jerolan/slack-poll/domain/entity"

func (uc *UseCase) CreatePoll(poll *entity.Poll) error {
	if poll.Mode == 0 {
		poll.Mode = entity.PollModeSingle
	}

	err := uc.pollService.CreatePoll(poll)

	if err != nil {
		return err
	}

	return nil
}
