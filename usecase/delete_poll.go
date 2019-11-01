package usecase

import (
	"errors"
)

func (uc *UseCase) DeletePoll(ownerID string, pollId string) error {
	poll, err := uc.pollService.GetPollByID(pollId)

	if err != nil {
		return err
	}

	if poll.Owner != ownerID {
		return errors.New("oh no")
	}

	uc.pollService.DeletePoll(pollId)
	return nil
}
