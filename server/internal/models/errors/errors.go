package cerrors

import "errors"

var (
	ErrUsernameTaken = errors.New("USERNAME_ALREADY_TAKEN")
	ErrEmailTaken    = errors.New("EMAIL_ALREADY_TAKEN")
)
