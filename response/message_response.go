package response

// MessageResponse represents the structure of a response message, typically used in API responses.
type MessageResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

// NewMessageResponse creates a new MessageResponse instance with the given success status and message.
func NewMessageResponse(success bool, message string) *MessageResponse {
	return &MessageResponse{success, message}
}
