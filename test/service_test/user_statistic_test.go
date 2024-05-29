package service_test

import (
	"context"
	"testing"
	"time"

	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/test/wire"
	"github.com/stretchr/testify/assert"
)

func TestUserGetDataStatistic(t *testing.T) {
	var (
		injector             = wire.GetInjector()
		ctx                  = injector.Ctx
		userStatisticService = injector.UserStatisticService
		startDate            = time.Now().AddDate(0, 0, -1)
		endDate              = time.Now()
	)
	ctx = context.WithValue(ctx, config.UserIDKey, injector.UserDaoMock.RandomUserID().Hex())
	resp, err := userStatisticService.GetDataStatistic(ctx, &startDate, &endDate)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	t.Logf("Response Data: %+v", resp)
}
