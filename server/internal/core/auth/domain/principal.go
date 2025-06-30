package domain

import (
	namespace "github.com/server/internal/core/namespace/domain"
	serviceuser "github.com/server/internal/core/serviceuser/domain"
	user "github.com/server/internal/core/user/domain"
)

type Principal struct {
	// ID is the unique identifier of principal
	Id string
	// Type is the namespace of principal
	// E.g. app/user, app/serviceuser
	Type 		namespace.NameSpaceName

	Email 		Email

	User        *user.User
	ServiceUser *serviceuser.ServiceUser
}