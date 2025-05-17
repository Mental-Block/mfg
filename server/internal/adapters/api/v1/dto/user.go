package dto

type User struct {
	Id       int    `json:"id" example:"1" doc:"unique identifier"`
	Username string `json:"username" example:"bob" doc:"bob"`
}
