package dao

import (
	"context"

	"data-collection-hub-server/internal/pkg/models"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
)

type InstructionDataDao interface {
	GetInstructionDataById(instructionDataId string, mongoClient *qmgo.QmgoClient, ctx context.Context) (*models.InstructionDataModel, error)
	GetInstructionDataByRowId(rowId string, mongoClient *qmgo.QmgoClient, ctx context.Context) (*models.InstructionDataModel, error)
	GetInstructionDataList(mongoClient *qmgo.QmgoClient, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByUserUUID(mongoClient *qmgo.QmgoClient, userUUID string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByFuzzyQuery(mongoClient *qmgo.QmgoClient, query string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByFuzzyQueryAndTheme(mongoClient *qmgo.QmgoClient, query, theme string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByFuzzyQueryAndThemeAndStatusCode(mongoClient *qmgo.QmgoClient, query, theme, statusCode string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByFuzzyQueryAndThemeAndCreatedTime(mongoClient *qmgo.QmgoClient, query, theme, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByFuzzyQueryAndThemeAndStatusCodeAndCreatedTime(mongoClient *qmgo.QmgoClient, query, theme, statusCode, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByFuzzyQueryAndStatusCode(mongoClient *qmgo.QmgoClient, query, statusCode string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByFuzzyQueryAndStatusCodeAndCreatedTime(mongoClient *qmgo.QmgoClient, query, statusCode, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByFuzzyQueryAndCreatedTime(mongoClient *qmgo.QmgoClient, query, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByTheme(mongoClient *qmgo.QmgoClient, theme string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByThemeAndStatusCode(mongoClient *qmgo.QmgoClient, theme, statusCode string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByThemeAndCreatedTime(mongoClient *qmgo.QmgoClient, theme, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByThemeAndStatusCodeAndCreatedTime(mongoClient *qmgo.QmgoClient, theme, statusCode, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByStatusCode(mongoClient *qmgo.QmgoClient, statusCode string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByStatusCodeAndCreatedTime(mongoClient *qmgo.QmgoClient, statusCode, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByCreatedTime(mongoClient *qmgo.QmgoClient, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByUpdatedTime(mongoClient *qmgo.QmgoClient, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	InsertInstructionData(instructionData *models.InstructionDataModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error
	UpdateInstructionData(instructionData *models.InstructionDataModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error
	DeleteInstructionData(instructionData *models.InstructionDataModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error
}

type InstructionData struct{}

func (i InstructionData) GetInstructionDataById(instructionDataId string, mongoClient *qmgo.QmgoClient, ctx context.Context) (*models.InstructionDataModel, error) {
	var instructionData models.InstructionDataModel
	err := mongoClient.Find(ctx, bson.M{"_id": instructionDataId}).One(&instructionData)
	if err != nil {
		return nil, err
	} else {
		return &instructionData, nil
	}
}

func (i InstructionData) GetInstructionDataByRowId(rowId string, mongoClient *qmgo.QmgoClient, ctx context.Context) (*models.InstructionDataModel, error) {
	var instructionData models.InstructionDataModel
	err := mongoClient.Find(ctx, bson.M{"row_id": rowId}).One(&instructionData)
	if err != nil {
		return nil, err
	} else {
		return &instructionData, nil
	}
}

func (i InstructionData) GetInstructionDataList(mongoClient *qmgo.QmgoClient, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	if desc {
		err = mongoClient.Find(ctx, bson.M{}).Sort("-created_at").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = mongoClient.Find(ctx, bson.M{}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		return nil, err
	} else {
		return instructionDataList, nil
	}
}

func (i InstructionData) GetInstructionDataListByUserUUID(mongoClient *qmgo.QmgoClient, userUUID string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	if desc {
		err = mongoClient.Find(ctx, bson.M{"user_uuid": userUUID}).Sort("-created_at").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = mongoClient.Find(ctx, bson.M{"user_uuid": userUUID}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		return nil, err
	} else {
		return instructionDataList, nil
	}
}

func (i InstructionData) GetInstructionDataListByFuzzyQuery(mongoClient *qmgo.QmgoClient, query string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	if desc {
		err = mongoClient.Find(ctx, bson.M{"$text": bson.M{"$search": query}}).Sort("-created_at").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = mongoClient.Find(ctx, bson.M{"$text": bson.M{"$search": query}}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		return nil, err
	} else {
		return instructionDataList, nil
	}
}

func (i InstructionData) GetInstructionDataListByFuzzyQueryAndTheme(mongoClient *qmgo.QmgoClient, query, theme string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	if desc {
		err = mongoClient.Find(ctx, bson.M{"$text": bson.M{"$search": query}, "theme": theme}).Sort("-created_at").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = mongoClient.Find(ctx, bson.M{"$text": bson.M{"$search": query}, "theme": theme}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		return nil, err
	} else {
		return instructionDataList, nil
	}
}

func (i InstructionData) GetInstructionDataListByFuzzyQueryAndThemeAndStatusCode(mongoClient *qmgo.QmgoClient, query, theme, statusCode string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	if desc {
		err = mongoClient.Find(ctx, bson.M{"$text": bson.M{"$search": query}, "theme": theme, "status_code": statusCode}).Sort("-created_at").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = mongoClient.Find(ctx, bson.M{"$text": bson.M{"$search": query}, "theme": theme, "status_code": statusCode}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		return nil, err
	} else {
		return instructionDataList, nil
	}
}

func (i InstructionData) GetInstructionDataListByFuzzyQueryAndThemeAndCreatedTime(mongoClient *qmgo.QmgoClient, query, theme, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	if desc {
		err = mongoClient.Find(
			ctx, bson.M{"$text": bson.M{"$search": query}, "theme": theme, "created_at": bson.M{"$gte": startTime, "$lte": endTime}},
		).Sort("-created_at").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = mongoClient.Find(
			ctx, bson.M{"$text": bson.M{"$search": query}, "theme": theme, "created_at": bson.M{"$gte": startTime, "$lte": endTime}},
		).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		return nil, err
	} else {
		return instructionDataList, nil
	}
}

func (i InstructionData) GetInstructionDataListByFuzzyQueryAndThemeAndStatusCodeAndCreatedTime(mongoClient *qmgo.QmgoClient, query, theme, statusCode, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	if desc {
		err = mongoClient.Find(
			ctx, bson.M{"$text": bson.M{"$search": query}, "theme": theme, "status_code": statusCode, "created_at": bson.M{"$gte": startTime, "$lte": endTime}},
		).Sort("-created_at").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = mongoClient.Find(
			ctx, bson.M{"$text": bson.M{"$search": query}, "theme": theme, "status_code": statusCode, "created_at": bson.M{"$gte": startTime, "$lte": endTime}},
		).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		return nil, err
	} else {
		return instructionDataList, nil
	}
}

func (i InstructionData) GetInstructionDataListByFuzzyQueryAndStatusCode(mongoClient *qmgo.QmgoClient, query, statusCode string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	if desc {
		err = mongoClient.Find(ctx, bson.M{"$text": bson.M{"$search": query}, "status_code": statusCode}).Sort("-created_at").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = mongoClient.Find(ctx, bson.M{"$text": bson.M{"$search": query}, "status_code": statusCode}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		return nil, err
	} else {
		return instructionDataList, nil
	}
}

func (i InstructionData) GetInstructionDataListByFuzzyQueryAndStatusCodeAndCreatedTime(mongoClient *qmgo.QmgoClient, query, statusCode, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	if desc {
		err = mongoClient.Find(
			ctx, bson.M{"$text": bson.M{"$search": query}, "status_code": statusCode, "created_at": bson.M{"$gte": startTime, "$lte": endTime}},
		).Sort("-created_at").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = mongoClient.Find(
			ctx, bson.M{"$text": bson.M{"$search": query}, "status_code": statusCode, "created_at": bson.M{"$gte": startTime, "$lte": endTime}},
		).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		return nil, err
	} else {
		return instructionDataList, nil
	}
}

func (i InstructionData) GetInstructionDataListByFuzzyQueryAndCreatedTime(mongoClient *qmgo.QmgoClient, query, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	if desc {
		err = mongoClient.Find(
			ctx, bson.M{"$text": bson.M{"$search": query}, "created_at": bson.M{"$gte": startTime, "$lte": endTime}},
		).Sort("-created_at").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = mongoClient.Find(
			ctx, bson.M{"$text": bson.M{"$search": query}, "created_at": bson.M{"$gte": startTime, "$lte": endTime}},
		).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		return nil, err
	} else {
		return instructionDataList, nil
	}
}

func (i InstructionData) GetInstructionDataListByTheme(mongoClient *qmgo.QmgoClient, theme string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	if desc {
		err = mongoClient.Find(ctx, bson.M{"theme": theme}).Sort("-created_at").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = mongoClient.Find(ctx, bson.M{"theme": theme}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		return nil, err
	} else {
		return instructionDataList, nil
	}
}

func (i InstructionData) GetInstructionDataListByThemeAndStatusCode(mongoClient *qmgo.QmgoClient, theme, statusCode string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	if desc {
		err = mongoClient.Find(ctx, bson.M{"theme": theme, "status_code": statusCode}).Sort("-created_at").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = mongoClient.Find(ctx, bson.M{"theme": theme, "status_code": statusCode}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		return nil, err
	} else {
		return instructionDataList, nil
	}
}

func (i InstructionData) GetInstructionDataListByThemeAndCreatedTime(mongoClient *qmgo.QmgoClient, theme, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	if desc {
		err = mongoClient.Find(
			ctx, bson.M{"theme": theme, "created_at": bson.M{"$gte": startTime, "$lte": endTime}},
		).Sort("-created_at").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = mongoClient.Find(
			ctx, bson.M{"theme": theme, "created_at": bson.M{"$gte": startTime, "$lte": endTime}},
		).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		return nil, err
	} else {
		return instructionDataList, nil
	}
}

func (i InstructionData) GetInstructionDataListByThemeAndStatusCodeAndCreatedTime(mongoClient *qmgo.QmgoClient, theme, statusCode, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	if desc {
		err = mongoClient.Find(
			ctx, bson.M{"theme": theme, "status_code": statusCode, "created_at": bson.M{"$gte": startTime, "$lte": endTime}},
		).Sort("-created_at").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = mongoClient.Find(
			ctx, bson.M{"theme": theme, "status_code": statusCode, "created_at": bson.M{"$gte": startTime, "$lte": endTime}},
		).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		return nil, err
	} else {
		return instructionDataList, nil
	}
}

func (i InstructionData) GetInstructionDataListByStatusCode(mongoClient *qmgo.QmgoClient, statusCode string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	if desc {
		err = mongoClient.Find(ctx, bson.M{"status_code": statusCode}).Sort("-created_at").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = mongoClient.Find(ctx, bson.M{"status_code": statusCode}).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		return nil, err
	} else {
		return instructionDataList, nil
	}
}

func (i InstructionData) GetInstructionDataListByStatusCodeAndCreatedTime(mongoClient *qmgo.QmgoClient, statusCode, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	if desc {
		err = mongoClient.Find(
			ctx, bson.M{"status_code": statusCode, "created_at": bson.M{"$gte": startTime, "$lte": endTime}},
		).Sort("-created_at").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = mongoClient.Find(
			ctx, bson.M{"status_code": statusCode, "created_at": bson.M{"$gte": startTime, "$lte": endTime}},
		).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		return nil, err
	} else {
		return instructionDataList, nil
	}
}

func (i InstructionData) GetInstructionDataListByCreatedTime(mongoClient *qmgo.QmgoClient, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	if desc {
		err = mongoClient.Find(
			ctx, bson.M{"created_at": bson.M{"$gte": startTime, "$lte": endTime}},
		).Sort("-created_at").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = mongoClient.Find(
			ctx, bson.M{"created_at": bson.M{"$gte": startTime, "$lte": endTime}},
		).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		return nil, err
	} else {
		return instructionDataList, nil
	}
}

func (i InstructionData) GetInstructionDataListByUpdatedTime(mongoClient *qmgo.QmgoClient, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	var instructionDataList []models.InstructionDataModel
	var err error
	if desc {
		err = mongoClient.Find(
			ctx, bson.M{"updated_at": bson.M{"$gte": startTime, "$lte": endTime}},
		).Sort("-created_at").Skip(offset).Limit(limit).All(&instructionDataList)
	} else {
		err = mongoClient.Find(
			ctx, bson.M{"updated_at": bson.M{"$gte": startTime, "$lte": endTime}},
		).Skip(offset).Limit(limit).All(&instructionDataList)
	}
	if err != nil {
		return nil, err
	} else {
		return instructionDataList, nil
	}
}

func (i InstructionData) InsertInstructionData(instructionData *models.InstructionDataModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error {
	_, err := mongoClient.InsertOne(ctx, instructionData)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (i InstructionData) UpdateInstructionData(instructionData *models.InstructionDataModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error {
	err := mongoClient.UpdateOne(ctx, bson.M{"_id": instructionData.InstructionDataID}, bson.M{"$set": instructionData})
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (i InstructionData) DeleteInstructionData(instructionData *models.InstructionDataModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error {
	err := mongoClient.RemoveId(ctx, instructionData.InstructionDataID)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func NewInstructionDataDao() InstructionDataDao {
	var _ InstructionDataDao = new(InstructionData)
	return &InstructionData{}
}
