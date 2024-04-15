package middleware

import (
	"fmt"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/services"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"time"
)

type authMiddleware struct {
	s services.Service
}

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := utils.GetTokenFromHeader(c)
		if accessToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		}
		jwtAccess, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return []byte(os.Getenv("SECRET")), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		}
		expTime, err := jwtAccess.Claims.GetExpirationTime()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		}
		if expTime.Unix() <= time.Now().Unix() {
			fmt.Println(expTime.Unix() >= time.Now().Unix())
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Token is expired"})
		}
	}
}
