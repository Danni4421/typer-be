package dto

type StoreWords struct {
	Words []string `json:"words" validate:"required,min=1,max=1000,dive,required,min=1,max=50"`
}

func (dto StoreWords) ErrorMessages() map[string]string {
	return map[string]string{
		"Words.required": "Words array is required",
		"Words.min":      "At least 1 word is required",
		"Words.max":      "Maximum 1000 words allowed",
		"Words.dive":     "Each word must be valid",
	}
}
