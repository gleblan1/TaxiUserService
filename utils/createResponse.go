package utils

import (
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/models"
	"github.com/gin-gonic/gin"
)

func DefineResponse(c *gin.Context, code int, err error, response ...interface{}) {
	var Response models.Response
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	} else {
		errMsg = ""
	}
	Response = models.Response{
		StatusCode: code,
		Message:    errMsg,
		Response:   response,
	}
	c.JSON(code, Response)
}
