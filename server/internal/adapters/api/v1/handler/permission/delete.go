package permission

import (
	"context"
	"net/http"
	"strconv"

	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/adapters/api/v1/util"
)

type DeletePermissionRequest struct{
	Id string `json:"id" path:"permission-id" example:"1" doc:"unique identifier"`
}

type DeletePermissionResponse struct {
	Body struct {
		Id int `json:"id" example:"1" doc:"unique identifier"`
	}
}

func (s *PermissionHandler) deletePermission(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"permission"},
		OperationID:   "delete-permissions",
		Summary:       "Delete permission",
		Description:   "deletes permission by id",
		Path:          "/{permission-id}/",
		Method:        http.MethodDelete,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, req *DeletePermissionRequest) (*DeletePermissionResponse, error) {
		
		if req.Id == "" {
			return nil, huma.Error400BadRequest("permission-id can't be empty")
		}

		permissionId, err := strconv.Atoi(req.Id)

		if err != nil {
			return nil, huma.Error400BadRequest("permission-id has to be a number")
		}

		id, err := s.permissionService.Remove(ctx, permissionId)

		if (err != nil) {
			return nil, util.HumaError(err)
		}

		resp := &DeletePermissionResponse{}
		resp.Body.Id = int(*id)

		return resp, nil
	})
}