package auth

import (
	desc "chat-server/pkg/access_v1"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

const servicePort = 50051

type Client struct {
	ctx context.Context
	md  metadata.MD
}

func AccessCheck(ctx context.Context, token string, url string) error {
	md := metadata.Pairs("authorization", token)
	ctx = metadata.NewOutgoingContext(ctx, md)

	conn, err := grpc.Dial(
		fmt.Sprintf(":%d", servicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return fmt.Errorf("failed to dial GRPC client: %w", err)
	}
	defer func() { _ = conn.Close() }()

	cl := desc.NewAccessV1Client(conn)

	_, err = cl.Check(ctx, &desc.CheckRequest{
		EndpointAddress: url,
	})

	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	return nil
}
