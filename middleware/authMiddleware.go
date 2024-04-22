package middleware

import (
	"errors"
	handler "github.com/GO-Trainee/GlebL-innotaxi-userservice/endpoints"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/services"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthMiddleware struct {
	s services.Service
}

func NewAuthMiddleware(s services.Service) *AuthMiddleware {
	return &AuthMiddleware{s: s}
}

func (s *AuthMiddleware) ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := utils.GetTokenFromHeader(c)
		if accessToken == "" {
			handler.DefineResponse(c, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}
	}
}
