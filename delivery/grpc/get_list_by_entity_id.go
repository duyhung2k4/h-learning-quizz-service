package grpchandle

import (
	"app/generated/grpc/servicegrpc"
	"app/generated/grpc/sharedgrpc"
	"context"
)

func (h *grpcHandle) GetListByEntityId(ctx context.Context, req *sharedgrpc.ID) (*servicegrpc.GetListByEntityIdResponse, error) {
	return nil, nil
}
