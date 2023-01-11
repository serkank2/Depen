package repo

import (
	"context"
	"depen/dto"
	"depen/model"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	mongoCollection *mongo.Collection
}

func (u *UserRepo) RegisterUser(model *model.RegisterUser) *dto.UserRegisterDto {
	check := CheckUser(model.Email, u.mongoCollection)
	if !check {
		return &dto.UserRegisterDto{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         []string{"Email already exists"},
			Register:       "failed",
			Data:           nil,
		}
	}
	insertID, err := u.mongoCollection.InsertOne(context.TODO(), bson.D{
		{Key: "email", Value: model.Email},
		{Key: "password", Value: model.Password},
		{Key: "created_at", Value: time.Now()},
		{Key: "updated", Value: time.Now()},
	})
	if err != nil {
		return &dto.UserRegisterDto{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         []string{err.Error()},
			Data:           nil,
		}
	}
	return &dto.UserRegisterDto{
		HttpStatusCode: http.StatusCreated,
		Errors:         nil,
		Register:       "success",
		Data:           bson.M{"insertID": insertID.InsertedID},
	}
}
func (u *UserRepo) LoginUser(model *model.LoginUser) *dto.UserLoginDto {
	var user dto.User
	filter := bson.D{
		{Key: "$and", Value: bson.A{
			bson.D{{Key: "email", Value: model.Email}},
			bson.D{{Key: "password", Value: model.Password}},
		}},
	}
	err := u.mongoCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return &dto.UserLoginDto{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         []string{"User not found"},
			Data:           nil,
			Login:          "failed",
		}
	}
	if err != nil {
		return &dto.UserLoginDto{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         []string{err.Error()},
			Data:           nil,
			Login:          "failed",
		}
	}
	return &dto.UserLoginDto{
		HttpStatusCode: http.StatusOK,
		Errors:         nil,
		Login:          "Login success",
		Data: bson.M{
			"Email":    user.Email,
			"CretedAt": user.Created_at,
		},
	}
}

func NewUserRepo(mongoCollection *mongo.Collection) *UserRepo {
	return &UserRepo{mongoCollection: mongoCollection}
}
