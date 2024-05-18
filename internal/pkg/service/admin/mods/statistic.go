package mods

import (
	"context"
	"time"

	"data-collection-hub-server/internal/pkg/config"
	dao "data-collection-hub-server/internal/pkg/dao/mods"
	"data-collection-hub-server/internal/pkg/schema/admin"
	"data-collection-hub-server/internal/pkg/service"
	"data-collection-hub-server/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StatisticService interface {
	GetDataStatistic(ctx context.Context, startDate, endDate *time.Time) (*admin.GetDataStatisticResponse, error)
	GetUserStatistic(ctx context.Context, userID *primitive.ObjectID) (*admin.GetUserStatisticResponse, error)
	GetUserStatisticList(
		ctx context.Context, page, pageSize *int64, loginBefore, loginAfter, createdBefore, createdAfter *time.Time,
	) (*admin.GetUserStatisticListResponse, error)
}

type StatisticServiceImpl struct {
	service            *service.Core
	instructionDataDao dao.InstructionDataDao
	userDao            dao.UserDao
}

func NewStatisticService(
	s *service.Core, instructionDataDao dao.InstructionDataDao, userDao dao.UserDao,
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
	pendingStatus := config.InstructionDataStatusPending
	approvedStatus := config.InstructionDataStatusApproved
	rejectedStatus := config.InstructionDataStatusRejected
	themeField := "theme"
	total, err := s.instructionDataDao.CountInstructionData(
		ctx, nil, nil, nil, nil,
		nil, nil, nil,
	)
	if err != nil {
		return nil, errors.MongoError(errors.ReadError(err))
	}

	pendingCount, err := s.instructionDataDao.CountInstructionData(
		ctx, nil, nil, &pendingStatus, nil,
		nil, nil, nil,
	)
	if err != nil {
		return nil, errors.MongoError(errors.ReadError(err))
	}

	approvedCount, err := s.instructionDataDao.CountInstructionData(
		ctx, nil, nil, &approvedStatus, nil,
		nil, nil, nil,
	)
	if err != nil {
		return nil, errors.MongoError(errors.ReadError(err))
	}

	rejectedCount, err := s.instructionDataDao.CountInstructionData(
		ctx, nil, nil, &rejectedStatus, nil,
		nil, nil, nil,
	)
	if err != nil {
		return nil, errors.MongoError(errors.ReadError(err))
	}

	themeCount, err := s.instructionDataDao.AggregateCountInstructionData(ctx, &themeField)
	if err != nil {
		return nil, errors.MongoError(errors.ReadError(err))
	}

	timeRangeStatistic := make([]*admin.TimeRangeStatistic, 0, int(endDate.Sub(*startDate).Hours()/24)+1)
	for i := 0; i <= int(endDate.Sub(*startDate).Hours()/24); i++ {
		start := startDate.AddDate(0, 0, i)
		end := startDate.AddDate(0, 0, i+1)
		_total, err := s.instructionDataDao.CountInstructionData(
			ctx, nil, nil, nil, &start, &end,
			nil, nil,
		)
		if err != nil {
			return nil, errors.MongoError(errors.ReadError(err))
		}

		_pendingCount, err := s.instructionDataDao.CountInstructionData(
			ctx, nil, nil, &pendingStatus, &start, &end,
			nil, nil,
		)
		if err != nil {
			return nil, errors.MongoError(errors.ReadError(err))
		}

		_approvedCount, err := s.instructionDataDao.CountInstructionData(
			ctx, nil, nil, &approvedStatus, &start, &end,
			nil, nil,
		)
		if err != nil {
			return nil, errors.MongoError(errors.ReadError(err))
		}

		_rejectedCount, err := s.instructionDataDao.CountInstructionData(
			ctx, nil, nil, &rejectedStatus, &start, &end,
			nil, nil,
		)
		if err != nil {
			return nil, errors.MongoError(errors.ReadError(err))
		}

		_themeCount, err := s.instructionDataDao.AggregateCountInstructionData(ctx, &themeField)
		if err != nil {
			return nil, errors.MongoError(errors.ReadError(err))
		}

		timeRangeStatistic = append(
			timeRangeStatistic, &admin.TimeRangeStatistic{
				Date:          start.Format(time.RFC3339),
				Total:         *_total,
				PendingCount:  *_pendingCount,
				ApprovedCount: *_approvedCount,
				RejectedCount: *_rejectedCount,
				ThemeCount:    _themeCount,
			},
		)
	}
	return &admin.GetDataStatisticResponse{
		Total:              *total,
		PendingCount:       *pendingCount,
		ApprovedCount:      *approvedCount,
		RejectedCount:      *rejectedCount,
		ThemeCount:         themeCount,
		TimeRangeStatistic: timeRangeStatistic,
	}, nil
}

func (s StatisticServiceImpl) GetUserStatistic(
	ctx context.Context, userID *primitive.ObjectID,
) (*admin.GetUserStatisticResponse, error) {
	pendingStatus := config.InstructionDataStatusPending
	approvedStatus := config.InstructionDataStatusApproved
	rejectedStatus := config.InstructionDataStatusRejected
	user, err := s.userDao.GetUserById(ctx, *userID)
	if err != nil {
		return nil, errors.MongoError(errors.ReadError(err))
	}
	total, err := s.instructionDataDao.CountInstructionData(
		ctx, userID, nil, nil, nil,
		nil, nil, nil,
	)
	if err != nil {
		return nil, errors.MongoError(errors.ReadError(err))
	}

	pendingCount, err := s.instructionDataDao.CountInstructionData(
		ctx, userID, nil, &pendingStatus, nil,
		nil, nil, nil,
	)
	if err != nil {
		return nil, errors.MongoError(errors.ReadError(err))
	}

	approvedCount, err := s.instructionDataDao.CountInstructionData(
		ctx, userID, nil, &approvedStatus, nil,
		nil, nil, nil,
	)
	if err != nil {
		return nil, errors.MongoError(errors.ReadError(err))
	}

	rejectedCount, err := s.instructionDataDao.CountInstructionData(
		ctx, userID, nil, &rejectedStatus, nil,
		nil, nil, nil,
	)
	if err != nil {
		return nil, errors.MongoError(errors.ReadError(err))
	}

	return &admin.GetUserStatisticResponse{
		Username: user.Username,
		Data: admin.UserStatistic{
			Total:         *total,
			PendingCount:  *pendingCount,
			ApprovedCount: *approvedCount,
			RejectedCount: *rejectedCount,
		},
	}, nil
}

func (s StatisticServiceImpl) GetUserStatisticList(
	ctx context.Context, page, pageSize *int64, loginBefore, loginAfter, createdBefore, createdAfter *time.Time,
) (*admin.GetUserStatisticListResponse, error) {
	pendingStatus := config.InstructionDataStatusPending
	approvedStatus := config.InstructionDataStatusApproved
	rejectedStatus := config.InstructionDataStatusRejected
	offset := (*page - 1) * *pageSize
	users, count, err := s.userDao.GetUserList(
		ctx, offset, *pageSize, false, nil, nil, createdBefore, createdAfter,
		nil, nil, loginBefore, loginAfter, nil,
	)
	if err != nil {
		return nil, errors.MongoError(errors.ReadError(err))
	}
	resp := make([]*admin.GetUserStatisticResponse, 0, len(users))
	for _, user := range users {
		total, err := s.instructionDataDao.CountInstructionData(
			ctx, &user.UserID, nil, nil, nil,
			nil, nil, nil,
		)
		if err != nil {
			return nil, errors.MongoError(errors.ReadError(err))
		}

		pendingCount, err := s.instructionDataDao.CountInstructionData(
			ctx, &user.UserID, nil, &pendingStatus, nil,
			nil, nil, nil,
		)
		if err != nil {
			return nil, errors.MongoError(errors.ReadError(err))
		}

		approvedCount, err := s.instructionDataDao.CountInstructionData(
			ctx, &user.UserID, nil, &approvedStatus, nil,
			nil, nil, nil,
		)
		if err != nil {
			return nil, errors.MongoError(errors.ReadError(err))
		}

		rejectedCount, err := s.instructionDataDao.CountInstructionData(
			ctx, &user.UserID, nil, &rejectedStatus, nil,
			nil, nil, nil,
		)
		if err != nil {
			return nil, errors.MongoError(errors.ReadError(err))
		}

		resp = append(
			resp, &admin.GetUserStatisticResponse{
				Username: user.Username,
				Data: admin.UserStatistic{
					Total:         *total,
					PendingCount:  *pendingCount,
					ApprovedCount: *approvedCount,
					RejectedCount: *rejectedCount,
				},
			},
		)
	}

	return &admin.GetUserStatisticListResponse{
		Total:             *count,
		UserStatisticList: resp,
	}, nil
}
