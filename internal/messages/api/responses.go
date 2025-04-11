package messagesapi

import (
	"github.com/kimanimichael/mk-device-manager/internal/messages"
	"time"
)

type MessageResponse struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UID       string    `json:"uid"`
}

func messageToMessageResponse(message messages.Message) MessageResponse {
	return MessageResponse{
		ID:        message.ID,
		CreatedAt: message.CreatedAt,
		UID:       message.DeviceUID,
	}
}
