package mods

import (
	"context"
	"time"

	dao "data-collection-hub-server/internal/pkg/dal/mods"
	"data-collection-hub-server/internal/pkg/schema/admin"
	"data-collection-hub-server/internal/pkg/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StatisticService interface {
	GetDataStatistic(ctx context.Context, startDate, endDate *time.Time) (*admin.GetDataStatisticResponse, error)
	GetUserStatistic(ctx context.Context, userID *primitive.ObjectID) (*admin.GetUserStatisticResponse, error)
	GetUserStatisticList(
		ctx context.Context, page *int, loginBefore, loginAfter, createdBefore, createdAfter *time.Time,
	) (*admin.GetUserStatisticListResponse, error)
}

type StatisticServiceImpl struct {
	service            *service.Service
	instructionDataDao dao.InstructionDataDao
	userDao            dao.UserDao
}

func NewStatisticService(
	s *service.Service, instructionDataDao dao.InstructionDataDao, userDao dao.UserDao,
) StatisticService {
	return &StatisticServiceImpl{
		service:            s,
		instructionDataDao: instructionDataDao,
		userDao:            userDao,
	}
}

func (s StatisticServiceImpl) GetDataStatistic(
	ctx context.Context, startDate, endDate *time.Time,
) (*admin.GetDataStatisticResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (s StatisticServiceImpl) GetUserStatistic(
	ctx context.Context, userID *primitive.ObjectID,
) (*admin.GetUserStatisticResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (s StatisticServiceImpl) GetUserStatisticList(
	ctx context.Context, page *int, loginBefore, loginAfter, createdBefore, createdAfter *time.Time,
) (*admin.GetUserStatisticListResponse, error) {
	// TODO implement me
	panic("implement me")
}
