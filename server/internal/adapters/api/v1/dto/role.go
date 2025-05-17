package dto

type Role struct {
	Id       int    `json:"id" example:"1" doc:"unique identifier"`
	Name 	 string  `json:"name" example:"admin" doc:"role associated with user account"`
}
