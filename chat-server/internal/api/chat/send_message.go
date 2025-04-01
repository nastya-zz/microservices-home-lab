package chat

import (
	"chat-server/internal/converter"
	desc "chat-server/pkg/chat_v1"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (c *Implementation) SendMessage(ctx context.Context, req *desc.SendRequest) (*emptypb.Empty, error) {
	log.Printf("Received message SendMessageRequest %+v", req.GetMessage())

	createMsg := converter.ToMessageFromDesc(req.GetMessage())
	if err := c.chatService.SendMessage(ctx, createMsg); err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"send message failed")
	}

	return &emptypb.Empty{}, nil
}
