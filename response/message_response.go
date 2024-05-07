package response

type MessageResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

func NewMessageResponse(success bool, message string) *MessageResponse {
	return &MessageResponse{success, message}
}
