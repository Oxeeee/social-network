package responses

type Response struct {
	Message string         `json:"message,omitempty"`
	Error   error          `json:"error,omitempty"`
	Details map[string]any `json:"details,omitempty"`
}
