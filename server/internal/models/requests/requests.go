package requests

// Register представляет данные для регистрации нового пользователя
type Register struct {
	// Email пользователя, должен быть уникальным
	Email string `json:"email" validate:"required,email" example:"user@example.com"`
	// Имя пользователя
	Name string `json:"name" validate:"required" example:"Иван"`
	// Фамилия пользователя
	Surname string `json:"surname" validate:"required" example:"Иванов"`
	// Имя пользователя, должно быть уникальным
	Username string `json:"username" validate:"required" example:"ivan2024"`
	// Пароль пользователя
	Password string `json:"password" validate:"required,min=8" example:"Password123!"`
	// Фото юзера (base64)
	PhotoEncrypted string `json:"photo" example:"base64url"`
}

// Login представляет данные для входа пользователя
type Login struct {
	// Email пользователя
	Email string `json:"email" validate:"required,email" example:"user@example.com"`
	// Пароль пользователя
	Password string `json:"password" validate:"required" example:"Password123!"`
}
