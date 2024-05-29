package service_test

import (
	"testing"
	"time"

	"data-collection-hub-server/test/wire"
	"github.com/stretchr/testify/assert"
)

func TestGetDataStatistic(t *testing.T) {
	var (
		injector         = wire.GetInjector()
		ctx              = injector.Ctx
		statisticService = injector.AdminStatisticService
		startDate        = time.Now().AddDate(0, 0, -6)
		endDate          = time.Now()
	)
	resp, err := statisticService.GetDataStatistic(ctx, &startDate, &endDate)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	t.Logf("Response Data: %+v", resp)

	resp, err = statisticService.GetDataStatistic(ctx, nil, nil)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.TimeRangeStatistic)
	assert.Equal(t, 7, len(resp.TimeRangeStatistic))

	t.Logf("Response Data: %+v", resp)
}

func TestGetUserStatistic(t *testing.T) {
	var (
		injector         = wire.GetInjector()
		ctx              = injector.Ctx
		statisticService = injector.AdminStatisticService
		userID           = injector.UserDaoMock.RandomUserID()
	)
	resp, err := statisticService.GetUserStatistic(ctx, &userID)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	t.Logf("Response Data: %+v", resp)
}

func TestGetUserStatisticList(t *testing.T) {
	var (
		injector         = wire.GetInjector()
		ctx              = injector.Ctx
		statisticService = injector.AdminStatisticService
		page             = int64(1)
		pageSize         = int64(10)
		loginTimeStart   = time.Now().AddDate(0, 0, -1)
		loginTimeEnd     = time.Now()
		createTimeStart  = time.Now().AddDate(0, 0, -1)
		createTimeEnd    = time.Now()
	)
	resp, err := statisticService.GetUserStatisticList(
		ctx, &page, &pageSize, &loginTimeStart, &loginTimeEnd, &createTimeStart, &createTimeEnd,
	)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	resp, err = statisticService.GetUserStatisticList(ctx, &page, &pageSize, nil, nil, nil, nil)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.UserStatisticList)
	assert.Equal(t, pageSize, int64(len(resp.UserStatisticList)))

	t.Logf("Response Data: %+v", resp)
}
