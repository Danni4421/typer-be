package services

import (
	"database/sql"
	"log"
	"strings"
	"typer/package/exceptions"
)

type WordService struct {
	DB *sql.DB
}

func (s *WordService) StoreWords(words []string, languageID int) error {
	query := `INSERT INTO words (word, language_id, length) VALUES ($1, $2, $3)`

	for _, word := range words {
		// Converting words to lowercase and calculating their length
		wordLength := len(word)
		wordLower := strings.ToLower(word)

		_, err := s.DB.Exec(query, wordLower, languageID, wordLength)
		if err != nil {
			log.Println("Error inserting word:", word, "Error:", err)
			return &exceptions.ServerError{
				Code:    500,
				Message: "Failed to create words",
			}
		}
	}
	return nil
}

func (s *WordService) GetWordsByLanguage(languageID int) ([]string, error) {
	query := `SELECT word FROM words WHERE language_id = $1`
	rows, err := s.DB.Query(query, languageID)
	if err != nil {
		return nil, &exceptions.ServerError{
			Code:    500,
			Message: "Failed to retrieve words",
		}
	}
	defer rows.Close()

	var words []string
	for rows.Next() {
		var word string
		if err := rows.Scan(&word); err != nil {
			return nil, &exceptions.ServerError{
				Code:    500,
				Message: "Failed to scan word",
			}
		}
		words = append(words, word)
	}

	if err := rows.Err(); err != nil {
		return nil, &exceptions.ServerError{
			Code:    500,
			Message: "Error occurred while processing words",
		}
	}

	return words, nil
}

func (s *WordService) GetRandomWords(id int, limit int) ([]string, error) {
	query := `SELECT word FROM words WHERE language_id = $1 ORDER BY RANDOM() LIMIT $2`
	rows, err := s.DB.Query(query, id, limit)
	if err != nil {
		return nil, &exceptions.ServerError{
			Code:    500,
			Message: "Failed to retrieve random words",
		}
	}
	defer rows.Close()

	var words []string
	for rows.Next() {
		var word string
		if err := rows.Scan(&word); err != nil {
			return nil, &exceptions.ServerError{
				Code:    500,
				Message: "Failed to scan word",
			}
		}
		words = append(words, word)
	}

	if err := rows.Err(); err != nil {
		return nil, &exceptions.ServerError{
			Code:    500,
			Message: "Error occurred while processing random words",
		}
	}

	return words, nil
}
