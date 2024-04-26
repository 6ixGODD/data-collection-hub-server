package modules

import (
	"context"

	"data-collection-hub-server/dal"
	"data-collection-hub-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type UserDao interface {
	GetUserById(userId string, ctx context.Context) (*models.UserModel, error)
	GetUserList(offset, limit int64, desc bool, ctx context.Context) ([]models.UserModel, error)
	GetUserListByOrganization(organization string, offset, limit int64, desc bool, ctx context.Context) ([]models.UserModel, error)
	GetUserListByRole(role string, offset, limit int64, desc bool, ctx context.Context) ([]models.UserModel, error)
	GetUserListByCreatedTime(startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.UserModel, error)
	GetUserListByUpdatedTime(startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.UserModel, error)
	GetUserListByFuzzyQuery(query string, offset, limit int64, desc bool, ctx context.Context) ([]models.UserModel, error)
	InsertUser(user *models.UserModel, ctx context.Context) error
	UpdateUser(user *models.UserModel, ctx context.Context) error
	DeleteUser(user *models.UserModel, ctx context.Context) error
}

type UserDaoImpl struct{ *dal.Dao }

func NewUserDao(dao *dal.Dao) UserDao {
	var _ UserDao = new(UserDaoImpl)
	return &UserDaoImpl{dao}
}

func (u UserDaoImpl) GetUserById(userId string, ctx context.Context) (*models.UserModel, error) {
	var user models.UserModel
	collection := u.Dao.MongoDB.Collection("user")
	err := collection.Find(ctx, bson.M{"user_id": userId}).One(&user)
	if err != nil {
		u.Dao.Logger.Error(
			"UserDaoImpl.GetUserById",
			zap.Field{Key: "userId", Type: zapcore.StringType, String: userId},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		u.Dao.Logger.Info(
			"UserDaoImpl.GetUserById",
			zap.Field{Key: "userId", Type: zapcore.StringType, String: userId},
		)
		return &user, nil
	}
}

func (u UserDaoImpl) GetUserList(offset, limit int64, desc bool, ctx context.Context) ([]models.UserModel, error) {
	var users []models.UserModel
	var err error
	collection := u.Dao.MongoDB.Collection("user")
	if desc {
		err = collection.Find(ctx, bson.M{}).Sort("-created_time").Skip(offset).Limit(limit).All(&users)
	} else {
		err = collection.Find(ctx, bson.M{}).Skip(offset).Limit(limit).All(&users)
	}
	if err != nil {
		u.Dao.Logger.Error(
			"UserDaoImpl.GetUserList",
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		return users, nil
	}
}

func (u UserDaoImpl) GetUserListByOrganization(organization string, offset, limit int64, desc bool, ctx context.Context) ([]models.UserModel, error) {
	var users []models.UserModel
	var err error
	collection := u.Dao.MongoDB.Collection("user")
	if desc {
		err = collection.Find(ctx, bson.M{"organization": organization}).Sort("-created_time").Skip(offset).Limit(limit).All(&users)
	} else {
		err = collection.Find(ctx, bson.M{"organization": organization}).Skip(offset).Limit(limit).All(&users)
	}
	if err != nil {
		u.Dao.Logger.Error(
			"UserDaoImpl.GetUserListByOrganization",
			zap.Field{Key: "organization", Type: zapcore.StringType, String: organization},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		u.Dao.Logger.Info(
			"UserDaoImpl.GetUserListByOrganization",
			zap.Field{Key: "organization", Type: zapcore.StringType, String: organization},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return users, nil
	}
}

func (u UserDaoImpl) GetUserListByRole(role string, offset, limit int64, desc bool, ctx context.Context) ([]models.UserModel, error) {
	var users []models.UserModel
	var err error
	collection := u.Dao.MongoDB.Collection("user")
	if desc {
		err = collection.Find(ctx, bson.M{"role": role}).Sort("-created_time").Skip(offset).Limit(limit).All(&users)
	} else {
		err = collection.Find(ctx, bson.M{"role": role}).Skip(offset).Limit(limit).All(&users)
	}
	if err != nil {
		u.Dao.Logger.Error(
			"UserDaoImpl.GetUserListByRole",
			zap.Field{Key: "role", Type: zapcore.StringType, String: role},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		u.Dao.Logger.Info(
			"UserDaoImpl.GetUserListByRole",
			zap.Field{Key: "role", Type: zapcore.StringType, String: role},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return users, nil
	}
}

func (u UserDaoImpl) GetUserListByCreatedTime(startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.UserModel, error) {
	var users []models.UserModel
	var err error
	collection := u.Dao.MongoDB.Collection("user")
	if desc {
		err = collection.Find(ctx, bson.M{"created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Sort("-created_time").Skip(offset).Limit(limit).All(&users)
	} else {
		err = collection.Find(ctx, bson.M{"created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Skip(offset).Limit(limit).All(&users)
	}
	if err != nil {
		u.Dao.Logger.Error(
			"UserDaoImpl.GetUserListByCreatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		u.Dao.Logger.Info(
			"UserDaoImpl.GetUserListByCreatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return users, nil
	}
}

func (u UserDaoImpl) GetUserListByUpdatedTime(startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.UserModel, error) {
	var users []models.UserModel
	var err error
	collection := u.Dao.MongoDB.Collection("user")
	if desc {
		err = collection.Find(ctx, bson.M{"updated_time": bson.M{"$gte": startTime, "$lte": endTime}}).Sort("-created_time").Skip(offset).Limit(limit).All(&users)
	} else {
		err = collection.Find(ctx, bson.M{"updated_time": bson.M{"$gte": startTime, "$lte": endTime}}).Skip(offset).Limit(limit).All(&users)
	}
	if err != nil {
		u.Dao.Logger.Error(
			"UserDaoImpl.GetUserListByUpdatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		u.Dao.Logger.Info(
			"UserDaoImpl.GetUserListByUpdatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return users, nil
	}
}

func (u UserDaoImpl) GetUserListByFuzzyQuery(query string, offset, limit int64, desc bool, ctx context.Context) ([]models.UserModel, error) {
	var users []models.UserModel
	var err error
	collection := u.Dao.MongoDB.Collection("user")
	if desc {
		err = collection.Find(ctx, bson.M{"$or": []bson.M{
			{"user_id": bson.M{"$regex": query}},
			{"user_name": bson.M{"$regex": query}},
			{"organization": bson.M{"$regex": query}},
			{"role": bson.M{"$regex": query}},
		}}).Sort("-created_time").Skip(offset).Limit(limit).All(&users)
	} else {
		err = collection.Find(ctx, bson.M{"$or": []bson.M{
			{"user_id": bson.M{"$regex": query}},
			{"user_name": bson.M{"$regex": query}},
			{"organization": bson.M{"$regex": query}},
			{"role": bson.M{"$regex": query}},
		}}).Skip(offset).Limit(limit).All(&users)
	}
	if err != nil {
		u.Dao.Logger.Error(
			"UserDaoImpl.GetUserListByFuzzyQuery",
			zap.Field{Key: "query", Type: zapcore.StringType, String: query},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		u.Dao.Logger.Info(
			"UserDaoImpl.GetUserListByFuzzyQuery",
			zap.Field{Key: "query", Type: zapcore.StringType, String: query},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Interface: desc},
		)
		return users, nil
	}
}

func (u UserDaoImpl) InsertUser(user *models.UserModel, ctx context.Context) error {
	collection := u.Dao.MongoDB.Collection("user")
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		u.Dao.Logger.Error(
			"UserDaoImpl.InsertUser",
			zap.Field{Key: "user", Type: zapcore.ObjectMarshalerType, Interface: user},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		u.Dao.Logger.Info(
			"UserDaoImpl.InsertUser",
			zap.Field{Key: "user", Type: zapcore.ObjectMarshalerType, Interface: user},
			zap.Field{Key: "result", Type: zapcore.ObjectMarshalerType, Interface: result},
		)
	}
	return err
}

func (u UserDaoImpl) UpdateUser(user *models.UserModel, ctx context.Context) error {
	collection := u.Dao.MongoDB.Collection("user")
	err := collection.UpdateOne(ctx, bson.M{"_id": user.UserID}, bson.M{"$set": user})
	if err != nil {
		u.Dao.Logger.Error(
			"UserDaoImpl.UpdateUser",
			zap.Field{Key: "user", Type: zapcore.ObjectMarshalerType, Interface: user},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		u.Dao.Logger.Info(
			"UserDaoImpl.UpdateUser",
			zap.Field{Key: "user", Type: zapcore.ObjectMarshalerType, Interface: user},
		)
	}
	return err
}

func (u UserDaoImpl) DeleteUser(user *models.UserModel, ctx context.Context) error {
	collection := u.Dao.MongoDB.Collection("user")
	err := collection.RemoveId(ctx, user.UserID)
	if err != nil {
		u.Dao.Logger.Error(
			"UserDaoImpl.DeleteUser",
			zap.Field{Key: "user", Type: zapcore.ObjectMarshalerType, Interface: user},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		u.Dao.Logger.Info(
			"UserDaoImpl.DeleteUser",
			zap.Field{Key: "user", Type: zapcore.ObjectMarshalerType, Interface: user},
		)
	}
	return err
}
