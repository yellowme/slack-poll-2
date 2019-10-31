package service

import (
	"strings"

	"github.com/jerolan/slack-poll/core/entity"
	"github.com/jerolan/slack-poll/infra/database"
)

type GormPollService struct{}

func (tps GormPollService) GetPollByID(pollID string) (entity.Poll, error) {
	var poll entity.Poll

	err := database.Database.DB.Where("id = ?", pollID).First(&poll).Error

	if err != nil {
		return poll, err
	}

	return poll, nil
}

func (tps GormPollService) CreatePoll(poll *entity.Poll) error {
	createdPoll := database.PollModel{
		ID:       poll.ID,
		Question: poll.Question,
		Owner:    poll.Owner,
		Options:  strings.Join(poll.Options, ","),
		Mode:     poll.Mode.String(),
	}

	err := database.Database.DB.Create(&createdPoll).Error
	if err != nil {
		return err
	}

	poll.ID = createdPoll.ID

	return nil
}

func (tps GormPollService) DeletePoll(pollID string) error {
	err := database.Database.DB.Where("id = ?", pollID).Delete(&database.PollModel{}).Error

	if err != nil {
		return err
	}

	return nil
}

func (tps GormPollService) FindPollAnswers(pollID string) ([]entity.PollAnswer, error) {
	var pollAnswers []entity.PollAnswer
	err := database.Database.DB.Where("poll_id = ?", pollID).Find(&pollAnswers).Error

	if err != nil {
		return pollAnswers, err
	}

	return pollAnswers, nil
}

func (tps GormPollService) CreatePollAnswer(pollAnswer *entity.PollAnswer) error {
	createdPollAnswer := database.PollAnswerModel{
		Option: pollAnswer.Option,
		Owner:  pollAnswer.Owner,
		PollID: pollAnswer.PollID,
	}

	err := database.Database.DB.Create(&createdPollAnswer).Error
	if err != nil {
		return err
	}

	return nil
}

func (tps GormPollService) DeletePollAnswer(pollAnswerID string) error {
	err := database.Database.DB.Where("id = ?", pollAnswerID).Delete(&database.PollAnswerModel{}).Error

	if err != nil {
		return err
	}

	return nil
}
