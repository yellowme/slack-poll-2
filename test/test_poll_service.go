package test

import (
	"strings"

	"github.com/jerolan/slack-poll/domain/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type testPollService struct {
	gormDB *gorm.DB
}

func NewTestPollService(gormDB *gorm.DB) *testPollService {
	return &testPollService{
		gormDB: gormDB,
	}
}

func (tps *testPollService) GetPollByID(pollID string) (entity.Poll, error) {
	var poll entity.Poll

	err := tps.gormDB.Where("id = ?", pollID).First(&poll).Error

	if err != nil {
		return poll, err
	}

	return poll, nil
}

func (tps *testPollService) CreatePoll(poll *entity.Poll) error {
	createdPoll := PollModel{
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

func (tps *testPollService) DeletePoll(pollID string) error {
	err := tps.gormDB.Where("id = ?", pollID).Delete(&PollModel{}).Error

	if err != nil {
		return err
	}

	return nil
}

func (tps *testPollService) FindPollAnswers(pollID string) ([]entity.PollAnswer, error) {
	var pollAnswers []entity.PollAnswer
	err := tps.gormDB.Where("poll_id = ?", pollID).Find(&pollAnswers).Error

	if err != nil {
		return pollAnswers, err
	}

	return pollAnswers, nil
}

func (tps *testPollService) CreatePollAnswer(pollAnswer *entity.PollAnswer) error {
	createdPollAnswer := PollAnswerModel{
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

func (tps *testPollService) DeletePollAnswer(pollAnswerID string) error {
	err := tps.gormDB.Where("id = ?", pollAnswerID).Delete(&PollAnswerModel{}).Error

	if err != nil {
		return err
	}

	return nil
}
