package services

import (
	"database/sql"
	"time"
	"typer/app/dto"
	"typer/package/exceptions"
	"typer/package/utils"
)

type JWTService struct {
	DB *sql.DB
}

func (s *JWTService) GenerateTokens(id string) (*dto.LoggedInUser, error) {
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

	return &dto.LoggedInUser{
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}
