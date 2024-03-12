package response

type MessageResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func NewMessageResponse(success bool, message string) *MessageResponse {
	return &MessageResponse{success, message}
}
