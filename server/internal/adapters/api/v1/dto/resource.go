package dto

type Resource struct {
	Id       int     `json:"id" example:"1" doc:"unique identifier"`
	Name 	 string  `json:"name" example:"document" doc:"resource name to be accessed. Such as file, line, job, etc..."`
}
