package role

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/adapters/api/v1/dto"
	"github.com/server/internal/adapters/api/v1/util"
)


type RolesResponse struct {
	Body struct {
		Roles []dto.Role
	}
}

func (s *RoleHandler) getRoles(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"role"},
		OperationID:   "get-roles",
		Summary:       "Gets roles",
		Description:   "gets roles",
		Path:          "/",
		Method:        http.MethodGet,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, req *struct{}) (*RolesResponse, error) {

		data, err := s.roleService.GetRoles(ctx)

		if (err != nil) {
			return nil, util.HumaError(err)
		}

		roles := make([]dto.Role, len(data))

		for i, v := range data {
			roles[i].Id = int(v.Id)
			roles[i].Name = string(v.Name)
		}

		resp := &RolesResponse{}
		resp.Body.Roles = roles

		return resp, nil
	})
}
