package exceptions

import "fmt"

type ServerError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewServerError(code int, message string) *ServerError {
	return &ServerError{
		Code:    code,
		Message: message,
	}
}

func (e *ServerError) Error() string {
	return e.Message
}

func (e *ServerError) StatusCode() int {
	return e.Code
}

func (e *ServerError) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"code":    e.Code,
		"message": e.Message,
	}
}

func (e *ServerError) ToJSON() string {
	return `{"code": ` + fmt.Sprint(e.Code) + `, "message": "` + e.Message + `"}`
}