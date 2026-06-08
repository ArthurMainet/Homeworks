package auth

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"required"`
}

type PhoneLoginRequest struct {
	Phone string `json:"phone" validate:"required"`
}
