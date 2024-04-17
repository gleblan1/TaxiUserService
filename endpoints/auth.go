package handler

import (
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/models"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) SignUp(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := h.s.SignUp(req.Name, req.PhoneNumber, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *Handler) LogIn(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.s.Login(c, req.Name, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (h *Handler) LogOut(c *gin.Context) {
	claims, err := utils.ExtractClaims(utils.GetTokenFromHeader(c))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := strconv.Atoi(claims.Audience)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.s.LogOut(c, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func (h *Handler) Refresh(c *gin.Context) {
	refreshTokenString := getTokenData(c).RefreshToken
	tokens, err := h.s.Refresh(c, refreshTokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tokens)
}
