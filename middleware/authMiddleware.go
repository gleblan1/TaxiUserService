package middleware

import (
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
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			return
		}
		isTokenExists, err := s.s.ValidateToken(c, accessToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token: " + err.Error()})
			return
		}
		if !isTokenExists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
			return
		}
	}
}
