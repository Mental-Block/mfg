package domain

import "errors"

var (
	ErrResourcePermissionNotFound  =  errors.New("resource permission does not exist") 
 	ErrRoleResourcePermissionNotFound = errors.New("role resource permission does not exist")
 	ErrUserResourcePermissionNotFound = errors.New("user resource permission does not exist")
)

type ResourcePermission struct {
	Id Id
	Permission Permission
	Resource Resource
} 