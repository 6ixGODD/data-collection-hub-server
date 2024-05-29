package service_test

import (
	"testing"
	"time"

	"data-collection-hub-server/test/mock"
	"data-collection-hub-server/test/wire"
	"github.com/stretchr/testify/assert"
)

func TestGetLoginLog(t *testing.T) {
	var (
		injector    = wire.GetInjector()
		ctx         = injector.Ctx
		logsService = injector.AdminLogsService
		loginLogID  = injector.LoginLogDaoMock.RandomLoginLogID()
	)
	resp, err := logsService.GetLoginLog(ctx, &loginLogID)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	t.Logf("Response Data: %+v", resp)
}

func TestGetLoginLogList(t *testing.T) {
	var (
		injector        = wire.GetInjector()
		ctx             = injector.Ctx
		logsService     = injector.AdminLogsService
		page            = int64(1)
		pageSize        = int64(10)
		desc            = false
		createTimeStart = time.Now().AddDate(0, 0, -1)
		createTimeEnd   = time.Now()
		query           = "a"
	)
	resp, err := logsService.GetLoginLogList(ctx, &page, &pageSize, &desc, &query, &createTimeStart, &createTimeEnd)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	resp, err = logsService.GetLoginLogList(ctx, &page, &pageSize, &desc, nil, nil, nil)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.LoginLogList)
	assert.Equal(t, pageSize, int64(len(resp.LoginLogList)))

	t.Logf("Response Data: %+v", resp)
}

func TestGetOperationLog(t *testing.T) {
	var (
		injector       = wire.GetInjector()
		ctx            = injector.Ctx
		logsService    = injector.AdminLogsService
		operationLogID = injector.OperationLogDaoMock.RandomOperationLogID()
	)
	resp, err := logsService.GetOperationLog(ctx, &operationLogID)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	t.Logf("Response Data: %+v", resp)
}

func TestGetOperationLogList(t *testing.T) {
	var (
		injector        = wire.GetInjector()
		ctx             = injector.Ctx
		logsService     = injector.AdminLogsService
		page            = int64(1)
		pageSize        = int64(10)
		desc            = false
		createTimeStart = time.Now().AddDate(0, 0, -1)
		createTimeEnd   = time.Now()
		query           = "a"
		operation       = mock.RandomEnum([]string{"CREATE", "UPDATE", "DELETE"})
	)
	resp, err := logsService.GetOperationLogList(
		ctx, &page, &pageSize, &desc, &query, &operation, &createTimeStart, &createTimeEnd,
	)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	resp, err = logsService.GetOperationLogList(ctx, &page, &pageSize, &desc, nil, nil, nil, nil)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.OperationLogList)
	assert.Equal(t, pageSize, int64(len(resp.OperationLogList)))

	t.Logf("Response Data: %+v", resp)
}
