package role

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/adapters/api/v1/dto"
	"github.com/server/internal/adapters/api/v1/util"
)


type CreateRoleRequest struct{
	Body struct {
		Name string `json:"name" example:"admin" doc:"role"`
	}
}

type CreateRoleResponse struct {
	Body dto.Role
}

func (s *RoleHandler) createRole(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"role"},
		OperationID:   "create-role",
		Summary:       "Create role",
		Description:   "Create a new role",
		Path:          "/",
		Method:        http.MethodPost,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, req *CreateRoleRequest) (*CreateRoleResponse, error) {
		
		role, err := s.roleService.New(ctx, req.Body.Name)

		if (err != nil) {
			return nil, util.HumaError(err)
		}

		resp := &CreateRoleResponse{
			Body: dto.Role{
				Id: int(role.Id),
				Name: role.Name,
			},
		}

		return resp, nil
	})
}
