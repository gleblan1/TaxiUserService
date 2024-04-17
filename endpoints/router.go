package handler

import (
	"encoding/json"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Router struct {
	authMiddleware IAuthMiddleware
	handler        Handler
}

func NewRouter(authMiddleware IAuthMiddleware, handler Handler) *Router {
	return &Router{authMiddleware: authMiddleware, handler: handler}
}

type IAuthMiddleware interface {
	ValidateToken() gin.HandlerFunc
}

func (r *Router) InitRoutes() *gin.Engine {

	router := gin.Default()
	auth := router.Group("/auth")
	auth.POST("/login", r.handler.LogIn)
	auth.POST("/sign-up", r.handler.SignUp)
	auth.POST("/refresh", r.handler.Refresh)
	auth.Use(r.authMiddleware.ValidateToken()).POST("/logout", r.handler.LogOut)
	return router
}

func getTokenData(c *gin.Context) models.RefreshRequestBody {
	refreshToken := models.RefreshRequestBody{}
	values, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if err := json.Unmarshal(values, &refreshToken); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	return refreshToken
}
