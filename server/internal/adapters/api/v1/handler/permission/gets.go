package permission

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/adapters/api/v1/dto"
	"github.com/server/internal/adapters/api/v1/util"
)

type PermissionsResponse struct {
	Body struct {
		Permissions []dto.Permission `json:"permissions"  doc:"array of pemissions"`
	}
}

func (s *PermissionHandler) getPermissions(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"permission"},
		OperationID:   "get-permissions",
		Summary:       "Gets permissions",
		Description:   "gets permissions",
		Path:          "/",
		Method:        http.MethodGet,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, req *struct{}) (*PermissionsResponse, error) {

		data, err := s.permissionService.GetPermissions(ctx)

		if (err != nil) {
			return nil, util.HumaError(err)
		}

		permissions := make([]dto.Permission, len(data))

		for i, v := range data {
			permissions[i].Id = int(v.Id)
			permissions[i].Name = string(v.Name)
		}

		resp := &PermissionsResponse{}
		resp.Body.Permissions = permissions

		return resp, nil
	})
}