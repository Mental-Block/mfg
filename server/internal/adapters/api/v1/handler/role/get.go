package role

import (
	"context"
	"net/http"
	"strconv"

	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/adapters/api/v1/dto"
	"github.com/server/internal/adapters/api/v1/util"
)

type RoleRequest struct{
	Id string `json:"id" path:"role-id" example:"1" doc:"unique identifier"`
}

type RoleResponse struct {
	Body *dto.Role
}

func (s *RoleHandler) getRole(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"role"},
		OperationID:   "get-role",
		Summary:       "Get role",
		Description:   "Gets role by id",
		Path:          "/{role-id}/",
		Method:        http.MethodGet,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, req *RoleRequest) (*RoleResponse, error) {
		
		if req.Id == "" {
			return nil, huma.Error400BadRequest("role-id can't be empty")
		}

		roleId, err := strconv.Atoi(req.Id)

		if err != nil {
			return nil, huma.Error400BadRequest("role-id has to be a number")
		}

		role, err := s.roleService.GetRole(ctx, roleId)

		if (err != nil) {
			return nil, util.HumaError(err)
		}

		resp := &RoleResponse{}
	 	resp.Body = &dto.Role{
			Id: int(role.Id),
			Name: role.Name,
		}

		return resp, nil
	})
}
