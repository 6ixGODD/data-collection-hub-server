package dal__test

import (
	"context"
	"testing"

	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/internal/pkg/dal"
	"data-collection-hub-server/internal/pkg/dal/mods"
	"data-collection-hub-server/pkg/mongo"
	"data-collection-hub-server/pkg/redis"
	"data-collection-hub-server/pkg/zap"
	"github.com/stretchr/testify/assert"
)

func TestDao__Admin_UserDao(t *testing.T) {
	ctx := context.Background()
	cfg, err := config.New()
	assert.NoError(t, err)
	mg, err := mongo.New(ctx, cfg.MongoConfig.GetQmgoConfig(), cfg.MongoConfig.PingTimeoutS, cfg.MongoConfig.Database)
	assert.NoError(t, err)
	rd, err := redis.New(ctx, cfg.RedisConfig.GetRedisOptions())
	assert.NoError(t, err)
	zp, err := zap.New(cfg.ZapConfig.GetZapConfig())
	assert.NoError(t, err)
	dao := dal.New(ctx, mg, rd, zp, *cfg)
	assert.NotNil(t, dao)
	userDao, err := mods.NewUserDao(dao)
	assert.NoError(t, err)
	assert.NotNil(t, userDao)

	userID, err := userDao.InsertUser(ctx, "username", "email", "password", "role", "organization")
	assert.NoError(t, err)
	assert.NotNil(t, userID)
	t.Logf("userID: %v", userID)

	user, err := userDao.GetUserById(ctx, userID)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	t.Logf("user: %v", user)
}

func TestDao__Admin_InstructionDataDao(t *testing.T) {
	ctx := context.Background()
	cfg, err := config.New()
	assert.NoError(t, err)
	mg, err := mongo.New(ctx, cfg.MongoConfig.GetQmgoConfig(), cfg.MongoConfig.PingTimeoutS, cfg.MongoConfig.Database)
	assert.NoError(t, err)
	rd, err := redis.New(ctx, cfg.RedisConfig.GetRedisOptions())
	assert.NoError(t, err)
	zp, err := zap.New(cfg.ZapConfig.GetZapConfig())
	assert.NoError(t, err)
	dao := dal.New(ctx, mg, rd, zp, *cfg)
	assert.NotNil(t, dao)
	userDao, err := mods.NewUserDao(dao)
	assert.NotNil(t, userDao)
	instructionDataDao := mods.NewInstructionDataDao(dao, userDao)
	assert.NotNil(t, instructionDataDao)
}
