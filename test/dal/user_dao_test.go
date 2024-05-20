package dal_test

import (
	"context"
	"os"
	"testing"
	"time"

	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/internal/pkg/dao"
	"data-collection-hub-server/internal/pkg/dao/mods"
	"data-collection-hub-server/pkg/mongo"
	"data-collection-hub-server/pkg/redis"
	"data-collection-hub-server/pkg/zap"
	"github.com/stretchr/testify/assert"
)

var (
	ctx     context.Context
	userDao mods.UserDao
)

func TestMain(m *testing.M) {
	ctx = context.Background()
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}
	mg, err := mongo.New(ctx, cfg.MongoConfig.GetQmgoConfig(), cfg.MongoConfig.PingTimeoutS, cfg.MongoConfig.Database)
	if err != nil {
		panic(err)
	}
	rd, err := redis.New(ctx, cfg.CacheConfig.GetRedisOptions())
	if err != nil {
		panic(err)
	}
	zp, err := zap.New(cfg.ZapConfig.GetZapConfig())
	if err != nil {
		panic(err)
	}
	dao := dao.NewCore(ctx, mg, rd, zp, *cfg)
	userDao, err = mods.NewUserDao(dao)
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestAdminUserDao(t *testing.T) {
	userID, err := userDao.InsertUser(
		ctx, "username", "email", "password", "role", "organization",
	)
	assert.NoError(t, err)
	assert.NotNil(t, userID)
	t.Logf("userID: %v", userID)

	user, err := userDao.GetUserById(ctx, userID)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	t.Logf("user: %v", user)

	email := "email"
	_user, err := userDao.GetUserByEmail(ctx, &email)
	assert.NoError(t, err)
	assert.NotNil(t, _user)
	t.Logf("_user: %v", _user)
	assert.Equal(t, user.UserID, _user.UserID)

	newUsername := "newUsername"
	err = userDao.UpdateUser(ctx, userID, &newUsername, nil, nil, nil, nil)
	assert.NoError(t, err)

	newEmail := "newEmail"
	err = userDao.UpdateUser(ctx, userID, nil, &newEmail, nil, nil, nil)
	assert.NoError(t, err)

	_, err = userDao.InsertUser(
		ctx, "newUsername", "newEmail", "newPassword", "newRole", "newOrganization",
	)
	assert.Error(t, err)

	err = userDao.DeleteUser(ctx, userID)
	assert.NoError(t, err)

	userFooID, err := userDao.InsertUser(
		ctx, "foo", "foo@gmail.com", "password", "ADMIN", "FOO",
	)
	assert.NoError(t, err)
	assert.NotNil(t, userFooID)
	t.Logf("userFooID: %v", userFooID)

	userBarID, err := userDao.InsertUser(
		ctx, "bar", "bar@gmail.com", "password", "ADMIN", "BAR",
	)
	assert.NoError(t, err)
	assert.NotNil(t, userBarID)
	t.Logf("userBarID: %v", userBarID)

	userBazID, err := userDao.InsertUser(
		ctx, "baz", "baz@gmail.com", "password", "ADMIN", "BAZ",
	)
	assert.NoError(t, err)
	assert.NotNil(t, userBazID)
	t.Logf("userBazID: %v", userBazID)

	userQuxID, err := userDao.InsertUser(
		ctx, "qux", "qux@gmail.com", "password", "ADMIN", "QUX",
	)
	assert.NoError(t, err)
	assert.NotNil(t, userQuxID)
	t.Logf("userQuxID: %v", userQuxID)

	// Test Multi
	userTest1ID, err := userDao.InsertUser(
		ctx, "test1", "test1@test.com", "password", "USER", "FOO",
	)
	assert.NoError(t, err)
	assert.NotNil(t, userTest1ID)
	t.Logf("userTest1ID: %v", userTest1ID)

	userTest2ID, err := userDao.InsertUser(
		ctx, "test2", "test2@test.com", "password", "USER", "FOO",
	)
	assert.NoError(t, err)
	assert.NotNil(t, userTest2ID)
	t.Logf("userTest2ID: %v", userTest2ID)

	userTest3ID, err := userDao.InsertUser(
		ctx, "test3", "test3@test.com", "password", "USER", "BAR",
	)
	assert.NoError(t, err)
	assert.NotNil(t, userTest3ID)
	t.Logf("userTest3ID: %v", userTest3ID)

	userTest4ID, err := userDao.InsertUser(
		ctx, "test4", "test4@test.com", "password", "USER", "QUX",
	)
	assert.NoError(t, err)
	assert.NotNil(t, userTest4ID)
	t.Logf("userTest4ID: %v", userTest4ID)

	userTest5ID, err := userDao.InsertUser(
		ctx, "test5", "test5@test.com", "password", "USER", "QUX",
	)
	assert.NoError(t, err)
	assert.NotNil(t, userTest5ID)
	t.Logf("userTest5ID: %v", userTest5ID)

	userTest6ID, err := userDao.InsertUser(
		ctx, "test6", "test6@test.com", "password", "USER", "QUX",
	)
	assert.NoError(t, err)
	assert.NotNil(t, userTest6ID)
	t.Logf("userTest6ID: %v", userTest6ID)

	userTest7ID, err := userDao.InsertUser(
		ctx, "test6", "test7@test.com", "password", "USER", "BAR",
	)
	assert.NoError(t, err)
	assert.NotNil(t, userTest7ID)
	t.Logf("userTest7ID: %v", userTest7ID)

	userTest8ID, err := userDao.InsertUser(
		ctx, "test8", "test8@test.com", "password", "USER", "BAZ",
	)
	assert.NoError(t, err)
	assert.NotNil(t, userTest8ID)
	t.Logf("userTest8ID: %v", userTest8ID)

	userTest9ID, err := userDao.InsertUser(
		ctx, "test9", "test9@test.com", "password", "USER", "BAZ",
	)
	assert.NoError(t, err)
	assert.NotNil(t, userTest9ID)
	t.Logf("userTest9ID: %v", userTest9ID)

	userTest10ID, err := userDao.InsertUser(
		ctx, "test10", "test10@test.com", "password", "USER", "BAR",
	)
	assert.NoError(t, err)
	assert.NotNil(t, userTest10ID)
	t.Logf("userTest10ID: %v", userTest10ID)

	users, total, err := userDao.GetUserList(
		ctx, 0, 10, false, nil, nil, nil, nil,
		nil, nil, nil, nil, nil,
	)
	assert.NoError(t, err)
	assert.NotNil(t, users)
	t.Logf("total: %v", *total)
	t.Logf("users: %v", users)
	t.Logf("--------------------------")

	descUser, total, err := userDao.GetUserList(
		ctx, 0, 10, true, nil, nil, nil, nil,
		nil, nil, nil, nil, nil,
	)
	assert.NoError(t, err)
	assert.NotNil(t, descUser)
	t.Logf("total: %v", *total)
	t.Logf("descUser: %v", descUser)
	t.Logf("--------------------------")

	fooOrg := "FOO"
	fooOrgUsers, total, err := userDao.GetUserList(
		ctx, 0, 10, false, &fooOrg, nil, nil, nil, nil,
		nil, nil, nil, nil,
	)
	assert.NoError(t, err)
	assert.NotNil(t, fooOrgUsers)
	t.Logf("total: %v", *total)
	t.Logf("fooOrgUsers: %v", fooOrgUsers)
	t.Logf("--------------------------")

	fooQuery := "foo"
	queryUsers, total, err := userDao.GetUserList(
		ctx, 0, 10, false, nil, nil, nil, nil,
		nil, nil, nil, nil, &fooQuery,
	)
	assert.NoError(t, err)
	assert.NotNil(t, queryUsers)
	t.Logf("total: %v", *total)
	t.Logf("queryUsers: %v", queryUsers)
	t.Logf("--------------------------")

	adminRole := "ADMIN"
	adminUsers, total, err := userDao.GetUserList(
		ctx, 0, 10, false, nil, &adminRole, nil, nil,
		nil, nil, nil, nil, nil,
	)
	assert.NoError(t, err)
	assert.NotNil(t, adminUsers)
	t.Logf("total: %v", *total)
	t.Logf("adminUsers: %v", adminUsers)
	t.Logf("--------------------------")

	timeRange := []time.Time{time.Now().Add(-time.Hour * 24), time.Now()}
	createdBetweenUsers, total, err := userDao.GetUserList(
		ctx, 0, 10, false, nil, nil, &timeRange[0], &timeRange[1],
		nil, nil, nil, nil, nil,
	)
	assert.NoError(t, err)
	assert.NotNil(t, createdBetweenUsers)
	t.Logf("total: %v", *total)
	t.Logf("createdBetweenUsers: %v", createdBetweenUsers)
	t.Logf("--------------------------")

	total, err = userDao.CountUser(
		ctx, nil, nil, nil, nil, nil, nil,
		nil, nil,
	)
	assert.NoError(t, err)
	assert.NotNil(t, total)
	t.Logf("total: %v", *total)
	t.Logf("--------------------------")

	fooTotal, err := userDao.CountUser(
		ctx, &fooOrg, nil, nil, nil, nil, nil,
		nil, nil,
	)
	assert.NoError(t, err)
	assert.NotNil(t, fooTotal)
	t.Logf("fooTotal: %v", *fooTotal)
	t.Logf("--------------------------")

	adminTotal, err := userDao.CountUser(
		ctx, nil, &adminRole, nil, nil, nil, nil,
		nil, nil,
	)
	assert.NoError(t, err)
	assert.NotNil(t, adminTotal)
	t.Logf("adminTotal: %v", *adminTotal)
	t.Logf("--------------------------")

	createdBetweenTotal, err := userDao.CountUser(
		ctx, nil, nil, &timeRange[0], &timeRange[1], nil, nil,
		nil, nil,
	)
	assert.NoError(t, err)
	assert.NotNil(t, createdBetweenTotal)
	t.Logf("createdBetweenTotal: %v", *createdBetweenTotal)
	t.Logf("--------------------------")

	err = userDao.SoftDeleteUser(ctx, userFooID)
	assert.NoError(t, err)
	userFoo, err := userDao.GetUserById(ctx, userFooID)
	assert.Error(t, err)
	assert.Nil(t, userFoo)
	newTotal, err := userDao.CountUser(
		ctx, nil, nil, nil, nil, nil, nil,
		nil, nil,
	)
	assert.NoError(t, err)
	assert.NotNil(t, newTotal)
	assert.Equal(t, *total-1, *newTotal)
	t.Logf("newTotal: %v", *newTotal)
	t.Logf("--------------------------")

	err = userDao.DeleteUser(ctx, userBarID)
	assert.NoError(t, err)
	userBar, err := userDao.GetUserById(ctx, userBarID)
	assert.Error(t, err)
	assert.Nil(t, userBar)
	newTotal, err = userDao.CountUser(
		ctx, nil, nil, nil, nil, nil, nil,
		nil, nil,
	)
	assert.NoError(t, err)
	assert.NotNil(t, newTotal)
	assert.Equal(t, *total-2, *newTotal)
	t.Logf("newTotal: %v", *newTotal)
	t.Logf("--------------------------")

	count, err := userDao.SoftDeleteUserList(
		ctx, nil, nil, nil, nil, nil, nil,
		nil, nil,
	)
	assert.NoError(t, err)
	assert.NotNil(t, count)
	t.Logf("count: %v", *count)
	t.Logf("--------------------------")

	users, total, err = userDao.GetUserList(
		ctx, 0, 10, false, nil, nil, nil, nil,
		nil, nil, nil, nil, nil,
	)
	assert.NoError(t, err)
	assert.Empty(t, users)

	count, err = userDao.DeleteUserList(
		ctx, nil, nil, nil, nil, nil, nil,
		nil, nil,
	)
	assert.NoError(t, err)
	assert.NotNil(t, count)
	t.Logf("count: %v", *count)
	t.Logf("--------------------------")
}
