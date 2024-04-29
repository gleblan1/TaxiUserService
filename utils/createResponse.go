package utils

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message,omitempty"`
	Response   interface{} `json:"response,omitempty"`
}

func DefineResponse(c *gin.Context, code int, err error, response ...interface{}) {
	var resp Response
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	} else {
		errMsg = ""
	}
	resp = Response{
		StatusCode: code,
		Message:    errMsg,
		Response:   response,
	}
	c.JSON(code, resp)
}
