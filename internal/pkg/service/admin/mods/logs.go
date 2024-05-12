package mods

import (
	"context"
	"time"

	dao "data-collection-hub-server/internal/pkg/dal/mods"
	"data-collection-hub-server/internal/pkg/schema/admin"
	"data-collection-hub-server/internal/pkg/service"
	"data-collection-hub-server/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	GetErrorLog(ctx context.Context, errorLogID *primitive.ObjectID) (*admin.GetErrorLogResponse, error)
	GetErrorLogList(
		ctx context.Context, page, pageSize *int64, desc *bool, requestURL, errorCode *string,
		createTimeStart, createTimeEnd *time.Time,
	) (*admin.GetErrorLogListResponse, error)
}

type LogsServiceImpl struct {
	service         *service.Service
	loginLogDao     dao.LoginLogDao
	operationLogDao dao.OperationLogDao
	errorLogDao     dao.ErrorLogDao
}

func NewLogsService(
	s *service.Service, loginLogDao dao.LoginLogDao, operationLogDao dao.OperationLogDao, errorLogDao dao.ErrorLogDao,
) LogsService {
	return &LogsServiceImpl{
		service:         s,
		loginLogDao:     loginLogDao,
		operationLogDao: operationLogDao,
		errorLogDao:     errorLogDao,
	}
}

func (l LogsServiceImpl) GetLoginLog(
	ctx context.Context, loginLogID *primitive.ObjectID,
) (*admin.GetLoginLogResponse, error) {
	loginLog, err := l.loginLogDao.GetLoginLogById(ctx, *loginLogID)
	if err != nil {
		return nil, errors.MongoError(errors.ReadError(err))
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
		return nil, errors.MongoError(errors.ReadError(err))
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
	operationLog, err := l.operationLogDao.GetOperationLogById(ctx, *operationLogID)
	if err != nil {
		return nil, errors.MongoError(errors.ReadError(err))
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
		return nil, errors.MongoError(errors.ReadError(err))
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

func (l LogsServiceImpl) GetErrorLog(
	ctx context.Context, errorLogID *primitive.ObjectID,
) (*admin.GetErrorLogResponse, error) {
	errorLog, err := l.errorLogDao.GetErrorLogById(ctx, *errorLogID)
	if err != nil {
		return nil, errors.MongoError(errors.ReadError(err))
	}
	return &admin.GetErrorLogResponse{
		ErrorLogID:     errorLog.ErrorLogID.Hex(),
		UserID:         errorLog.UserID.Hex(),
		Username:       errorLog.Username,
		IPAddress:      errorLog.IPAddress,
		UserAgent:      errorLog.UserAgent,
		RequestURL:     errorLog.RequestURL,
		RequestMethod:  errorLog.RequestMethod,
		RequestPayload: errorLog.RequestPayload,
		ErrorCode:      errorLog.ErrorCode,
		ErrorMsg:       errorLog.ErrorMsg,
		Stack:          errorLog.Stack,
		CreatedAt:      errorLog.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (l LogsServiceImpl) GetErrorLogList(
	ctx context.Context, page, pageSize *int64, desc *bool, requestURL, errorCode *string,
	createTimeStart, createTimeEnd *time.Time,
) (*admin.GetErrorLogListResponse, error) {
	offset := (*page - 1) * *pageSize
	errorLogs, total, err := l.errorLogDao.GetErrorLogList(
		ctx, offset, *pageSize, *desc, createTimeStart, createTimeEnd, nil,
		nil, requestURL, errorCode, nil,
	)
	if err != nil {
		return nil, errors.MongoError(errors.ReadError(err))
	}
	errorLogList := make([]*admin.GetErrorLogResponse, 0, len(errorLogs))
	for _, errorLog := range errorLogs {
		errorLogList = append(
			errorLogList, &admin.GetErrorLogResponse{
				ErrorLogID:     errorLog.ErrorLogID.Hex(),
				UserID:         errorLog.UserID.Hex(),
				Username:       errorLog.Username,
				IPAddress:      errorLog.IPAddress,
				UserAgent:      errorLog.UserAgent,
				RequestURL:     errorLog.RequestURL,
				RequestMethod:  errorLog.RequestMethod,
				RequestPayload: errorLog.RequestPayload,
				ErrorCode:      errorLog.ErrorCode,
				ErrorMsg:       errorLog.ErrorMsg,
				Stack:          errorLog.Stack,
				CreatedAt:      errorLog.CreatedAt.Format(time.RFC3339),
			},
		)
	}
	return &admin.GetErrorLogListResponse{
		Total:        *total,
		ErrorLogList: errorLogList,
	}, nil

}
