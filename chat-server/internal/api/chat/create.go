package chat

import (
	"chat-server/internal/converter"
	desc "chat-server/pkg/chat_v1"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (c *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Received user CreateRequest %+v", req.GetChat())

	createChat := converter.ToCreateChatFromDesc(req.GetChat())
	id, err := c.chatService.Create(ctx, createChat)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Chat creation failed")
	}

	return &desc.CreateResponse{Id: id}, nil
}
