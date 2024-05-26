package dao_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInsertLoginLog(t *testing.T) {
	// t.Skip("Skip TestInsertLoginLog")
	var (
		userID    = mockUser.RandomUserID()
		ipAddress = "123.456.789.012"
		userAgent = "User Agent"
	)

	loginLogID, err = loginLogDao.InsertLoginLog(loginLogDaoCtx, userID, ipAddress, userAgent)
	assert.NoError(t, err)
	assert.NotEmpty(t, loginLogID)

	loginLog, err := loginLogDao.GetLoginLogByID(loginLogDaoCtx, loginLogID)
	assert.NoError(t, err)
	assert.NotNil(t, loginLog)
	assert.NotEmpty(t, loginLog.LoginLogID)
	assert.NotEmpty(t, loginLog.UserID)
	assert.NotEmpty(t, loginLog.IPAddress)
	assert.NotEmpty(t, loginLog.UserAgent)
	assert.NotEmpty(t, loginLog.CreatedAt)
}

func TestCacheLoginLog(t *testing.T) {
	// t.Skip("Skip TestCacheLoginLog")
	var (
		userID    = mockUser.RandomUserID()
		ipAddress = "cache 123.456.789.012"
		userAgent = "cache User Agent"
	)

	err := loginLogDao.CacheLoginLog(loginLogDaoCtx, userID, ipAddress, userAgent)
	assert.NoError(t, err)
	pop, err := cache.LeftPop(loginLogDaoCtx, "log:login")
	assert.NoError(t, err)
	assert.NotNil(t, pop)
	assert.NotEmpty(t, *pop)
	t.Logf("Login log: %s", *pop)
	t.Logf("=====================================")
	err = loginLogDao.CacheLoginLog(loginLogDaoCtx, userID, ipAddress, userAgent)
	assert.NoError(t, err)
}

func TestSyncLoginLog(t *testing.T) {
	var (
		ipAddress = "cache 123.456.789.012"
		userAgent = "cache User Agent"
	)
	// t.Skip("Skip TestSyncLoginLog")
	loginLogDao.SyncLoginLog(loginLogDaoCtx)
	loginLogList, count, err := loginLogDao.GetLoginLogList(
		loginLogDaoCtx, 0, 10, false, nil, nil, nil, &ipAddress, &userAgent, nil,
	)
	assert.NoError(t, err)
	assert.NotNil(t, count)
	assert.NotEmpty(t, *count)
	assert.NotEmpty(t, loginLogList)
	assert.Equal(t, 1, len(loginLogList))
	t.Logf("IP address: %s", ipAddress)
	t.Logf("User agent: %s", userAgent)
	t.Logf("Login log count: %d", *count)
	t.Logf("Login log list: %v", loginLogList)
	t.Logf("=====================================")
}

func TestGetLoginLogList(t *testing.T) {
	// t.Skip("Skip TestGetLoginLogList")
	var (
		userID    = mockUser.RandomUserID()
		startTime = time.Now().AddDate(0, 0, -1)
		endTime   = time.Now()
		ipAddress = "123.456.789.012"
		userAgent = "User Agent"
		query     = "a"
	)
	for i := 0; i < 100; i++ {
		_ = mockLoginLog.GenerateLoginLogWithUserID(userID)
		_ = mockLoginLog.GenerateLoginLogWithIpAddress(ipAddress)
		_ = mockLoginLog.GenerateLoginLogWithUserAgent(userAgent)
	}

	loginLogList, count, err := loginLogDao.GetLoginLogList(
		loginLogDaoCtx, 0, 10, false, nil, nil, nil, nil, nil, nil,
	)
	assert.NoError(t, err)
	assert.NotNil(t, count)
	assert.NotEmpty(t, *count)
	assert.NotEmpty(t, loginLogList)
	assert.Equal(t, 10, len(loginLogList))
	t.Logf("Login log count: %d", *count)
	t.Logf("Login log list: %v", loginLogList)
	t.Logf("=====================================")

	loginLogList, count, err = loginLogDao.GetLoginLogList(
		loginLogDaoCtx, 0, 10, false, &startTime, &endTime, nil, nil, nil, nil,
	)
	assert.NoError(t, err)
	assert.NotNil(t, count)
	t.Logf("Start time: %s", startTime)
	t.Logf("End time: %s", endTime)
	t.Logf("Login log count: %d", *count)
	t.Logf("Login log list: %v", loginLogList)
	t.Logf("=====================================")

	loginLogList, count, err = loginLogDao.GetLoginLogList(
		loginLogDaoCtx, 0, 10, false, nil, nil, &userID, nil, nil, nil,
	)
	assert.NoError(t, err)
	assert.NotNil(t, count)
	t.Logf("User ID: %s", userID)
	t.Logf("Login log count: %d", *count)
	t.Logf("Login log list: %v", loginLogList)
	t.Logf("=====================================")

	loginLogList, count, err = loginLogDao.GetLoginLogList(
		loginLogDaoCtx, 0, 10, false, nil, nil, nil, &ipAddress, nil, nil,
	)
	assert.NoError(t, err)
	assert.NotNil(t, count)
	t.Logf("IP address: %s", ipAddress)
	t.Logf("Login log count: %d", *count)
	t.Logf("Login log list: %v", loginLogList)
	t.Logf("=====================================")

	loginLogList, count, err = loginLogDao.GetLoginLogList(
		loginLogDaoCtx, 0, 10, false, nil, nil, nil, nil, &userAgent, nil,
	)
	assert.NoError(t, err)
	assert.NotNil(t, count)
	t.Logf("User agent: %s", userAgent)
	t.Logf("Login log count: %d", *count)
	t.Logf("Login log list: %v", loginLogList)
	t.Logf("=====================================")

	loginLogList, count, err = loginLogDao.GetLoginLogList(
		loginLogDaoCtx, 0, 10, false, nil, nil, nil, nil, nil, &query,
	)
	assert.NoError(t, err)
	assert.NotNil(t, count)
	t.Logf("Query: %s", query)
	t.Logf("Login log count: %d", *count)
	t.Logf("Login log list: %v", loginLogList)
	t.Logf("=====================================")

	loginLogList, count, err = loginLogDao.GetLoginLogList(
		loginLogDaoCtx, 0, 10, false, &startTime, &endTime, &userID, &ipAddress, &userAgent, &query,
	)
	assert.NoError(t, err)
	assert.NotNil(t, count)
	t.Logf("Start time: %s", startTime)
	t.Logf("End time: %s", endTime)
	t.Logf("User ID: %s", userID)
	t.Logf("IP address: %s", ipAddress)
	t.Logf("User agent: %s", userAgent)
	t.Logf("Query: %s", query)
	t.Logf("Login log count: %d", *count)
	t.Logf("Login log list: %v", loginLogList)
	t.Logf("=====================================")
}

func TestDeleteLoginLog(t *testing.T) {
	// t.Skip("Skip TestDeleteLoginLog")
	err := loginLogDao.DeleteLoginLog(loginLogDaoCtx, loginLogID)
	assert.NoError(t, err)

	loginLog, err := loginLogDao.GetLoginLogByID(loginLogDaoCtx, loginLogID)
	assert.Error(t, err)
	assert.Nil(t, loginLog)
}

func TestDeleteLoginLogList(t *testing.T) {
	// t.Skip("Skip TestDeleteLoginLogList")
	var (
		userID    = mockUser.RandomUserID()
		ipAddress = "123.456.789.012"
		userAgent = "User Agent"
	)
	for i := 0; i < 100; i++ {
		_ = mockLoginLog.GenerateLoginLogWithUserID(userID)
		_ = mockLoginLog.GenerateLoginLogWithIpAddress(ipAddress)
		_ = mockLoginLog.GenerateLoginLogWithUserAgent(userAgent)
	}

	count, err := loginLogDao.DeleteLoginLogList(loginLogDaoCtx, nil, nil, &userID, nil, nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, count)
	loginLogList, _, err := loginLogDao.GetLoginLogList(loginLogDaoCtx, 0, 10, false, nil, nil, &userID, nil, nil, nil)
	assert.NoError(t, err)
	assert.Empty(t, loginLogList)
	t.Logf("User ID: %s", userID)
	t.Logf("Delete count: %d", count)
	t.Logf("=====================================")

	count, err = loginLogDao.DeleteLoginLogList(loginLogDaoCtx, nil, nil, nil, &ipAddress, nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, count)
	loginLogList, _, err = loginLogDao.GetLoginLogList(
		loginLogDaoCtx, 0, 10, false, nil, nil, nil, &ipAddress, nil, nil,
	)
	assert.NoError(t, err)
	assert.Empty(t, loginLogList)
	t.Logf("IP address: %s", ipAddress)
	t.Logf("Delete count: %d", count)
	t.Logf("=====================================")

	count, err = loginLogDao.DeleteLoginLogList(loginLogDaoCtx, nil, nil, nil, nil, &userAgent)
	assert.NoError(t, err)
	assert.NotEmpty(t, count)
	loginLogList, _, err = loginLogDao.GetLoginLogList(
		loginLogDaoCtx, 0, 10, false, nil, nil, nil, nil, &userAgent, nil,
	)
	assert.NoError(t, err)
	assert.Empty(t, loginLogList)
	t.Logf("User agent: %s", userAgent)
	t.Logf("Delete count: %d", count)
	t.Logf("=====================================")
}
