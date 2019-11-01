package usecase

import (
	"github.com/jerolan/slack-poll/domain/port"
	"github.com/jerolan/slack-poll/domain/service"
)

type UseCase struct {
	pollService service.PollService
	uuid        port.UUIDPort
}

func NewUseCase(pS service.PollService, uuid port.UUIDPort) *UseCase {
	return &UseCase{
		pollService: pS,
		uuid:        uuid,
	}
}
