package core

import "github.com/jerolan/slack-poll/core/entity"

type PollRepository interface {
	Create(entity.Poll) (poll entity.Poll)
}
