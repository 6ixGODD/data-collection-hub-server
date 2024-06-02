package mods

import (
	"context"
	e "errors"
	"fmt"
	"time"

	dao "data-collection-hub-server/internal/pkg/dao/mods"
	"data-collection-hub-server/internal/pkg/domain/vo/admin"
	"data-collection-hub-server/internal/pkg/service"
	"data-collection-hub-server/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LogsService interface {
	GetLoginLog(ctx context.Context, loginLogID *primitive.ObjectID) (*admin.GetLoginLogResponse, error)
	GetLoginLogList(
		ctx context.Context, page, pageSize *int64, desc *bool, query *string,
		createTimeStart, createTimeEnd *time.Time,
	) (*admin.GetLoginLogListResponse, error)
	GetOperationLog(ctx context.Context, operationLogID *primitive.ObjectID) (*admin.GetOperationLogResponse, error)
	GetOperationLogList(
		ctx context.Context, page, pageSize *int64, desc *bool, query, operation *string,
		createTimeStart, createTimeEnd *time.Time,
	) (*admin.GetOperationLogListResponse, error)
}

type LogsServiceImpl struct {
	core            *service.Core
	loginLogDao     dao.LoginLogDao
	operationLogDao dao.OperationLogDao
}

func NewLogsService(
	core *service.Core, loginLogDao dao.LoginLogDao, operationLogDao dao.OperationLogDao,
) LogsService {
	return &LogsServiceImpl{
		core:            core,
		loginLogDao:     loginLogDao,
		operationLogDao: operationLogDao,
	}
}

func (l LogsServiceImpl) GetLoginLog(
	ctx context.Context, loginLogID *primitive.ObjectID,
) (*admin.GetLoginLogResponse, error) {
	loginLog, err := l.loginLogDao.GetLoginLogByID(ctx, *loginLogID)
	if err != nil {
		if e.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.NotFound(fmt.Errorf("login log (id: %s) not found", loginLogID.Hex()))
		} else {
			return nil, errors.OperationFailed(fmt.Errorf("failed to get login log (id: %s)", loginLogID.Hex()))
		}
	}
	return &admin.GetLoginLogResponse{
		LoginLogID: loginLog.LoginLogID.Hex(),
		UserID:     loginLog.UserID.Hex(),
		Username:   loginLog.Username,
		Email:      loginLog.Email,
		IPAddress:  loginLog.IPAddress,
		UserAgent:  loginLog.UserAgent,
		CreatedAt:  loginLog.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (l LogsServiceImpl) GetLoginLogList(
	ctx context.Context, page, pageSize *int64, desc *bool, query *string, createTimeStart, createTimeEnd *time.Time,
) (*admin.GetLoginLogListResponse, error) {
	offset := (*page - 1) * *pageSize
	loginLogs, total, err := l.loginLogDao.GetLoginLogList(
		ctx, offset, *pageSize, *desc, createTimeStart, createTimeEnd, nil, nil, nil, query,
	)
	if err != nil {
		return nil, errors.OperationFailed(fmt.Errorf("failed to get login log list"))
	}
	loginLogList := make([]*admin.GetLoginLogResponse, 0, len(loginLogs))
	for _, loginLog := range loginLogs {
		loginLogList = append(
			loginLogList, &admin.GetLoginLogResponse{
				LoginLogID: loginLog.LoginLogID.Hex(),
				UserID:     loginLog.UserID.Hex(),
				Username:   loginLog.Username,
				Email:      loginLog.Email,
				IPAddress:  loginLog.IPAddress,
				UserAgent:  loginLog.UserAgent,
				CreatedAt:  loginLog.CreatedAt.Format(time.RFC3339),
			},
		)
	}
	return &admin.GetLoginLogListResponse{
		Total:        *total,
		LoginLogList: loginLogList,
	}, nil
}

func (l LogsServiceImpl) GetOperationLog(
	ctx context.Context, operationLogID *primitive.ObjectID,
) (*admin.GetOperationLogResponse, error) {
	operationLog, err := l.operationLogDao.GetOperationLogByID(ctx, *operationLogID)
	if err != nil {
		if e.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.NotFound(fmt.Errorf("operation log (id: %s) not found", operationLogID.Hex()))
		} else {
			return nil, errors.OperationFailed(fmt.Errorf("failed to get operation log (id: %s)", operationLogID.Hex()))
		}
	}
	return &admin.GetOperationLogResponse{
		OperationLogID: operationLog.OperationLogID.Hex(),
		UserID:         operationLog.UserID.Hex(),
		Username:       operationLog.Username,
		Email:          operationLog.Email,
		IPAddress:      operationLog.IPAddress,
		UserAgent:      operationLog.UserAgent,
		Operation:      operationLog.Operation,
		EntityID:       operationLog.EntityID.Hex(),
		CreatedAt:      operationLog.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (l LogsServiceImpl) GetOperationLogList(
	ctx context.Context, page, pageSize *int64, desc *bool, query, operation *string,
	createTimeStart, createTimeEnd *time.Time,
) (*admin.GetOperationLogListResponse, error) {
	offset := (*page - 1) * *pageSize
	operationLogs, total, err := l.operationLogDao.GetOperationLogList(
		ctx, offset, *pageSize, *desc, createTimeStart, createTimeEnd, nil, nil,
		nil, operation, nil, nil, query,
	)
	if err != nil {
		return nil, errors.OperationFailed(fmt.Errorf("failed to get operation log list"))
	}
	operationLogList := make([]*admin.GetOperationLogResponse, 0, len(operationLogs))
	for _, operationLog := range operationLogs {
		operationLogList = append(
			operationLogList, &admin.GetOperationLogResponse{
				OperationLogID: operationLog.OperationLogID.Hex(),
				UserID:         operationLog.UserID.Hex(),
				Username:       operationLog.Username,
				Email:          operationLog.Email,
				IPAddress:      operationLog.IPAddress,
				UserAgent:      operationLog.UserAgent,
				Operation:      operationLog.Operation,
				EntityID:       operationLog.EntityID.Hex(),
				CreatedAt:      operationLog.CreatedAt.Format(time.RFC3339),
			},
		)
	}
	return &admin.GetOperationLogListResponse{
		Total:            *total,
		OperationLogList: operationLogList,
	}, nil
}
