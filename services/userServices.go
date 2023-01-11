package services

import (
	"depen/dto"
	"depen/model"
	"depen/repo"
	"depen/validation"
	"net/http"
)

type UserServices struct {
	userRepo *repo.UserRepo
}

func (u *UserServices) RegisterUser(model *model.RegisterUser) *dto.UserDto {

	validate := validation.ValidateStruct(*model)
	if validate != nil {
		var errMessage []string
		for _, err := range validate {
			errMessage = append(errMessage, err.FailedField+" "+err.Tag+" "+err.Value)
		}
		return &dto.UserDto{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         errMessage,
			Data:           nil,
		}
	}

	u.userRepo.RegisterUser(model)
	return &dto.UserDto{
		HttpStatusCode: http.StatusOK,
		Errors:         nil,
		Data:           model,
	}

}
func (u *UserServices) LoginUser() error {
	return nil
}

func NewUserServices(userRepo *repo.UserRepo) *UserServices {
	return &UserServices{userRepo: userRepo}
}
