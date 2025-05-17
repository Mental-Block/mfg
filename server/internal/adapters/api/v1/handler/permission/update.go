package permission

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/adapters/api/v1/dto"
	"github.com/server/internal/adapters/api/v1/util"
)

type UpdatePermissionRequest struct{
	Body dto.Role
}

type UpdatePermissionResponse struct {
	Body dto.Role
}

func (s *PermissionHandler) updatePermission(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"permission"},
		OperationID:   "update-permission",
		Summary:       "Update permission",
		Description:   "",
		Path:          "/",
		Method:        http.MethodPatch,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, req *UpdatePermissionRequest) (*UpdatePermissionResponse, error) {

		role, err := s.permissionService.Update(ctx, req.Body.Id, req.Body.Name)

		if (err != nil) {
			return nil, util.HumaError(err)
		}

		resp := &UpdatePermissionResponse{}
		resp.Body = dto.Role{
			Id: int(role.Id),
			Name: role.Name,
		}

		return resp, nil
	})
}