package providers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	Port   string
	Server *gin.Engine
}

func InitServer(port string, server *gin.Engine) error {
	err := server.Run(":" + port)
	if err != nil {
		return err
	}
	return nil
}

func NewServer(options ...func(*Server)) *Server {
	server := &Server{}
	for _, option := range options {
		option(server)
	}
	return server
}

func WithPort(port string) func(*Server) {
	return func(server *Server) {
		server.Port = port
	}
}

func WithServer(srv *gin.Engine) func(*Server) {
	return func(server *Server) {
		server.Server = srv
	}
}
