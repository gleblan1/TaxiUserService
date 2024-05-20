package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/services"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/utils"
)

var (
	unauthorizedErr = errors.New("unauthorized")
	invalidTokenErr = errors.New("invalid token")
)

type Middleware struct {
	service *services.Service
}

func NewMiddleware(options ...func(*Middleware)) *Middleware {
	repo := &Middleware{}
	for _, option := range options {
		option(repo)
	}
	return repo
}

func WithAuthMiddleware(service *services.Service) func(*Middleware) {
	return func(middleware *Middleware) {
		middleware.service = service
	}
}

func (m *Middleware) ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := utils.GetTokenFromHeader(c)
		if accessToken == "" {
			utils.DefineResponse(c, http.StatusUnauthorized, unauthorizedErr)
			return
		}
		isTokenValid, err := m.service.ValidateToken(c, accessToken)
		if err != nil {
			utils.DefineResponse(c, http.StatusUnauthorized, invalidTokenErr)
			return
		}
		if !isTokenValid {
			utils.DefineResponse(c, http.StatusUnauthorized, unauthorizedErr)
			return
		}
		c.Next()
	}
}
