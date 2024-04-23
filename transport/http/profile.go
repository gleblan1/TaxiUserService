package http

import (
	"net/http"
	"strconv"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/config"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/utils"
	"github.com/gin-gonic/gin"
)

type PatchRequest struct {
	Name        string `json:"name" binding:"omitempty,min=4,max=20"`
	PhoneNumber string `json:"phone_number" binding:"omitempty,phoneValid"`
	Email       string `json:"email" binding:"omitempty,emailValid"`
}

type GetAccountInfoRequest struct {
	Id int
}

type UpdateProfileRequest struct {
	Id      int
	NewData PatchRequest
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
	requestBody := config.GetAccountInfoRequest{
		Id: id,
	}
	response, err := h.e.GetAccountInfo(c, requestBody)
	if err != nil {

		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}
	utils.DefineResponse(c, http.StatusOK, err, response)
	return
}

func (h *Handler) UpdateProfile(c *gin.Context) {
	var newData config.PatchRequest
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

	requestBody := config.UpdateProfileRequest{
		Id:      id,
		NewData: newData,
	}

	profile, err := h.e.UpdateProfile(c, requestBody)
	if err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}
	utils.DefineResponse(c, http.StatusOK, nil, profile)
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
	requestBody := config.DeleteProfileRequest{
		Id: id,
	}
	if _, err := h.e.DeleteProfile(c, requestBody); err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}
	utils.DefineResponse(c, http.StatusOK, nil, "Profile deleted successfully")
	return
}
