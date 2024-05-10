package mods

import (
	"context"
	"time"

	"data-collection-hub-server/internal/pkg/dal"
	"data-collection-hub-server/internal/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const instructionDataCollectionName = "instruction_data"

type InstructionDataDao interface {
	GetInstructionDataById(instructionDataID primitive.ObjectID, ctx context.Context) (
		*models.InstructionDataModel, error,
	)
	GetInstructionDataList(offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByUserID(
		userID primitive.ObjectID, offset, limit int64, desc bool, ctx context.Context,
	) ([]models.InstructionDataModel, error)
	GetInstructionDataListByFuzzyQuery(
		query string, offset, limit int64, desc bool, ctx context.Context,
	) ([]models.InstructionDataModel, error)
	GetInstructionDataListByFuzzyQueryAndTheme(
		query, theme string, offset, limit int64, desc bool, ctx context.Context,
	) ([]models.InstructionDataModel, error)
	GetInstructionDataListByFuzzyQueryAndThemeAndStatusCode(
		query, theme, statusCode string, offset, limit int64, desc bool, ctx context.Context,
	) ([]models.InstructionDataModel, error)
	GetInstructionDataListByFuzzyQueryAndThemeAndCreatedTime(
		query, theme string, startTime, endTime time.Time, offset, limit int64, desc bool, ctx context.Context,
	) ([]models.InstructionDataModel, error)
	GetInstructionDataListByFuzzyQueryAndThemeAndStatusCodeAndCreatedTime(
		query, theme, statusCode string, startTime, endTime time.Time, offset, limit int64, desc bool,
		ctx context.Context,
	) ([]models.InstructionDataModel, error)
	GetInstructionDataListByFuzzyQueryAndStatusCode(
		query, statusCode string, offset, limit int64, desc bool, ctx context.Context,
	) ([]models.InstructionDataModel, error)
	GetInstructionDataListByFuzzyQueryAndStatusCodeAndCreatedTime(
		query, statusCode string, startTime, endTime time.Time, offset, limit int64, desc bool, ctx context.Context,
	) ([]models.InstructionDataModel, error)
	GetInstructionDataListByFuzzyQueryAndCreatedTime(
		query string, startTime, endTime time.Time, offset, limit int64, desc bool, ctx context.Context,
	) ([]models.InstructionDataModel, error)
	GetInstructionDataListByTheme(
		theme string, offset, limit int64, desc bool, ctx context.Context,
	) ([]models.InstructionDataModel, error)
	GetInstructionDataListByThemeAndStatusCode(
		theme, statusCode string, offset, limit int64, desc bool, ctx context.Context,
	) ([]models.InstructionDataModel, error)
	GetInstructionDataListByThemeAndCreatedTime(
		theme string, startTime, endTime time.Time, offset, limit int64, desc bool, ctx context.Context,
	) ([]models.InstructionDataModel, error)
	GetInstructionDataListByThemeAndStatusCodeAndCreatedTime(
		theme, statusCode string, startTime, endTime time.Time, offset, limit int64, desc bool, ctx context.Context,
	) ([]models.InstructionDataModel, error)
	GetInstructionDataListByStatusCode(
		statusCode string, offset, limit int64, desc bool, ctx context.Context,
	) ([]models.InstructionDataModel, error)
	GetInstructionDataListByStatusCodeAndCreatedTime(
		statusCode string, startTime, endTime time.Time, offset, limit int64, desc bool, ctx context.Context,
	) ([]models.InstructionDataModel, error)
	GetInstructionDataListByCreatedTime(
		startTime, endTime time.Time, offset, limit int64, desc bool, ctx context.Context,
	) ([]models.InstructionDataModel, error)
	GetInstructionDataListByUpdatedTime(
		startTime, endTime time.Time, offset, limit int64, desc bool, ctx context.Context,
	) ([]models.InstructionDataModel, error)
	InsertInstructionData(
		userID primitive.ObjectID,
		username, rowInstruction, rowInput, rowOutput, theme, source, note, statusCode, statusMessage string,
		ctx context.Context,
	) (primitive.ObjectID, error)
	UpdateInstructionData(
		instructionDataID, userID primitive.ObjectID,
		username, rowInstruction, rowInput, rowOutput, theme, source, note, statusCode, statusMessage string,
		ctx context.Context,
	) error
	SoftDeleteInstructionData(instructionDataID primitive.ObjectID, ctx context.Context) error
	DeleteInstructionData(instructionDataID primitive.ObjectID, ctx context.Context) error
}

type InstructionDataDaoImpl struct{ *dal.Dao }

func NewInstructionDataDao(dao *dal.Dao) InstructionDataDao {
	var _ InstructionDataDao = (*InstructionDataDaoImpl)(nil) // Ensure that the interface is implemented
	return &InstructionDataDaoImpl{dao}
}

func (i InstructionDataDaoImpl) GetInstructionDataById(
	instructionDataID primitive.ObjectID, ctx context.Context,
) (*models.InstructionDataModel, error) {
	var instructionData models.InstructionDataModel
	collection := i.Dao.Mongo.MongoDatabase.Collection(instructionDataCollectionName)
	err := collection.Find(
		ctx,
		bson.M{
			"instruction_data_id": instructionDataID,
			"deleted":             false,
		},
	).One(&instructionData)
	if err != nil {
		i.Dao.Zap.Logger.Error(
			"InstructionDataDaoImpl.GetInstructionDataById",
			zap.Field{Key: "instructionDataID", Type: zapcore.StringType, String: instructionDataID.Hex()},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		i.Dao.Zap.Logger.Info(
			"InstructionDataDaoImpl.GetInstructionDataById",
			zap.Field{Key: "instructionDataID", Type: zapcore.StringType, String: instructionDataID.Hex()},
		)
		return &instructionData, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataList(
	offset, limit int64, desc bool, ctx context.Context,
) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.Mongo.MongoDatabase.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(
			ctx,
			bson.M{
				"deleted": false,
			},
		).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(
			ctx, bson.M{
				"deleted": false,
			},
		).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		i.Dao.Zap.Logger.Error(
			"InstructionDataDaoImpl.GetInstructionDataList",
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		i.Dao.Zap.Logger.Info(
			"InstructionDataDaoImpl.GetInstructionDataList",
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByUserID(
	userID primitive.ObjectID, offset, limit int64, desc bool, ctx context.Context,
) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.Mongo.MongoDatabase.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(
			ctx,
			bson.M{
				"user_id": userID,
				"deleted": false,
			},
		).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(
			ctx,
			bson.M{
				"user_id": userID,
				"deleted": false,
			},
		).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		i.Dao.Zap.Logger.Error(
			"InstructionDataDaoImpl.GetInstructionDataListByUserID",
			zap.Field{Key: "userID", Type: zapcore.StringType, String: userID.Hex()},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		i.Dao.Zap.Logger.Info(
			"InstructionDataDaoImpl.GetInstructionDataListByUserID",
			zap.Field{Key: "userID", Type: zapcore.StringType, String: userID.Hex()},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByFuzzyQuery(
	query string, offset, limit int64, desc bool, ctx context.Context,
) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.Mongo.MongoDatabase.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(
			ctx,
			bson.M{
				"$text":   bson.M{"$search": query},
				"deleted": false,
			},
		).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(
			ctx,
			bson.M{
				"$text":   bson.M{"$search": query},
				"deleted": false,
			},
		).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		i.Dao.Zap.Logger.Error(
			"InstructionDataDaoImpl.GetInstructionDataListByFuzzyQuery",
			zap.Field{Key: "query", Type: zapcore.StringType, String: query},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		i.Dao.Zap.Logger.Info(
			"InstructionDataDaoImpl.GetInstructionDataListByFuzzyQuery",
			zap.Field{Key: "query", Type: zapcore.StringType, String: query},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByFuzzyQueryAndTheme(
	query, theme string, offset, limit int64, desc bool, ctx context.Context,
) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.Mongo.MongoDatabase.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(
			ctx,
			bson.M{
				"$text":   bson.M{"$search": query},
				"theme":   theme,
				"deleted": false,
			},
		).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(
			ctx,
			bson.M{
				"$text":   bson.M{"$search": query},
				"theme":   theme,
				"deleted": false,
			},
		).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		i.Dao.Zap.Logger.Error(
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
		i.Dao.Zap.Logger.Info(
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

func (i InstructionDataDaoImpl) GetInstructionDataListByFuzzyQueryAndThemeAndStatusCode(
	query, theme, statusCode string, offset, limit int64, desc bool, ctx context.Context,
) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.Mongo.MongoDatabase.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(
			ctx,
			bson.M{
				"$text":       bson.M{"$search": query},
				"theme":       theme,
				"status_code": statusCode,
				"deleted":     false,
			},
		).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(
			ctx,
			bson.M{
				"$text":       bson.M{"$search": query},
				"theme":       theme,
				"status_code": statusCode,
				"deleted":     false,
			},
		).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		i.Dao.Zap.Logger.Error(
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
		i.Dao.Zap.Logger.Info(
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

func (i InstructionDataDaoImpl) GetInstructionDataListByFuzzyQueryAndThemeAndCreatedTime(
	query, theme string, startTime, endTime time.Time, offset, limit int64, desc bool, ctx context.Context,
) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.Mongo.MongoDatabase.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(
			ctx,
			bson.M{
				"$text":        bson.M{"$search": query},
				"theme":        theme,
				"created_time": bson.M{"$gte": startTime, "$lte": endTime},
				"deleted":      false,
			},
		).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(
			ctx, bson.M{
				"$text":        bson.M{"$search": query},
				"theme":        theme,
				"created_time": bson.M{"$gte": startTime, "$lte": endTime},
				"deleted":      false,
			},
		).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		i.Dao.Zap.Logger.Error(
			"InstructionDataDaoImpl.GetInstructionDataListByFuzzyQueryAndThemeAndCreatedTime",
			zap.Field{Key: "query", Type: zapcore.StringType, String: query},
			zap.Field{Key: "theme", Type: zapcore.StringType, String: theme},
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		i.Dao.Zap.Logger.Info(
			"InstructionDataDaoImpl.GetInstructionDataListByFuzzyQueryAndThemeAndCreatedTime",
			zap.Field{Key: "query", Type: zapcore.StringType, String: query},
			zap.Field{Key: "theme", Type: zapcore.StringType, String: theme},
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByFuzzyQueryAndThemeAndStatusCodeAndCreatedTime(
	query, theme, statusCode string, startTime, endTime time.Time, offset, limit int64, desc bool, ctx context.Context,
) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.Mongo.MongoDatabase.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(
			ctx,
			bson.M{
				"$text":        bson.M{"$search": query},
				"theme":        theme,
				"status_code":  statusCode,
				"created_time": bson.M{"$gte": startTime, "$lte": endTime},
				"deleted":      false,
			},
		).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(
			ctx, bson.M{
				"$text": bson.M{"$search": query}, "theme": theme, "status_code": statusCode,
				"created_time": bson.M{"$gte": startTime, "$lte": endTime},
			},
		).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		i.Dao.Zap.Logger.Error(
			"InstructionDataDaoImpl.GetInstructionDataListByFuzzyQueryAndThemeAndStatusCodeAndCreatedTime",
			zap.Field{Key: "query", Type: zapcore.StringType, String: query},
			zap.Field{Key: "theme", Type: zapcore.StringType, String: theme},
			zap.Field{Key: "statusCode", Type: zapcore.StringType, String: statusCode},
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		i.Dao.Zap.Logger.Info(
			"InstructionDataDaoImpl.GetInstructionDataListByFuzzyQueryAndThemeAndStatusCodeAndCreatedTime",
			zap.Field{Key: "query", Type: zapcore.StringType, String: query},
			zap.Field{Key: "theme", Type: zapcore.StringType, String: theme},
			zap.Field{Key: "statusCode", Type: zapcore.StringType, String: statusCode},
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByFuzzyQueryAndStatusCode(
	query, statusCode string, offset, limit int64, desc bool, ctx context.Context,
) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.Mongo.MongoDatabase.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(
			ctx,
			bson.M{
				"$text":       bson.M{"$search": query},
				"status_code": statusCode,
				"deleted":     false,
			},
		).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(
			ctx, bson.M{"$text": bson.M{"$search": query}, "status_code": statusCode},
		).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		i.Dao.Zap.Logger.Error(
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
		i.Dao.Zap.Logger.Info(
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

func (i InstructionDataDaoImpl) GetInstructionDataListByFuzzyQueryAndStatusCodeAndCreatedTime(
	query, statusCode string, startTime, endTime time.Time, offset, limit int64, desc bool, ctx context.Context,
) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.Mongo.MongoDatabase.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(
			ctx,
			bson.M{
				"$text":        bson.M{"$search": query},
				"status_code":  statusCode,
				"created_time": bson.M{"$gte": startTime, "$lte": endTime},
				"deleted":      false,
			},
		).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(
			ctx,
			bson.M{
				"$text":        bson.M{"$search": query},
				"status_code":  statusCode,
				"created_time": bson.M{"$gte": startTime, "$lte": endTime},
				"deleted":      false,
			},
		).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		i.Dao.Zap.Logger.Error(
			"InstructionDataDaoImpl.GetInstructionDataListByFuzzyQueryAndStatusCodeAndCreatedTime",
			zap.Field{Key: "query", Type: zapcore.StringType, String: query},
			zap.Field{Key: "statusCode", Type: zapcore.StringType, String: statusCode},
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		i.Dao.Zap.Logger.Info(
			"InstructionDataDaoImpl.GetInstructionDataListByFuzzyQueryAndStatusCodeAndCreatedTime",
			zap.Field{Key: "query", Type: zapcore.StringType, String: query},
			zap.Field{Key: "statusCode", Type: zapcore.StringType, String: statusCode},
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByFuzzyQueryAndCreatedTime(
	query string, startTime, endTime time.Time, offset, limit int64, desc bool, ctx context.Context,
) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.Mongo.MongoDatabase.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(
			ctx,
			bson.M{
				"$text":        bson.M{"$search": query},
				"created_time": bson.M{"$gte": startTime, "$lte": endTime},
				"deleted":      false,
			},
		).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(
			ctx,
			bson.M{
				"$text":        bson.M{"$search": query},
				"created_time": bson.M{"$gte": startTime, "$lte": endTime},
				"deleted":      false,
			},
		).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		i.Dao.Zap.Logger.Error(
			"InstructionDataDaoImpl.GetInstructionDataListByFuzzyQueryAndCreatedTime",
			zap.Field{Key: "query", Type: zapcore.StringType, String: query},
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		i.Dao.Zap.Logger.Info(
			"InstructionDataDaoImpl.GetInstructionDataListByFuzzyQueryAndCreatedTime",
			zap.Field{Key: "query", Type: zapcore.StringType, String: query},
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByTheme(
	theme string, offset, limit int64, desc bool, ctx context.Context,
) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.Mongo.MongoDatabase.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(
			ctx, bson.M{"theme": theme},
		).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(ctx, bson.M{"theme": theme}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		i.Dao.Zap.Logger.Error(
			"InstructionDataDaoImpl.GetInstructionDataListByTheme",
			zap.Field{Key: "theme", Type: zapcore.StringType, String: theme},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		i.Dao.Zap.Logger.Info(
			"InstructionDataDaoImpl.GetInstructionDataListByTheme",
			zap.Field{Key: "theme", Type: zapcore.StringType, String: theme},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByThemeAndStatusCode(
	theme, statusCode string, offset, limit int64, desc bool, ctx context.Context,
) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.Mongo.MongoDatabase.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(
			ctx,
			bson.M{
				"theme":       theme,
				"status_code": statusCode,
				"deleted":     false,
			},
		).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(
			ctx,
			bson.M{
				"theme":       theme,
				"status_code": statusCode,
				"deleted":     false,
			},
		).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		i.Dao.Zap.Logger.Error(
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
		i.Dao.Zap.Logger.Info(
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

func (i InstructionDataDaoImpl) GetInstructionDataListByThemeAndCreatedTime(
	theme string, startTime, endTime time.Time, offset, limit int64, desc bool, ctx context.Context,
) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.Mongo.MongoDatabase.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(
			ctx,
			bson.M{
				"theme":        theme,
				"created_time": bson.M{"$gte": startTime, "$lte": endTime},
				"deleted":      false,
			},
		).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(
			ctx,
			bson.M{
				"theme":        theme,
				"created_time": bson.M{"$gte": startTime, "$lte": endTime},
				"deleted":      false,
			},
		).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		i.Dao.Zap.Logger.Error(
			"InstructionDataDaoImpl.GetInstructionDataListByThemeAndCreatedTime",
			zap.Field{Key: "theme", Type: zapcore.StringType, String: theme},
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		i.Dao.Zap.Logger.Info(
			"InstructionDataDaoImpl.GetInstructionDataListByThemeAndCreatedTime",
			zap.Field{Key: "theme", Type: zapcore.StringType, String: theme},
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByThemeAndStatusCodeAndCreatedTime(
	theme, statusCode string, startTime, endTime time.Time, offset, limit int64, desc bool, ctx context.Context,
) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.Mongo.MongoDatabase.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(
			ctx,
			bson.M{
				"theme":        theme,
				"status_code":  statusCode,
				"created_time": bson.M{"$gte": startTime, "$lte": endTime},
				"deleted":      false,
			},
		).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(
			ctx,
			bson.M{
				"theme":        theme,
				"status_code":  statusCode,
				"created_time": bson.M{"$gte": startTime, "$lte": endTime},
				"deleted":      false,
			},
		).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		i.Dao.Zap.Logger.Error(
			"InstructionDataDaoImpl.GetInstructionDataListByThemeAndStatusCodeAndCreatedTime",
			zap.Field{Key: "theme", Type: zapcore.StringType, String: theme},
			zap.Field{Key: "statusCode", Type: zapcore.StringType, String: statusCode},
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		i.Dao.Zap.Logger.Info(
			"InstructionDataDaoImpl.GetInstructionDataListByThemeAndStatusCodeAndCreatedTime",
			zap.Field{Key: "theme", Type: zapcore.StringType, String: theme},
			zap.Field{Key: "statusCode", Type: zapcore.StringType, String: statusCode},
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByStatusCode(
	statusCode string, offset, limit int64, desc bool, ctx context.Context,
) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.Mongo.MongoDatabase.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(
			ctx,
			bson.M{
				"status_code": statusCode,
				"deleted":     false,
			},
		).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(
			ctx,
			bson.M{
				"status_code": statusCode,
				"deleted":     false,
			},
		).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		i.Dao.Zap.Logger.Error(
			"InstructionDataDaoImpl.GetInstructionDataListByStatusCode",
			zap.Field{Key: "statusCode", Type: zapcore.StringType, String: statusCode},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		i.Dao.Zap.Logger.Info(
			"InstructionDataDaoImpl.GetInstructionDataListByStatusCode",
			zap.Field{Key: "statusCode", Type: zapcore.StringType, String: statusCode},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByStatusCodeAndCreatedTime(
	statusCode string, startTime, endTime time.Time, offset, limit int64, desc bool, ctx context.Context,
) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.Mongo.MongoDatabase.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(
			ctx,
			bson.M{
				"status_code":  statusCode,
				"created_time": bson.M{"$gte": startTime, "$lte": endTime},
				"deleted":      false,
			},
		).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(
			ctx,
			bson.M{
				"status_code":  statusCode,
				"created_time": bson.M{"$gte": startTime, "$lte": endTime},
				"deleted":      false,
			},
		).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		i.Dao.Zap.Logger.Error(
			"InstructionDataDaoImpl.GetInstructionDataListByStatusCodeAndCreatedTime",
			zap.Field{Key: "statusCode", Type: zapcore.StringType, String: statusCode},
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		i.Dao.Zap.Logger.Info(
			"InstructionDataDaoImpl.GetInstructionDataListByStatusCodeAndCreatedTime",
			zap.Field{Key: "statusCode", Type: zapcore.StringType, String: statusCode},
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByCreatedTime(
	startTime, endTime time.Time, offset, limit int64, desc bool, ctx context.Context,
) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.Mongo.MongoDatabase.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(
			ctx,
			bson.M{
				"created_time": bson.M{"$gte": startTime, "$lte": endTime},
				"deleted":      false,
			},
		).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(
			ctx,
			bson.M{
				"created_time": bson.M{"$gte": startTime, "$lte": endTime},
				"deleted":      false,
			},
		).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		i.Dao.Zap.Logger.Error(
			"InstructionDataDaoImpl.GetInstructionDataListByCreatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		i.Dao.Zap.Logger.Info(
			"InstructionDataDaoImpl.GetInstructionDataListByCreatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByUpdatedTime(
	startTime, endTime time.Time, offset, limit int64, desc bool, ctx context.Context,
) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.Mongo.MongoDatabase.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(
			ctx,
			bson.M{
				"updated_time": bson.M{"$gte": startTime, "$lte": endTime},
				"deleted":      false,
			},
		).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(
			ctx,
			bson.M{
				"updated_time": bson.M{"$gte": startTime, "$lte": endTime},
				"deleted":      false,
			},
		).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		i.Dao.Zap.Logger.Error(
			"InstructionDataDaoImpl.GetInstructionDataListByUpdatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		i.Dao.Zap.Logger.Info(
			"InstructionDataDaoImpl.GetInstructionDataListByUpdatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) InsertInstructionData(
	userID primitive.ObjectID,
	username, rowInstruction, rowInput, rowOutput, theme, source, note, statusCode, statusMessage string,
	ctx context.Context,
) (primitive.ObjectID, error) {
	instructionData := &models.InstructionDataModel{
		UserID:   userID,
		Username: username,
		Row: struct {
			Instruction string `json:"instruction" bson:"instruction"`
			Input       string `json:"input" bson:"input"`
			Output      string `json:"output" bson:"output"`
		}{Instruction: rowInstruction, Input: rowInput, Output: rowOutput},
		Theme:  theme,
		Source: source,
		Note:   note,
		Status: struct {
			Code    string `json:"code" bson:"code"`
			Message string `json:"message" bson:"message"`
		}{Code: statusCode, Message: statusMessage},
		Deleted:   false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: time.Time{},
	}
	collection := i.Dao.Mongo.MongoDatabase.Collection(instructionDataCollectionName)
	result, err := collection.InsertOne(ctx, instructionData)
	if err != nil {
		i.Dao.Zap.Logger.Error(
			"InstructionDataDaoImpl.InsertInstructionData",
			zap.Field{Key: "instructionData", Type: zapcore.ObjectMarshalerType, Interface: instructionData},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		i.Dao.Zap.Logger.Info(
			"InstructionDataDaoImpl.InsertInstructionData",
			zap.Field{Key: "instructionData", Type: zapcore.ObjectMarshalerType, Interface: instructionData},
		)
	}
	return result.InsertedID.(primitive.ObjectID), err

}

func (i InstructionDataDaoImpl) UpdateInstructionData(
	instructionDataID, userID primitive.ObjectID,
	username, rowInstruction, rowInput, rowOutput, theme, source, note, statusCode, statusMessage string,
	ctx context.Context,
) error {
	collection := i.Dao.Mongo.MongoDatabase.Collection(instructionDataCollectionName)
	err := collection.UpdateId(
		ctx,
		instructionDataID,
		bson.M{
			"$set": bson.M{
				"user_id":  userID,
				"username": username,
				"row": bson.M{
					"instruction": rowInstruction,
					"input":       rowInput,
					"output":      rowOutput,
				},
				"theme":  theme,
				"source": source,
				"note":   note,
				"status": bson.M{
					"code":    statusCode,
					"message": statusMessage,
				},
				"updated_at": time.Now(),
			},
		},
	)
	if err != nil {
		i.Dao.Zap.Logger.Error(
			"InstructionDataDaoImpl.UpdateInstructionData",
			zap.Field{Key: "instructionDataID", Type: zapcore.StringType, String: instructionDataID.Hex()},
			zap.Field{Key: "rowInstruction", Type: zapcore.StringType, String: rowInstruction},
			zap.Field{Key: "rowInput", Type: zapcore.StringType, String: rowInput},
			zap.Field{Key: "rowOutput", Type: zapcore.StringType, String: rowOutput},
			zap.Field{Key: "theme", Type: zapcore.StringType, String: theme},
			zap.Field{Key: "source", Type: zapcore.StringType, String: source},
			zap.Field{Key: "note", Type: zapcore.StringType, String: note},
			zap.Field{Key: "statusCode", Type: zapcore.StringType, String: statusCode},
			zap.Field{Key: "statusMessage", Type: zapcore.StringType, String: statusMessage},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		i.Dao.Zap.Logger.Info(
			"InstructionDataDaoImpl.UpdateInstructionData",
			zap.Field{Key: "instructionDataID", Type: zapcore.StringType, String: instructionDataID.Hex()},
			zap.Field{Key: "rowInstruction", Type: zapcore.StringType, String: rowInstruction},
			zap.Field{Key: "rowInput", Type: zapcore.StringType, String: rowInput},
			zap.Field{Key: "rowOutput", Type: zapcore.StringType, String: rowOutput},
			zap.Field{Key: "theme", Type: zapcore.StringType, String: theme},
			zap.Field{Key: "source", Type: zapcore.StringType, String: source},
			zap.Field{Key: "note", Type: zapcore.StringType, String: note},
			zap.Field{Key: "statusCode", Type: zapcore.StringType, String: statusCode},
			zap.Field{Key: "statusMessage", Type: zapcore.StringType, String: statusMessage},
		)
	}
	return err
}

func (i InstructionDataDaoImpl) SoftDeleteInstructionData(
	instructionDataID primitive.ObjectID, ctx context.Context,
) error {
	collection := i.Dao.Mongo.MongoDatabase.Collection(instructionDataCollectionName)
	err := collection.UpdateId(
		ctx,
		instructionDataID,
		bson.M{
			"$set": bson.M{
				"deleted":    true,
				"deleted_at": time.Now(),
			},
		},
	)
	if err != nil {
		i.Dao.Zap.Logger.Error(
			"InstructionDataDaoImpl.SoftDeleteInstructionData",
			zap.Field{Key: "instructionDataID", Type: zapcore.StringType, String: instructionDataID.Hex()},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		i.Dao.Zap.Logger.Info(
			"InstructionDataDaoImpl.SoftDeleteInstructionData",
			zap.Field{Key: "instructionDataID", Type: zapcore.StringType, String: instructionDataID.Hex()},
		)
	}
	return err
}

func (i InstructionDataDaoImpl) DeleteInstructionData(instructionDataID primitive.ObjectID, ctx context.Context) error {
	collection := i.Dao.Mongo.MongoDatabase.Collection(instructionDataCollectionName)
	err := collection.RemoveId(ctx, instructionDataID)
	if err != nil {
		i.Dao.Zap.Logger.Error(
			"InstructionDataDaoImpl.DeleteInstructionData",
			zap.Field{Key: "instructionDataID", Type: zapcore.StringType, String: instructionDataID.Hex()},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		i.Dao.Zap.Logger.Info(
			"InstructionDataDaoImpl.DeleteInstructionData",
			zap.Field{Key: "instructionDataID", Type: zapcore.StringType, String: instructionDataID.Hex()},
		)
	}
	return err
}
