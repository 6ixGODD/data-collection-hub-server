package mods

import (
	"context"
	e "errors"
	"fmt"
	"time"

	"data-collection-hub-server/internal/pkg/config"
	dao "data-collection-hub-server/internal/pkg/dao/mods"
	"data-collection-hub-server/internal/pkg/domain/vo/admin"
	"data-collection-hub-server/internal/pkg/service"
	"data-collection-hub-server/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type StatisticService interface {
	GetDataStatistic(ctx context.Context, startDate, endDate *time.Time) (*admin.GetDataStatisticResponse, error)
	GetUserStatistic(ctx context.Context, userID *primitive.ObjectID) (*admin.GetUserStatisticResponse, error)
	GetUserStatisticList(
		ctx context.Context, page, pageSize *int64,
		loginStartTime, loginEndTime, createdBefore, createdAfter *time.Time,
	) (*admin.GetUserStatisticListResponse, error)
}

type StatisticServiceImpl struct {
	core               *service.Core
	instructionDataDao dao.InstructionDataDao
	userDao            dao.UserDao
}

func NewStatisticService(
	core *service.Core, instructionDataDao dao.InstructionDataDao, userDao dao.UserDao,
) StatisticService {
	return &StatisticServiceImpl{
		core:               core,
		instructionDataDao: instructionDataDao,
		userDao:            userDao,
	}
}

func (s StatisticServiceImpl) GetDataStatistic(
	ctx context.Context, startDate, endDate *time.Time,
) (*admin.GetDataStatisticResponse, error) {
	var (
		pendingStatus  = config.InstructionDataStatusPending
		approvedStatus = config.InstructionDataStatusApproved
		rejectedStatus = config.InstructionDataStatusRejected
		themeField     = "theme"
	)

	total, err := s.instructionDataDao.CountInstructionData(
		ctx, nil, nil, nil, nil,
		nil, nil, nil,
	)
	if err != nil {
		return nil, errors.OperationFailed(fmt.Errorf("failed to count instruction data"))
	}

	pendingCount, err := s.instructionDataDao.CountInstructionData(
		ctx, nil, nil, &pendingStatus, nil,
		nil, nil, nil,
	)
	if err != nil {
		return nil, errors.OperationFailed(fmt.Errorf("failed to count instruction data with status %s", pendingStatus))
	}

	approvedCount, err := s.instructionDataDao.CountInstructionData(
		ctx, nil, nil, &approvedStatus, nil,
		nil, nil, nil,
	)
	if err != nil {
		return nil, errors.OperationFailed(
			fmt.Errorf(
				"failed to count instruction data with status %s", approvedStatus,
			),
		)
	}

	rejectedCount, err := s.instructionDataDao.CountInstructionData(
		ctx, nil, nil, &rejectedStatus, nil,
		nil, nil, nil,
	)
	if err != nil {
		return nil, errors.OperationFailed(
			fmt.Errorf(
				"failed to count instruction data with status %s", rejectedStatus,
			),
		)
	}

	themeCount, err := s.instructionDataDao.AggregateCountInstructionData(ctx, &themeField)
	if err != nil {
		return nil, errors.OperationFailed(fmt.Errorf("failed to count instruction data with field %s", themeField))
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
	timeRangeStatistic := make([]*admin.TimeRangeStatistic, 0, int(endDate.Sub(*startDate).Hours()/24)+1)
	for i := 0; i <= int(endDate.Sub(*startDate).Hours()/24); i++ {
		start := startDate.AddDate(0, 0, i)
		end := startDate.AddDate(0, 0, i+1)
		_total, err := s.instructionDataDao.CountInstructionData(
			ctx, nil, nil, nil, &start, &end,
			nil, nil,
		)
		if err != nil {
			return nil, errors.OperationFailed(
				fmt.Errorf(
					"failed to count instruction data in time range %s - %s", start.Format(time.RFC3339),
					end.Format(time.RFC3339),
				),
			)
		}

		_pendingCount, err := s.instructionDataDao.CountInstructionData(
			ctx, nil, nil, &pendingStatus, &start, &end,
			nil, nil,
		)
		if err != nil {
			return nil, errors.OperationFailed(
				fmt.Errorf(
					"failed to count instruction data with status %s in time range %s - %s", pendingStatus,
					start.Format(time.RFC3339), end.Format(time.RFC3339),
				),
			)
		}

		_approvedCount, err := s.instructionDataDao.CountInstructionData(
			ctx, nil, nil, &approvedStatus, &start, &end,
			nil, nil,
		)
		if err != nil {
			return nil, errors.OperationFailed(
				fmt.Errorf(
					"failed to count instruction data with status %s in time range %s - %s", approvedStatus,
					start.Format(time.RFC3339), end.Format(time.RFC3339),
				),
			)
		}

		_rejectedCount, err := s.instructionDataDao.CountInstructionData(
			ctx, nil, nil, &rejectedStatus, &start, &end,
			nil, nil,
		)
		if err != nil {
			return nil, errors.OperationFailed(
				fmt.Errorf(
					"failed to count instruction data with status %s in time range %s - %s", rejectedStatus,
					start.Format(time.RFC3339), end.Format(time.RFC3339),
				),
			)
		}

		_themeCount, err := s.instructionDataDao.AggregateCountInstructionData(ctx, &themeField)
		if err != nil {
			return nil, errors.OperationFailed(
				fmt.Errorf(
					"failed to count instruction data with field %s in time range %s - %s", themeField,
					start.Format(time.RFC3339), end.Format(time.RFC3339),
				),
			)
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
	user, err := s.userDao.GetUserByID(ctx, *userID)
	if err != nil {
		if e.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.NotFound(fmt.Errorf("user with id %s not found", userID.Hex()))
		}
		return nil, errors.OperationFailed(fmt.Errorf("failed to get user with id %s", userID.Hex()))
	}
	total, err := s.instructionDataDao.CountInstructionData(
		ctx, userID, nil, nil, nil,
		nil, nil, nil,
	)
	if err != nil {
		return nil, errors.OperationFailed(fmt.Errorf("failed to count instruction data for user %s", userID.Hex()))
	}

	pendingCount, err := s.instructionDataDao.CountInstructionData(
		ctx, userID, nil, &pendingStatus, nil,
		nil, nil, nil,
	)
	if err != nil {
		return nil, errors.OperationFailed(
			fmt.Errorf(
				"failed to count instruction data with status %s for user %s", pendingStatus, userID.Hex(),
			),
		)
	}

	approvedCount, err := s.instructionDataDao.CountInstructionData(
		ctx, userID, nil, &approvedStatus, nil,
		nil, nil, nil,
	)
	if err != nil {
		return nil, errors.OperationFailed(
			fmt.Errorf(
				"failed to count instruction data with status %s for user %s", approvedStatus, userID.Hex(),
			),
		)
	}

	rejectedCount, err := s.instructionDataDao.CountInstructionData(
		ctx, userID, nil, &rejectedStatus, nil,
		nil, nil, nil,
	)
	if err != nil {
		return nil, errors.OperationFailed(
			fmt.Errorf(
				"failed to count instruction data with status %s for user %s", pendingStatus, userID.Hex(),
			),
		)
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
		return nil, errors.OperationFailed(fmt.Errorf("failed to get user list"))
	}
	resp := make([]*admin.GetUserStatisticResponse, 0, len(users))
	for _, user := range users {
		total, err := s.instructionDataDao.CountInstructionData(
			ctx, &user.UserID, nil, nil, nil,
			nil, nil, nil,
		)
		if err != nil {
			return nil, errors.OperationFailed(
				fmt.Errorf(
					"failed to count instruction data for user %s", user.UserID.Hex(),
				),
			)
		}

		pendingCount, err := s.instructionDataDao.CountInstructionData(
			ctx, &user.UserID, nil, &pendingStatus, nil,
			nil, nil, nil,
		)
		if err != nil {
			return nil, errors.OperationFailed(
				fmt.Errorf(
					"failed to count instruction data with status %s for user %s", pendingStatus, user.UserID.Hex(),
				),
			)
		}

		approvedCount, err := s.instructionDataDao.CountInstructionData(
			ctx, &user.UserID, nil, &approvedStatus, nil,
			nil, nil, nil,
		)
		if err != nil {
			return nil, errors.OperationFailed(
				fmt.Errorf(
					"failed to count instruction data with status %s for user %s", approvedStatus, user.UserID.Hex(),
				),
			)
		}

		rejectedCount, err := s.instructionDataDao.CountInstructionData(
			ctx, &user.UserID, nil, &rejectedStatus, nil,
			nil, nil, nil,
		)
		if err != nil {
			return nil, errors.OperationFailed(
				fmt.Errorf(
					"failed to count instruction data with status %s for user %s", rejectedStatus, user.UserID.Hex(),
				),
			)
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
