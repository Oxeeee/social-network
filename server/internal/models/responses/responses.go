package responses

// Response представляет стандартный формат ответа API
type Response[Data any] struct {
	// Сообщение об успешном выполнении операции | ЕСЛИ ОПЕРАЦИЯ НЕ ВЫПОЛНЕНА, ТО НЕ ВОЗВРАЩАЕТСЯ
	Message string `json:"message,omitempty" example:"user registered successfully"`
	// Сообщение об ошибке | ЕСЛИ ОШИБКИ НЕТУ, ТО НЕ ВОЗВРАЩАЕТСЯ
	Error string `json:"error,omitempty" example:"USERNAME_ALREADY_TAKEN"`
	// Дополнительные данные ответа | ЕСЛИ ДАННЫХ НЕТ, ТО НЕ ВОЗВРАЩАЕТСЯ
	Data Data `json:"data,omitempty"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"-"`
	Username     string `json:"username"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	Photo        string `json:"photo,omitempty"`
}
