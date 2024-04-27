package dao

import (
	"context"

	"data-collection-hub-server/internal/pkg/models"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
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
	var user models.UserModel
	err := mongoClient.Find(ctx, bson.M{"user_id": userId}).One(&user)
	if err != nil {
		return nil, err
	} else {
		return &user, nil
	}
}

func (u UserDaoImpl) GetUserList(mongoClient *qmgo.QmgoClient, offset, limit int64, desc bool, ctx context.Context) ([]models.UserModel, error) {
	var users []models.UserModel
	err := mongoClient.Find(ctx, bson.M{}).Skip(offset).Limit(limit).All(&users)
	if err != nil {
		return nil, err
	} else {
		return users, nil
	}
}

func (u UserDaoImpl) GetUserListByOrganization(mongoClient *qmgo.QmgoClient, organization string, offset, limit int64, desc bool, ctx context.Context) ([]models.UserModel, error) {
	var users []models.UserModel
	err := mongoClient.Find(ctx, bson.M{"organization": organization}).Skip(offset).Limit(limit).All(&users)
	if err != nil {
		return nil, err
	} else {
		return users, nil
	}
}

func (u UserDaoImpl) GetUserListByRole(mongoClient *qmgo.QmgoClient, role string, offset, limit int64, desc bool, ctx context.Context) ([]models.UserModel, error) {
	var users []models.UserModel
	err := mongoClient.Find(ctx, bson.M{"role": role}).Skip(offset).Limit(limit).All(&users)
	if err != nil {
		return nil, err
	} else {
		return users, nil
	}
}

func (u UserDaoImpl) GetUserListByCreatedTime(mongoClient *qmgo.QmgoClient, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.UserModel, error) {
	var users []models.UserModel
	err := mongoClient.Find(ctx, bson.M{"created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Skip(offset).Limit(limit).All(&users)
	if err != nil {
		return nil, err
	} else {
		return users, nil
	}
}

func (u UserDaoImpl) GetUserListByUpdatedTime(mongoClient *qmgo.QmgoClient, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.UserModel, error) {
	var users []models.UserModel
	err := mongoClient.Find(ctx, bson.M{"updated_time": bson.M{"$gte": startTime, "$lte": endTime}}).Skip(offset).Limit(limit).All(&users)
	if err != nil {
		return nil, err
	} else {
		return users, nil
	}
}

func (u UserDaoImpl) GetUserListByFuzzyQuery(mongoClient *qmgo.QmgoClient, query string, offset, limit int64, desc bool, ctx context.Context) ([]models.UserModel, error) {
	var users []models.UserModel
	err := mongoClient.Find(ctx, bson.M{"$or": []bson.M{
		{"_id": bson.M{"$regex": query}},
		{"username": bson.M{"$regex": query}},
		{"organization:": bson.M{"$regex": query}},
		{"role": bson.M{"$regex": query}},
	}}).Skip(offset).Limit(limit).All(&users)
	if err != nil {
		return nil, err
	} else {
		return users, nil
	}
}

func (u UserDaoImpl) InsertUser(user *models.UserModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error {
	_, err := mongoClient.InsertOne(ctx, user)
	return err
}

func (u UserDaoImpl) UpdateUser(user *models.UserModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error {
	err := mongoClient.UpdateOne(ctx, bson.M{"user_id": user.UserID}, bson.M{"$set": user})
	return err
}

func (u UserDaoImpl) DeleteUser(user *models.UserModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error {
	err := mongoClient.RemoveId(ctx, user.UserID)
	return err
}

func NewUserDao() UserDao {
	var _ UserDao = new(UserDaoImpl)
	return &UserDaoImpl{}
}
