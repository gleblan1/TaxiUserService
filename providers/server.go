package providers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	Port   string
	Server *gin.Engine
}

type ServerOption func(*Server)

func NewServer(options ...ServerOption) *Server {
	server := &Server{}
	for _, option := range options {
		option(server)
	}
	return server
}

func WithPort(port string) ServerOption {
	return func(server *Server) {
		server.Port = port
	}
}

func WithServer(srv *gin.Engine) ServerOption {
	return func(server *Server) {
		server.Server = srv
	}
}
