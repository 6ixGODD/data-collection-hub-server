package dao_test

import (
	"testing"
	"time"

	"data-collection-hub-server/internal/pkg/domain/entity"
	"data-collection-hub-server/pkg/utils/crypt"
	"data-collection-hub-server/test/dao/mock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertUser(t *testing.T) {
	// t.Skip("Skip TestInsertUser")
	username = "Admin"
	email = "admin@admin.com"
	password, err := crypt.Hash("Admin@123")
	assert.NoError(t, err)
	role := "ADMIN"
	org := "Data Collection Hub"
	userID, err = userDao.InsertUser(userDaoCtx, username, email, password, role, org)
	assert.NoError(t, err)
	assert.NotEmpty(t, userID)

	user, err := userDao.GetUserByID(userDaoCtx, userID)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, username, user.Username)
	assert.Equal(t, email, user.Email)
	assert.Equal(t, role, user.Role)
	assert.Equal(t, org, user.Organization)
	assert.True(t, crypt.Compare("Admin@123", user.Password))
}

func TestGetUser(t *testing.T) {
	// t.Skip("Skip TestGetUser")
	user, err := userDao.GetUserByID(userDaoCtx, userID)
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

	_, _ = userDao.GetUserByID(userDaoCtx, userID)

	userCache, err := cache.Get(userDaoCtx, "dao:user:userID:"+user.UserID.Hex())
	assert.NoError(t, err)
	assert.NotNil(t, userCache)
	assert.NotEmpty(t, userCache)
	t.Logf("User cache: %v", *userCache)

	user, err = userDao.GetUserByEmail(userDaoCtx, email)
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
	_, _ = userDao.GetUserByEmail(userDaoCtx, email)

	user, err = userDao.GetUserByUsername(userDaoCtx, username)
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
	_, _ = userDao.GetUserByUsername(userDaoCtx, username)

	userNil, err := userDao.GetUserByID(userDaoCtx, primitive.NewObjectID())
	assert.Error(t, err)
	assert.Nil(t, userNil)

	userNil, err = userDao.GetUserByEmail(userDaoCtx, "")
	assert.Error(t, err)
	assert.Nil(t, userNil)

	userNil, err = userDao.GetUserByUsername(userDaoCtx, "")
	assert.Error(t, err)
	assert.Nil(t, userNil)
}

func TestGetUserList(t *testing.T) {
	// t.Skip("Skip TestGetUserList")
	var (
		organization       = "FOO"
		role               = "USER"
		createTimeStart    = time.Now().Add(-time.Hour)
		createTimeEnd      = time.Now()
		updateTimeStart    = time.Now().Add(-time.Hour)
		updateTimeEnd      = time.Now()
		lastLoginTimeStart = time.Now().Add(-time.Hour)
		lastLoginTimeEnd   = time.Now()
		query              = "Fo"
	)
	userList, count, err := userDao.GetUserList(
		userDaoCtx, 0, 10, false, nil, nil, nil,
		nil, nil, nil, nil,
		nil, nil,
	)
	assert.NoError(t, err)
	assert.NotNil(t, count)
	assert.NotEmpty(t, *count)
	assert.NotEmpty(t, userList)
	assert.NotNil(t, userList)
	assert.Equal(t, 10, len(userList))
	t.Logf("User count: %d", *count)
	t.Logf("User list: %v", userList)
	t.Logf("=====================================")

	userList, count, err = userDao.GetUserList(
		userDaoCtx, 0, 10, false, &organization, nil, nil,
		nil, nil, nil, nil,
		nil, nil,
	)
	assert.NoError(t, err)
	assert.NotNil(t, count)
	t.Logf("Organization: %s", organization)
	t.Logf("User count: %d", *count)
	t.Logf("User list: %v", userList)
	t.Logf("=====================================")

	userList, count, err = userDao.GetUserList(
		userDaoCtx, 0, 10, false, &organization, &role, nil,
		nil, nil, nil, nil,
		nil, nil,
	)
	assert.NoError(t, err)
	assert.NotNil(t, count)
	assert.NotNil(t, userList)
	t.Logf("Organization: %s", organization)
	t.Logf("Role: %s", role)
	t.Logf("User count: %d", *count)
	t.Logf("User list: %v", userList)
	t.Logf("=====================================")

	userList, count, err = userDao.GetUserList(
		userDaoCtx, 0, 10, false, nil, nil, &createTimeStart,
		&createTimeEnd, nil, nil, nil,
		nil, nil,
	)
	assert.NoError(t, err)
	assert.NotNil(t, count)
	assert.NotNil(t, userList)
	t.Logf("Create time start: %v", createTimeStart)
	t.Logf("Create time end: %v", createTimeEnd)
	t.Logf("User count: %d", *count)
	t.Logf("User list: %v", userList)
	t.Logf("=====================================")

	userList, count, err = userDao.GetUserList(
		userDaoCtx, 0, 10, false, nil, nil, nil,
		nil, &updateTimeStart, &updateTimeEnd, nil,
		nil, nil,
	)
	assert.NoError(t, err)
	assert.NotNil(t, count)
	assert.NotNil(t, userList)
	t.Logf("Update time start: %v", updateTimeStart)
	t.Logf("Update time end: %v", updateTimeEnd)
	t.Logf("User count: %d", *count)
	t.Logf("User list: %v", userList)
	t.Logf("=====================================")

	userList, count, err = userDao.GetUserList(
		userDaoCtx, 0, 10, false, nil, nil, nil,
		nil, nil, nil, &lastLoginTimeStart,
		&lastLoginTimeEnd, nil,
	)
	assert.NoError(t, err)
	assert.NotNil(t, count)
	assert.NotNil(t, userList)
	t.Logf("Last login time start: %v", lastLoginTimeStart)
	t.Logf("Last login time end: %v", lastLoginTimeEnd)
	t.Logf("User count: %d", *count)
	t.Logf("User list: %v", userList)
	t.Logf("=====================================")

	userList, count, err = userDao.GetUserList(
		userDaoCtx, 0, 10, false, nil, nil, nil,
		nil, nil, nil, nil,
		nil, &query,
	)
	assert.NoError(t, err)
	assert.NotNil(t, count)
	assert.NotNil(t, userList)
	t.Logf("Query: %s", query)
	t.Logf("User count: %d", *count)
	t.Logf("User list: %v", userList)
	t.Logf("=====================================")

	userList, count, err = userDao.GetUserList(
		userDaoCtx, 0, 10, false, &organization, &role, &createTimeStart,
		&createTimeEnd, &updateTimeStart, &updateTimeEnd, &lastLoginTimeStart,
		&lastLoginTimeEnd, &query,
	)
	assert.NoError(t, err)
	assert.NotNil(t, count)
	assert.NotNil(t, userList)
	t.Logf("Organization: %s", organization)
	t.Logf("Role: %s", role)
	t.Logf("Create time start: %v", createTimeStart)
	t.Logf("Create time end: %v", createTimeEnd)
	t.Logf("Update time start: %v", updateTimeStart)
	t.Logf("Update time end: %v", updateTimeEnd)
	t.Logf("Last login time start: %v", lastLoginTimeStart)
	t.Logf("Last login time end: %v", lastLoginTimeEnd)
	t.Logf("Query: %s", query)
	t.Logf("User count: %d", *count)
	t.Logf("User list: %v", userList)
	t.Logf("=====================================")
}

func TestUpdateUser(t *testing.T) {
	// t.Skip("Skip TestUpdateUser")
	username := "User"
	email := "user@user.com"
	role := "USER"
	org := "Data Collection Hub X"
	password, err := crypt.Hash("User@123")
	assert.NoError(t, err)
	err = userDao.UpdateUser(userDaoCtx, userID, &username, &email, &password, &role, &org)
	assert.NoError(t, err)

	user, err := userDao.GetUserByID(userDaoCtx, userID)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, username, user.Username)
	assert.Equal(t, email, user.Email)
	assert.Equal(t, role, user.Role)
	assert.Equal(t, org, user.Organization)
	assert.True(t, crypt.Compare("User@123", user.Password))
}

func TestDeleteUser(t *testing.T) {
	// t.Skip("Skip TestDeleteUser")
	err := userDao.SoftDeleteUser(userDaoCtx, userID)
	assert.NoError(t, err)

	user, err := userDao.GetUserByID(userDaoCtx, userID)
	assert.Error(t, err)
	assert.Nil(t, user)

	err = userDao.DeleteUser(userDaoCtx, userID)
	assert.NoError(t, err)

	user, err = userDao.GetUserByID(userDaoCtx, userID)
	assert.Error(t, err)
	assert.Nil(t, user)
}

func TestDeleteUserList(t *testing.T) {
	// t.Skip("Skip TestDeleteUserList")
	var (
		organization = "FOO"
		role         = "USER"
	)

	count, err := userDao.SoftDeleteUserList(userDaoCtx, &organization, nil, nil, nil, nil, nil, nil, nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, count)
	t.Logf("Organization: %s", organization)
	t.Logf("Delete count: %d", *count)
	t.Logf("=====================================")

	userList, count, err := userDao.GetUserList(
		userDaoCtx, 0, 10, false, &organization, nil, nil,
		nil, nil, nil, nil,
		nil, nil,
	)
	assert.Empty(t, userList)

	count, err = userDao.DeleteUserList(userDaoCtx, nil, &role, nil, nil, nil, nil, nil, nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, count)
	t.Logf("Role: %s", role)
	t.Logf("Delete count: %d", *count)
	t.Logf("=====================================")

	userList, count, err = userDao.GetUserList(
		userDaoCtx, 0, 10, false, nil, &role, nil,
		nil, nil, nil, nil,
		nil, nil,
	)
	assert.Empty(t, userList)

	count, err = userDao.SoftDeleteUserList(userDaoCtx, nil, &role, nil, nil, nil, nil, nil, nil)
	assert.NoError(t, err)
	assert.NotNil(t, count)
	t.Logf("Role: %s", role)
	t.Logf("Delete count: %d", *count)
	t.Logf("=====================================")

	userList, count, err = userDao.GetUserList(
		userDaoCtx, 0, 10, false, nil, &role, nil,
		nil, nil, nil, nil,
		nil, nil,
	)
	assert.Empty(t, userList)

	count, err = userDao.DeleteUserList(userDaoCtx, nil, &role, nil, nil, nil, nil, nil, nil)
	assert.NoError(t, err)
	assert.NotNil(t, count)
	t.Logf("Role: %s", role)
	t.Logf("Delete count: %d", *count)
	t.Logf("=====================================")

	userList, count, err = userDao.GetUserList(
		userDaoCtx, 0, 10, false, nil, &role, nil,
		nil, nil, nil, nil,
		nil, nil,
	)
	assert.Empty(t, userList)
}

func BenchmarkInsertUser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		usernameMock, emailMock, password, role, org := mock.GenerateUser()
		b.StartTimer()
		InsertUserID, err := userDao.InsertUser(userDaoCtx, usernameMock, emailMock, password, role, org)
		b.StopTimer()
		assert.NoError(b, err)
		assert.NotEmpty(b, InsertUserID)
		mockUser.Create(
			&entity.UserModel{
				UserID:       InsertUserID,
				Username:     usernameMock,
				Email:        emailMock,
				Password:     password,
				Role:         role,
				Organization: org,
			},
		)
		b.StartTimer()
	}
}

func BenchmarkGetUser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		userID := mockUser.RandomUserID()
		user, err := userDao.GetUserByID(userDaoCtx, userID)
		assert.NoError(b, err)
		assert.NotNil(b, user)
	}

	for i := 0; i < b.N; i++ {
		email := mockUser.UserMap[mockUser.RandomUserID()].Email
		user, err := userDao.GetUserByEmail(userDaoCtx, email)
		assert.NoError(b, err)
		assert.NotNil(b, user)
	}

	for i := 0; i < b.N; i++ {
		username := mockUser.UserMap[mockUser.RandomUserID()].Username
		user, err := userDao.GetUserByUsername(userDaoCtx, username)
		assert.NoError(b, err)
		assert.NotNil(b, user)
	}
}

func BenchmarkUpdateUser(b *testing.B) {

}

func BenchmarkDeleteUser(b *testing.B) {

}
