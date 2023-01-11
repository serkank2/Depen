package dto

import "time"

type UserRegisterDto struct {
	HttpStatusCode int         `json:"http_status_code"`
	Register       string      `json:"register"`
	Errors         []string    `json:"errors"`
	Data           interface{} `json:"data"`
}

type UserLoginDto struct {
	HttpStatusCode int         `json:"http_status_code"`
	Errors         []string    `json:"errors"`
	Login          string      `json:"login"`
	AccessToken    string      `json:"accesstoken"`
	RefreshToken   string      `json:"refreshtoken"`
	Data           interface{} `json:"data"`
}

type User struct {
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Created_at time.Time `json:"created_at"`
	Updated    time.Time `json:"updated"`
}
