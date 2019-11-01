package service

import (
	"strings"

	"github.com/jerolan/slack-poll/domain/entity"
	"github.com/jerolan/slack-poll/infra/database"
	"github.com/jinzhu/gorm"
)

type gormPollService struct {
	gormDB *gorm.DB
}

func NewGormPollService(gormDB *gorm.DB) *gormPollService {
	return &gormPollService{
		gormDB: gormDB,
	}
}

func (tps *gormPollService) GetPollByID(pollID string) (entity.Poll, error) {
	var poll entity.Poll

	err := tps.gormDB.Where("id = ?", pollID).First(&poll).Error

	if err != nil {
		return poll, err
	}

	return poll, nil
}

func (tps *gormPollService) CreatePoll(poll *entity.Poll) error {
	createdPoll := database.PollModel{
		ID:       poll.ID,
		Question: poll.Question,
		Owner:    poll.Owner,
		Options:  strings.Join(poll.Options, ","),
		Mode:     poll.Mode.String(),
	}

	err := tps.gormDB.Create(&createdPoll).Error
	if err != nil {
		return err
	}

	poll.ID = createdPoll.ID

	return nil
}

func (tps *gormPollService) DeletePoll(pollID string) error {
	err := tps.gormDB.Where("id = ?", pollID).Delete(&database.PollModel{}).Error

	if err != nil {
		return err
	}

	return nil
}

func (tps *gormPollService) FindPollAnswers(pollID string) ([]entity.PollAnswer, error) {
	var pollAnswers []entity.PollAnswer
	err := tps.gormDB.Where("poll_id = ?", pollID).Find(&pollAnswers).Error

	if err != nil {
		return pollAnswers, err
	}

	return pollAnswers, nil
}

func (tps *gormPollService) CreatePollAnswer(pollAnswer *entity.PollAnswer) error {
	createdPollAnswer := database.PollAnswerModel{
		Option: pollAnswer.Option,
		Owner:  pollAnswer.Owner,
		PollID: pollAnswer.PollID,
	}

	err := tps.gormDB.Create(&createdPollAnswer).Error
	if err != nil {
		return err
	}

	return nil
}

func (tps *gormPollService) DeletePollAnswer(pollAnswerID string) error {
	err := tps.gormDB.Where("id = ?", pollAnswerID).Delete(&database.PollAnswerModel{}).Error

	if err != nil {
		return err
	}

	return nil
}
