package messages

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	sqlcdatabase "github.com/kimanimichael/mk-device-manager/internal/adapters/database/sqlc/gensql"
	"time"
)

type MessageRepositorySQL struct {
	DB *sqlcdatabase.Queries
}

var _ MessageRepository = (*MessageRepositorySQL)(nil)

func NewMessageRepositorySQL(db *sqlcdatabase.Queries) *MessageRepositorySQL {
	return &MessageRepositorySQL{
		DB: db,
	}
}

func (r *MessageRepositorySQL) CreateMessage(ctx context.Context, msg *Message) (*Message, error) {
	message, err := r.DB.CreateMessage(ctx, sqlcdatabase.CreateMessageParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		Payload:   msg.Payload,
		DeviceUid: msg.DeviceUID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create device: %v", err)
	}
	return &Message{
		message.ID.String(),
		message.CreatedAt,
		message.Payload,
		message.DeviceUid,
	}, nil
}

func (r *MessageRepositorySQL) GetMessageByID(ctx context.Context, ID string) (*Message, error) {
	messageID, err := uuid.Parse(ID)
	if err != nil {
		return nil, fmt.Errorf("failed to parse message ID: %v", err)
	}
	message, err := r.DB.GetMessageByID(ctx, messageID)
	if err != nil {
		return nil, fmt.Errorf("failed to create device: %v", err)
	}
	return &Message{
		message.ID.String(),
		message.CreatedAt,
		message.Payload,
		message.DeviceUid,
	}, nil
}
