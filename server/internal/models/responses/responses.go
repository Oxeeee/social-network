package responses

// Response представляет стандартный формат ответа API
type Response struct {
	// Сообщение об успешном выполнении операции | ЕСЛИ ОПЕРАЦИЯ НЕ ВЫПОЛНЕНА, ТО НЕ ВОЗВРАЩАЕТСЯ
	Message string `json:"message,omitempty" example:"user registered successfully"`
	// Сообщение об ошибке | ЕСЛИ ОШИБКИ НЕТУ, ТО НЕ ВОЗВРАЩАЕТСЯ
	Error string `json:"error,omitempty" example:"USERNAME_ALREADY_TAKEN"`
	// Дополнительные данные ответа | ЕСЛИ ДАННЫХ НЕТ, ТО НЕ ВОЗВРАЩАЕТСЯ
	Details map[string]interface{} `json:"details,omitempty"`
}
