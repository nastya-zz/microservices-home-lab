package access

import (
	"auth/internal/utils"
	desc "auth/pkg/access_v1"
	"context"
	"errors"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
	"strings"
)

func (i *Implementation) Check(ctx context.Context, req *desc.CheckRequest) (*emptypb.Empty, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("metadata is not provided")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return nil, errors.New("authorization header is not provided")
	}

	if !strings.HasPrefix(authHeader[0], utils.AuthPrefix) {
		return nil, errors.New("invalid authorization header format")
	}

	accessToken := strings.TrimPrefix(authHeader[0], utils.AuthPrefix)

	claims, err := utils.VerifyToken(accessToken, []byte(utils.AccessTokenSecretKey))
	if err != nil {
		return nil, errors.New("access token is invalid")
	}

	ok, err = i.accessService.Check(ctx, req.GetEndpointAddress(), claims)
	if !ok || err != nil {
		return nil, errors.New("access DENIED")
	}

	return &emptypb.Empty{}, nil
}
