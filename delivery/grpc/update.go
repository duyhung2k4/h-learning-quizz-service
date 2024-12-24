package grpchandle

import (
	"app/generated/grpc/servicegrpc"
	"app/generated/grpc/sharedgrpc"
	"context"
)

func (h *grpcHandle) UpdateQuizz(ctx context.Context, req *servicegrpc.UpdateQuizzRequest) (*sharedgrpc.Empty, error) {
	return nil, nil
}
