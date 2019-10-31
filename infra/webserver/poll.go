package webserver

import (
	"github.com/gin-gonic/gin"
	"github.com/jerolan/slack-poll/adapter"
	"github.com/jerolan/slack-poll/adapter/controller"
	"github.com/jerolan/slack-poll/adapter/service"
)

func SetUpPollHanlder(r *gin.Engine) {
	r.POST("/poll", CreatePostPoll)
}

func CreatePostPoll(c *gin.Context) {
	var pollBody struct {
		Command string `json:"command" binding:"required"`
	}

	c.BindJSON(&pollBody)

	uuid := &adapter.UUID{}
	pollService := &service.GormPollService{}

	owner := "Test User"

	controller.CreatePoll(pollService, uuid, owner, pollBody.Command)
}
