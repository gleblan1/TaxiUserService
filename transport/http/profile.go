package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/models"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/requests"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/utils"
)

type PatchRequest struct {
	Id          int    `json:"id"`
	Name        string `json:"name,omitempty" binding:"omitempty,alpha,min=5,max=20"`
	PhoneNumber string `json:"phone_number,omitempty" binding:"omitempty,e164,len=13"`
	Email       string `json:"email,omitempty" binding:"omitempty,email"`
}

type GetAccountInfoRequest struct {
	Id int
}

type userInfoResponse struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	PhoneNumber string  `json:"phone_number"`
	Email       string  `json:"email"`
	Rating      float32 `json:"rating"`
}

type DeleteProfileRequest struct {
	Id int
}

func (h *Handler) GetAccountInfo(c *gin.Context) {
	token := utils.GetTokenFromHeader(c)
	claims, err := utils.ExtractClaims(token)
	if err != nil {

		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}
	id, _ := strconv.Atoi(claims.Audience)
	requestBody := requests.GetAccountInfoRequest{
		Id: id,
	}
	accountInfo, err := h.e.GetAccountInfo(c, requestBody)
	response := userInfoResponse{
		Id:          accountInfo.(models.User).Id,
		Name:        accountInfo.(models.User).Name,
		PhoneNumber: accountInfo.(models.User).PhoneNumber,
		Email:       accountInfo.(models.User).Email,
		Rating:      accountInfo.(models.User).Rating,
	}
	if err != nil {

		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}
	utils.DefineResponse(c, http.StatusOK, err, response)
	return
}

func (h *Handler) UpdateProfile(c *gin.Context) {
	var newData PatchRequest
	if err := c.ShouldBindJSON(&newData); err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}
	token := utils.GetTokenFromHeader(c)
	claims, err := utils.ExtractClaims(token)
	if err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}
	id, _ := strconv.Atoi(claims.Audience)

	requestBody := requests.UpdateProfileRequest{
		Id:          id,
		Name:        newData.Name,
		PhoneNumber: newData.PhoneNumber,
		Email:       newData.Email,
	}

	profile, err := h.e.UpdateProfile(c, requestBody)
	response := userInfoResponse{
		Id:          profile.(models.User).Id,
		Name:        profile.(models.User).Name,
		PhoneNumber: profile.(models.User).PhoneNumber,
		Email:       profile.(models.User).Email,
		Rating:      profile.(models.User).Rating,
	}
	if err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}
	utils.DefineResponse(c, http.StatusOK, nil, response)
	return
}

func (h *Handler) DeleteProfile(c *gin.Context) {
	token := utils.GetTokenFromHeader(c)
	claims, err := utils.ExtractClaims(token)
	if err != nil {

		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}
	id, _ := strconv.Atoi(claims.Audience)
	requestBody := requests.DeleteProfileRequest{
		Id: id,
	}
	if _, err := h.e.DeleteProfile(c, requestBody); err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}
	utils.DefineResponse(c, http.StatusOK, nil, "Profile deleted successfully")
	return
}
