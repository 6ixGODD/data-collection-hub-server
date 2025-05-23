package mods

import (
	"context"
	"fmt"
	"time"

	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/internal/pkg/dao"
	"data-collection-hub-server/internal/pkg/domain/entity"
	"data-collection-hub-server/pkg/utils/common"
	"github.com/goccy/go-json"
	"github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type InstructionDataDao interface {
	GetInstructionDataByID(
		ctx context.Context, instructionDataID primitive.ObjectID,
	) (*entity.InstructionDataModel, error)
	GetInstructionDataList(
		ctx context.Context, offset, limit int64, desc bool, userID *primitive.ObjectID, theme, statusCode *string,
		createStartTime, createEndTime, updateStartTime, updateEndTime *time.Time, query *string,
	) ([]entity.InstructionDataModel, *int64, error)
	CountInstructionData(
		ctx context.Context, userID *primitive.ObjectID, theme, statusCode *string,
		createStartTime, createEndTime, updateStartTime, updateEndTime *time.Time,
	) (*int64, error)
	AggregateCountInstructionData(
		ctx context.Context, groupBy *string, createStartTime, createEndTime *time.Time,
	) (map[string]int64, error)
	InsertInstructionData(
		ctx context.Context,
		userID primitive.ObjectID,
		rowInstruction, rowInput, rowOutput, theme, source, note, statusCode, statusMessage string,
	) (primitive.ObjectID, error)
	UpdateInstructionData(
		ctx context.Context, instructionDataID primitive.ObjectID, userID *primitive.ObjectID,
		rowInstruction, rowInput, rowOutput, theme, source, note, statusCode, statusMessage *string,
	) error
	SoftDeleteInstructionData(ctx context.Context, instructionDataID primitive.ObjectID) error
	SoftDeleteInstructionDataList(
		ctx context.Context, userID *primitive.ObjectID, theme, statusCode *string,
		createStartTime, createEndTime, updateStartTime, updateEndTime *time.Time,
	) (*int64, error)
	DeleteInstructionData(ctx context.Context, instructionDataID primitive.ObjectID) error
	DeleteInstructionDataList(
		ctx context.Context, userID *primitive.ObjectID, theme, statusCode *string,
		createStartTime, createEndTime, updateStartTime, updateEndTime *time.Time,
	) (*int64, error)
}

type InstructionDataDaoImpl struct {
	Dao     *dao.Core
	UserDao UserDao
}

func NewInstructionDataDao(ctx context.Context, core *dao.Core, userDao UserDao) (InstructionDataDao, error) {
	var _ InstructionDataDao = (*InstructionDataDaoImpl)(nil) // Ensure that the interface is implemented
	collection := core.Mongo.MongoClient.Database(core.Mongo.DatabaseName).Collection(config.InstructionDataCollectionName)
	err := collection.CreateIndexes(
		ctx, []options.IndexModel{
			{Key: []string{"user_id"}}, {Key: []string{"theme"}}, {Key: []string{"status.code"}},
			{Key: []string{"created_at"}}, {Key: []string{"updated_at"}},
		},
	)
	if err != nil {
		core.Logger.Error(
			fmt.Sprintf("Failed to create indexes for %s", config.InstructionDataCollectionName),
			zap.Error(err),
		)
		return nil, err
	}
	return &InstructionDataDaoImpl{core, userDao}, nil
}

func (i *InstructionDataDaoImpl) GetInstructionDataByID(
	ctx context.Context, instructionDataID primitive.ObjectID,
) (*entity.InstructionDataModel, error) {
	var instructionData entity.InstructionDataModel
	collection := i.Dao.Mongo.MongoClient.Database(i.Dao.Mongo.DatabaseName).Collection(config.InstructionDataCollectionName)
	err := collection.Find(ctx, bson.M{"_id": instructionDataID, "deleted": false}).One(&instructionData)
	if err != nil {
		i.Dao.Logger.Error(
			"InstructionDataDaoImpl.GetInstructionDataByID: failed to find instruction data",
			zap.Error(err), zap.String("instructionDataID", instructionDataID.Hex()),
		)
		return nil, err
	} else {
		i.Dao.Logger.Info(
			"InstructionDataDaoImpl.GetInstructionDataByID: success",
			zap.String("instructionDataID", instructionDataID.Hex()),
		)
		return &instructionData, nil
	}
}

func (i *InstructionDataDaoImpl) GetInstructionDataList(
	ctx context.Context, offset, limit int64, desc bool, userID *primitive.ObjectID,
	theme, statusCode *string,
	createStartTime, createEndTime, updateStartTime, updateEndTime *time.Time, query *string,
) ([]entity.InstructionDataModel, *int64, error) {
	var instructionDataList []entity.InstructionDataModel
	var err error

	collection := i.Dao.Mongo.MongoClient.Database(i.Dao.Mongo.DatabaseName).Collection(config.InstructionDataCollectionName)
	doc := bson.M{"deleted": false}
	if userID != nil {
		doc["user_id"] = *userID
	}
	if theme != nil {
		doc["theme"] = *theme
	}
	if statusCode != nil {
		doc["status.code"] = *statusCode
	}
	if createStartTime != nil && createEndTime != nil {
		doc["created_at"] = bson.M{"$gte": *createStartTime, "$lte": *createEndTime}
	}
	if updateStartTime != nil && updateEndTime != nil {
		doc["updated_at"] = bson.M{"$gte": *updateStartTime, "$lte": *updateEndTime}
	}
	if query != nil {
		safetyQuery := common.EscapeSpecialChars(*query)
		pattern := fmt.Sprintf(".*%s.*", safetyQuery)
		doc["$or"] = []bson.M{
			{"username": bson.M{"$regex": primitive.Regex{Pattern: pattern, Options: "i"}}},
		}
	}
	docJSON, _ := json.Marshal(doc)
	cursor := collection.Find(ctx, doc)
	count, err := cursor.Count()
	if err != nil {
		i.Dao.Logger.Error(
			"InstructionDataDaoImpl.GetInstructionDataList: failed to count instruction data list",
			zap.ByteString(config.InstructionDataCollectionName, docJSON), zap.Error(err),
		)
		return nil, nil, err
	}
	if desc {
		err = cursor.Sort("-created_at").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = cursor.Skip(offset).Limit(limit).All(&instructionDataList)
	}

	if err != nil {
		i.Dao.Logger.Error(
			"InstructionDataDaoImpl.GetInstructionDataList: failed to find instruction data list",
			zap.ByteString(config.InstructionDataCollectionName, docJSON), zap.Error(err),
		)
		return nil, nil, err
	}

	i.Dao.Logger.Info(
		"InstructionDataDaoImpl.GetInstructionDataList",
		zap.Int64("count", count), zap.ByteString(config.InstructionDataCollectionName, docJSON),
	)
	return instructionDataList, &count, nil
}

func (i *InstructionDataDaoImpl) CountInstructionData(
	ctx context.Context, userID *primitive.ObjectID, theme, statusCode *string,
	createStartTime, createEndTime, updateStartTime, updateEndTime *time.Time,
) (*int64, error) {
	collection := i.Dao.Mongo.MongoClient.Database(i.Dao.Mongo.DatabaseName).Collection(config.InstructionDataCollectionName)
	doc := bson.M{"deleted": false}
	if userID != nil {
		doc["user_id"] = *userID
	}
	if theme != nil {
		doc["theme"] = *theme
	}
	if statusCode != nil {
		doc["status.code"] = *statusCode
	}
	if createStartTime != nil && createEndTime != nil {
		doc["created_at"] = bson.M{"$gte": *createStartTime, "$lte": *createEndTime}
	}
	if updateStartTime != nil && updateEndTime != nil {
		doc["updated_at"] = bson.M{"$gte": *updateStartTime, "$lte": *updateEndTime}
	}
	docJSON, _ := json.Marshal(doc)

	count, err := collection.Find(ctx, doc).Count()
	if err != nil {
		i.Dao.Logger.Error(
			"InstructionDataDaoImpl.CountInstructionData: failed to count instruction data",
			zap.Error(err), zap.ByteString(config.InstructionDataCollectionName, docJSON),
		)
	} else {
		i.Dao.Logger.Info(
			"InstructionDataDaoImpl.CountInstructionData: success",
			zap.Int64("count", count), zap.ByteString(config.InstructionDataCollectionName, docJSON),
		)
	}
	return &count, err
}

func (i *InstructionDataDaoImpl) AggregateCountInstructionData(
	ctx context.Context, groupBy *string, createStartTime, createEndTime *time.Time,
) (map[string]int64, error) {
	collection := i.Dao.Mongo.MongoClient.Database(i.Dao.Mongo.DatabaseName).Collection(config.InstructionDataCollectionName)
	match := bson.M{"deleted": false}
	if createStartTime != nil && createEndTime != nil {
		match["created_at"] = bson.M{"$gte": *createStartTime, "$lte": *createEndTime}
	}
	pipeline := []bson.M{
		{"$match": match},
		{"$group": bson.M{"_id": "$" + *groupBy, "count": bson.M{"$sum": 1}}},
	}
	cursor := collection.Aggregate(ctx, pipeline)
	var result []bson.M
	if err := cursor.All(&result); err != nil {
		i.Dao.Logger.Error(
			"InstructionDataDaoImpl.AggregateCountGetInstructionData: failed to aggregate instruction data",
			zap.Error(err), zap.String("groupBy", *groupBy),
		)
		return nil, err
	}
	countMap := make(map[string]int64, len(result))
	for _, item := range result {
		countMap[item["_id"].(string)] = int64(item["count"].(int32))
	}
	i.Dao.Logger.Info(
		"InstructionDataDaoImpl.AggregateCountGetInstructionData: success",
		zap.String("groupBy", *groupBy),
		zap.Any("countMap", countMap),
	)
	return countMap, nil
}

func (i *InstructionDataDaoImpl) InsertInstructionData(
	ctx context.Context,
	userID primitive.ObjectID,
	rowInstruction, rowInput, rowOutput, theme, source, note, statusCode, statusMessage string,
) (primitive.ObjectID, error) {
	user, err := i.UserDao.GetUserByID(ctx, userID)
	if err != nil {
		i.Dao.Logger.Error(
			"InstructionDataDaoImpl.InsertInstructionData: failed to GetUserByID",
			zap.String("userID", userID.Hex()), zap.Error(err),
		)
		return primitive.NilObjectID, err
	}
	username := user.Username
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
	collection := i.Dao.Mongo.MongoClient.Database(i.Dao.Mongo.DatabaseName).Collection(config.InstructionDataCollectionName)
	result, err := collection.InsertOne(ctx, doc)
	if err != nil {
		i.Dao.Logger.Error(
			"InstructionDataDaoImpl.InsertInstructionData: failed to insert instruction data",
			zap.Error(err), zap.ByteString(config.InstructionDataCollectionName, docJSON),
		)
	} else {
		i.Dao.Logger.Info(
			"InstructionDataDaoImpl.InsertInstructionData: success",
			zap.String("instructionDataID", result.InsertedID.(primitive.ObjectID).Hex()),
			zap.ByteString(config.InstructionDataCollectionName, docJSON),
		)
	}
	return result.InsertedID.(primitive.ObjectID), err
}

func (i *InstructionDataDaoImpl) UpdateInstructionData(
	ctx context.Context,
	instructionDataID primitive.ObjectID, userID *primitive.ObjectID,
	rowInstruction, rowInput, rowOutput, theme, source, note, statusCode, statusMessage *string,
) error {
	collection := i.Dao.Mongo.MongoClient.Database(i.Dao.Mongo.DatabaseName).Collection(config.InstructionDataCollectionName)
	doc := bson.M{"updated_at": time.Now()}
	if userID != nil {
		doc["user_id"] = *userID
		user, err := i.UserDao.GetUserByID(ctx, *userID)
		if err != nil {
			i.Dao.Logger.Error(
				"InstructionDataDaoImpl.UpdateInstructionData: failed to GetUserByID",
				zap.String("userID", userID.Hex()), zap.String("instructionDataID", instructionDataID.Hex()),
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
		i.Dao.Logger.Error(
			"InstructionDataDaoImpl.UpdateInstructionData: failed to update instruction data",
			zap.String("instructionDataID", instructionDataID.Hex()),
			zap.ByteString(config.InstructionDataCollectionName, docJSON), zap.Error(err),
		)
	} else {
		i.Dao.Logger.Info(
			"InstructionDataDaoImpl.UpdateInstructionData: success",
			zap.String("instructionDataID", instructionDataID.Hex()),
			zap.ByteString(config.InstructionDataCollectionName, docJSON),
		)
	}
	return err
}

func (i *InstructionDataDaoImpl) SoftDeleteInstructionData(
	ctx context.Context, instructionDataID primitive.ObjectID,
) error {
	collection := i.Dao.Mongo.MongoClient.Database(i.Dao.Mongo.DatabaseName).Collection(config.InstructionDataCollectionName)
	err := collection.UpdateId(
		ctx,
		instructionDataID,
		bson.M{"$set": bson.M{"deleted": true, "deleted_at": time.Now()}},
	)
	if err != nil {
		i.Dao.Logger.Error(
			"InstructionDataDaoImpl.SoftDeleteInstructionData: failed to delete instruction data",
			zap.Error(err), zap.String("instructionDataID", instructionDataID.Hex()),
		)
	} else {
		i.Dao.Logger.Info(
			"InstructionDataDaoImpl.SoftDeleteInstructionData: success",
			zap.String("instructionDataID", instructionDataID.Hex()),
		)
	}
	return err
}

func (i *InstructionDataDaoImpl) SoftDeleteInstructionDataList(
	ctx context.Context, userID *primitive.ObjectID, theme, statusCode *string,
	createStartTime, createEndTime, updateStartTime, updateEndTime *time.Time,
) (*int64, error) {
	collection := i.Dao.Mongo.MongoClient.Database(i.Dao.Mongo.DatabaseName).Collection(config.InstructionDataCollectionName)
	doc := bson.M{"deleted": false}
	if userID != nil {
		doc["user_id"] = *userID
	}
	if theme != nil {
		doc["theme"] = *theme
	}
	if statusCode != nil {
		doc["status.code"] = *statusCode
	}
	if createStartTime != nil && createEndTime != nil {
		doc["created_at"] = bson.M{"$gte": *createStartTime, "$lte": *createEndTime}
	}
	if updateStartTime != nil && updateEndTime != nil {
		doc["updated_at"] = bson.M{"$gte": *updateStartTime, "$lte": *updateEndTime}
	}
	docJSON, _ := json.Marshal(doc)

	result, err := collection.UpdateAll(ctx, doc, bson.M{"$set": bson.M{"deleted": true, "deleted_at": time.Now()}})
	if err != nil {
		i.Dao.Logger.Error(
			"InstructionDataDaoImpl.SoftDeleteInstructionDataList: failed to delete instruction data list",
			zap.Error(err), zap.ByteString(config.InstructionDataCollectionName, docJSON),
		)
	} else {
		i.Dao.Logger.Info(
			"InstructionDataDaoImpl.SoftDeleteInstructionDataList: success",
			zap.Int64("count", result.ModifiedCount), zap.ByteString(config.InstructionDataCollectionName, docJSON),
		)
	}
	return &result.ModifiedCount, err
}

func (i *InstructionDataDaoImpl) DeleteInstructionData(
	ctx context.Context, instructionDataID primitive.ObjectID,
) error {
	collection := i.Dao.Mongo.MongoClient.Database(i.Dao.Mongo.DatabaseName).Collection(config.InstructionDataCollectionName)
	err := collection.RemoveId(ctx, instructionDataID)
	if err != nil {
		i.Dao.Logger.Error(
			"InstructionDataDaoImpl.DeleteInstructionData: failed to delete instruction data",
			zap.Error(err), zap.String("instructionDataID", instructionDataID.Hex()),
		)
	} else {
		i.Dao.Logger.Info(
			"InstructionDataDaoImpl.DeleteInstructionData: success",
			zap.String("instructionDataID", instructionDataID.Hex()),
		)
	}
	return err
}

func (i *InstructionDataDaoImpl) DeleteInstructionDataList(
	ctx context.Context, userID *primitive.ObjectID, theme, statusCode *string,
	createStartTime, createEndTime, updateStartTime, updateEndTime *time.Time,
) (*int64, error) {
	coll := i.Dao.Mongo.MongoClient.Database(i.Dao.Mongo.DatabaseName).Collection(config.InstructionDataCollectionName)
	doc := bson.M{}
	if userID != nil {
		doc["user_id"] = *userID
	}
	if theme != nil {
		doc["theme"] = *theme
	}
	if statusCode != nil {
		doc["status.code"] = *statusCode
	}
	if createStartTime != nil && createEndTime != nil {
		doc["created_at"] = bson.M{"$gte": *createStartTime, "$lte": *createEndTime}
	}
	if updateStartTime != nil && updateEndTime != nil {
		doc["updated_at"] = bson.M{"$gte": *updateStartTime, "$lte": *updateEndTime}
	}
	docJSON, _ := json.Marshal(doc)

	result, err := coll.RemoveAll(ctx, doc)
	if err != nil {
		i.Dao.Logger.Error(
			"InstructionDataDaoImpl.DeleteInstructionDataList: failed to delete instruction data list",
			zap.Error(err), zap.ByteString(config.InstructionDataCollectionName, docJSON),
		)
		return nil, err
	}
	i.Dao.Logger.Info(
		"InstructionDataDaoImpl.DeleteInstructionDataList: success",
		zap.Int64("count", result.DeletedCount), zap.ByteString(config.InstructionDataCollectionName, docJSON),
	)
	return &result.DeletedCount, nil
}
