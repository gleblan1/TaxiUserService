package http

import (
	"encoding/json"
	"net/http"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/endpoints"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/utils"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	e *endpoints.Endpoints
}

func NewHandler(options ...func(*Handler)) *Handler {
	handler := &Handler{}
	for _, option := range options {
		option(handler)
	}
	return handler
}

func WithAuthService(e *endpoints.Endpoints) func(*Handler) {
	return func(h *Handler) {
		h.e = e
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
