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

func (s *messageService) GetMessagesByUID(ctx context.Context, UID string) ([]Message, error) {
	messages, err := s.repo.GetMessagesByUID(ctx, UID)
	if err != nil {
		return nil, err
	}
	return messages, nil
}
