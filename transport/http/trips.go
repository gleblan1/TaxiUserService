package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/config"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/models"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/utils"
	"github.com/gin-gonic/gin"
)

type RateTripRequest struct {
	UserId int `json:"id"`
	Rate   int `json:"rate"`
	TripId int `json:"trip_id"`
}

type TripHistoryRequest struct {
	UserId int `json:"id"`
}

type Trip struct {
	Id       int    `json:"id"`
	TaxiType int    `json:"taxi_type"`
	From     string `json:"from"`
	To       string `json:"to"`
	Rate     int    `json:"rate"`
}

type TripHistoryResponse struct {
	Trips []Trip      `json:"trips"`
	User  models.User `json:"user"`
}

func (h *Handler) RateTrip(c *gin.Context) {
	var req RateTripRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}

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

	requestBody := config.RateTripRequest{
		UserId: id,
		Rate:   req.Rate,
	}

	result, err := h.e.RateTrip(c, requestBody)
	if err != nil {

		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}
	utils.DefineResponse(c, http.StatusOK, err, result)
	return

}

func (h *Handler) GetTripsHistory(c *gin.Context) {
	var req TripHistoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}

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
	fmt.Println(id)

	requestBody := config.GetHistoryRequest{
		UserId: id,
	}

	result, err := h.e.GetTripsHistory(c, requestBody)
	if err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}

	response := TripHistoryResponse{
		Trips: make([]Trip, len(result.(models.TripHistory).Trips)),
		User:  result.(models.TripHistory).User,
	}
	for i, v := range result.(models.TripHistory).Trips {
		response.Trips[i] = Trip{
			Id:       v.Id,
			TaxiType: v.TaxiType,
			From:     v.From,
			To:       v.To,
			Rate:     v.Rate,
		}
	}
	utils.DefineResponse(c, http.StatusOK, err, response)
	return

}
