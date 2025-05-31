package services

import (
	"database/sql"
	"typer/app/models"
	"typer/package/exceptions"

	"github.com/lib/pq"
)

type LanguageService struct {
	DB *sql.DB
}

func (s *LanguageService) CreateLanguage(name string, code string) error {
	query := `INSERT INTO languages (name, code) VALUES ($1, $2)`

	_, err := s.DB.Exec(query, name, code)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return &exceptions.ClientError{
				Code:    409,
				Message: "Language name or code already exists",
			}
		}

		return &exceptions.ServerError{
			Code:    500,
			Message: "Failed to create language",
		}
	}
	return nil
}

func (s *LanguageService) GetAllLanguages() ([]models.Language, error) {
	languages := []models.Language{}

	query := `SELECT id, name, code FROM languages`
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, &exceptions.ServerError{
			Code:    500,
			Message: "Failed to retrieve languages",
		}
	}
	defer rows.Close()

	for rows.Next() {
		var lang models.Language
		if err := rows.Scan(&lang.ID, &lang.Name, &lang.Code); err != nil {
			return nil, &exceptions.ServerError{
				Code:    500,
				Message: "Failed to scan language name",
			}
		}
		languages = append(languages, lang)
	}

	if err := rows.Err(); err != nil {
		return nil, &exceptions.ServerError{
			Code:    500,
			Message: "Error occurred while processing languages",
		}
	}

	return languages, nil
}

func (s *LanguageService) GetLanguageByName(name string) (*models.Language, error) {
	var language models.Language

	query := `SELECT id, name, code FROM languages WHERE name = $1`
	err := s.DB.QueryRow(query, name).Scan(&language.ID, &language.Name, &language.Code)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &exceptions.ClientError{
				Code:    404,
				Message: "Language not found",
			}
		}
		return nil, &exceptions.ServerError{
			Code:    500,
			Message: "Failed to retrieve language",
		}
	}
	return &language, nil
}

func (s *LanguageService) GetLanguageByCode(code string) (*models.Language, error) {
	query := `SELECT id, name, code FROM languages WHERE code = $1`
	var language models.Language
	err := s.DB.QueryRow(query, code).Scan(&language.ID, &language.Name, &language.Code)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &exceptions.ClientError{
				Code:    404,
				Message: "Language not found",
			}
		}
		return nil, &exceptions.ServerError{
			Code:    500,
			Message: "Failed to retrieve language",
		}
	}
	return &language, nil
}

func (s *LanguageService) DeleteLanguageByCode(code string) error {
	query := `DELETE FROM languages WHERE code = $1`
	result, err := s.DB.Exec(query, code)
	if err != nil {
		return &exceptions.ServerError{
			Code:    500,
			Message: "Failed to delete language",
		}
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return &exceptions.ServerError{
			Code:    500,
			Message: "Failed to check rows affected",
		}
	}

	if rowsAffected == 0 {
		return &exceptions.ClientError{
			Code:    404,
			Message: "Language not found",
		}
	}

	return nil
}
