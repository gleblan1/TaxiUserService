package http

import (
	"encoding/json"
	"net/http"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/utils"
	"github.com/gin-gonic/gin"
)

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
