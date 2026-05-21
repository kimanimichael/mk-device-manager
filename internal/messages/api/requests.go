package messagesapi

type CreateMessageRequest struct {
	EventType string `json:"event_type"`
	UID       string `json:"uid"`
}
