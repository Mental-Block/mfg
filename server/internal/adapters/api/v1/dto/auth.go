package dto

type Authorization struct {
	Id       int    `json:"id" example:"1" doc:"unique identifier"`
	Username string `json:"username" example:"bob" doc:"bob"`
	Roles 	 []string `json:"roles" example:"resource:action" doc:"roles and permissions (RBAC)"`
}