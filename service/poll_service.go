package service

import "github.com/jerolan/slack-poll/core/entity"

type PollService interface {
	Find(string) entity.Poll
	Create(entity.Poll) entity.Poll
	Delete(string) entity.Poll
	AddOrUpdateAnswer(string, string) entity.PollAnswer
}
