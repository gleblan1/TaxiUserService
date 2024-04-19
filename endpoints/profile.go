package handler

import (
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/models"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetAccountInfo(c *gin.Context) {
	token := utils.GetTokenFromHeader(c)
	claims, err := utils.ExtractClaims(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	id, _ := strconv.Atoi(claims.Audience)
	response, err := h.s.GetAccountInfo(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (h *Handler) UpdateProfile(c *gin.Context) {
	var newData models.PatchRequest
	if err := c.ShouldBindJSON(&newData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token := utils.GetTokenFromHeader(c)
	claims, err := utils.ExtractClaims(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	id, _ := strconv.Atoi(claims.Audience)
	profile, err := h.s.UpdateProfile(id, newData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, profile)
}

func (h *Handler) DeleteProfile(c *gin.Context) {
	token := utils.GetTokenFromHeader(c)
	claims, err := utils.ExtractClaims(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, _ := strconv.Atoi(claims.Audience)
	if err := h.s.DeleteProfile(c, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusNoContent, gin.H{"message": "Profile deleted"})
}
