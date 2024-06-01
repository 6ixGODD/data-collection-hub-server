package mods

import (
	"context"
	"fmt"
	"time"

	"data-collection-hub-server/internal/pkg/config"
	dao "data-collection-hub-server/internal/pkg/dao/mods"
	"data-collection-hub-server/internal/pkg/domain/vo/user"
	"data-collection-hub-server/internal/pkg/service"
	"data-collection-hub-server/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StatisticService interface {
	GetDataStatistic(ctx context.Context, startDate, endDate *time.Time) (*user.GetDataStatisticResponse, error)
}

type statisticServiceImpl struct {
	core               *service.Core
	instructionDataDao dao.InstructionDataDao
}

func NewStatisticService(core *service.Core, instructionDataDao dao.InstructionDataDao) StatisticService {
	return &statisticServiceImpl{
		core:               core,
		instructionDataDao: instructionDataDao,
	}
}

func (s statisticServiceImpl) GetDataStatistic(
	ctx context.Context, startDate, endDate *time.Time,
) (*user.GetDataStatisticResponse, error) {
	var (
		pendingStatus  = config.InstructionDataStatusPending
		approvedStatus = config.InstructionDataStatusApproved
		rejectedStatus = config.InstructionDataStatusRejected
		themeField     = "theme"
	)
	userIDHex, ok := ctx.Value(config.UserIDKey).(string)
	if !ok {
		return nil, errors.UserNotFound(errors.UserNotFound(fmt.Errorf("user id not found in context")))
	}
	userID, err := primitive.ObjectIDFromHex(userIDHex)
	if err != nil {
		return nil, errors.UserNotFound(err) // TODO: change error type
	}
	total, err := s.instructionDataDao.CountInstructionData(
		ctx, &userID, nil, nil, nil,
		nil, nil, nil,
	)
	if err != nil {
		return nil, errors.DBError(errors.ReadError(err))
	}

	pendingCount, err := s.instructionDataDao.CountInstructionData(
		ctx, &userID, nil, &pendingStatus, nil,
		nil, nil, nil,
	)
	if err != nil {
		return nil, errors.DBError(errors.ReadError(err))
	}

	approvedCount, err := s.instructionDataDao.CountInstructionData(
		ctx, &userID, nil, &approvedStatus, nil,
		nil, nil, nil,
	)
	if err != nil {
		return nil, errors.DBError(errors.ReadError(err))
	}

	rejectedCount, err := s.instructionDataDao.CountInstructionData(
		ctx, &userID, nil, &rejectedStatus, nil,
		nil, nil, nil,
	)
	if err != nil {
		return nil, errors.DBError(errors.ReadError(err))
	}

	themeCount, err := s.instructionDataDao.AggregateCountInstructionData(ctx, &themeField)
	if err != nil {
		return nil, errors.DBError(errors.ReadError(err))
	}

	if startDate == nil && endDate == nil {
		__startDate := time.Now().AddDate(0, 0, -6)
		__endDate := time.Now()
		startDate = &__startDate
		endDate = &__endDate
	} else if startDate == nil {
		__startDate := endDate.AddDate(0, 0, -6)
		startDate = &__startDate
	} else if endDate == nil {
		__endDate := startDate.AddDate(0, 0, 6)
		endDate = &__endDate
	}

	timeRangeStatistic := make([]*user.TimeRangeStatistic, 0, int(endDate.Sub(*startDate).Hours()/24)+1)
	for i := 0; i <= int(endDate.Sub(*startDate).Hours()/24); i++ {
		start := startDate.Add(time.Duration(i) * time.Hour * 24)
		end := startDate.Add(time.Duration(i+1) * time.Hour * 24)
		_total, err := s.instructionDataDao.CountInstructionData(
			ctx, &userID, nil, nil, &start, &end,
			nil, nil,
		)
		if err != nil {
			return nil, errors.DBError(errors.ReadError(err))
		}

		_pendingCount, err := s.instructionDataDao.CountInstructionData(
			ctx, &userID, nil, &pendingStatus, &start, &end,
			nil, nil,
		)
		if err != nil {
			return nil, errors.DBError(errors.ReadError(err))
		}

		_approvedCount, err := s.instructionDataDao.CountInstructionData(
			ctx, &userID, nil, &approvedStatus, &start, &end,
			nil, nil,
		)
		if err != nil {
			return nil, errors.DBError(errors.ReadError(err))
		}

		_rejectedCount, err := s.instructionDataDao.CountInstructionData(
			ctx, &userID, nil, &rejectedStatus, &start, &end,
			nil, nil,
		)
		if err != nil {
			return nil, errors.DBError(errors.ReadError(err))
		}

		timeRangeStatistic = append(
			timeRangeStatistic, &user.TimeRangeStatistic{
				Date:          start.Format(time.RFC3339),
				Total:         *_total,
				PendingCount:  *_pendingCount,
				ApprovedCount: *_approvedCount,
				RejectedCount: *_rejectedCount,
			},
		)
	}

	return &user.GetDataStatisticResponse{
		Total:              *total,
		PendingCount:       *pendingCount,
		ApprovedCount:      *approvedCount,
		RejectedCount:      *rejectedCount,
		ThemeCount:         themeCount,
		TimeRangeStatistic: timeRangeStatistic,
	}, nil
}
