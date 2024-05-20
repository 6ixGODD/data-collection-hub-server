package dal_test

import (
	"context"
	"testing"

	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/internal/pkg/dao"
	"data-collection-hub-server/internal/pkg/dao/mods"
	"data-collection-hub-server/internal/pkg/wire"
	"data-collection-hub-server/pkg/utils/crypt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ctx     context.Context
	userDao mods.UserDao
	userID  primitive.ObjectID
)

func TestMain(m *testing.M) {
	ctx = context.Background()
	cfg := config.New()
	cfg.MongoConfig.Database = "data-collection-hub"
	cfg.CacheConfig.RedisConfig.Password = "root"
	mongo, err := wire.InitializeMongo(ctx, cfg)
	if err != nil {
		panic(err)
	}
	redis, err := wire.InitializeRedis(ctx, cfg)
	if err != nil {
		panic(err)
	}
	zap, err := wire.InitializeZap(cfg)
	if err != nil {
		panic(err)
	}
	core, err := dao.NewCore(ctx, mongo, zap)
	if err != nil {
		panic(err)
	}
	cache := dao.NewCache(redis, cfg)
	if err != nil {
		panic(err)
	}
	userDao, err = mods.NewUserDao(ctx, core, cache)
	if err != nil {
		panic(err)
	}
	m.Run()
}

func TestInsertUser(t *testing.T) {
	username := "Admin"
	email := "6goddddddd@gmail.com"
	password, err := crypt.Hash("Admin@123")
	role := "ADMIN"
	org := "Data Collection Hub"
	assert.NoError(t, err)
	userID, err = userDao.InsertUser(ctx, username, email, password, role, org)
	assert.NoError(t, err)
	assert.NotEmpty(t, userID)

	user, err := userDao.GetUserByID(ctx, userID)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, username, user.Username)
	assert.Equal(t, email, user.Email)
	assert.Equal(t, role, user.Role)
	assert.Equal(t, org, user.Organization)
}
