package user

import "github.com/server/internal/core/domain/entity"

type User struct {
	Username Username
}

type Username = string

type UserEntity struct {
	Id        entity.Id
	Username  Username
	CreatedBy entity.CreatedBy
	CreatedDT entity.CreatedDT
	UpdatedBy entity.UpdatedBy
	UpdatedDT entity.UpdatedDT
}
