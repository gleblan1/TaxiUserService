package handler

import (
	"errors"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/models"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) SignUp(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		DefineResponse(c, http.StatusBadRequest, err)
		return
	}
	result, err := h.s.SignUp(req.Name, req.PhoneNumber, req.Email, req.Password)
	if err != nil {
		DefineResponse(c, http.StatusBadRequest, err)
		return
	}
	DefineResponse(c, http.StatusOK, err, result)
	return
}

func (h *Handler) LogIn(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		DefineResponse(c, http.StatusBadRequest, err)
		return
	}

	response, err := h.s.Login(c, req.PhoneNumber, req.Password)
	if err != nil {
		DefineResponse(c, http.StatusBadRequest, err)
		return
	}
	DefineResponse(c, http.StatusOK, err, response)
	return
}

func (h *Handler) LogOut(c *gin.Context) {
	claims, err := utils.ExtractClaims(utils.GetTokenFromHeader(c))
	if err != nil {
		DefineResponse(c, http.StatusBadRequest, err)
		return
	}
	id, err := strconv.Atoi(claims.Audience)
	if err != nil {
		DefineResponse(c, http.StatusBadRequest, err)
		return
	}
	session, err := strconv.Atoi(claims.Session)
	if err != nil {
		DefineResponse(c, http.StatusBadRequest, err)
		return
	}
	if err := h.s.LogOut(c, session, id); err != nil {
		DefineResponse(c, http.StatusBadRequest, errors.New("already log outed "))
		return
	}
	return
}

func (h *Handler) Refresh(c *gin.Context) {
	refreshTokenString := h.getTokenData(c).RefreshToken
	tokens, err := h.s.Refresh(c, refreshTokenString)
	if err != nil {
		DefineResponse(c, http.StatusUnauthorized, err)
		return
	}
	DefineResponse(c, http.StatusOK, err, tokens)
	return
}
