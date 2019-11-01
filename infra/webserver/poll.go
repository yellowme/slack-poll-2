package webserver

import (
	"github.com/gin-gonic/gin"
	"github.com/jerolan/slack-poll/adapter"
	"github.com/jerolan/slack-poll/adapter/controller"
	"github.com/jerolan/slack-poll/adapter/service"
	"github.com/jerolan/slack-poll/infra/database"
	"github.com/jerolan/slack-poll/usecase"
)

type pollHandler struct {
	pollController *controller.PollController
}

func (server *GinServer) SetupPollHandler() {
	uuid := adapter.NewUUID()
	pollService := service.NewGormPollService(database.GormDatabase)

	useCase := usecase.NewUseCase(pollService, uuid)
	pollController := controller.NewPollController(useCase)

	handler := &pollHandler{
		pollController: pollController,
	}

	server.Router.POST("/poll", handler.createPostPoll)
}

func (h *pollHandler) createPostPoll(c *gin.Context) {
	var pollBody struct {
		Command string `json:"command" binding:"required"`
	}

	c.BindJSON(&pollBody)

	owner := "Test User"

	h.pollController.CreatePoll(owner, pollBody.Command)
}
