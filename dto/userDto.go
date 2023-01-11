package dto

type UserDto struct {
	HttpStatusCode int         `json:"http_status_code"`
	Errors         []string    `json:"errors"`
	Data           interface{} `json:"data"`
}
