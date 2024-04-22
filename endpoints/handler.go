package handler

import (
	"encoding/json"
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

func DefineResponse(c *gin.Context, code int, err error, response ...interface{}) {
	var Response models.Response
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	} else {
		errMsg = ""
	}
	Response = models.Response{
		Code:     code,
		Message:  errMsg,
		Response: response,
	}
	c.JSON(code, Response)
}

func (h *Handler) getTokenData(c *gin.Context) models.RefreshRequestBody {
	refreshToken := models.RefreshRequestBody{}
	values, err := c.GetRawData()
	if err != nil {
		DefineResponse(c, http.StatusBadRequest, err, values)
		return refreshToken
	}
	if err := json.Unmarshal(values, &refreshToken); err != nil {
		DefineResponse(c, http.StatusBadRequest, err, values)
		return refreshToken
	}
	return refreshToken
}
