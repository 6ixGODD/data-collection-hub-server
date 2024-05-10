package mods

import (
	"context"
	"time"

	dao "data-collection-hub-server/internal/pkg/dal/mods"
	"data-collection-hub-server/internal/pkg/schema/user"
	"data-collection-hub-server/internal/pkg/service"
)

type StatisticService interface {
	GetDataStatistic(ctx context.Context, startDate, endDate *time.Time) (*user.GetDataStatisticResponse, error)
}

type StatisticServiceImpl struct {
	core               *service.Service
	instructionDataDao dao.InstructionDataDao
}

func NewStatisticService(s *service.Service, instructionDataDao dao.InstructionDataDao) StatisticService {
	return &StatisticServiceImpl{
		core:               s,
		instructionDataDao: instructionDataDao,
	}
}

func (s StatisticServiceImpl) GetDataStatistic(
	ctx context.Context, startDate, endDate *time.Time,
) (*user.GetDataStatisticResponse, error) {
	// TODO implement me
	panic("implement me")
}
