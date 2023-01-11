package model

type RegisterUser struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=8,max=20" json:"password`
}
