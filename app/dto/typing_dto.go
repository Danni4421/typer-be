package dto

type StoreUserLogDto struct {
	Calculation TypeCalculation `json:"calculation" validate:"required"`
}

type TypingCalculationDto struct {
	Text       string `json:"text" validate:"required,min=1"`
	FailedText string `json:"failed_text" validate:"required,min=1"`
}

func (dto TypingCalculationDto) ErrorMessages() map[string]string {
	return map[string]string{
		"Text.required":       "Text is required",
		"Text.min":            "Text must have at least 1 character",
		"FailedText.required": "Failed text is required",
		"FailedText.min":      "Failed text must have at least 1 character",
	}
}

type TypeCalculation struct {
	WordsCount     int16   `json:"words_count" validate:"required,min=1,max=1000"`
	CharacterCount int16   `json:"character_count" validate:"required,min=1,max=10000"`
	Accuracy       float32 `json:"accuracy" validate:"required,min=0,max=100"`
	WPM            int16   `json:"wpm" validate:"required,min=0"`
}

func (dto StoreUserLogDto) ErrorMessages() map[string]string {
	return map[string]string{
		"Calculation.required":                "Calculation data is required",
		"Calculation.WordsCount.required":     "Words count is required",
		"Calculation.WordsCount.min":          "At least 1 word is required",
		"Calculation.WordsCount.max":          "Maximum 1000 words allowed",
		"Calculation.CharacterCount.required": "Character count is required",
		"Calculation.CharacterCount.min":      "At least 1 character is required",
		"Calculation.CharacterCount.max":      "Maximum 10000 characters allowed",
		"Calculation.Accuracy.required":       "Accuracy is required",
		"Calculation.Accuracy.min":            "Accuracy must be at least 0%",
		"Calculation.Accuracy.max":            "Accuracy cannot exceed 100%",
		"Calculation.WPM.required":            "WPM is required",
		"Calculation.WPM.min":                 "WPM cannot be negative",
	}
}
