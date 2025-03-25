package cerrors

import "errors"

var (
	ErrUsernameTaken           = errors.New("USERNAME_ALREADY_TAKEN")
	ErrEmailTaken              = errors.New("EMAIL_ALREADY_TAKEN")
	ErrInvalidEmail            = errors.New("INVALID_EMAIL")
	ErrInvalidPassword         = errors.New("INVALID_PASSWORD")
	ErrMissingToken            = errors.New("MISSING_AUTHORIZATION_TOKEN")
	ErrInvalidAuthHeaderFormat = errors.New("INVALID_AUTHORIZATION_HEADER_FORMAT")
	ErrUnexpectedSigningMethod = errors.New("UNEXPECTED_SIGNING_METHOD")
	ErrInvalidExpToken         = errors.New("INVALID_OR_EXPIRED_TOKEN")
	ErrExpToken                = errors.New("EXPIRED_ACCESS_TOKEN")
	ErrInvalidPayload          = errors.New("EXPIRED_TOKEN_PAYLOAD")
)
