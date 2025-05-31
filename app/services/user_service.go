package services

import (
	"database/sql"
	"time"
	"typer/app/dto"
	"typer/app/models"
	"typer/package/exceptions"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	DB *sql.DB
}

func (s *UserService) CreateUser(user *models.User) (*dto.RegisterUserResponse, error) {
	var exists bool
	err := s.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = $1 OR email = $2)", user.Username, user.Email).Scan(&exists)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, &exceptions.ClientError{
			Code:    409,
			Message: "Username or email already exists",
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	var id int
	err = s.DB.QueryRow(
		"INSERT INTO users (username, name, email, password) VALUES ($1, $2, $3, $4) RETURNING id",
		user.Username, user.Name, user.Email, hashedPassword,
	).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &dto.RegisterUserResponse{
		ID:        id,
		Username:  user.Username,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}, nil
}

func (s *UserService) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := s.DB.QueryRow("SELECT id, username, name, email, created_at FROM users WHERE username = $1", username).Scan(&user.ID, &user.Username, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &exceptions.ClientError{
				Code:    404,
				Message: "User not found",
			}
		}
		return nil, err
	}
	return &user, nil
}

func (s *UserService) GetUserByID(userID int) (*models.User, error) {
	var user models.User

	err := s.DB.QueryRow(
		`SELECT id, username, name, email, created_at FROM users WHERE id = $1`,
		userID,
	).Scan(&user.ID, &user.Username, &user.Name, &user.Email, &user.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &exceptions.ClientError{
				Code:    404,
				Message: "User not found",
			}
		}
		return nil, err
	}

	return &user, nil
}
