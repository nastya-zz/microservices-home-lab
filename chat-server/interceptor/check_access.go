package interceptor

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"

	auth "chat-server/internal/client/auth"
)

func AccessInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Получаем метаданные из контекста
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	authHeader, ok := md["authorization"]

	if !ok || len(authHeader) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	token := authHeader[0]

	fullMethod := info.FullMethod

	err := auth.AccessCheck(ctx, token, fullMethod)

	log.Println("AccessInterceptor", fmt.Errorf("%w", err))

	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "role not allowed")
	}

	return handler(ctx, req)
}
