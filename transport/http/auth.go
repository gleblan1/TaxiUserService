package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/models"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/requests"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/utils"
)

type RefreshTokensRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RegisterRequest struct {
	Name        string `json:"name" binding:"required,min=4,max=20,alpha"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phone_number" binding:"required,e164,len=13"`
	Password    string `json:"password" binding:"required,min=8"`
}

type RegisterResponse struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	PhoneNumber string  `json:"phone_number"`
	Password    string  `json:"password"`
	Rating      float32 `json:"rating"`
}

type SignInRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required,e164"`
	Password    string `json:"password" binding:"required,min=8"`
}

type SignInResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
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

	response := RegisterResponse{
		Id:          result.(models.User).Id,
		Name:        result.(models.User).Name,
		Email:       result.(models.User).Email,
		PhoneNumber: result.(models.User).PhoneNumber,
		Password:    result.(models.User).Password,
		Rating:      result.(models.User).Rating,
	}
	utils.DefineResponse(c, http.StatusOK, err, response)
	return
}

func (h *Handler) SignIn(c *gin.Context) {
	var req SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}

	requestBody := requests.SignInRequest{
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
	}

	tokens, err := h.e.SignIn(c, requestBody)
	if err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}
	response := SignInResponse{
		AccessToken:  tokens.(models.JwtToken).AccessToken,
		RefreshToken: tokens.(models.JwtToken).RefreshToken,
	}
	utils.DefineResponse(c, http.StatusOK, err, response)
	return
}

func (h *Handler) SignOut(c *gin.Context) {
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
	if err, _ := h.e.SignOut(c, requestBody); err != nil {
		return
	}
	return
}

func (h *Handler) RefreshTokens(c *gin.Context) {
	refreshTokenString := h.getTokenData(c).RefreshToken
	req := RefreshTokensRequest{
		RefreshToken: refreshTokenString,
	}

	requestBody := requests.RefreshTokensRequest{
		RefreshToken: req.RefreshToken,
	}

	tokens, err := h.e.RefreshTokens(c, requestBody)
	if err != nil {
		utils.DefineResponse(c, http.StatusUnauthorized, err)
		return
	}
	response := SignInResponse{
		AccessToken:  tokens.(models.JwtToken).AccessToken,
		RefreshToken: tokens.(models.JwtToken).RefreshToken,
	}
	utils.DefineResponse(c, http.StatusOK, err, response)
	return
}
