package repository

import (
	"chat-server/internal/model"
	"context"
)

type ChatRepository interface {
	SendMessage(ctx context.Context, message *model.CreateMessage) error
	ListMessages(ctx context.Context, chatId int64) ([]*model.Message, error)
	Create(ctx context.Context, chat *model.CreateChat) (int64, error)
	Delete(ctx context.Context, id int64) error
}
