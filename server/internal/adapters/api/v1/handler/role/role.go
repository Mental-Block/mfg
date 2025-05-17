package role

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/core/ports"
)

type RoleHandler struct {
	roleService ports.RoleService
}

func NewRoleHandler(service ports.RoleService) *RoleHandler {
	return &RoleHandler{
		roleService: service,
	}
}

func (s *RoleHandler) Routes(parrentGrp *huma.Group) {
	roleGrp := huma.NewGroup(parrentGrp, "/roles")

	s.createRole(roleGrp)
	s.deleteRole(roleGrp)
	s.updateRole(roleGrp)
	s.getRole(roleGrp)
	s.getRoles(roleGrp)
}
