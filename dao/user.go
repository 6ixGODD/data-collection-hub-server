package dao

import (
	"context"
	"data-collection-hub-server/models"
	"github.com/qiniu/qmgo"
)

type UserDao interface {
	GetUserById(userId string, mongoClient *qmgo.QmgoClient, ctx context.Context) (*models.UserModel, error)
	GetUserList(mongoClient *qmgo.QmgoClient, offset, limit int64, desc bool, ctx context.Context) ([]models.UserModel, error)
	GetUserListByOrganization(mongoClient *qmgo.QmgoClient, organization string, offset, limit int64, desc bool, ctx context.Context) ([]models.UserModel, error)
	GetUserListByRole(mongoClient *qmgo.QmgoClient, role string, offset, limit int64, desc bool, ctx context.Context) ([]models.UserModel, error)
	GetUserListByCreatedTime(mongoClient *qmgo.QmgoClient, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.UserModel, error)
	GetUserListByUpdatedTime(mongoClient *qmgo.QmgoClient, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.UserModel, error)
	GetUserListByFuzzyQuery(mongoClient *qmgo.QmgoClient, query string, offset, limit int64, desc bool, ctx context.Context) ([]models.UserModel, error)
	InsertUser(user *models.UserModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error
	UpdateUser(user *models.UserModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error
	DeleteUser(user *models.UserModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error
}

type UserDaoImpl struct{}

func (u UserDaoImpl) GetUserById(userId string, mongoClient *qmgo.QmgoClient, ctx context.Context) (*models.UserModel, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserDaoImpl) GetUserList(mongoClient *qmgo.QmgoClient, offset, limit int64, desc bool, ctx context.Context) ([]models.UserModel, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserDaoImpl) GetUserListByOrganization(mongoClient *qmgo.QmgoClient, organization string, offset, limit int64, desc bool, ctx context.Context) ([]models.UserModel, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserDaoImpl) GetUserListByRole(mongoClient *qmgo.QmgoClient, role string, offset, limit int64, desc bool, ctx context.Context) ([]models.UserModel, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserDaoImpl) GetUserListByCreatedTime(mongoClient *qmgo.QmgoClient, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.UserModel, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserDaoImpl) GetUserListByUpdatedTime(mongoClient *qmgo.QmgoClient, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.UserModel, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserDaoImpl) GetUserListByFuzzyQuery(mongoClient *qmgo.QmgoClient, query string, offset, limit int64, desc bool, ctx context.Context) ([]models.UserModel, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserDaoImpl) InsertUser(user *models.UserModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (u UserDaoImpl) UpdateUser(user *models.UserModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (u UserDaoImpl) DeleteUser(user *models.UserModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func NewUserDao() UserDao {
	var _ UserDao = new(UserDaoImpl) // TODO: remove this line when deploying
	return &UserDaoImpl{}
}
