package mods

import (
	"data-collection-hub-server/internal/pkg/dao/mods"
	"data-collection-hub-server/internal/pkg/service"
)

type LogsService interface {
	InsertLoginLog()
	InsertOperationLog()
}

type logsServiceImpl struct {
	service         *service.Core
	loginLogDao     mods.LoginLogDao
	operationLogDao mods.OperationLogDao
}

func NewLogsService(s *service.Core) LogsService {
	return &logsServiceImpl{
		service:         s,
		loginLogDao:     mods.NewLoginLogDao(s),
		operationLogDao: mods.NewOperationLogDao(s),
	}
}

func (l logsServiceImpl) InsertLoginLog() {
	// TODO implement me
	panic("implement me")
}

func (l logsServiceImpl) InsertOperationLog() {
	// TODO implement me
	panic("implement me")
}
