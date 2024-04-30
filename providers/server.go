package providers

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	serverNotConfiguredError = errors.New("server is not configured")
)

type Server struct {
	Port   string
	Server *http.Server
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
		if srv == nil {
			server.Server = &http.Server{
				Addr:    ":" + server.Port,
				Handler: gin.Default(),
			}
		} else {
			server.Server = &http.Server{
				Addr:    ":" + server.Port,
				Handler: srv,
			}
		}
	}
}

func (server *Server) Run() error {
	if server.Server == nil {
		return serverNotConfiguredError
	}
	err := server.Server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (server *Server) Stop(ctx context.Context) error {
	return server.Server.Shutdown(ctx)
}
