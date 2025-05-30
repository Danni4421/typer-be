package exceptions

import "fmt"

type ClientError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewClientError(code int, message string) *ClientError {
	return &ClientError{
		Code:    code,
		Message: message,
	}
}

func (e *ClientError) Error() string {
	return e.Message
}

func (e *ClientError) StatusCode() int {
	return e.Code
}

func (e *ClientError) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"code":    e.Code,
		"message": e.Message,
	}
}

func (e *ClientError) ToJSON() string {
	return `{"code": ` + fmt.Sprint(e.Code) + `, "message": "` + e.Message + `"}`
}