package main

import (
	"auth/internal/app"
	desc "auth/pkg/user_v1"
	"context"
	"log"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedUserV1Server
}

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
