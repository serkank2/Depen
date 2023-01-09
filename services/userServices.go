package services

import (
	"depen/model"
	"depen/repo"
)

type UserServices struct {
	userRepo *repo.UserRepo
}

func (u *UserServices) RegisterUser(model *model.RegisterUser) error {
	u.userRepo.RegisterUser(model)
	return nil
}
func (u *UserServices) LoginUser() error {
	return nil
}

func NewUserServices(userRepo *repo.UserRepo) *UserServices {
	return &UserServices{userRepo: userRepo}
}
