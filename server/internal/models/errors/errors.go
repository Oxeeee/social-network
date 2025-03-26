package cerrors

import "errors"

var (
	// ErrUsernameTaken возникает, когда запрашиваемое имя пользователя уже занято
	ErrUsernameTaken = errors.New("USERNAME_ALREADY_TAKEN")
	// ErrEmailTaken возникает, когда запрашиваемый email уже используется
	ErrEmailTaken = errors.New("EMAIL_ALREADY_TAKEN")
	// ErrInvalidEmail возникает, когда указан несуществующий email
	ErrInvalidEmail = errors.New("INVALID_EMAIL")
	// ErrInvalidPassword возникает, когда указан неверный пароль
	ErrInvalidPassword = errors.New("INVALID_PASSWORD")
	// ErrMissingToken возникает, когда отсутствует токен авторизации
	ErrMissingToken = errors.New("MISSING_AUTHORIZATION_TOKEN")
	// ErrInvalidAuthHeaderFormat возникает, когда неверный формат заголовка авторизации
	ErrInvalidAuthHeaderFormat = errors.New("INVALID_AUTHORIZATION_HEADER_FORMAT")
	// ErrUnexpectedSigningMethod возникает, когда используется неожиданный метод подписи
	ErrUnexpectedSigningMethod = errors.New("UNEXPECTED_SIGNING_METHOD")
	// ErrInvalidExpToken возникает, когда токен недействителен или истек
	ErrInvalidExpToken = errors.New("INVALID_OR_EXPIRED_TOKEN")
	// ErrExpToken возникает, когда истек срок действия access токена
	ErrExpToken = errors.New("EXPIRED_ACCESS_TOKEN")
	// ErrInvalidPayload возникает, когда недействителен payload токена
	ErrInvalidPayload = errors.New("EXPIRED_TOKEN_PAYLOAD")
)
