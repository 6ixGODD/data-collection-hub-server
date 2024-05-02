package admin

import (
	dao "data-collection-hub-server/internal/pkg/dal/modules"
	"data-collection-hub-server/internal/pkg/service"
)

type LogsService interface {
}

type LogsServiceImpl struct {
	Service         *service.Service
	LoginLogDao     dao.LoginLogDao
	OperationLogDao dao.OperationLogDao
	ErrorLogDao     dao.ErrorLogDao
}

func NewLogsService(s *service.Service, loginLogDaoImpl *dao.LoginLogDaoImpl, operationLogDaoImpl *dao.OperationLogDaoImpl, errorLogDaoImpl *dao.ErrorLogDaoImpl) LogsService {
	return &LogsServiceImpl{
		Service:         s,
		LoginLogDao:     loginLogDaoImpl,
		OperationLogDao: operationLogDaoImpl,
		ErrorLogDao:     errorLogDaoImpl,
	}
}
