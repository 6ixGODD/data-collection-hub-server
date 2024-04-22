package modules

import (
	"context"

	"data-collection-hub-server/models"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
)

type LoginLogDao interface {
	GetLoginLogById(loginLogId string, ctx context.Context) (*models.LoginLogModel, error)
	GetLoginLogList(offset, limit int64, ctx context.Context) ([]models.LoginLogModel, error)
	GetLoginLogListByCreatedTime(startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.LoginLogModel, error)
	GetLoginLogListByUserUUID(userUUID string, offset, limit int64, ctx context.Context) ([]models.LoginLogModel, error)
	GetLoginLogListByIPAddress(ipAddress string, offset, limit int64, ctx context.Context) ([]models.LoginLogModel, error)
	InsertLoginLog(loginLog *models.LoginLogModel, ctx context.Context) error
	DeleteLoginLog(loginLog *models.LoginLogModel, ctx context.Context) error
}

type LoginLogDaoImpl struct {
	loginLogClient *qmgo.Collection
}

func NewLoginLogDao(mongoDatabase *qmgo.Database) LoginLogDao {
	var _ LoginLogDao = new(LoginLogDaoImpl) // Ensure that the interface is implemented
	return &LoginLogDaoImpl{loginLogClient: mongoDatabase.Collection("login_log")}
}

func (l *LoginLogDaoImpl) GetLoginLogById(loginLogId string, ctx context.Context) (*models.LoginLogModel, error) {
	var loginLog models.LoginLogModel
	err := l.loginLogClient.Find(ctx, bson.M{"_id": loginLogId}).One(&loginLog)
	if err != nil {
		return nil, err
	} else {
		return &loginLog, nil
	}
}
func (l *LoginLogDaoImpl) GetLoginLogList(offset, limit int64, ctx context.Context) ([]models.LoginLogModel, error) {
	var loginLogList []models.LoginLogModel
	err := l.loginLogClient.Find(ctx, bson.M{}).Skip(offset).Limit(limit).All(&loginLogList)
	if err != nil {
		return nil, err
	} else {
		return loginLogList, nil
	}
}
func (l *LoginLogDaoImpl) GetLoginLogListByCreatedTime(startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.LoginLogModel, error) {
	var loginLogList []models.LoginLogModel
	err := l.loginLogClient.Find(
		ctx, bson.M{"created_at": bson.M{"$gte": startTime, "$lte": endTime}},
	).Skip(offset).Limit(limit).All(&loginLogList)
	if err != nil {
		return nil, err
	} else {
		return loginLogList, nil
	}
}
func (l *LoginLogDaoImpl) GetLoginLogListByUserUUID(userUUID string, offset, limit int64, ctx context.Context) ([]models.LoginLogModel, error) {
	var loginLogList []models.LoginLogModel
	err := l.loginLogClient.Find(ctx, bson.M{"user_uuid": userUUID}).Skip(offset).Limit(limit).All(&loginLogList)
	if err != nil {
		return nil, err
	} else {
		return loginLogList, nil
	}
}
func (l *LoginLogDaoImpl) GetLoginLogListByIPAddress(ipAddress string, offset, limit int64, ctx context.Context) ([]models.LoginLogModel, error) {
	var loginLogList []models.LoginLogModel
	err := l.loginLogClient.Find(ctx, bson.M{"ip_address": ipAddress}).Skip(offset).Limit(limit).All(&loginLogList)
	if err != nil {
		return nil, err
	} else {
		return loginLogList, nil
	}
}
func (l *LoginLogDaoImpl) InsertLoginLog(loginLog *models.LoginLogModel, ctx context.Context) error {
	_, err := l.loginLogClient.InsertOne(ctx, loginLog)
	return err
}
func (l *LoginLogDaoImpl) DeleteLoginLog(loginLog *models.LoginLogModel, ctx context.Context) error {
	err := l.loginLogClient.RemoveId(ctx, loginLog.LoginLogID)
	return err
}
