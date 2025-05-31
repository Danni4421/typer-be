package services

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	"typer/app/dto"
	"typer/package/exceptions"
	"typer/package/utils"
)

type JWTService struct {
	DB *sql.DB
}

func (s *JWTService) GenerateTokens(id string) (*dto.LoggedInUser, error) {
	fmt.Println("Generating tokens for user ID:", id)

	token, err := utils.GenerateJWT(id, time.Hour)
	if err != nil {
		return nil, &exceptions.ServerError{
			Code:    500,
			Message: "Failed to generate your access token",
		}
	}

	refreshToken, err := utils.GenerateJWT(id, time.Hour*24*30)
	if err != nil {
		return nil, &exceptions.ServerError{
			Code:    500,
			Message: "Failed to generate your refresh token",
		}
	}

	deleteQuery := `DELETE FROM sessions WHERE user_id = $1`
	_, err = s.DB.Exec(deleteQuery, id)
	if err != nil {
		return nil, &exceptions.ServerError{
			Code:    500,
			Message: "Failed to clean up existing sessions",
		}
	}

	storeQuery := `INSERT INTO sessions (user_id, token) VALUES ($1, $2)`
	_, err = s.DB.Exec(storeQuery, id, refreshToken)

	if err != nil {
		return nil, &exceptions.ServerError{
			Code:    500,
			Message: "Failed to store your refresh token",
		}
	}

	return &dto.LoggedInUser{
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}

func (s *JWTService) ValidateRefreshToken(token string) (string, error) {
	log.Println("Validating refresh token:", token)

	userID, err := utils.ParseJWT(token)
	if err != nil {
		return "", &exceptions.ClientError{
			Code:    401,
			Message: "Invalid refresh token",
		}
	}

	var exists bool
	err = s.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM sessions WHERE user_id = $1 AND token = $2)", userID, token).Scan(&exists)
	if err != nil {
		return "", &exceptions.ServerError{
			Code:    500,
			Message: "Failed to validate refresh token",
		}
	}

	if !exists {
		return "", &exceptions.ClientError{
			Code:    401,
			Message: "Refresh token not found",
		}
	}

	return userID, nil
}

func (s *JWTService) RenewTokens(userID string) (*dto.LoggedInUser, error) {
	log.Println("Renewing tokens for user ID:", userID)

	newToken, err := utils.GenerateJWT(userID, time.Hour)
	if err != nil {
		return nil, &exceptions.ServerError{
			Code:    500,
			Message: "Failed to generate new access token",
		}
	}

	newRefreshToken, err := utils.GenerateJWT(userID, time.Hour*24*30)
	if err != nil {
		return nil, &exceptions.ServerError{
			Code:    500,
			Message: "Failed to generate new refresh token",
		}
	}

	updateQuery := `UPDATE sessions SET token = $1 WHERE user_id = $2`
	_, err = s.DB.Exec(updateQuery, newRefreshToken, userID)
	if err != nil {
		return nil, &exceptions.ServerError{
			Code:    500,
			Message: "Failed to update refresh token",
		}
	}

	return &dto.LoggedInUser{
		Token:        newToken,
		RefreshToken: newRefreshToken,
	}, nil
}
