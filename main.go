package main

import (
	"github.com/jerolan/slack-poll/database"
	"github.com/jerolan/slack-poll/server"
)

func main() {
	database.InitializeTestDatabase()
	ginServer := server.CreateServer(3000, server.DebugMode)
	ginServer.Start()
}
