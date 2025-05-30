package exceptions

type ValidationError struct {
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors"`
}

func (e ValidationError) Error() string {
	return e.Message
}

func (e ValidationError) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"message": e.Message,
		"errors":  e.Errors,
	}
}
