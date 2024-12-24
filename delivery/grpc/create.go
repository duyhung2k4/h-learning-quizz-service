package grpchandle

import (
	"app/generated/grpc/servicegrpc"
	"app/generated/grpc/sharedgrpc"
	constant "app/internal/constants"
	"app/internal/entity"
	logapp "app/pkg/log"
	"context"
)

func (h *grpcHandle) CreateQuizz(ctx context.Context, req *servicegrpc.CreateQuizzRequest) (*sharedgrpc.ID, error) {

	newQuizz, err := h.service.QueryQuizzService.Create(entity.Quizz{
		Ask:    req.Quizz.Ask,
		Result: req.Quizz.Result,
		Option: req.Quizz.Option,
		Time:   int(req.Quizz.Time),

		EntityType: entity.ENTITY_TYPE(req.Quizz.EntityType.String()),
		EntityId:   uint(req.Quizz.EntityId),
	})

	if err != nil {
		logapp.Logger("create-quizz", err.Error(), constant.ERROR_LOG)
		return nil, err
	}

	return &sharedgrpc.ID{
		Id: uint64(newQuizz.ID),
	}, nil
}
