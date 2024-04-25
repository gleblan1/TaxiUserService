package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/requests"
	"github.com/gin-gonic/gin"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/utils"
)

type RefreshRequestBody struct {
	RefreshToken string `json:"refresh_token"`
}

type RegisterRequest struct {
	Name        string `json:"name" binding:"required,min=4,max=20,alpha"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phone_number" binding:"required,e164,len=13"`
	Password    string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required,e164"`
	Password    string `json:"password" binding:"required,min=8"`
}

func (h *Handler) SignUp(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}

	requestBody := requests.RegisterRequest{
		Name:        req.Name,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
	}

	result, err := h.e.SignUp(c, requestBody)
	if err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}
	fmt.Println(result)
	utils.DefineResponse(c, http.StatusOK, err, result)
	return
}

func (h *Handler) LogIn(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}

	requestBody := requests.LoginRequest{
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
	}

	response, err := h.e.Login(c, requestBody)
	if err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}
	utils.DefineResponse(c, http.StatusOK, err, response)
	return
}

func (h *Handler) LogOut(c *gin.Context) {
	claims, err := utils.ExtractClaims(utils.GetTokenFromHeader(c))
	if err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}
	id, err := strconv.Atoi(claims.Audience)
	if err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}
	session, err := strconv.Atoi(claims.Session)
	if err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}
	requestBody := requests.LogoutRequest{
		SessionId: session,
		UserId:    id,
	}
	if err, _ := h.e.LogOut(c, requestBody); err != nil {
		return
	}
	return
}

func (h *Handler) Refresh(c *gin.Context) {
	refreshTokenString := h.getTokenData(c).RefreshToken
	req := RefreshRequestBody{
		RefreshToken: refreshTokenString,
	}

	requestBody := requests.RefreshRequestBody{
		RefreshToken: req.RefreshToken,
	}

	tokens, err := h.e.Refresh(c, requestBody)
	if err != nil {
		utils.DefineResponse(c, http.StatusUnauthorized, err)
		return
	}
	utils.DefineResponse(c, http.StatusOK, err, tokens)
	return
}
