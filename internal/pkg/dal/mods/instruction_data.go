package mods

import (
	"context"
	"time"

	"data-collection-hub-server/internal/pkg/dal"
	"data-collection-hub-server/internal/pkg/models"
	"github.com/goccy/go-json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type InstructionDataDao interface {
	GetInstructionDataById(
		ctx context.Context, instructionDataID *primitive.ObjectID,
	) (*models.InstructionDataModel, error)
	GetInstructionDataList(
		ctx context.Context, offset, limit int64, desc bool, userID *primitive.ObjectID, theme, statusCode *string,
		createTimeStart, createTimeEnd, updateTimeStart, updateTimeEnd *time.Time, query *string,
	) ([]models.InstructionDataModel, *int64, error)
	CountInstructionData(
		ctx context.Context, userID *primitive.ObjectID, theme, statusCode *string,
		createTimeStart, createTimeEnd, updateTimeStart, updateTimeEnd *time.Time,
	) (*int64, error)
	AggregateCountInstructionData(ctx context.Context, groupBy *string) (map[string]int64, error)
	InsertInstructionData(
		ctx context.Context,
		userID primitive.ObjectID,
		username, rowInstruction, rowInput, rowOutput, theme, source, note, statusCode, statusMessage string,
	) (primitive.ObjectID, error)
	UpdateInstructionData(
		ctx context.Context, instructionDataID primitive.ObjectID, userID *primitive.ObjectID,
		rowInstruction, rowInput, rowOutput, theme, source, note, statusCode, statusMessage *string,
	) error
	SoftDeleteInstructionData(ctx context.Context, instructionDataID primitive.ObjectID) error
	SoftDeleteInstructionDataList(
		ctx context.Context, userID *primitive.ObjectID, theme, statusCode *string,
		createTimeStart, createTimeEnd, updateTimeStart, updateTimeEnd *time.Time,
	) (*int64, error)
	DeleteInstructionData(ctx context.Context, instructionDataID primitive.ObjectID) error
	DeleteInstructionDataList(
		ctx context.Context, userID *primitive.ObjectID, theme, statusCode *string,
		createTimeStart, createTimeEnd, updateTimeStart, updateTimeEnd *time.Time,
	) (*int64, error)
}

type InstructionDataDaoImpl struct {
	Dao     *dal.Dao
	UserDao UserDao
}

func NewInstructionDataDao(dao *dal.Dao, userDao UserDao) InstructionDataDao {
	var _ InstructionDataDao = (*InstructionDataDaoImpl)(nil) // Ensure that the interface is implemented
	return &InstructionDataDaoImpl{dao, userDao}
}

func (i *InstructionDataDaoImpl) GetInstructionDataById(
	ctx context.Context, instructionDataID *primitive.ObjectID,
) (*models.InstructionDataModel, error) {
	var instructionData models.InstructionDataModel
	collection := i.Dao.Mongo.MongoDatabase.Collection(instructionDataCollectionName)
	err := collection.Find(ctx, bson.M{"_id": instructionDataID, "deleted": false}).One(&instructionData)
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

func (i *InstructionDataDaoImpl) GetInstructionDataList(
	ctx context.Context, offset, limit int64, desc bool, userID *primitive.ObjectID,
	theme, statusCode *string,
	createTimeStart, createTimeEnd, updateTimeStart, updateTimeEnd *time.Time, query *string,
) ([]models.InstructionDataModel, *int64, error) {
	var instructionDataList []models.InstructionDataModel
	var err error

	collection := i.Dao.Mongo.MongoDatabase.Collection(instructionDataCollectionName)
	doc := bson.M{"deleted": false}
	if userID != nil {
		doc["user_id"] = *userID
	}
	if theme != nil {
		doc["theme"] = *theme
	}
	if statusCode != nil {
		doc["status_code"] = *statusCode
	}
	if createTimeStart != nil && createTimeEnd != nil {
		doc["created_time"] = bson.M{"$gte": *createTimeStart, "$lte": *createTimeEnd}
	}
	if updateTimeStart != nil && updateTimeEnd != nil {
		doc["updated_time"] = bson.M{"$gte": *updateTimeStart, "$lte": *updateTimeEnd}
	}
	if query != nil {
		doc["$text"] = bson.M{"$search": *query}
	}
	docJSON, _ := json.Marshal(doc)

	if desc {
		err = collection.Find(ctx, doc).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(ctx, doc).Skip(offset).Limit(limit).All(&instructionDataList)
	}

	if err != nil {
		i.Dao.Zap.Logger.Error(
			"InstructionDataDaoImpl.GetInstructionDataList",
			zap.Int64("offset", offset), zap.Int64("limit", limit), zap.Bool("desc", desc),
			zap.ByteString(instructionDataCollectionName, docJSON), zap.Error(err),
		)
		return nil, nil, err
	}
	count, err := collection.Find(ctx, doc).Count()
	if err != nil {
		i.Dao.Zap.Logger.Error(
			"InstructionDataDaoImpl.GetInstructionDataList",
			zap.Int64("offset", offset), zap.Int64("limit", limit), zap.Bool("desc", desc),
			zap.ByteString(instructionDataCollectionName, docJSON), zap.Error(err),
		)
		return nil, nil, err
	}
	i.Dao.Zap.Logger.Info(
		"InstructionDataDaoImpl.GetInstructionDataList",
		zap.Int64("offset", offset), zap.Int64("limit", limit), zap.Bool("desc", desc),
		zap.ByteString(instructionDataCollectionName, docJSON), zap.Int64("count", count),
	)
	return instructionDataList, &count, nil
}

func (i *InstructionDataDaoImpl) CountInstructionData(
	ctx context.Context, userID *primitive.ObjectID, theme, statusCode *string,
	createTimeStart, createTimeEnd, updateTimeStart, updateTimeEnd *time.Time,
) (*int64, error) {
	collection := i.Dao.Mongo.MongoDatabase.Collection(instructionDataCollectionName)
	doc := bson.M{"deleted": false}
	if userID != nil {
		doc["user_id"] = *userID
	}
	if theme != nil {
		doc["theme"] = *theme
	}
	if statusCode != nil {
		doc["status_code"] = *statusCode
	}
	if createTimeStart != nil && createTimeEnd != nil {
		doc["created_time"] = bson.M{"$gte": *createTimeStart, "$lte": *createTimeEnd}
	}
	if updateTimeStart != nil && updateTimeEnd != nil {
		doc["updated_time"] = bson.M{"$gte": *updateTimeStart, "$lte": *updateTimeEnd}
	}
	docJSON, _ := json.Marshal(doc)

	count, err := collection.Find(ctx, doc).Count()
	if err != nil {
		i.Dao.Zap.Logger.Error(
			"InstructionDataDaoImpl.CountInstructionData",
			zap.ByteString(instructionDataCollectionName, docJSON),
			zap.Error(err),
		)
	} else {
		i.Dao.Zap.Logger.Info(
			"InstructionDataDaoImpl.CountInstructionData",
			zap.ByteString(instructionDataCollectionName, docJSON),
			zap.Int64("count", count),
		)
	}
	return &count, err
}

func (i *InstructionDataDaoImpl) AggregateCountInstructionData(
	ctx context.Context, groupBy *string,
) (map[string]int64, error) {
	collection := i.Dao.Mongo.MongoDatabase.Collection(instructionDataCollectionName)
	pipeline := []bson.M{
		{"$match": bson.M{"deleted": false}},
		{"$group": bson.M{"_id": "$" + *groupBy, "count": bson.M{"$sum": 1}}},
	}
	cursor := collection.Aggregate(ctx, pipeline)
	var result []bson.M
	if err := cursor.All(&result); err != nil {
		i.Dao.Zap.Logger.Error(
			"InstructionDataDaoImpl.AggregateCountGetInstructionData",
			zap.String("groupBy", *groupBy),
			zap.Error(err),
		)
		return nil, err
	}
	countMap := make(map[string]int64, len(result))
	for _, item := range result {
		countMap[item["_id"].(string)] = item["count"].(int64)
	}
	i.Dao.Zap.Logger.Info(
		"InstructionDataDaoImpl.AggregateCountGetInstructionData",
		zap.String("groupBy", *groupBy),
		zap.Any("countMap", countMap),
	)
	return countMap, nil
}

func (i *InstructionDataDaoImpl) InsertInstructionData(
	ctx context.Context,
	userID primitive.ObjectID,
	username, rowInstruction, rowInput, rowOutput, theme, source, note, statusCode, statusMessage string,
) (primitive.ObjectID, error) {
	doc := bson.M{
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
		"created_at": time.Now(),
		"updated_at": time.Now(),
		"deleted":    false,
		"deleted_at": nil,
	}
	docJSON, _ := bson.Marshal(doc)
	collection := i.Dao.Mongo.MongoDatabase.Collection(instructionDataCollectionName)
	result, err := collection.InsertOne(ctx, doc)
	if err != nil {
		i.Dao.Zap.Logger.Error(
			"InstructionDataDaoImpl.InsertInstructionData",
			zap.ByteString(instructionDataCollectionName, docJSON),
			zap.Error(err),
		)
	} else {
		i.Dao.Zap.Logger.Info(
			"InstructionDataDaoImpl.InsertInstructionData",
			zap.ByteString(instructionDataCollectionName, docJSON),
			zap.String("instructionDataID", result.InsertedID.(primitive.ObjectID).Hex()),
		)
	}
	return result.InsertedID.(primitive.ObjectID), err
}

func (i *InstructionDataDaoImpl) UpdateInstructionData(
	ctx context.Context,
	instructionDataID primitive.ObjectID, userID *primitive.ObjectID,
	rowInstruction, rowInput, rowOutput, theme, source, note, statusCode, statusMessage *string,
) error {
	collection := i.Dao.Mongo.MongoDatabase.Collection(instructionDataCollectionName)
	doc := bson.M{"updated_at": time.Now()}
	if userID != nil {
		doc["user_id"] = *userID
		user, err := i.UserDao.GetUserById(ctx, *userID)
		if err != nil {
			i.Dao.Zap.Logger.Error(
				"InstructionDataDaoImpl.UpdateInstructionData",
				zap.String("instructionDataID", instructionDataID.Hex()),
				zap.String("userID", userID.Hex()),
				zap.Error(err),
			)
			return err
		}
		doc["username"] = user.Username
	}
	if rowInstruction != nil {
		doc["row.instruction"] = *rowInstruction
	}
	if rowInput != nil {
		doc["row.input"] = *rowInput
	}
	if rowOutput != nil {
		doc["row.output"] = *rowOutput
	}
	if theme != nil {
		doc["theme"] = *theme
	}
	if source != nil {
		doc["source"] = *source
	}
	if note != nil {
		doc["note"] = *note
	}
	if statusCode != nil {
		doc["status.code"] = *statusCode
	}
	if statusMessage != nil {
		doc["status.message"] = *statusMessage
	}
	docJSON, _ := json.Marshal(doc)

	err := collection.UpdateId(ctx, instructionDataID, bson.M{"$set": doc})
	if err != nil {
		i.Dao.Zap.Logger.Error(
			"InstructionDataDaoImpl.UpdateInstructionData",
			zap.String("instructionDataID", instructionDataID.Hex()),
			zap.ByteString(instructionDataCollectionName, docJSON),
			zap.Error(err),
		)
	} else {
		i.Dao.Zap.Logger.Info(
			"InstructionDataDaoImpl.UpdateInstructionData",
			zap.String("instructionDataID", instructionDataID.Hex()),
			zap.ByteString(instructionDataCollectionName, docJSON),
		)
	}
	return err
}

func (i *InstructionDataDaoImpl) SoftDeleteInstructionData(
	ctx context.Context, instructionDataID primitive.ObjectID,
) error {
	collection := i.Dao.Mongo.MongoDatabase.Collection(instructionDataCollectionName)
	err := collection.UpdateId(
		ctx,
		instructionDataID,
		bson.M{"$set": bson.M{"deleted": true, "deleted_at": time.Now()}},
	)
	if err != nil {
		i.Dao.Zap.Logger.Error(
			"InstructionDataDaoImpl.SoftDeleteInstructionData",
			zap.String("instructionDataID", instructionDataID.Hex()),
			zap.Error(err),
		)
	} else {
		i.Dao.Zap.Logger.Info(
			"InstructionDataDaoImpl.SoftDeleteInstructionData",
			zap.String("instructionDataID", instructionDataID.Hex()),
		)
	}
	return err
}

func (i *InstructionDataDaoImpl) SoftDeleteInstructionDataList(
	ctx context.Context, userID *primitive.ObjectID, theme, statusCode *string,
	createTimeStart, createTimeEnd, updateTimeStart, updateTimeEnd *time.Time,
) (*int64, error) {
	collection := i.Dao.Mongo.MongoDatabase.Collection(instructionDataCollectionName)
	doc := bson.M{"deleted": false}
	if userID != nil {
		doc["user_id"] = *userID
	}
	if theme != nil {
		doc["theme"] = *theme
	}
	if statusCode != nil {
		doc["status_code"] = *statusCode
	}
	if createTimeStart != nil && createTimeEnd != nil {
		doc["created_time"] = bson.M{"$gte": *createTimeStart, "$lte": *createTimeEnd}
	}
	if updateTimeStart != nil && updateTimeEnd != nil {
		doc["updated_time"] = bson.M{"$gte": *updateTimeStart, "$lte": *updateTimeEnd}
	}
	docJSON, _ := json.Marshal(doc)

	result, err := collection.UpdateAll(ctx, doc, bson.M{"$set": bson.M{"deleted": true, "deleted_at": time.Now()}})
	if err != nil {
		i.Dao.Zap.Logger.Error(
			"InstructionDataDaoImpl.SoftDeleteInstructionDataList",
			zap.ByteString(instructionDataCollectionName, docJSON),
			zap.Error(err),
		)
	} else {
		i.Dao.Zap.Logger.Info(
			"InstructionDataDaoImpl.SoftDeleteInstructionDataList",
			zap.ByteString(instructionDataCollectionName, docJSON),
			zap.Int64("deletedCount", result.ModifiedCount),
		)
	}
	return &result.ModifiedCount, err
}

func (i *InstructionDataDaoImpl) DeleteInstructionData(
	ctx context.Context, instructionDataID primitive.ObjectID,
) error {
	collection := i.Dao.Mongo.MongoDatabase.Collection(instructionDataCollectionName)
	err := collection.RemoveId(ctx, instructionDataID)
	if err != nil {
		i.Dao.Zap.Logger.Error(
			"InstructionDataDaoImpl.DeleteInstructionData",
			zap.String("instructionDataID", instructionDataID.Hex()),
			zap.Error(err),
		)
	} else {
		i.Dao.Zap.Logger.Info(
			"InstructionDataDaoImpl.DeleteInstructionData",
			zap.String("instructionDataID", instructionDataID.Hex()),
		)
	}
	return err
}

func (i *InstructionDataDaoImpl) DeleteInstructionDataList(
	ctx context.Context, userID *primitive.ObjectID, theme, statusCode *string,
	createTimeStart, createTimeEnd, updateTimeStart, updateTimeEnd *time.Time,
) (*int64, error) {
	collection := i.Dao.Mongo.MongoDatabase.Collection(instructionDataCollectionName)
	doc := bson.M{"deleted": false}
	if userID != nil {
		doc["user_id"] = *userID
	}
	if theme != nil {
		doc["theme"] = *theme
	}
	if statusCode != nil {
		doc["status_code"] = *statusCode
	}
	if createTimeStart != nil && createTimeEnd != nil {
		doc["created_time"] = bson.M{"$gte": *createTimeStart, "$lte": *createTimeEnd}
	}
	if updateTimeStart != nil && updateTimeEnd != nil {
		doc["updated_time"] = bson.M{"$gte": *updateTimeStart, "$lte": *updateTimeEnd}
	}
	docJSON, _ := json.Marshal(doc)

	result, err := collection.RemoveAll(ctx, doc)
	if err != nil {
		i.Dao.Zap.Logger.Error(
			"InstructionDataDaoImpl.DeleteInstructionDataList",
			zap.ByteString(instructionDataCollectionName, docJSON),
			zap.Error(err),
		)
	} else {
		i.Dao.Zap.Logger.Info(
			"InstructionDataDaoImpl.DeleteInstructionDataList",
			zap.ByteString(instructionDataCollectionName, docJSON),
			zap.Int64("deletedCount", result.DeletedCount),
		)
	}
	return &result.DeletedCount, err
}
