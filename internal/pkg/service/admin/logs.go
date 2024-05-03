package admin

import (
	dao "data-collection-hub-server/internal/pkg/dal/modules"
	"data-collection-hub-server/internal/pkg/service"
)

type LogsService interface {
}

type logsServiceImpl struct {
	service         *service.Service
	loginLogDao     dao.LoginLogDao
	operationLogDao dao.OperationLogDao
	errorLogDao     dao.ErrorLogDao
}

func NewLogsService(s *service.Service, loginLogDao dao.LoginLogDao, operationLogDao dao.OperationLogDao, errorLogDao dao.ErrorLogDao) LogsService {
	return &logsServiceImpl{
		service:         s,
		loginLogDao:     loginLogDao,
		operationLogDao: operationLogDao,
		errorLogDao:     errorLogDao,
	}
}
