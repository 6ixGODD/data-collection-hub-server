package modules

import (
	"context"

	"data-collection-hub-server/dal"
	"data-collection-hub-server/models"
	"go.mongodb.org/mongo-driver/bson"
)

var errorLogCollectionName = "error_log"

type ErrorLogDao interface {
	GetErrorLogById(errorLogId string, ctx context.Context) (*models.ErrorLogModel, error)
	GetErrorLogList(offset, limit int64, desc bool, ctx context.Context) ([]models.ErrorLogModel, error)
	GetErrorLogListByCreatedTime(startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.ErrorLogModel, error)
	GetErrorLogListByUserUUID(userUUID string, offset, limit int64, desc bool, ctx context.Context) ([]models.ErrorLogModel, error)
	GetErrorLogListByIPAddress(ipAddress string, offset, limit int64, ctx context.Context) ([]models.ErrorLogModel, error)
	GetErrorLogListByRequestURL(requestURL string, offset, limit int64, ctx context.Context) ([]models.ErrorLogModel, error)
	GetErrorLogListByErrorCode(errorCode string, offset, limit int64, ctx context.Context) ([]models.ErrorLogModel, error)
	GetErrorLogListByFuzzyQuery(query string, offset, limit int64, ctx context.Context) ([]models.ErrorLogModel, error)
	InsertErrorLog(errorLog *models.ErrorLogModel, ctx context.Context) error
	DeleteErrorLog(errorLog *models.ErrorLogModel, ctx context.Context) error
}

type ErrorLogDaoImpl struct{ *dal.Dao }

func NewErrorLogDao(dao *dal.Dao) ErrorLogDao {
	var _ ErrorLogDao = new(ErrorLogDaoImpl) // Ensure that the interface is implemented
	return &ErrorLogDaoImpl{dao}
}

func (e *ErrorLogDaoImpl) GetErrorLogById(errorLogId string, ctx context.Context) (*models.ErrorLogModel, error) {
	var errorLog models.ErrorLogModel
	collection := e.Dao.MongoDB.Collection(errorLogCollectionName)
	err := collection.Find(ctx, bson.M{"_id": errorLogId}).One(&errorLog)
	if err != nil {
		return nil, err
	} else {
		return &errorLog, nil
	}
}

func (e *ErrorLogDaoImpl) GetErrorLogList(offset, limit int64, desc bool, ctx context.Context) ([]models.ErrorLogModel, error) {
	var errorLogList []models.ErrorLogModel
	var err error
	collection := e.Dao.MongoDB.Collection(errorLogCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{}).Sort("-created_at").Skip(offset).Limit(limit).All(&errorLogList)
	} else {
		err = collection.Find(ctx, bson.M{}).Skip(offset).Limit(limit).All(&errorLogList)
	}
	if err != nil {
		return nil, err
	} else {
		return errorLogList, nil
	}
}

func (e *ErrorLogDaoImpl) GetErrorLogListByCreatedTime(startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.ErrorLogModel, error) {
	var errorLogList []models.ErrorLogModel
	var err error
	collection := e.Dao.MongoDB.Collection(errorLogCollectionName)
	if desc {
		err = collection.Find(
			ctx, bson.M{"created_at": bson.M{"$gte": startTime, "$lte": endTime}},
		).Sort("-created_at").Skip(offset).Limit(limit).All(&errorLogList)
	} else {
		err = collection.Find(
			ctx, bson.M{"created_at": bson.M{"$gte": startTime, "$lte": endTime}},
		).Skip(offset).Limit(limit).All(&errorLogList)
	}
	if err != nil {
		return nil, err
	} else {
		return errorLogList, nil
	}
}

func (e *ErrorLogDaoImpl) GetErrorLogListByUserUUID(userUUID string, offset, limit int64, desc bool, ctx context.Context) ([]models.ErrorLogModel, error) {
	var errorLogList []models.ErrorLogModel
	var err error
	collection := e.Dao.MongoDB.Collection(errorLogCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{"user_uuid": userUUID}).Sort("-created_at").Skip(offset).Limit(limit).All(&errorLogList)
	} else {
		err = collection.Find(ctx, bson.M{"user_uuid": userUUID}).Skip(offset).Limit(limit).All(&errorLogList)
	}
	if err != nil {
		return nil, err
	} else {
		return errorLogList, nil
	}
}

func (e *ErrorLogDaoImpl) GetErrorLogListByIPAddress(ipAddress string, offset, limit int64, ctx context.Context) ([]models.ErrorLogModel, error) {
	var errorLogList []models.ErrorLogModel
	collection := e.Dao.MongoDB.Collection(errorLogCollectionName)
	err := collection.Find(ctx, bson.M{"ip_address": ipAddress}).Skip(offset).Limit(limit).All(&errorLogList)
	if err != nil {
		return nil, err
	} else {
		return errorLogList, nil
	}
}

func (e *ErrorLogDaoImpl) GetErrorLogListByRequestURL(requestURL string, offset, limit int64, ctx context.Context) ([]models.ErrorLogModel, error) {
	var errorLogList []models.ErrorLogModel
	collection := e.Dao.MongoDB.Collection(errorLogCollectionName)
	err := collection.Find(ctx, bson.M{"request_url": requestURL}).Skip(offset).Limit(limit).All(&errorLogList)
	if err != nil {
		return nil, err
	} else {
		return errorLogList, nil
	}
}

func (e *ErrorLogDaoImpl) GetErrorLogListByErrorCode(errorCode string, offset, limit int64, ctx context.Context) ([]models.ErrorLogModel, error) {
	var errorLogList []models.ErrorLogModel
	collection := e.Dao.MongoDB.Collection(errorLogCollectionName)
	err := collection.Find(ctx, bson.M{"error_code": errorCode}).Skip(offset).Limit(limit).All(&errorLogList)
	if err != nil {
		return nil, err
	} else {
		return errorLogList, nil
	}
}

func (e *ErrorLogDaoImpl) GetErrorLogListByFuzzyQuery(query string, offset, limit int64, ctx context.Context) ([]models.ErrorLogModel, error) {
	var errorLogList []models.ErrorLogModel
	collection := e.Dao.MongoDB.Collection(errorLogCollectionName)
	err := collection.Find(ctx, bson.M{"$or": []bson.M{
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

func (e *ErrorLogDaoImpl) InsertErrorLog(errorLog *models.ErrorLogModel, ctx context.Context) error {
	collection := e.Dao.MongoDB.Collection(errorLogCollectionName)
	_, err := collection.InsertOne(ctx, errorLog)
	return err
}

func (e *ErrorLogDaoImpl) DeleteErrorLog(errorLog *models.ErrorLogModel, ctx context.Context) error {
	collection := e.Dao.MongoDB.Collection(errorLogCollectionName)
	err := collection.RemoveId(ctx, errorLog.ErrorLogID)
	return err
}
