package permission

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/adapters/api/v1/dto"
	"github.com/server/internal/adapters/api/v1/util"
)


type CreatePermissionRequest struct{
	Body struct {
		Name string `json:"name" example:"create" doc:"permission"`
	}
}

type CreatePermissionResponse struct {
	Body dto.Permission
}

func (s *PermissionHandler) createPermission(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"permission"},
		OperationID:   "create-permission",
		Summary:       "Create permission",
		Description:   "Create a new permission",
		Path:          "/",
		Method:        http.MethodPost,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, req *CreatePermissionRequest) (*CreatePermissionResponse, error) {
		
		permission, err := s.permissionService.New(ctx, req.Body.Name)

		if (err != nil) {
			return nil, util.HumaError(err)
		}

		resp := &CreatePermissionResponse{
			Body: dto.Permission{
				Id: int(permission.Id),
				Name: permission.Name,
			},
		}

		return resp, nil
	})
}