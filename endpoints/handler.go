package handler

import (
	"encoding/json"
	"net/http"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/services"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/utils"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	s *services.Service
}

func NewHandler(options ...func(*Handler)) *Handler {
	handler := &Handler{}
	for _, option := range options {
		option(handler)
	}
	return handler
}

func WithAuthService(service *services.Service) func(*Handler) {
	return func(h *Handler) {
		h.s = service
	}
}

func (h *Handler) getTokenData(c *gin.Context) RefreshRequestBody {
	refreshToken := RefreshRequestBody{}
	values, err := c.GetRawData()
	if err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err, values)
		return refreshToken
	}
	if err := json.Unmarshal(values, &refreshToken); err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err, values)
		return refreshToken
	}
	return refreshToken
}
