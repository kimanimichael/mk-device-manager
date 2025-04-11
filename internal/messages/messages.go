package messages

import (
	"context"
	"encoding/json"
	"time"
)

type Message struct {
	ID        string
	CreatedAt time.Time
	Payload   json.RawMessage
	DeviceUID string
}

type page struct {
	Offset uint32
	Total  uint32
}

type MessagePage struct {
	page
	Messages []Message
}

type MessageService interface {
	CreateMessage(ctx context.Context, msg *Message) (*Message, error)
	GetMessageByID(ctx context.Context, ID string) (*Message, error)
	GetMessagesByUID(ctx context.Context, UID string, offset, limit uint32) (*MessagePage, error)
	GetAllMessages(ctx context.Context, offset uint32, limit uint32) (*MessagePage, error)
	//DeleteMessage(ctx context.Context, ID string) error
}

type MessageRepository interface {
	CreateMessage(ctx context.Context, msg *Message) (*Message, error)
	GetMessageByID(ctx context.Context, ID string) (*Message, error)
	GetMessagesByUID(ctx context.Context, UID string, offset, limit uint32) (*MessagePage, error)
	GetAllMessages(ctx context.Context, offset uint32, limit uint32) (*MessagePage, error)
	//DeleteMessage(ctx context.Context, ID string) error
}
