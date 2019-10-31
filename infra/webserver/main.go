package webserver

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type GinServer struct {
	port   int
	Router *gin.Engine
}

func CreateServer(port int, mode GinServerMode) *GinServer {
	server := &GinServer{}
	server.port = port

	server.Router = gin.New()

	switch mode {
	case DebugMode:
		gin.SetMode(gin.DebugMode)
	case TestMode:
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	server.Router.Use(gin.Recovery())

	return server
}

func (s *GinServer) Start() {
	s.Router.Run(":" + strconv.Itoa(int(s.port)))
}
