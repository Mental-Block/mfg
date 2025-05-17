package role

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/adapters/api/v1/dto"
	"github.com/server/internal/adapters/api/v1/util"
)

type UpdateRoleRequest struct{
	Body dto.Role
}

type UpdateRoleResponse struct {
	Body dto.Role
}

func (s *RoleHandler) updateRole(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"role"},
		OperationID:   "update-role",
		Summary:       "Update role",
		Description:   "",
		Path:          "/",
		Method:        http.MethodPatch,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, req *UpdateRoleRequest) (*UpdateRoleResponse, error) {

		role, err := s.roleService.Update(ctx, req.Body.Id, req.Body.Name)

		if (err != nil) {
			return nil, util.HumaError(err)
		}

		resp := &UpdateRoleResponse{}
		resp.Body = dto.Role{
			Id: int(role.Id),
			Name: role.Name,
		}

		return resp, nil
	})
}