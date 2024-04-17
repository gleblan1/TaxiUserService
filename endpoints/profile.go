package handler

import (
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) GetAccountInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	//handle error
	h.s.GetAccountInfo(id)
}

func (h *Handler) UpdateProfile(c *gin.Context) {
	var user models.User
	id, _ := strconv.Atoi(c.Param("id"))
	h.s.UpdateProfile(id, user)
}

func (h *Handler) DeleteProfile(c *gin.Context) {
}
