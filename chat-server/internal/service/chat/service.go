package chat

import (
	"chat-server/internal/client/db"
	"chat-server/internal/model"
	"chat-server/internal/repository"
	"chat-server/internal/service"
	"context"
)

type serv struct {
	chatRepository repository.ChatRepository
	txManager      db.TxManager
}

func NewService(
	chatRepository repository.ChatRepository,
	txManager db.TxManager,
) service.ChatService {
	return &serv{
		chatRepository: chatRepository,
		txManager:      txManager,
	}
}

func (s serv) SendMessage(ctx context.Context, message *model.CreateMessage) error {
	err := s.chatRepository.SendMessage(ctx, message)
	if err != nil {
		return err
	}

	return nil
}

func (s serv) Create(ctx context.Context, chat *model.CreateChat) (int64, error) {
	id, err := s.chatRepository.Create(ctx, chat)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s serv) Delete(ctx context.Context, id int64) error {
	err := s.chatRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
