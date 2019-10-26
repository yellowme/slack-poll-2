package core

import "github.com/jerolan/slack-poll/core/entity"

type PollRepository interface {
	Find(string) entity.Poll
	Create(entity.Poll) entity.Poll
	Delete(string) entity.Poll
}
