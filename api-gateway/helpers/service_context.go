package helpers

import (
	"context"
	"time"

	grpcMetadata "google.golang.org/grpc/metadata"
)

func NewServiceContext() (context.Context, context.CancelFunc, error) {
	token, err := SignJWTForGRPC()
	if err != nil {
		return nil, nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ctxWithAuth := grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)

	return ctxWithAuth, cancel, nil
}
