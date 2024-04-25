package models

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message,omitempty"`
	Response   interface{} `json:"response,omitempty"`
}
