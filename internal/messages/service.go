package messages

import "context"

type messageService struct {
	repo MessageRepository
}

func NewMessageService(repo MessageRepository) MessageService {
	return &messageService{
		repo: repo,
	}
}

func (s *messageService) CreateMessage(ctx context.Context, msg *Message) (*Message, error) {
	message, err := s.repo.CreateMessage(ctx, msg)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (s *messageService) GetMessageByID(ctx context.Context, ID string) (*Message, error) {
	message, err := s.repo.GetMessageByID(ctx, ID)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (s *messageService) GetMessagesByUID(ctx context.Context, UID string, offset, limit uint32) (*MessagePage, error) {
	messages, err := s.repo.GetMessagesByUID(ctx, UID, offset, limit)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (s *messageService) GetAllMessages(ctx context.Context, offset, limit uint32) (*MessagePage, error) {
	messages, err := s.repo.GetAllMessages(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	return messages, nil
}
