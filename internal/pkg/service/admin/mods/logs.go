package mods

import (
	"context"
	"time"

	dao "data-collection-hub-server/internal/pkg/dal/mods"
	"data-collection-hub-server/internal/pkg/schema/admin"
	"data-collection-hub-server/internal/pkg/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LogsService interface {
	GetLoginLog(ctx context.Context, loginLogID *primitive.ObjectID) (*admin.GetLoginLogResponse, error)
	GetLoginLogList(
		ctx context.Context, page *int, query *string, createdBefore, createdAfter *time.Time,
	) (*admin.GetLoginLogListResponse, error)
	GetOperationLog(ctx context.Context, operationLogID *primitive.ObjectID) (*admin.GetOperationLogResponse, error)
	GetOperationLogList(
		ctx context.Context, page *int, query, operation *string, createdBefore, createdAfter *time.Time,
	) (*admin.GetOperationLogListResponse, error)
	GetErrorLog(ctx context.Context, errorLogID *primitive.ObjectID) (*admin.GetErrorLogResponse, error)
	GetErrorLogList(
		ctx context.Context, page *int, requestURL *string, errorCode *int, createdBefore, createdAfter *time.Time,
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

func (l LogsServiceImpl) GetLoginLog(ctx context.Context, loginLogID *primitive.ObjectID) (
	*admin.GetLoginLogResponse, error,
) {
	// TODO implement me
	panic("implement me")
}

func (l LogsServiceImpl) GetLoginLogList(
	ctx context.Context, page *int, query *string, createdBefore, createdAfter *time.Time,
) (*admin.GetLoginLogListResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (l LogsServiceImpl) GetOperationLog(
	ctx context.Context, operationLogID *primitive.ObjectID,
) (*admin.GetOperationLogResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (l LogsServiceImpl) GetOperationLogList(
	ctx context.Context, page *int, query, operation *string, createdBefore, createdAfter *time.Time,
) (*admin.GetOperationLogListResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (l LogsServiceImpl) GetErrorLog(ctx context.Context, errorLogID *primitive.ObjectID) (
	*admin.GetErrorLogResponse, error,
) {
	// TODO implement me
	panic("implement me")
}

func (l LogsServiceImpl) GetErrorLogList(
	ctx context.Context, page *int, requestURL *string, errorCode *int, createdBefore, createdAfter *time.Time,
) (*admin.GetErrorLogListResponse, error) {
	// TODO implement me
	panic("implement me")
}
