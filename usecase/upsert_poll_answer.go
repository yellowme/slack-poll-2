package usecase

import "github.com/jerolan/slack-poll/domain/entity"

func (uc *UseCase) UpsertPollAnswer(pollAnswer *entity.PollAnswer) error {
	_, err := uc.pollService.GetPollByID(pollAnswer.PollID)

	if err != nil {
		return err
	}

	err = uc.pollService.CreatePollAnswer(pollAnswer)

	if err != nil {
		return err
	}

	return nil
}
