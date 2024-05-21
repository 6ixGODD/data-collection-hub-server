package mock

import (
	"math/rand"

	"data-collection-hub-server/internal/pkg/domain/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserDaoMock is a mock for UserDao
type UserDaoMock struct {
	UserMap map[primitive.ObjectID]*entity.UserModel
}

// NewUserDaoMock returns a new UserDaoMock
func NewUserDaoMock() *UserDaoMock {
	return &UserDaoMock{
		UserMap: make(map[primitive.ObjectID]*entity.UserModel),
	}
}

// NewUserDaoMockWithRandomData returns a new UserDaoMock with n random users
func NewUserDaoMockWithRandomData(n int) *UserDaoMock {
	userDaoMock := NewUserDaoMock()
	for i := 0; i < n; i++ {
		user := GenerateUserModel()
		userDaoMock.UserMap[user.UserID] = user
	}
	return userDaoMock
}

// Create mocks the Create method
func (m *UserDaoMock) Create(user *entity.UserModel) error {
	m.UserMap[user.UserID] = user
	return nil
}

// Get mocks the Get method
func (m *UserDaoMock) Get(userID primitive.ObjectID) (*entity.UserModel, error) {
	user, ok := m.UserMap[userID]
	if !ok {
		return nil, nil
	}
	return user, nil
}

// GenerateUserModel generates a new UserModel
func GenerateUserModel() *entity.UserModel {
	return &entity.UserModel{
		UserID:       primitive.NewObjectID(),
		Email:        randomString(10) + "@test.com",
		Password:     randomString(10),
		Role:         randomEnum([]string{"ADMIN", "USER"}),
		Organization: randomEnum([]string{"FOO", "BAR", "BAZ", "QUX"}),
	}
}

// GenerateUser generates a new user, returns username, email, password, role, organization
func GenerateUser() (string, string, string, string, string) {
	return randomString(10),
		randomString(10) + "@test.com",
		randomString(10),
		randomEnum([]string{"ADMIN", "USER"}),
		randomEnum([]string{"FOO", "BAR", "BAZ", "QUX"})
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func randomEnum(enum []string) string {
	return enum[rand.Intn(len(enum))]
}
