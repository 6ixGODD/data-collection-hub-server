package mods

import (
	"context"
	"time"

	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/internal/pkg/dal"
	"data-collection-hub-server/internal/pkg/models"
	"github.com/goccy/go-json"
	"github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	opt "go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type UserDao interface {
	GetUserById(ctx context.Context, userID primitive.ObjectID) (*models.UserModel, error)
	GetUserByEmail(ctx context.Context, email *string) (*models.UserModel, error)
	GetUserList(
		ctx context.Context,
		offset, limit int64, desc bool, organization, role *string,
		createStartTime, createEndTime, updateStartTime, updateEndTime, lastLoginStartTime, lastLoginEndTime *time.Time,
		query *string,
	) ([]models.UserModel, *int64, error)
	CountUser(
		ctx context.Context, organization, role *string,
		createStartTime, createEndTime, updateStartTime, updateEndTime, lastLoginStartTime, lastLoginEndTime *time.Time,
	) (*int64, error)
	InsertUser(ctx context.Context, username, email, password, role, organization string) (primitive.ObjectID, error)
	UpdateUser(
		ctx context.Context, userID primitive.ObjectID, username, email, password, role, organization *string,
	) error
	SoftDeleteUser(ctx context.Context, userID primitive.ObjectID) error
	SoftDeleteUserList(
		ctx context.Context, organization, role *string,
		createStartTime, createEndTime, updateStartTime, updateEndTime, lastLoginStartTime, lastLoginEndTime *time.Time,
	) (*int64, error)
	DeleteUser(ctx context.Context, userID primitive.ObjectID) error
	DeleteUserList(
		ctx context.Context, organization, role *string,
		createStartTime, createEndTime, updateStartTime, updateEndTime, lastLoginStartTime, lastLoginEndTime *time.Time,
	) (*int64, error)
}

type UserDaoImpl struct{ *dal.Dao }

func NewUserDao(dao *dal.Dao) (UserDao, error) {
	var _ UserDao = (*UserDaoImpl)(nil)
	collection := dao.Mongo.MongoDatabase.Collection(config.UserCollectionName)
	err := collection.CreateIndexes(
		context.Background(), []options.IndexModel{
			{
				Key:          []string{"username", "email"},
				IndexOptions: opt.Index().SetUnique(true),
			},
			{
				Key: []string{"created_at", "updated_at"},
			},
		},
	)
	if err != nil {
		return nil, err
	}
	return &UserDaoImpl{dao}, nil
}

func (u *UserDaoImpl) GetUserById(ctx context.Context, userID primitive.ObjectID) (*models.UserModel, error) {
	var user models.UserModel
	collection := u.Dao.Mongo.MongoDatabase.Collection(config.UserCollectionName)
	err := collection.Find(
		ctx,
		bson.M{
			"_id":     userID,
			"deleted": false,
		},
	).One(&user)
	if err != nil {
		u.Dao.Zap.Logger.Error(
			"UserDaoImpl.GetUserById",
			zap.Field{Key: "userID", Type: zapcore.StringType, String: userID.Hex()},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		u.Dao.Zap.Logger.Info(
			"UserDaoImpl.GetUserById",
			zap.Field{Key: "userID", Type: zapcore.StringType, String: userID.Hex()},
		)
		return &user, nil
	}
}

func (u *UserDaoImpl) GetUserByEmail(ctx context.Context, email *string) (*models.UserModel, error) {
	var user models.UserModel
	collection := u.Dao.Mongo.MongoDatabase.Collection(config.UserCollectionName)
	err := collection.Find(
		ctx,
		bson.M{
			"email":   *email,
			"deleted": false,
		},
	).One(&user)
	if err != nil {
		u.Dao.Zap.Logger.Error(
			"UserDaoImpl.GetUserByEmail",
			zap.String("email", *email),
			zap.Error(err),
		)
		return nil, err
	} else {
		u.Dao.Zap.Logger.Info(
			"UserDaoImpl.GetUserByEmail",
			zap.String("email", *email),
		)
		return &user, nil
	}
}

func (u *UserDaoImpl) GetUserList(
	ctx context.Context,
	offset, limit int64, desc bool, organization, role *string,
	createStartTime, createEndTime, updateStartTime, updateEndTime, lastLoginStartTime, lastLoginEndTime *time.Time,
	query *string,
) ([]models.UserModel, *int64, error) {
	collection := u.Dao.Mongo.MongoDatabase.Collection(config.UserCollectionName)
	var users []models.UserModel
	var err error
	doc := bson.M{"deleted": false}
	if organization != nil {
		doc["organization"] = *organization
	}
	if role != nil {
		doc["role"] = *role
	}
	if createStartTime != nil && createEndTime != nil {
		doc["created_at"] = bson.M{"$gte": createStartTime, "$lte": createEndTime}
	}
	if updateStartTime != nil && updateEndTime != nil {
		doc["updated_at"] = bson.M{"$gte": updateStartTime, "$lte": updateEndTime}
	}
	if lastLoginStartTime != nil && lastLoginEndTime != nil {
		doc["last_login_at"] = bson.M{"$gte": lastLoginStartTime, "$lte": lastLoginEndTime}
	}
	if query != nil {
		doc["$or"] = []bson.M{
			{"user_id": bson.M{"$regex": *query}},
			{"user_name": bson.M{"$regex": *query}},
			{"organization": bson.M{"$regex": *query}},
			{"role": bson.M{"$regex": *query}},
		}
	}
	docJSON, _ := json.Marshal(doc)
	if desc {
		err = collection.Find(ctx, doc).Sort("-created_at").Skip(offset).Limit(limit).All(&users)
	} else {
		err = collection.Find(ctx, doc).Skip(offset).Limit(limit).All(&users)
	}
	if err != nil {
		u.Dao.Zap.Logger.Error(
			"UserDaoImpl.GetUserList",
			zap.Int64("offset", offset), zap.Int64("limit", limit), zap.Bool("desc", desc),
			zap.ByteString(config.UserCollectionName, docJSON), zap.Error(err),
		)
		return nil, nil, err
	}
	count, err := collection.Find(ctx, doc).Count()
	if err != nil {
		u.Dao.Zap.Logger.Error(
			"UserDaoImpl.GetUserList",
			zap.Int64("offset", offset), zap.Int64("limit", limit), zap.Bool("desc", desc),
			zap.ByteString(config.UserCollectionName, docJSON), zap.Error(err),
		)
		return nil, nil, err
	}
	u.Dao.Zap.Logger.Info(
		"UserDaoImpl.GetUserList",
		zap.Int64("offset", offset), zap.Int64("limit", limit), zap.Bool("desc", desc),
		zap.ByteString(config.UserCollectionName, docJSON), zap.Int64("count", count),
	)
	return users, &count, nil
}

func (u *UserDaoImpl) CountUser(
	ctx context.Context,
	organization, role *string,
	createStartTime, createEndTime, updateStartTime, updateEndTime, lastLoginStartTime, lastLoginEndTime *time.Time,
) (*int64, error) {
	collection := u.Dao.Mongo.MongoDatabase.Collection(config.UserCollectionName)
	doc := bson.M{"deleted": false}
	if organization != nil {
		doc["organization"] = *organization
	}
	if role != nil {
		doc["role"] = *role
	}
	if createStartTime != nil && createEndTime != nil {
		doc["created_at"] = bson.M{"$gte": createStartTime, "$lte": createEndTime}
	}
	if updateStartTime != nil && updateEndTime != nil {
		doc["updated_at"] = bson.M{"$gte": updateStartTime, "$lte": updateEndTime}
	}
	if lastLoginStartTime != nil && lastLoginEndTime != nil {
		doc["last_login_at"] = bson.M{"$gte": lastLoginStartTime, "$lte": lastLoginEndTime}
	}
	docJSON, _ := json.Marshal(doc)
	count, err := collection.Find(ctx, doc).Count()
	if err != nil {
		u.Dao.Zap.Logger.Error(
			"UserDaoImpl.CountUser",
			zap.ByteString(config.UserCollectionName, docJSON),
			zap.Error(err),
		)
	} else {
		u.Dao.Zap.Logger.Info(
			"UserDaoImpl.CountUser",
			zap.ByteString(config.UserCollectionName, docJSON),
			zap.Int64("count", count),
		)
	}
	return &count, err
}

func (u *UserDaoImpl) InsertUser(
	ctx context.Context,
	username, email, password, role, organization string,
) (primitive.ObjectID, error) {
	collection := u.Dao.Mongo.MongoDatabase.Collection(config.UserCollectionName)
	doc := bson.M{
		"username":      username,
		"email":         email,
		"password":      password,
		"role":          role,
		"organization":  organization,
		"last_login_at": nil,
		"deleted":       false,
		"created_at":    time.Now(),
		"updated_at":    time.Now(),
		"deleted_at":    nil,
	}
	docJSON, _ := json.Marshal(doc)
	result, err := collection.InsertOne(ctx, doc)
	if err != nil {
		u.Dao.Zap.Logger.Error(
			"UserDaoImpl.InsertUser",
			zap.ByteString(config.UserCollectionName, docJSON),
			zap.Error(err),
		)
	} else {
		u.Dao.Zap.Logger.Info(
			"UserDaoImpl.InsertUser",
			zap.ByteString(config.UserCollectionName, docJSON),
			zap.String("userID", result.InsertedID.(primitive.ObjectID).Hex()),
		)
	}
	return result.InsertedID.(primitive.ObjectID), err
}

func (u *UserDaoImpl) UpdateUser(
	ctx context.Context, userID primitive.ObjectID, username, email, password, role, organization *string,
) error {
	collection := u.Dao.Mongo.MongoDatabase.Collection(config.UserCollectionName)
	doc := bson.M{"updated_at": time.Now()}
	if username != nil {
		doc["username"] = *username
	}
	if email != nil {
		doc["email"] = *email
	}
	if password != nil {
		doc["password"] = *password
	}
	if role != nil {
		doc["role"] = *role
	}
	if organization != nil {
		doc["organization"] = *organization
	}
	docJSON, _ := json.Marshal(doc)
	err := collection.UpdateId(ctx, userID, bson.M{"$set": doc})
	if err != nil {
		u.Dao.Zap.Logger.Error(
			"UserDaoImpl.UpdateUser",
			zap.String("userID", userID.Hex()),
			zap.ByteString(config.UserCollectionName, docJSON),
			zap.Error(err),
		)
	} else {
		u.Dao.Zap.Logger.Info(
			"UserDaoImpl.UpdateUser",
			zap.String("userID", userID.Hex()),
			zap.ByteString(config.UserCollectionName, docJSON),
		)
	}
	return err
}

func (u *UserDaoImpl) SoftDeleteUser(ctx context.Context, userID primitive.ObjectID) error {
	collection := u.Dao.Mongo.MongoDatabase.Collection(config.UserCollectionName)
	err := collection.UpdateId(ctx, userID, bson.M{"$set": bson.M{"deleted": true, "deleted_at": time.Now()}})
	if err != nil {
		u.Dao.Zap.Logger.Error(
			"UserDaoImpl.DeleteUser",
			zap.String("userID", userID.Hex()),
			zap.Error(err),
		)
	} else {
		u.Dao.Zap.Logger.Info(
			"UserDaoImpl.DeleteUser",
			zap.String("userID", userID.Hex()),
		)
	}
	return err
}

func (u *UserDaoImpl) SoftDeleteUserList(
	ctx context.Context, organization, role *string,
	createStartTime, createEndTime, updateStartTime, updateEndTime, lastLoginStartTime, lastLoginEndTime *time.Time,
) (*int64, error) {
	collection := u.Dao.Mongo.MongoDatabase.Collection(config.UserCollectionName)
	doc := bson.M{"deleted": false}
	if organization != nil {
		doc["organization"] = *organization
	}
	if role != nil {
		doc["role"] = *role
	}
	if createStartTime != nil && createEndTime != nil {
		doc["created_at"] = bson.M{"$gte": createStartTime, "$lte": createEndTime}
	}
	if updateStartTime != nil && updateEndTime != nil {
		doc["updated_at"] = bson.M{"$gte": updateStartTime, "$lte": updateEndTime}
	}
	if lastLoginStartTime != nil && lastLoginEndTime != nil {
		doc["last_login_at"] = bson.M{"$gte": lastLoginStartTime, "$lte": lastLoginEndTime}
	}
	docJSON, _ := json.Marshal(doc)
	result, err := collection.UpdateAll(ctx, doc, bson.M{"$set": bson.M{"deleted": true, "deleted_at": time.Now()}})
	if err != nil {
		u.Dao.Zap.Logger.Error(
			"UserDaoImpl.DeleteUserList",
			zap.ByteString(config.UserCollectionName, docJSON),
			zap.Error(err),
		)
	} else {
		u.Dao.Zap.Logger.Info(
			"UserDaoImpl.DeleteUserList",
			zap.ByteString(config.UserCollectionName, docJSON),
			zap.Int64("count", result.ModifiedCount),
		)
	}
	return &result.ModifiedCount, err
}

func (u *UserDaoImpl) DeleteUser(ctx context.Context, userID primitive.ObjectID) error {
	collection := u.Dao.Mongo.MongoDatabase.Collection(config.UserCollectionName)
	err := collection.RemoveId(ctx, userID)
	if err != nil {
		u.Dao.Zap.Logger.Error(
			"UserDaoImpl.DeleteUser",
			zap.String("userID", userID.Hex()),
			zap.Error(err),
		)
	} else {
		u.Dao.Zap.Logger.Info(
			"UserDaoImpl.DeleteUser",
			zap.String("userID", userID.Hex()),
		)
	}
	return err
}

func (u *UserDaoImpl) DeleteUserList(
	ctx context.Context, organization, role *string,
	createStartTime, createEndTime, updateStartTime, updateEndTime, lastLoginStartTime, lastLoginEndTime *time.Time,
) (*int64, error) {
	collection := u.Dao.Mongo.MongoDatabase.Collection(config.UserCollectionName)
	doc := bson.M{}
	if organization != nil {
		doc["organization"] = *organization
	}
	if role != nil {
		doc["role"] = *role
	}
	if createStartTime != nil && createEndTime != nil {
		doc["created_at"] = bson.M{"$gte": createStartTime, "$lte": createEndTime}
	}
	if updateStartTime != nil && updateEndTime != nil {
		doc["updated_at"] = bson.M{"$gte": updateStartTime, "$lte": updateEndTime}
	}
	if lastLoginStartTime != nil && lastLoginEndTime != nil {
		doc["last_login_at"] = bson.M{"$gte": lastLoginStartTime, "$lte": lastLoginEndTime}
	}
	docJSON, _ := json.Marshal(doc)
	result, err := collection.RemoveAll(ctx, doc)
	if err != nil {
		u.Dao.Zap.Logger.Error(
			"UserDaoImpl.DeleteUserList",
			zap.ByteString(config.UserCollectionName, docJSON),
			zap.Error(err),
		)
	} else {
		u.Dao.Zap.Logger.Info(
			"UserDaoImpl.DeleteUserList",
			zap.ByteString(config.UserCollectionName, docJSON),
			zap.Int64("count", result.DeletedCount),
		)
	}
	return &result.DeletedCount, err
}
