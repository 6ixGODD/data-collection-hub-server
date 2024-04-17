package dao

import (
	"context"
	"data-collection-hub-server/models"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
)

type ErrorLogDao interface {
	GetErrorLogById(errorLogId string, mongoClient *qmgo.QmgoClient, ctx context.Context) (*models.ErrorLogModel, error)
	GetErrorLogList(mongoClient *qmgo.QmgoClient, offset, limit int64, desc bool, ctx context.Context) ([]models.ErrorLogModel, error)
	GetErrorLogListByCreatedTime(mongoClient *qmgo.QmgoClient, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.ErrorLogModel, error)
	GetErrorLogListByUserUUID(mongoClient *qmgo.QmgoClient, userUUID string, offset, limit int64, desc bool, ctx context.Context) ([]models.ErrorLogModel, error)
	GetErrorLogListByIPAddress(mongoClient *qmgo.QmgoClient, ipAddress string, offset, limit int64, ctx context.Context) ([]models.ErrorLogModel, error)
	GetErrorLogListByRequestURL(mongoClient *qmgo.QmgoClient, requestURL string, offset, limit int64, ctx context.Context) ([]models.ErrorLogModel, error)
	GetErrorLogListByErrorCode(mongoClient *qmgo.QmgoClient, errorCode string, offset, limit int64, ctx context.Context) ([]models.ErrorLogModel, error)
	GetErrorLogListByFuzzyQuery(mongoClient *qmgo.QmgoClient, query string, offset, limit int64, ctx context.Context) ([]models.ErrorLogModel, error)
	InsertErrorLog(errorLog *models.ErrorLogModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error
	DeleteErrorLog(errorLog *models.ErrorLogModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error
}

type ErrorLogDaoImpl struct{}

func (errorLogDao *ErrorLogDaoImpl) GetErrorLogById(errorLogId string, mongoClient *qmgo.QmgoClient, ctx context.Context) (*models.ErrorLogModel, error) {
	var errorLog models.ErrorLogModel
	err := mongoClient.Find(ctx, bson.M{"_id": errorLogId}).One(&errorLog)
	if err != nil {
		return nil, err
	} else {
		return &errorLog, nil
	}
}

func (errorLogDao *ErrorLogDaoImpl) GetErrorLogList(mongoClient *qmgo.QmgoClient, offset, limit int64, desc bool, ctx context.Context) ([]models.ErrorLogModel, error) {
	var errorLogList []models.ErrorLogModel
	var err error
	if desc {
		err = mongoClient.Find(ctx, bson.M{}).Sort("-created_at").Skip(offset).Limit(limit).All(&errorLogList)
	} else {
		err = mongoClient.Find(ctx, bson.M{}).Skip(offset).Limit(limit).All(&errorLogList)
	}
	if err != nil {
		return nil, err
	} else {
		return errorLogList, nil
	}
}

func (errorLogDao *ErrorLogDaoImpl) GetErrorLogListByCreatedTime(mongoClient *qmgo.QmgoClient, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.ErrorLogModel, error) {
	var errorLogList []models.ErrorLogModel
	var err error
	if desc {
		err = mongoClient.Find(
			ctx, bson.M{"created_at": bson.M{"$gte": startTime, "$lte": endTime}},
		).Sort("-created_at").Skip(offset).Limit(limit).All(&errorLogList)
	} else {
		err = mongoClient.Find(
			ctx, bson.M{"created_at": bson.M{"$gte": startTime, "$lte": endTime}},
		).Skip(offset).Limit(limit).All(&errorLogList)
	}
	if err != nil {
		return nil, err
	} else {
		return errorLogList, nil
	}
}

func (errorLogDao *ErrorLogDaoImpl) GetErrorLogListByUserUUID(mongoClient *qmgo.QmgoClient, userUUID string, offset, limit int64, desc bool, ctx context.Context) ([]models.ErrorLogModel, error) {
	var errorLogList []models.ErrorLogModel
	var err error
	if desc {
		err = mongoClient.Find(ctx, bson.M{"user_uuid": userUUID}).Sort("-created_at").Skip(offset).Limit(limit).All(&errorLogList)
	} else {
		err = mongoClient.Find(ctx, bson.M{"user_uuid": userUUID}).Skip(offset).Limit(limit).All(&errorLogList)
	}
	if err != nil {
		return nil, err
	} else {
		return errorLogList, nil
	}
}

func (errorLogDao *ErrorLogDaoImpl) GetErrorLogListByIPAddress(mongoClient *qmgo.QmgoClient, ipAddress string, offset, limit int64, ctx context.Context) ([]models.ErrorLogModel, error) {
	var errorLogList []models.ErrorLogModel
	err := mongoClient.Find(ctx, bson.M{"ip_address": ipAddress}).Skip(offset).Limit(limit).All(&errorLogList)
	if err != nil {
		return nil, err
	} else {
		return errorLogList, nil
	}
}

func (errorLogDao *ErrorLogDaoImpl) GetErrorLogListByRequestURL(mongoClient *qmgo.QmgoClient, requestURL string, offset, limit int64, ctx context.Context) ([]models.ErrorLogModel, error) {
	var errorLogList []models.ErrorLogModel
	err := mongoClient.Find(ctx, bson.M{"request_url": requestURL}).Skip(offset).Limit(limit).All(&errorLogList)
	if err != nil {
		return nil, err
	} else {
		return errorLogList, nil
	}
}

func (errorLogDao *ErrorLogDaoImpl) GetErrorLogListByErrorCode(mongoClient *qmgo.QmgoClient, errorCode string, offset, limit int64, ctx context.Context) ([]models.ErrorLogModel, error) {
	var errorLogList []models.ErrorLogModel
	err := mongoClient.Find(ctx, bson.M{"error_code": errorCode}).Skip(offset).Limit(limit).All(&errorLogList)
	if err != nil {
		return nil, err
	} else {
		return errorLogList, nil
	}
}

func (errorLogDao *ErrorLogDaoImpl) GetErrorLogListByFuzzyQuery(mongoClient *qmgo.QmgoClient, query string, offset, limit int64, ctx context.Context) ([]models.ErrorLogModel, error) {
	var errorLogList []models.ErrorLogModel
	err := mongoClient.Find(ctx, bson.M{"$or": []bson.M{
		{"user_uuid": bson.M{"$regex": query}},
		{"username": bson.M{"$regex": query}},
		{"ip_address": bson.M{"$regex": query}},
		{"request_url": bson.M{"$regex": query}},
		{"error_code": bson.M{"$regex": query}},
		{"error_msg": bson.M{"$regex": query}},
		{"stack": bson.M{"$regex": query}},
	}}).Skip(offset).Limit(limit).All(&errorLogList)
	if err != nil {
		return nil, err
	} else {
		return errorLogList, nil
	}
}

func (errorLogDao *ErrorLogDaoImpl) InsertErrorLog(errorLog *models.ErrorLogModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error {
	_, err := mongoClient.InsertOne(ctx, errorLog)
	return err
}

func (errorLogDao *ErrorLogDaoImpl) DeleteErrorLog(errorLog *models.ErrorLogModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error {
	panic("implement me") // TODO: Implement
}

func NewErrorLogDaoImpl() *ErrorLogDaoImpl {
	var _ ErrorLogDao = new(ErrorLogDaoImpl) // Ensure that the interface is implemented
	return &ErrorLogDaoImpl{}
}
