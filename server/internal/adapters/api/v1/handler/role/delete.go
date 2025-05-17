package role

import (
	"context"
	"net/http"
	"strconv"

	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/adapters/api/v1/util"
)

type DeleteRoleRequest struct{
	Id string `json:"id" path:"role-id" example:"1" doc:"unique identifier"`
}

type DeleteRoleResponse struct {
	Body struct {
		Id int `json:"id" example:"1" doc:"unique identifier"`
	}
}

func (s *RoleHandler) deleteRole(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"role"},
		OperationID:   "delete-roles",
		Summary:       "Delete role",
		Description:   "deletes role by id",
		Path:          "/{role-id}/",
		Method:        http.MethodDelete,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, req *DeleteRoleRequest) (*DeleteRoleResponse, error) {
		
		if req.Id == "" {
			return nil, huma.Error400BadRequest("role-id can't be empty")
		}

		roleId, err := strconv.Atoi(req.Id)

		if err != nil {
			return nil, huma.Error400BadRequest("role-id has to be a number")
		}

		id, err := s.roleService.Remove(ctx, roleId)

		if (err != nil) {
			return nil, util.HumaError(err)
		}

		resp := &DeleteRoleResponse{}
		resp.Body.Id = int(*id)

		return resp, nil
	})
}