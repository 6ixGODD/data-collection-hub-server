package dao

import (
	"context"
	"data-collection-hub-server/models"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
)

type LoginLogDao interface {
	GetLoginLogById(loginLogId string, mongoClient *qmgo.QmgoClient, ctx context.Context) (*models.LoginLogModel, error)
	GetLoginLogList(mongoClient *qmgo.QmgoClient, offset, limit int64, ctx context.Context) ([]models.LoginLogModel, error)
	GetLoginLogListByCreatedTime(mongoClient *qmgo.QmgoClient, startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.LoginLogModel, error)
	GetLoginLogListByUserUUID(mongoClient *qmgo.QmgoClient, userUUID string, offset, limit int64, ctx context.Context) ([]models.LoginLogModel, error)
	GetLoginLogListByIPAddress(mongoClient *qmgo.QmgoClient, ipAddress string, offset, limit int64, ctx context.Context) ([]models.LoginLogModel, error)
	InsertLoginLog(loginLog *models.LoginLogModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error
	DeleteLoginLog(loginLogId string, mongoClient *qmgo.QmgoClient, ctx context.Context) error
}

type LoginLogDaoImpl struct{}

func NewLoginLogDaoImpl() *LoginLogDaoImpl {
	var _ LoginLogDao = new(LoginLogDaoImpl) // Ensure that the interface is implemented
	return &LoginLogDaoImpl{}
}

func (loginLogDao *LoginLogDaoImpl) GetLoginLogById(loginLogId string, mongoClient *qmgo.QmgoClient, ctx context.Context) (*models.LoginLogModel, error) {
	var loginLog models.LoginLogModel
	err := mongoClient.Find(ctx, bson.M{"_id": loginLogId}).One(&loginLog)
	if err != nil {
		return nil, err
	} else {
		return &loginLog, nil
	}
}
func (loginLogDao *LoginLogDaoImpl) GetLoginLogList(mongoClient *qmgo.QmgoClient, offset, limit int64, ctx context.Context) ([]models.LoginLogModel, error) {
	var loginLogList []models.LoginLogModel
	err := mongoClient.Find(ctx, bson.M{}).Skip(offset).Limit(limit).All(&loginLogList)
	if err != nil {
		return nil, err
	} else {
		return loginLogList, nil
	}
}
func (loginLogDao *LoginLogDaoImpl) GetLoginLogListByCreatedTime(mongoClient *qmgo.QmgoClient, startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.LoginLogModel, error) {
	var loginLogList []models.LoginLogModel
	err := mongoClient.Find(
		ctx, bson.M{"created_at": bson.M{"$gte": startTime, "$lte": endTime}},
	).Skip(offset).Limit(limit).All(&loginLogList)
	if err != nil {
		return nil, err
	} else {
		return loginLogList, nil
	}
}
func (loginLogDao *LoginLogDaoImpl) GetLoginLogListByUserUUID(mongoClient *qmgo.QmgoClient, userUUID string, offset, limit int64, ctx context.Context) ([]models.LoginLogModel, error) {
	var loginLogList []models.LoginLogModel
	err := mongoClient.Find(ctx, bson.M{"user_uuid": userUUID}).Skip(offset).Limit(limit).All(&loginLogList)
	if err != nil {
		return nil, err
	} else {
		return loginLogList, nil
	}
}
func (loginLogDao *LoginLogDaoImpl) GetLoginLogListByIPAddress(mongoClient *qmgo.QmgoClient, ipAddress string, offset, limit int64, ctx context.Context) ([]models.LoginLogModel, error) {
	var loginLogList []models.LoginLogModel
	err := mongoClient.Find(ctx, bson.M{"ip_address": ipAddress}).Skip(offset).Limit(limit).All(&loginLogList)
	if err != nil {
		return nil, err
	} else {
		return loginLogList, nil
	}
}
func (loginLogDao *LoginLogDaoImpl) InsertLoginLog(loginLog *models.LoginLogModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
func (loginLogDao *LoginLogDaoImpl) DeleteLoginLog(loginLogId string, mongoClient *qmgo.QmgoClient, ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
