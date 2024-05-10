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

const userCollectionName = "user"

type UserDao interface {
	GetUserById(userID primitive.ObjectID, ctx context.Context) (*models.UserModel, error)
	GetUserList(offset, limit int64, desc bool, ctx context.Context) ([]models.UserModel, error)
	GetUserListByOrganization(
		organization string, offset, limit int64, desc bool, ctx context.Context,
	) ([]models.UserModel, error)
	GetUserListByRole(role string, offset, limit int64, desc bool, ctx context.Context) ([]models.UserModel, error)
	GetUserListByCreatedTime(
		startTime, endTime time.Time, offset, limit int64, desc bool, ctx context.Context,
	) ([]models.UserModel, error)
	GetUserListByLastLoginTime(
		startTime, endTime time.Time, offset, limit int64, desc bool, ctx context.Context,
	) ([]models.UserModel, error)
	GetUserListByUpdatedTime(
		startTime, endTime time.Time, offset, limit int64, desc bool, ctx context.Context,
	) ([]models.UserModel, error)
	GetUserListByFuzzyQuery(query string, offset, limit int64, desc bool, ctx context.Context) (
		[]models.UserModel, error,
	)
	InsertUser(username, email, password, role, organization string, ctx context.Context) (primitive.ObjectID, error)
	UpdateUser(
		userID primitive.ObjectID, username, email, password, role, organization string, ctx context.Context,
	) error
	SoftDeleteUser(userID primitive.ObjectID, ctx context.Context) error
	DeleteUser(userID primitive.ObjectID, ctx context.Context) error
}

type UserDaoImpl struct{ *dal.Dao }

func NewUserDao(dao *dal.Dao) UserDao {
	var _ UserDao = (*UserDaoImpl)(nil) // Ensure that the interface is implemented
	return &UserDaoImpl{dao}
}

func (u UserDaoImpl) GetUserById(userID primitive.ObjectID, ctx context.Context) (*models.UserModel, error) {
	var user models.UserModel
	collection := u.Dao.Mongo.MongoDatabase.Collection(userCollectionName)
	err := collection.Find(
		ctx,
		bson.M{
			"user_id": userID,
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

func (u UserDaoImpl) GetUserList(offset, limit int64, desc bool, ctx context.Context) ([]models.UserModel, error) {
	var users []models.UserModel
	var err error
	collection := u.Dao.Mongo.MongoDatabase.Collection(userCollectionName)
	if desc {
		err = collection.Find(
			ctx,
			bson.M{
				"deleted": false,
			},
		).Sort("-created_time").Skip(offset).Limit(limit).All(&users)
	} else {
		err = collection.Find(
			ctx,
			bson.M{
				"deleted": false,
			},
		).Skip(offset).Limit(limit).All(&users)
	}
	if err != nil {
		u.Dao.Zap.Logger.Error(
			"UserDaoImpl.GetUserList",
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		u.Dao.Zap.Logger.Info(
			"UserDaoImpl.GetUserList",
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
	}
	return users, err
}

func (u UserDaoImpl) GetUserListByOrganization(
	organization string, offset, limit int64, desc bool, ctx context.Context,
) ([]models.UserModel, error) {
	var users []models.UserModel
	var err error
	collection := u.Dao.Mongo.MongoDatabase.Collection(userCollectionName)
	if desc {
		err = collection.Find(
			ctx,
			bson.M{
				"organization": organization,
				"deleted":      false,
			},
		).Sort("-created_time").Skip(offset).Limit(limit).All(&users)
	} else {
		err = collection.Find(
			ctx,
			bson.M{
				"organization": organization,
				"deleted":      false,
			},
		).Skip(offset).Limit(limit).All(&users)
	}
	if err != nil {
		u.Dao.Zap.Logger.Error(
			"UserDaoImpl.GetUserListByOrganization",
			zap.Field{Key: "organization", Type: zapcore.StringType, String: organization},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		u.Dao.Zap.Logger.Info(
			"UserDaoImpl.GetUserListByOrganization",
			zap.Field{Key: "organization", Type: zapcore.StringType, String: organization},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
	}
	return users, err
}

func (u UserDaoImpl) GetUserListByRole(
	role string, offset, limit int64, desc bool, ctx context.Context,
) ([]models.UserModel, error) {
	var users []models.UserModel
	var err error
	collection := u.Dao.Mongo.MongoDatabase.Collection(userCollectionName)
	if desc {
		err = collection.Find(
			ctx,
			bson.M{
				"role":    role,
				"deleted": false,
			},
		).Sort("-created_time").Skip(offset).Limit(limit).All(&users)
	} else {
		err = collection.Find(
			ctx,
			bson.M{
				"role":    role,
				"deleted": false,
			},
		).Skip(offset).Limit(limit).All(&users)
	}
	if err != nil {
		u.Dao.Zap.Logger.Error(
			"UserDaoImpl.GetUserListByRole",
			zap.Field{Key: "role", Type: zapcore.StringType, String: role},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		u.Dao.Zap.Logger.Info(
			"UserDaoImpl.GetUserListByRole",
			zap.Field{Key: "role", Type: zapcore.StringType, String: role},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
	}
	return users, err
}

func (u UserDaoImpl) GetUserListByCreatedTime(
	startTime, endTime time.Time, offset, limit int64, desc bool, ctx context.Context,
) ([]models.UserModel, error) {
	var users []models.UserModel
	var err error
	collection := u.Dao.Mongo.MongoDatabase.Collection(userCollectionName)
	if desc {
		err = collection.Find(
			ctx,
			bson.M{
				"created_time": bson.M{"$gte": startTime, "$lte": endTime},
				"deleted":      false,
			},
		).Sort("-created_time").Skip(offset).Limit(limit).All(&users)
	} else {
		err = collection.Find(
			ctx,
			bson.M{
				"created_time": bson.M{"$gte": startTime, "$lte": endTime},
				"deleted":      false,
			},
		).Skip(offset).Limit(limit).All(&users)
	}
	if err != nil {
		u.Dao.Zap.Logger.Error(
			"UserDaoImpl.GetUserListByCreatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		u.Dao.Zap.Logger.Info(
			"UserDaoImpl.GetUserListByCreatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return users, nil
	}
}

func (u UserDaoImpl) GetUserListByUpdatedTime(
	startTime, endTime time.Time, offset, limit int64, desc bool, ctx context.Context,
) ([]models.UserModel, error) {
	var users []models.UserModel
	var err error
	collection := u.Dao.Mongo.MongoDatabase.Collection(userCollectionName)
	if desc {
		err = collection.Find(
			ctx,
			bson.M{
				"updated_time": bson.M{"$gte": startTime, "$lte": endTime},
				"deleted":      false,
			},
		).Sort("-created_time").Skip(offset).Limit(limit).All(&users)
	} else {
		err = collection.Find(
			ctx,
			bson.M{
				"updated_time": bson.M{"$gte": startTime, "$lte": endTime},
				"deleted":      false,
			},
		).Skip(offset).Limit(limit).All(&users)
	}
	if err != nil {
		u.Dao.Zap.Logger.Error(
			"UserDaoImpl.GetUserListByUpdatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		u.Dao.Zap.Logger.Info(
			"UserDaoImpl.GetUserListByUpdatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return users, nil
	}
}

func (u UserDaoImpl) GetUserListByLastLoginTime(
	startTime, endTime time.Time, offset, limit int64, desc bool, ctx context.Context,
) ([]models.UserModel, error) {
	var users []models.UserModel
	var err error
	collection := u.Dao.Mongo.MongoDatabase.Collection(userCollectionName)
	if desc {
		err = collection.Find(
			ctx,
			bson.M{
				"last_login_time": bson.M{"$gte": startTime, "$lte": endTime},
				"deleted":         false,
			},
		).Sort("-created_time").Skip(offset).Limit(limit).All(&users)
	} else {
		err = collection.Find(
			ctx,
			bson.M{
				"last_login_time": bson.M{"$gte": startTime, "$lte": endTime},
				"deleted":         false,
			},
		).Skip(offset).Limit(limit).All(&users)
	}
	if err != nil {
		u.Dao.Zap.Logger.Error(
			"UserDaoImpl.GetUserListByLastLoginTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		u.Dao.Zap.Logger.Info(
			"UserDaoImpl.GetUserListByLastLoginTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return users, nil
	}
}

func (u UserDaoImpl) GetUserListByFuzzyQuery(
	query string, offset, limit int64, desc bool, ctx context.Context,
) ([]models.UserModel, error) {
	var users []models.UserModel
	var err error
	collection := u.Dao.Mongo.MongoDatabase.Collection(userCollectionName)
	if desc {
		err = collection.Find(
			ctx,
			bson.M{
				"$and": []bson.M{
					{"deleted": false},
					{
						"$or": []bson.M{
							{"user_id": bson.M{"$regex": query}},
							{"user_name": bson.M{"$regex": query}},
							{"organization": bson.M{"$regex": query}},
							{"role": bson.M{"$regex": query}},
						},
					},
				},
			},
		).Sort("-created_time").Skip(offset).Limit(limit).All(&users)
	} else {
		err = collection.Find(
			ctx,
			bson.M{
				"$and": []bson.M{
					{"deleted": false},
					{
						"$or": []bson.M{
							{"user_id": bson.M{"$regex": query}},
							{"user_name": bson.M{"$regex": query}},
							{"organization": bson.M{"$regex": query}},
							{"role": bson.M{"$regex": query}},
						},
					},
				},
			},
		).Skip(offset).Limit(limit).All(&users)
	}
	if err != nil {
		u.Dao.Zap.Logger.Error(
			"UserDaoImpl.GetUserListByFuzzyQuery",
			zap.Field{Key: "query", Type: zapcore.StringType, String: query},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		u.Dao.Zap.Logger.Info(
			"UserDaoImpl.GetUserListByFuzzyQuery",
			zap.Field{Key: "query", Type: zapcore.StringType, String: query},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return users, nil
	}
}

func (u UserDaoImpl) InsertUser(
	username, email, password, role, organization string, ctx context.Context,
) (primitive.ObjectID, error) {
	collection := u.Dao.Mongo.MongoDatabase.Collection(userCollectionName)
	user := models.UserModel{
		Username:     username,
		Email:        email,
		Password:     password,
		Role:         role,
		Organization: organization,
		Deleted:      false,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		u.Dao.Zap.Logger.Error(
			"UserDaoImpl.InsertUser",
			zap.Field{Key: userCollectionName, Type: zapcore.ObjectMarshalerType, Interface: user},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		u.Dao.Zap.Logger.Info(
			"UserDaoImpl.InsertUser",
			zap.Field{Key: "userID", Type: zapcore.StringType, String: result.InsertedID.(primitive.ObjectID).Hex()},
			zap.Field{Key: userCollectionName, Type: zapcore.ObjectMarshalerType, Interface: user},
		)
	}
	return result.InsertedID.(primitive.ObjectID), err
}

func (u UserDaoImpl) UpdateUser(
	userID primitive.ObjectID, username, email, password, role, organization string, ctx context.Context,
) error {
	collection := u.Dao.Mongo.MongoDatabase.Collection(userCollectionName)
	err := collection.UpdateId(
		ctx, userID, bson.M{
			"$set": bson.M{
				"username":     username,
				"email":        email,
				"password":     password,
				"role":         role,
				"organization": organization,
				"updated_at":   time.Now(),
			},
		},
	)
	if err != nil {
		u.Dao.Zap.Logger.Error(
			"UserDaoImpl.UpdateUser",
			zap.Field{Key: "userID", Type: zapcore.StringType, String: userID.Hex()},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		u.Dao.Zap.Logger.Info(
			"UserDaoImpl.UpdateUser",
			zap.Field{Key: "userID", Type: zapcore.StringType, String: userID.Hex()},
		)
	}
	return err
}

func (u UserDaoImpl) SoftDeleteUser(userID primitive.ObjectID, ctx context.Context) error {
	collection := u.Dao.Mongo.MongoDatabase.Collection(userCollectionName)
	err := collection.UpdateId(ctx, userID, bson.M{"$set": bson.M{"deleted": true, "deleted_at": time.Now()}})
	if err != nil {
		u.Dao.Zap.Logger.Error(
			"UserDaoImpl.DeleteUser",
			zap.Field{Key: "userID", Type: zapcore.StringType, String: userID.Hex()},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		u.Dao.Zap.Logger.Info(
			"UserDaoImpl.DeleteUser",
			zap.Field{Key: "userID", Type: zapcore.StringType, String: userID.Hex()},
		)
	}
	return err
}

func (u UserDaoImpl) DeleteUser(userID primitive.ObjectID, ctx context.Context) error {
	collection := u.Dao.Mongo.MongoDatabase.Collection(userCollectionName)
	err := collection.RemoveId(ctx, userID)
	if err != nil {
		u.Dao.Zap.Logger.Error(
			"UserDaoImpl.DeleteUser",
			zap.Field{Key: "userID", Type: zapcore.StringType, String: userID.Hex()},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		u.Dao.Zap.Logger.Info(
			"UserDaoImpl.DeleteUser",
			zap.Field{Key: "userID", Type: zapcore.StringType, String: userID.Hex()},
		)
	}
	return err
}
