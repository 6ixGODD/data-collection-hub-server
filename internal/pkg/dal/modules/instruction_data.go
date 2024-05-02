package modules

import (
	"context"

	"data-collection-hub-server/internal/pkg/dal"
	"data-collection-hub-server/internal/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var instructionDataCollectionName = "instruction_data"

type InstructionDataDao interface {
	GetInstructionDataById(instructionDataId string, ctx context.Context) (*models.InstructionDataModel, error)
	GetInstructionDataList(offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByUserUUID(userUUID string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByFuzzyQuery(query string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByFuzzyQueryAndTheme(query, theme string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByFuzzyQueryAndThemeAndStatusCode(query, theme, statusCode string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByFuzzyQueryAndThemeAndCreatedTime(query, theme, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByFuzzyQueryAndThemeAndStatusCodeAndCreatedTime(query, theme, statusCode, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByFuzzyQueryAndStatusCode(query, statusCode string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByFuzzyQueryAndStatusCodeAndCreatedTime(query, statusCode, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByFuzzyQueryAndCreatedTime(query, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByTheme(theme string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByThemeAndStatusCode(theme, statusCode string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByThemeAndCreatedTime(theme, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByThemeAndStatusCodeAndCreatedTime(theme, statusCode, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByStatusCode(statusCode string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByStatusCodeAndCreatedTime(statusCode, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByCreatedTime(startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByUpdatedTime(startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	InsertInstructionData(instructionData *models.InstructionDataModel, ctx context.Context) error
	UpdateInstructionData(instructionData *models.InstructionDataModel, ctx context.Context) error
	DeleteInstructionData(instructionData *models.InstructionDataModel, ctx context.Context) error
}

type InstructionDataDaoImpl struct{ *dal.Dao }

func NewInstructionDataDao(dao *dal.Dao) InstructionDataDao {
	var _ InstructionDataDao = new(InstructionDataDaoImpl)
	return &InstructionDataDaoImpl{
		Dao: dao,
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataById(instructionDataId string, ctx context.Context) (*models.InstructionDataModel, error) {
	var instructionData models.InstructionDataModel
	collection := i.Dao.MongoClient.MongoDatabase.Collection(instructionDataCollectionName)
	err := collection.Find(ctx, bson.M{"instruction_data_id": instructionDataId}).One(&instructionData)
	if err != nil {
		i.Dao.Logger.Logger.Error(
			"InstructionDataDaoImpl.GetInstructionDataById",
			zap.Field{Key: "instructionDataId", Type: zapcore.StringType, String: instructionDataId},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		i.Dao.Logger.Logger.Info(
			"InstructionDataDaoImpl.GetInstructionDataById",
			zap.Field{Key: "instructionDataId", Type: zapcore.StringType, String: instructionDataId},
		)
		return &instructionData, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataList(offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.MongoClient.MongoDatabase.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{}).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(ctx, bson.M{}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		i.Dao.Logger.Logger.Error(
			"InstructionDataDaoImpl.GetInstructionDataList",
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		i.Dao.Logger.Logger.Info(
			"InstructionDataDaoImpl.GetInstructionDataList",
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByUserUUID(userUUID string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.MongoClient.MongoDatabase.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{"user_uuid": userUUID}).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(ctx, bson.M{"user_uuid": userUUID}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		i.Dao.Logger.Logger.Error(
			"InstructionDataDaoImpl.GetInstructionDataListByUserUUID",
			zap.Field{Key: "userUUID", Type: zapcore.StringType, String: userUUID},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		i.Dao.Logger.Logger.Info(
			"InstructionDataDaoImpl.GetInstructionDataListByUserUUID",
			zap.Field{Key: "userUUID", Type: zapcore.StringType, String: userUUID},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByFuzzyQuery(query string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.MongoClient.MongoDatabase.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{"$text": bson.M{"$search": query}}).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(ctx, bson.M{"$text": bson.M{"$search": query}}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		i.Dao.Logger.Logger.Error(
			"InstructionDataDaoImpl.GetInstructionDataListByFuzzyQuery",
			zap.Field{Key: "query", Type: zapcore.StringType, String: query},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		i.Dao.Logger.Logger.Info(
			"InstructionDataDaoImpl.GetInstructionDataListByFuzzyQuery",
			zap.Field{Key: "query", Type: zapcore.StringType, String: query},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByFuzzyQueryAndTheme(query, theme string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.MongoClient.MongoDatabase.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{"$text": bson.M{"$search": query}, "theme": theme}).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(ctx, bson.M{"$text": bson.M{"$search": query}, "theme": theme}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		i.Dao.Logger.Logger.Error(
			"InstructionDataDaoImpl.GetInstructionDataListByFuzzyQueryAndTheme",
			zap.Field{Key: "query", Type: zapcore.StringType, String: query},
			zap.Field{Key: "theme", Type: zapcore.StringType, String: theme},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		i.Dao.Logger.Logger.Info(
			"InstructionDataDaoImpl.GetInstructionDataListByFuzzyQueryAndTheme",
			zap.Field{Key: "query", Type: zapcore.StringType, String: query},
			zap.Field{Key: "theme", Type: zapcore.StringType, String: theme},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByFuzzyQueryAndThemeAndStatusCode(query, theme, statusCode string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.MongoClient.MongoDatabase.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{"$text": bson.M{"$search": query}, "theme": theme, "status_code": statusCode}).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(ctx, bson.M{"$text": bson.M{"$search": query}, "theme": theme, "status_code": statusCode}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		i.Dao.Logger.Logger.Error(
			"InstructionDataDaoImpl.GetInstructionDataListByFuzzyQueryAndThemeAndStatusCode",
			zap.Field{Key: "query", Type: zapcore.StringType, String: query},
			zap.Field{Key: "theme", Type: zapcore.StringType, String: theme},
			zap.Field{Key: "statusCode", Type: zapcore.StringType, String: statusCode},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		i.Dao.Logger.Logger.Info(
			"InstructionDataDaoImpl.GetInstructionDataListByFuzzyQueryAndThemeAndStatusCode",
			zap.Field{Key: "query", Type: zapcore.StringType, String: query},
			zap.Field{Key: "theme", Type: zapcore.StringType, String: theme},
			zap.Field{Key: "statusCode", Type: zapcore.StringType, String: statusCode},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByFuzzyQueryAndThemeAndCreatedTime(query, theme, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.MongoClient.MongoDatabase.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{"$text": bson.M{"$search": query}, "theme": theme, "created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(ctx, bson.M{"$text": bson.M{"$search": query}, "theme": theme, "created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		i.Dao.Logger.Logger.Error(
			"InstructionDataDaoImpl.GetInstructionDataListByFuzzyQueryAndThemeAndCreatedTime",
			zap.Field{Key: "query", Type: zapcore.StringType, String: query},
			zap.Field{Key: "theme", Type: zapcore.StringType, String: theme},
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		i.Dao.Logger.Logger.Info(
			"InstructionDataDaoImpl.GetInstructionDataListByFuzzyQueryAndThemeAndCreatedTime",
			zap.Field{Key: "query", Type: zapcore.StringType, String: query},
			zap.Field{Key: "theme", Type: zapcore.StringType, String: theme},
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByFuzzyQueryAndThemeAndStatusCodeAndCreatedTime(query, theme, statusCode, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.MongoClient.MongoDatabase.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{"$text": bson.M{"$search": query}, "theme": theme, "status_code": statusCode, "created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(ctx, bson.M{"$text": bson.M{"$search": query}, "theme": theme, "status_code": statusCode, "created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		i.Dao.Logger.Logger.Error(
			"InstructionDataDaoImpl.GetInstructionDataListByFuzzyQueryAndThemeAndStatusCodeAndCreatedTime",
			zap.Field{Key: "query", Type: zapcore.StringType, String: query},
			zap.Field{Key: "theme", Type: zapcore.StringType, String: theme},
			zap.Field{Key: "statusCode", Type: zapcore.StringType, String: statusCode},
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		i.Dao.Logger.Logger.Info(
			"InstructionDataDaoImpl.GetInstructionDataListByFuzzyQueryAndThemeAndStatusCodeAndCreatedTime",
			zap.Field{Key: "query", Type: zapcore.StringType, String: query},
			zap.Field{Key: "theme", Type: zapcore.StringType, String: theme},
			zap.Field{Key: "statusCode", Type: zapcore.StringType, String: statusCode},
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByFuzzyQueryAndStatusCode(query, statusCode string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.MongoClient.MongoDatabase.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{"$text": bson.M{"$search": query}, "status_code": statusCode}).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(ctx, bson.M{"$text": bson.M{"$search": query}, "status_code": statusCode}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		i.Dao.Logger.Logger.Error(
			"InstructionDataDaoImpl.GetInstructionDataListByFuzzyQueryAndStatusCode",
			zap.Field{Key: "query", Type: zapcore.StringType, String: query},
			zap.Field{Key: "statusCode", Type: zapcore.StringType, String: statusCode},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		i.Dao.Logger.Logger.Info(
			"InstructionDataDaoImpl.GetInstructionDataListByFuzzyQueryAndStatusCode",
			zap.Field{Key: "query", Type: zapcore.StringType, String: query},
			zap.Field{Key: "statusCode", Type: zapcore.StringType, String: statusCode},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByFuzzyQueryAndStatusCodeAndCreatedTime(query, statusCode, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.MongoClient.MongoDatabase.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{"$text": bson.M{"$search": query}, "status_code": statusCode, "created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(ctx, bson.M{"$text": bson.M{"$search": query}, "status_code": statusCode, "created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		i.Dao.Logger.Logger.Error(
			"InstructionDataDaoImpl.GetInstructionDataListByFuzzyQueryAndStatusCodeAndCreatedTime",
			zap.Field{Key: "query", Type: zapcore.StringType, String: query},
			zap.Field{Key: "statusCode", Type: zapcore.StringType, String: statusCode},
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		i.Dao.Logger.Logger.Info(
			"InstructionDataDaoImpl.GetInstructionDataListByFuzzyQueryAndStatusCodeAndCreatedTime",
			zap.Field{Key: "query", Type: zapcore.StringType, String: query},
			zap.Field{Key: "statusCode", Type: zapcore.StringType, String: statusCode},
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByFuzzyQueryAndCreatedTime(query, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.MongoClient.MongoDatabase.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{"$text": bson.M{"$search": query}, "created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(ctx, bson.M{"$text": bson.M{"$search": query}, "created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		i.Dao.Logger.Logger.Error(
			"InstructionDataDaoImpl.GetInstructionDataListByFuzzyQueryAndCreatedTime",
			zap.Field{Key: "query", Type: zapcore.StringType, String: query},
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		i.Dao.Logger.Logger.Info(
			"InstructionDataDaoImpl.GetInstructionDataListByFuzzyQueryAndCreatedTime",
			zap.Field{Key: "query", Type: zapcore.StringType, String: query},
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByTheme(theme string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.MongoClient.MongoDatabase.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{"theme": theme}).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(ctx, bson.M{"theme": theme}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		i.Dao.Logger.Logger.Error(
			"InstructionDataDaoImpl.GetInstructionDataListByTheme",
			zap.Field{Key: "theme", Type: zapcore.StringType, String: theme},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		i.Dao.Logger.Logger.Info(
			"InstructionDataDaoImpl.GetInstructionDataListByTheme",
			zap.Field{Key: "theme", Type: zapcore.StringType, String: theme},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByThemeAndStatusCode(theme, statusCode string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.MongoClient.MongoDatabase.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{"theme": theme, "status_code": statusCode}).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(ctx, bson.M{"theme": theme, "status_code": statusCode}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		i.Dao.Logger.Logger.Error(
			"InstructionDataDaoImpl.GetInstructionDataListByThemeAndStatusCode",
			zap.Field{Key: "theme", Type: zapcore.StringType, String: theme},
			zap.Field{Key: "statusCode", Type: zapcore.StringType, String: statusCode},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		i.Dao.Logger.Logger.Info(
			"InstructionDataDaoImpl.GetInstructionDataListByThemeAndStatusCode",
			zap.Field{Key: "theme", Type: zapcore.StringType, String: theme},
			zap.Field{Key: "statusCode", Type: zapcore.StringType, String: statusCode},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByThemeAndCreatedTime(theme, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.MongoClient.MongoDatabase.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{"theme": theme, "created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(ctx, bson.M{"theme": theme, "created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		i.Dao.Logger.Logger.Error(
			"InstructionDataDaoImpl.GetInstructionDataListByThemeAndCreatedTime",
			zap.Field{Key: "theme", Type: zapcore.StringType, String: theme},
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		i.Dao.Logger.Logger.Info(
			"InstructionDataDaoImpl.GetInstructionDataListByThemeAndCreatedTime",
			zap.Field{Key: "theme", Type: zapcore.StringType, String: theme},
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByThemeAndStatusCodeAndCreatedTime(theme, statusCode, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.MongoClient.MongoDatabase.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{"theme": theme, "status_code": statusCode, "created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(ctx, bson.M{"theme": theme, "status_code": statusCode, "created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		i.Dao.Logger.Logger.Error(
			"InstructionDataDaoImpl.GetInstructionDataListByThemeAndStatusCodeAndCreatedTime",
			zap.Field{Key: "theme", Type: zapcore.StringType, String: theme},
			zap.Field{Key: "statusCode", Type: zapcore.StringType, String: statusCode},
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		i.Dao.Logger.Logger.Info(
			"InstructionDataDaoImpl.GetInstructionDataListByThemeAndStatusCodeAndCreatedTime",
			zap.Field{Key: "theme", Type: zapcore.StringType, String: theme},
			zap.Field{Key: "statusCode", Type: zapcore.StringType, String: statusCode},
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByStatusCode(statusCode string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.MongoClient.MongoDatabase.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{"status_code": statusCode}).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(ctx, bson.M{"status_code": statusCode}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		i.Dao.Logger.Logger.Error(
			"InstructionDataDaoImpl.GetInstructionDataListByStatusCode",
			zap.Field{Key: "statusCode", Type: zapcore.StringType, String: statusCode},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		i.Dao.Logger.Logger.Info(
			"InstructionDataDaoImpl.GetInstructionDataListByStatusCode",
			zap.Field{Key: "statusCode", Type: zapcore.StringType, String: statusCode},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByStatusCodeAndCreatedTime(statusCode, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.MongoClient.MongoDatabase.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{"status_code": statusCode, "created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(ctx, bson.M{"status_code": statusCode, "created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		i.Dao.Logger.Logger.Error(
			"InstructionDataDaoImpl.GetInstructionDataListByStatusCodeAndCreatedTime",
			zap.Field{Key: "statusCode", Type: zapcore.StringType, String: statusCode},
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		i.Dao.Logger.Logger.Info(
			"InstructionDataDaoImpl.GetInstructionDataListByStatusCodeAndCreatedTime",
			zap.Field{Key: "statusCode", Type: zapcore.StringType, String: statusCode},
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByCreatedTime(startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.MongoClient.MongoDatabase.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{"created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(ctx, bson.M{"created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		i.Dao.Logger.Logger.Error(
			"InstructionDataDaoImpl.GetInstructionDataListByCreatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		i.Dao.Logger.Logger.Info(
			"InstructionDataDaoImpl.GetInstructionDataListByCreatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByUpdatedTime(startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.MongoClient.MongoDatabase.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{"updated_time": bson.M{"$gte": startTime, "$lte": endTime}}).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(ctx, bson.M{"updated_time": bson.M{"$gte": startTime, "$lte": endTime}}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		i.Dao.Logger.Logger.Error(
			"InstructionDataDaoImpl.GetInstructionDataListByUpdatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		i.Dao.Logger.Logger.Info(
			"InstructionDataDaoImpl.GetInstructionDataListByUpdatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) InsertInstructionData(instructionData *models.InstructionDataModel, ctx context.Context) error {
	collection := i.Dao.MongoClient.MongoDatabase.Collection(instructionDataCollectionName)
	result, err := collection.InsertOne(ctx, instructionData)
	if err != nil {
		i.Dao.Logger.Logger.Error(
			"InstructionDataDaoImpl.InsertInstructionData",
			zap.Field{Key: "instructionData", Type: zapcore.ObjectMarshalerType, Interface: instructionData},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		i.Dao.Logger.Logger.Info(
			"InstructionDataDaoImpl.InsertInstructionData",
			zap.Field{Key: "instructionData", Type: zapcore.ObjectMarshalerType, Interface: instructionData},
			zap.Field{Key: "result", Type: zapcore.ObjectMarshalerType, Interface: result},
		)
	}
	return err
}

func (i InstructionDataDaoImpl) UpdateInstructionData(instructionData *models.InstructionDataModel, ctx context.Context) error {
	collection := i.Dao.MongoClient.MongoDatabase.Collection(instructionDataCollectionName)
	err := collection.UpdateOne(ctx, bson.M{"instruction_data_id": instructionData.InstructionDataID}, bson.M{"$set": instructionData})
	if err != nil {
		i.Dao.Logger.Logger.Error(
			"InstructionDataDaoImpl.UpdateInstructionData",
			zap.Field{Key: "instructionData", Type: zapcore.ObjectMarshalerType, Interface: instructionData},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		i.Dao.Logger.Logger.Info(
			"InstructionDataDaoImpl.UpdateInstructionData",
			zap.Field{Key: "instructionData", Type: zapcore.ObjectMarshalerType, Interface: instructionData},
		)
	}
	return err
}

func (i InstructionDataDaoImpl) DeleteInstructionData(instructionData *models.InstructionDataModel, ctx context.Context) error {
	collection := i.Dao.MongoClient.MongoDatabase.Collection(instructionDataCollectionName)
	err := collection.RemoveId(ctx, instructionData.InstructionDataID)
	if err != nil {
		i.Dao.Logger.Logger.Error(
			"InstructionDataDaoImpl.DeleteInstructionData",
			zap.Field{Key: "instructionData", Type: zapcore.ObjectMarshalerType, Interface: instructionData},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		i.Dao.Logger.Logger.Info(
			"InstructionDataDaoImpl.DeleteInstructionData",
			zap.Field{Key: "instructionData", Type: zapcore.ObjectMarshalerType, Interface: instructionData},
		)
	}
	return err
}
