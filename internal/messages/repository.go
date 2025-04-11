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

func (r *MessageRepositorySQL) GetMessagesByUID(ctx context.Context, UID string, offset, limit uint32) (*MessagePage, error) {
	messages, err := r.DB.GetMessagesByDeviceUID(ctx, sqlcdatabase.GetMessagesByDeviceUIDParams{
		DeviceUid: UID,
		Offset:    int32(offset),
		Limit:     int32(limit),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create device: %v", err)
	}
	var messagesToReturn []Message
	for _, message := range messages {
		messagesToReturn = append(messagesToReturn, Message{
			ID:        message.ID.String(),
			CreatedAt: message.CreatedAt,
			Payload:   message.Payload,
			DeviceUID: message.DeviceUid,
		})
	}
	totalMessages, err := r.DB.GetMessagesCount(ctx)
	if err != nil {
		return nil, fmt.Errorf("couldn't get total messages: %v", err)
	}
	messagePage := &MessagePage{
		page{
			Offset: 1,
			Total:  uint32(totalMessages),
		}, messagesToReturn,
	}
	return messagePage, nil
}

func (r *MessageRepositorySQL) GetAllMessages(ctx context.Context, offset, limit uint32) (*MessagePage, error) {
	messages, err := r.DB.GetAllMessages(ctx, sqlcdatabase.GetAllMessagesParams{
		Offset: int32(offset),
		Limit:  int32(limit),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch paged messages: %v", err)
	}
	var messagesToReturn []Message
	for _, message := range messages {
		messagesToReturn = append(messagesToReturn, Message{
			ID:        message.ID.String(),
			CreatedAt: message.CreatedAt,
			Payload:   message.Payload,
			DeviceUID: message.DeviceUid,
		})
	}
	totalMessages, err := r.DB.GetMessagesCount(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch total messages: %v", err)
	}
	pagedMessages := &MessagePage{
		page{
			Offset: offset,
			Total:  uint32(totalMessages),
		},
		messagesToReturn,
	}
	return pagedMessages, nil
}
