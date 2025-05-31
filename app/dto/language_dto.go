package dto

type CreateLanguage struct {
	Code string `json:"code" validate:"required,min=2,max=4,alphanum"`
	Name string `json:"name" validate:"required,min=1,max=100"`
}
