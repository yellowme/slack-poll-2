package test

import (
	"strings"

	"github.com/jerolan/slack-poll/core/entity"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type TestPollService struct{}

func (tps TestPollService) GetPollByID(pollID string) (entity.Poll, error) {
	var poll entity.Poll

	err := TestDatabase.Where("id = ?", pollID).First(&poll).Error

	if err != nil {
		return poll, err
	}

	return poll, nil
}

func (tps TestPollService) CreatePoll(poll *entity.Poll) error {
	createdPoll := PollModel{
		ID:       poll.ID,
		Question: poll.Question,
		Owner:    poll.Owner,
		Options:  strings.Join(poll.Options, ","),
		Mode:     poll.Mode.String(),
	}

	err := TestDatabase.Create(&createdPoll).Error
	if err != nil {
		return err
	}

	poll.ID = createdPoll.ID

	return nil
}

func (tps TestPollService) DeletePoll(pollID string) error {
	err := TestDatabase.Where("id = ?", pollID).Delete(&PollModel{}).Error

	if err != nil {
		return err
	}

	return nil
}

func (tps TestPollService) FindPollAnswers(pollID string) ([]entity.PollAnswer, error) {
	var pollAnswers []entity.PollAnswer
	err := TestDatabase.Where("poll_id = ?", pollID).Find(&pollAnswers).Error

	if err != nil {
		return pollAnswers, err
	}

	return pollAnswers, nil
}

func (tps TestPollService) CreatePollAnswer(pollAnswer *entity.PollAnswer) error {
	createdPollAnswer := PollAnswerModel{
		Option: pollAnswer.Option,
		Owner:  pollAnswer.Owner,
		PollID: pollAnswer.PollID,
	}

	err := TestDatabase.Create(&createdPollAnswer).Error
	if err != nil {
		return err
	}

	return nil
}

func (tps TestPollService) DeletePollAnswer(pollAnswerID string) error {
	err := TestDatabase.Where("id = ?", pollAnswerID).Delete(&PollAnswerModel{}).Error

	if err != nil {
		return err
	}

	return nil
}
