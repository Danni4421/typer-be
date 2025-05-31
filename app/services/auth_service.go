package services

import (
	"database/sql"
	"log"
	"typer/app/models"
	"typer/package/exceptions"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	DB *sql.DB
}

func (s *AuthService) ValidateCredentials(email, password string) (int, error) {

	user := new(models.User)

	// Check if user exists
	if err := s.DB.QueryRow("SELECT id, email, password FROM users WHERE email = $1", email).Scan(&user.ID, &user.Email, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return 0, &exceptions.ClientError{
				Code:    404,
				Message: "User not found",
			}
		}

		return 0, &exceptions.ClientError{
			Code:    500,
			Message: "Internal server error",
		}
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return 0, &exceptions.ClientError{
			Code:    401,
			Message: "Invalid credentials",
		}
	}

	return user.ID, nil
}

func (s *AuthService) Logout(userID int) error {
	log.Println("Logouting user sessions with id:", userID)

	deleteQuery := `DELETE FROM sessions WHERE user_id = $1`
	_, err := s.DB.Exec(deleteQuery, userID)
	if err != nil {
		return &exceptions.ServerError{
			Code:    500,
			Message: "Failed to invalidate tokens",
		}
	}

	return nil
}
