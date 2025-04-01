package converter

import (
	"chat-server/internal/model"
	desc "chat-server/pkg/chat_v1"
)

func ToCreateChatFromDesc(chat *desc.Chat) *model.CreateChat {
	return &model.CreateChat{
		Usernames: chat.Usernames,
	}
}

func ToMessageFromDesc(message *desc.Message) *model.CreateMessage {
	return &model.CreateMessage{
		Text:      message.Text,
		From:      message.From,
		Timestamp: message.Timestamp.AsTime(),
		ChatID:    message.ChatId,
	}
}
