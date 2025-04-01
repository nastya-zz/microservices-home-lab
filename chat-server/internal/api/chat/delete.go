package chat

import (
	desc "chat-server/pkg/chat_v1"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (c *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Received chat DeleteRequest %+v", req.GetId())

	err := c.chatService.Delete(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Chat delete failed")
	}

	return &emptypb.Empty{}, nil
}
