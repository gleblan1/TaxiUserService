package middleware

import (
	"errors"
	"net/http"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/services"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/utils"
	"github.com/gin-gonic/gin"
)

type Middleware struct {
	s *services.Service
}

func NewMiddleware(options ...func(*Middleware)) *Middleware {
	repo := &Middleware{}
	for _, option := range options {
		option(repo)
	}
	return repo
}

func WithAuthMiddleware(s *services.Service) func(*Middleware) {
	return func(m *Middleware) {
		m.s = s
	}
}

func (s *Middleware) ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := utils.GetTokenFromHeader(c)
		if accessToken == "" {
			utils.DefineResponse(c, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}
		isTokenValid, err := s.s.ValidateToken(c, accessToken)
		if err != nil {
			utils.DefineResponse(c, http.StatusUnauthorized, errors.New("invalid token"))
			return
		}
		if !isTokenValid {
			utils.DefineResponse(c, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}
	}
}
