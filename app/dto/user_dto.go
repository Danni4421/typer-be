package dto

type RegisterUser struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	Name     string `json:"name" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=50"`
}

func (r *RegisterUser) ErrorMessages() map[string]string {
	return map[string]string{
		"Username.required": "Username is required",
		"Username.min":      "Username must be at least 3 characters",
		"Username.max":      "Username cannot exceed 20 characters",
		"Name.required":     "Name is required",
		"Name.min":          "Name must be at least 3 characters",
		"Name.max":          "Name cannot exceed 50 characters",
		"Email.required":    "Email is required",
		"Email.email":       "Email must be a valid email address",
		"Password.required": "Password is required",
		"Password.min":      "Password must be at least 6 characters",
		"Password.max":      "Password cannot exceed 50 characters",
	}
}

type RegisterUserResponse struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}
