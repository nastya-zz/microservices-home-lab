package service

import (
	"chat-server/internal/model"
	"context"
)

type ChatService interface {
	SendMessage(ctx context.Context, message *model.CreateMessage) error
	Create(ctx context.Context, chat *model.CreateChat) (int64, error)
	Delete(ctx context.Context, id int64) error
}
