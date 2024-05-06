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
	var responseData interface{}
	if len(response) > 0 {
		responseData = response[0]
	} else {
		responseData = nil
	}

	resp = Response{
		StatusCode: code,
		Message:    errMsg,
		Response:   responseData,
	}
	c.AbortWithStatusJSON(code, resp)
}
