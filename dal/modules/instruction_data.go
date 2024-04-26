package modules

import (
	"context"

	"data-collection-hub-server/dal"
	"data-collection-hub-server/models"
	"go.mongodb.org/mongo-driver/bson"
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
	collection := i.Dao.MongoDB.Collection(instructionDataCollectionName)
	err := collection.Find(ctx, bson.M{"instruction_data_id": instructionDataId}).One(&instructionData)
	if err != nil {
		return nil, err
	} else {
		return &instructionData, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataList(offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.MongoDB.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{}).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(ctx, bson.M{}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		return nil, err
	} else {
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByUserUUID(userUUID string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.MongoDB.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{"user_uuid": userUUID}).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(ctx, bson.M{"user_uuid": userUUID}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		return nil, err
	} else {
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByFuzzyQuery(query string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.MongoDB.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{"$text": bson.M{"$search": query}}).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(ctx, bson.M{"$text": bson.M{"$search": query}}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		return nil, err
	} else {
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByFuzzyQueryAndTheme(query, theme string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.MongoDB.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{"$text": bson.M{"$search": query}, "theme": theme}).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(ctx, bson.M{"$text": bson.M{"$search": query}, "theme": theme}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		return nil, err
	} else {
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByFuzzyQueryAndThemeAndStatusCode(query, theme, statusCode string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.MongoDB.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{"$text": bson.M{"$search": query}, "theme": theme, "status_code": statusCode}).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(ctx, bson.M{"$text": bson.M{"$search": query}, "theme": theme, "status_code": statusCode}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		return nil, err
	} else {
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByFuzzyQueryAndThemeAndCreatedTime(query, theme, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.MongoDB.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{"$text": bson.M{"$search": query}, "theme": theme, "created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(ctx, bson.M{"$text": bson.M{"$search": query}, "theme": theme, "created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		return nil, err
	} else {
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByFuzzyQueryAndThemeAndStatusCodeAndCreatedTime(query, theme, statusCode, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.MongoDB.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{"$text": bson.M{"$search": query}, "theme": theme, "status_code": statusCode, "created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(ctx, bson.M{"$text": bson.M{"$search": query}, "theme": theme, "status_code": statusCode, "created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		return nil, err
	} else {
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByFuzzyQueryAndStatusCode(query, statusCode string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.MongoDB.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{"$text": bson.M{"$search": query}, "status_code": statusCode}).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(ctx, bson.M{"$text": bson.M{"$search": query}, "status_code": statusCode}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		return nil, err
	} else {
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByFuzzyQueryAndStatusCodeAndCreatedTime(query, statusCode, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.MongoDB.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{"$text": bson.M{"$search": query}, "status_code": statusCode, "created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(ctx, bson.M{"$text": bson.M{"$search": query}, "status_code": statusCode, "created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		return nil, err
	} else {
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByFuzzyQueryAndCreatedTime(query, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.MongoDB.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{"$text": bson.M{"$search": query}, "created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(ctx, bson.M{"$text": bson.M{"$search": query}, "created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		return nil, err
	} else {
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByTheme(theme string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.MongoDB.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{"theme": theme}).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(ctx, bson.M{"theme": theme}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		return nil, err
	} else {
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByThemeAndStatusCode(theme, statusCode string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.MongoDB.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{"theme": theme, "status_code": statusCode}).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(ctx, bson.M{"theme": theme, "status_code": statusCode}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		return nil, err
	} else {
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByThemeAndCreatedTime(theme, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.MongoDB.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{"theme": theme, "created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(ctx, bson.M{"theme": theme, "created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		return nil, err
	} else {
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByThemeAndStatusCodeAndCreatedTime(theme, statusCode, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.MongoDB.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{"theme": theme, "status_code": statusCode, "created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(ctx, bson.M{"theme": theme, "status_code": statusCode, "created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		return nil, err
	} else {
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByStatusCode(statusCode string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.MongoDB.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{"status_code": statusCode}).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(ctx, bson.M{"status_code": statusCode}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		return nil, err
	} else {
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByStatusCodeAndCreatedTime(statusCode, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.MongoDB.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{"status_code": statusCode, "created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(ctx, bson.M{"status_code": statusCode, "created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		return nil, err
	} else {
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByCreatedTime(startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.MongoDB.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{"created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(ctx, bson.M{"created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		return nil, err
	} else {
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) GetInstructionDataListByUpdatedTime(startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	collection := i.Dao.MongoDB.Collection(instructionDataCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{"updated_time": bson.M{"$gte": startTime, "$lte": endTime}}).Sort("-created_time").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = collection.Find(ctx, bson.M{"updated_time": bson.M{"$gte": startTime, "$lte": endTime}}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		return nil, err
	} else {
		return instructionDataList, nil
	}
}

func (i InstructionDataDaoImpl) InsertInstructionData(instructionData *models.InstructionDataModel, ctx context.Context) error {
	collection := i.Dao.MongoDB.Collection(instructionDataCollectionName)
	_, err := collection.InsertOne(ctx, instructionData)
	return err
}

func (i InstructionDataDaoImpl) UpdateInstructionData(instructionData *models.InstructionDataModel, ctx context.Context) error {
	collection := i.Dao.MongoDB.Collection(instructionDataCollectionName)
	err := collection.UpdateOne(ctx, bson.M{"instruction_data_id": instructionData.InstructionDataID}, bson.M{"$set": instructionData})
	return err
}

func (i InstructionDataDaoImpl) DeleteInstructionData(instructionData *models.InstructionDataModel, ctx context.Context) error {
	collection := i.Dao.MongoDB.Collection(instructionDataCollectionName)
	err := collection.RemoveId(ctx, instructionData.InstructionDataID)
	return err
}
