package messagesapi

type CreateMessageRequest struct {
	EventType string `json:"event_type"`
	UID       string `json:"uid"`
}

type GetMessageByIDRequest struct {
	ID        string `json:"id"`
	EventType string `json:"event_type"`
}

type GetMessagesRequest struct {
	Offset uint32 `json:"offset"`
	Limit  uint32 `json:"limit"`
}
