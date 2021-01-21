package types

// MessageResponse struct
type MessageResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// SendMessage struct
type SendMessage struct {
	From     string `json:"from" validate:"required"`
	Message  string `json:"message" validate:"required"`
	ToNumber string `json:"toNumber" validate:"required"`
}
