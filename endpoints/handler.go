package handler

import (
	"encoding/json"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/middleware"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/models"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	s *services.Service
}

func NewHandler(s *services.Service) *Handler {
	return &Handler{s: s}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()
	auth := r.Group("/auth")
	auth.POST("/login", h.LogIn)
	auth.POST("/sign-up", h.SignUp)
	auth.POST("/logout", h.LogOut).Use(middleware.ValidateToken())
	return r
}

func getUserData(c *gin.Context) models.User {
	user := models.User{}
	values, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if err := json.Unmarshal(values, &user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	return user
}
