package auth

import "errors"

var (
	ErrUserNotFound          = errors.New("user not found")
	ErrInvalidPassword       = errors.New("invalid password")
	ErrEmailVerified         = errors.New("email already verified")
	ErrUserAlreadyRegistered = errors.New("there is already a user with this email.")
	ErrReciveEmail           = errors.New("Recive mail error.")
)
