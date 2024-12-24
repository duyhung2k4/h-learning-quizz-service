package grpchandle

import (
	"app/generated/grpc/servicegrpc"
	"app/generated/grpc/sharedgrpc"
	"context"
)

func (s *grpcHandle) GetById(ctx context.Context, req *sharedgrpc.ID) (*servicegrpc.Quizz, error) {
	return nil, nil
}
