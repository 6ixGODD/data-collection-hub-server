package tasks

import (
	"context"

	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/internal/pkg/dao/mods"
	"data-collection-hub-server/pkg/cron"
	"data-collection-hub-server/pkg/jwt"
)

type Tasks struct {
	cron            *cron.Cron
	config          *config.Config
	loginLogDao     mods.LoginLogDao
	operationLogDao mods.OperationLogDao
	jwt             *jwt.Jwt
}

func New(
	ctx context.Context, config *config.Config, loginLogDao mods.LoginLogDao, operationLogDao mods.OperationLogDao,
) *Tasks {
	return &Tasks{
		cron:            cron.New(ctx),
		config:          config,
		loginLogDao:     loginLogDao,
		operationLogDao: operationLogDao,
	}
}

func (t *Tasks) syncLogs() {
	ctx := t.cron.Context()
	t.loginLogDao.SyncLoginLog(ctx)
	t.operationLogDao.SyncOperationLog(ctx)
}

func (t *Tasks) updateKey() {
	_ = t.jwt.UpdateKey()
}

func (t *Tasks) Start() error {
	if _, err := t.cron.AddFunc(t.config.TasksConfig.SyncLogsSpec, t.syncLogs); err != nil {
		return err
	}
	if _, err := t.cron.AddFunc(t.config.TasksConfig.UpdateKeySpec, t.updateKey); err != nil {
		return err
	}
	t.cron.Start()
	return nil
}

func (t *Tasks) Stop() {
	t.cron.Stop()
}
