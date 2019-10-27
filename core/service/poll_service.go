package service

import "github.com/jerolan/slack-poll/core/entity"

type PollService interface {
	GetPollByID(pollID string) (entity.Poll, error)
	CreatePoll(*entity.Poll) error
	DeletePoll(pollID string) error
	FindPollAnswers(pollID string) ([]entity.PollAnswer, error)
	CreatePollAnswer(*entity.PollAnswer) error
	DeletePollAnswer(pollAnswerID string) error
}
