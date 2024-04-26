package modules

import (
	"context"

	"data-collection-hub-server/dal"
	"data-collection-hub-server/models"
	"go.mongodb.org/mongo-driver/bson"
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
		return nil, err
	} else {
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
		return nil, err
	} else {
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
		return nil, err
	} else {
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
		return nil, err
	} else {
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
		return nil, err
	} else {
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
		return nil, err
	} else {
		return users, nil
	}
}

func (u UserDaoImpl) InsertUser(user *models.UserModel, ctx context.Context) error {
	collection := u.Dao.MongoDB.Collection("user")
	_, err := collection.InsertOne(ctx, user)
	return err
}

func (u UserDaoImpl) UpdateUser(user *models.UserModel, ctx context.Context) error {
	collection := u.Dao.MongoDB.Collection("user")
	err := collection.UpdateOne(ctx, bson.M{"_id": user.UserID}, bson.M{"$set": user})
	return err
}

func (u UserDaoImpl) DeleteUser(user *models.UserModel, ctx context.Context) error {
	collection := u.Dao.MongoDB.Collection("user")
	err := collection.RemoveId(ctx, user.UserID)
	return err
}
