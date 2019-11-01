package main

import (
	"github.com/jerolan/slack-poll/infra/database"
	"github.com/jerolan/slack-poll/infra/webserver"
)

func main() {
	database.InitializeDatabase()

	// TODO: Define server mode via ENV
	ginServer := webserver.CreateServer(3000, webserver.DebugMode)
	ginServer.Start()
}
