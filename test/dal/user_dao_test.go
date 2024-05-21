package dal_test

import (
	"context"
	"os"
	"testing"

	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/internal/pkg/dao"
	"data-collection-hub-server/internal/pkg/dao/mods"
	"data-collection-hub-server/internal/pkg/domain/entity"
	"data-collection-hub-server/internal/pkg/wire"
	"data-collection-hub-server/pkg/utils/crypt"
	"data-collection-hub-server/test/mock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ctx      context.Context
	userDao  mods.UserDao
	userID   primitive.ObjectID
	mockUser *mock.UserDaoMock
)

func TestMain(m *testing.M) {
	ctx = context.Background()
	cfg := config.New()
	cfg.MongoConfig.Database = "data-collection-hub"
	cfg.CacheConfig.RedisConfig.Password = "root"
	cfg.ZapConfig.Level = "error"
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
	mockUser = mock.NewUserDaoMockWithRandomData(1000)
	code := m.Run()
	err = mongo.Close(ctx)
	if err != nil {
		panic(err)
	}
	err = redis.Close()
	if err != nil {
		panic(err)
	}
	os.Exit(code)
}

func TestInsertUser(t *testing.T) {
	username := "Admin"
	email := "admin@admin.com"
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
	assert.True(t, crypt.Compare("Admin@123", user.Password))
}

func TestGetUser(t *testing.T) {
	user, err := userDao.GetUserByID(ctx, userID)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.UserID)
	assert.NotEmpty(t, user.Username)
	assert.NotEmpty(t, user.Email)
	assert.NotEmpty(t, user.Role)
	assert.NotEmpty(t, user.Organization)
	assert.NotEmpty(t, user.CreatedAt)
	assert.NotEmpty(t, user.UpdatedAt)
	assert.False(t, user.Deleted)
}

func TestUpdateUser(t *testing.T) {
	username := "User"
	email := "user@user.com"
	role := "USER"
	org := "Data Collection Hub X"
	password, err := crypt.Hash("User@123")
	assert.NoError(t, err)
	err = userDao.UpdateUser(ctx, userID, &username, &email, &password, &role, &org)
	assert.NoError(t, err)

	user, err := userDao.GetUserByID(ctx, userID)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, username, user.Username)
	assert.Equal(t, email, user.Email)
	assert.Equal(t, role, user.Role)
	assert.Equal(t, org, user.Organization)
	assert.True(t, crypt.Compare("User@123", user.Password))
}

func TestDeleteUser(t *testing.T) {
	err := userDao.SoftDeleteUser(ctx, userID)
	assert.NoError(t, err)

	user, err := userDao.GetUserByID(ctx, userID)
	assert.Error(t, err)
	assert.Nil(t, user)

	err = userDao.DeleteUser(ctx, userID)
	assert.NoError(t, err)

	user, err = userDao.GetUserByID(ctx, userID)
	assert.Error(t, err)
	assert.Nil(t, user)
}

func BenchmarkInsertUser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		username, email, password, role, org := mock.GenerateUser()
		InsertUserID, err := userDao.InsertUser(ctx, username, email, password, role, org)
		assert.NoError(b, err)
		assert.NotEmpty(b, InsertUserID)
		b.StopTimer()
		err = mockUser.Create(
			&entity.UserModel{
				UserID:       InsertUserID,
				Username:     username,
				Email:        email,
				Password:     password,
				Role:         role,
				Organization: org,
			},
		)
		assert.NoError(b, err)
		b.StartTimer()
	}
}

func BenchmarkGetUser(b *testing.B) {

}

func BenchmarkUpdateUser(b *testing.B) {

}

func BenchmarkDeleteUser(b *testing.B) {

}
