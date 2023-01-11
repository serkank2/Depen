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

func (u *UserServices) RegisterUser(model *model.RegisterUser) *dto.UserRegisterDto {

	validate := validation.ValidateRegisterStruct(*model)
	if validate != nil {
		var errMessage []string
		for _, err := range validate {
			errMessage = append(errMessage, err.FailedField+" "+err.Tag+" "+err.Value)
		}
		return &dto.UserRegisterDto{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         errMessage,
			Data:           nil,
		}
	}
	result := u.userRepo.RegisterUser(model)
	return result
}
func (u *UserServices) LoginUser(model *model.LoginUser) *dto.UserLoginDto {

	validate := validation.ValidateLoginStruct(*model)
	if validate != nil {
		var errMessage []string
		for _, err := range validate {
			errMessage = append(errMessage, err.FailedField+" "+err.Tag+" "+err.Value)
		}
		return &dto.UserLoginDto{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         errMessage,
			Data:           nil,
		}
	}
	result := u.userRepo.LoginUser(model)
	return result

}

func NewUserServices(userRepo *repo.UserRepo) *UserServices {
	return &UserServices{userRepo: userRepo}
}
