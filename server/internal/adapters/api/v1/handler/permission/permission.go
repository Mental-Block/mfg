package permission

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/core/ports"
)

type PermissionHandler struct {
	permissionService ports.PermissionService
}

func NewPermissionHandler(service ports.PermissionService) *PermissionHandler {
	return &PermissionHandler{
		permissionService: service,
	}
}

func (s *PermissionHandler) Routes(parrentGrp *huma.Group) {
	permiGrp := huma.NewGroup(parrentGrp, "/permissions")

	s.createPermission(permiGrp)
	s.deletePermission(permiGrp)
	s.getPermission(permiGrp)
	s.getPermissions(permiGrp)
	s.updatePermission(permiGrp)
}
