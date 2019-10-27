package core

import (
	"errors"

	"github.com/jerolan/slack-poll/core/service"
)

func DeletePoll(pollRepository service.PollService, ownerID string, pollId string) error {
	poll, err := pollRepository.GetPollByID(pollId)

	if err != nil {
		return err
	}

	if poll.Owner != ownerID {
		return errors.New("oh no")
	}

	pollRepository.DeletePoll(pollId)
	return nil
}
