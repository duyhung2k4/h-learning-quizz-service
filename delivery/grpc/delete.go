package grpchandle

import (
	"app/generated/grpc/sharedgrpc"
	"context"
)

func (h *grpcHandle) DeleteById(ctx context.Context, req *sharedgrpc.ID) (*sharedgrpc.Empty, error) {
	return nil, nil
}
