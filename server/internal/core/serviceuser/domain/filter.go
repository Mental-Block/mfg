package domain

import "github.com/server/internal/core/role/domain"

type CredentialFilter struct {
	Id string
	OrgId    string
	ServiceUserId string
	IsKey    bool
	IsSecret bool
	IsToken  bool
}

type ServiceUserFilter struct {
	OrgId    string
	Ids 	[]string
	State    domain.State
}
