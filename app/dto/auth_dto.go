package dto

type LoginUser struct {
	Email    string `json:"email" validate:"required,min=3,max=50"`
	Password string `json:"password" validate:"required,min=6,max=50"`
}

func (l *LoginUser) ErrorMessages() map[string]string {
	return map[string]string{
		"Email.required":    "Email is required",
		"Email.min":         "Email must be at least 3 characters",
		"Email.max":         "Email cannot exceed 50 characters",
		"Password.required": "Password is required",
		"Password.min":      "Password must be at least 6 characters",
		"Password.max":      "Password cannot exceed 50 characters",
	}
}

type LoggedInUser struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

func (r *RefreshTokenRequest) ErrorMessages() map[string]string {
	return map[string]string{
		"RefreshToken.required": "Refresh token is required",
	}
}
