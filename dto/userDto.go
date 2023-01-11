package dto

import "time"

type UserRegisterDto struct {
	HttpStatusCode int         `json:"http_status_code"`
	Register       string      `json:"register"`
	Errors         []string    `json:"errors"`
	Data           interface{} `json:"data"`
}

type UserLoginDto struct {
	HttpStatusCode int           `json:"http_status_code"`
	Errors         []string      `json:"errors"`
	Login          string        `json:"login"`
	AccessToken    string        `json:"accesstoken"`
	RefreshToken   string        `json:"refreshtoken"`
	Data           UserLoginData `json:"data"`
}
type UserLoginData struct {
	Email       string    `json:"email"`
	CretedAt    time.Time `json:"creted_at"`
	TokenStatus string    `json:"token_status"`
}

type UserLoginRedis struct {
	Email        string `json:"email"`
	AccessToken  string `json:"accesstoken"`
	RefreshToken string `json:"refreshtoken"`
}

type User struct {
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Created_at time.Time `json:"created_at"`
	Updated    time.Time `json:"updated"`
}
