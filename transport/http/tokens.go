package http

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/utils"
)

func (h *Handler) getTokenData(c *gin.Context) RefreshTokensRequest {
	refreshToken := RefreshTokensRequest{}
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
