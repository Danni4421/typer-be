package services

import (
	"database/sql"
	"fmt"
	"strings"
	"typer/app/dto"
	"typer/package/exceptions"
)

type TypingService struct {
	DB *sql.DB
}

func (s *TypingService) StoreTypingTestLog(userID uint, langID uint, calculation *dto.TypeCalculation) error {
	if calculation == nil {
		return &exceptions.ClientError{
			Code:    400,
			Message: "Calculation data cannot be nil",
		}
	}

	query := `INSERT INTO typing_logs (user_id, language_id, word_count, character_count, accuracy, wpm)
			  VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := s.DB.Exec(query, userID, langID, calculation.WordsCount, calculation.CharacterCount, calculation.Accuracy, calculation.WPM)

	if err != nil {
		fmt.Println("Error inserting typing test log:", err)
		return &exceptions.ServerError{
			Code:    500,
			Message: "Failed to store typing test log",
		}
	}

	return nil
}

func (s *TypingService) CalculateWPM(text string, failedText string) (*dto.TypeCalculation, error) {
	wordCount := len(strings.Split(text, " "))

	// Remove all whitespace characters to count characters accurately
	removedWhitespace := strings.ReplaceAll(text, " ", "")
	characterCount := len(removedWhitespace)

	// Remove all whitespace characters from failedText to count characters accurately
	removedWhitespaceFailed := strings.ReplaceAll(failedText, " ", "")
	failedCharacterCount := len(removedWhitespaceFailed)

	if characterCount == 0 {
		return nil, &exceptions.ClientError{
			Code:    400,
			Message: "Text cannot be empty",
		}
	}

	// Calculate WPM (Words Per Minute)
	// We assume that the test was completed
	// For now the supported WPM calculation is only for 1 minute
	wpm := int16(float64(characterCount) / 5.0)
	accuracy := float32(characterCount-failedCharacterCount) / float32(characterCount) * 100

	if accuracy < 0 {
		accuracy = 0
	}

	return &dto.TypeCalculation{
		WordsCount:     int16(wordCount),
		CharacterCount: int16(characterCount),
		Accuracy:       accuracy,
		WPM:            wpm,
	}, nil
}
