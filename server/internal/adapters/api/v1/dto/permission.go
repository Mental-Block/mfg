package dto

type Permission struct {
	Id       int    `json:"id" example:"1" doc:"unique identifier"`
	Name 	 string  `json:"name" example:"admin" doc:"permission associated with resource account"`
}
