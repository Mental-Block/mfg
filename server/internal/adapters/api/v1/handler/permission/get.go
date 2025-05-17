package permission

import (
	"context"
	"net/http"
	"strconv"

	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/adapters/api/v1/dto"
	"github.com/server/internal/adapters/api/v1/util"
)

type PermissionRequest struct{
	Id string `json:"id" path:"permission-id" example:"1" doc:"unique identifier"`
}

type PermissionResponse struct {
	Body *dto.Permission
}

func (s *PermissionHandler) getPermission(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"permission"},
		OperationID:   "get-permission",
		Summary:       "Get permission",
		Description:   "Gets permission by id",
		Path:          "/{permission-id}/",
		Method:        http.MethodGet,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, req *PermissionRequest) (*PermissionResponse, error) {
		
		if req.Id == "" {
			return nil, huma.Error400BadRequest("permission-id can't be empty")
		}

		id, err := strconv.Atoi(req.Id)

		if err != nil {
			return nil, huma.Error400BadRequest("permission-id has to be a number")
		}

		permission, err := s.permissionService.GetPermission(ctx, id)

		if (err != nil) {
			return nil, util.HumaError(err)
		}

		resp := &PermissionResponse{}
	 	resp.Body = &dto.Permission{
			Id: int(permission.Id),
			Name: permission.Name,
		}

		return resp, nil
	})
}
