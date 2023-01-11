package repo

import (
	"context"
	"depen/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	mongoCollection *mongo.Collection
}

func (u *UserRepo) RegisterUser(model *model.RegisterUser) error {
	u.mongoCollection.InsertOne(context.TODO(), model)
	return nil
}
func (u *UserRepo) LoginUser() error {
	return nil
}

func NewUserRepo(mongoCollection *mongo.Collection) *UserRepo {
	return &UserRepo{mongoCollection: mongoCollection}
}
