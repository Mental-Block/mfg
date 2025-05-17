package resource

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/adapters/api/v1/dto"
	"github.com/server/internal/adapters/api/v1/util"
)


type UpdateResourceRequest struct{
	Body dto.Resource 
}

type UpdateResourceResponse struct {
	Body dto.Resource
}

func (s *ResourceHandler) updateResource(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"resource"},
		OperationID:   "update-resource",
		Summary:       "Update resource",
		Description:   "",
		Path:          "/",
		Method:        http.MethodPatch,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, req *UpdateResourceRequest) (*UpdateResourceResponse, error) {

		role, err := s.resourceService.Update(ctx, req.Body.Id, req.Body.Name)

		if (err != nil) {
			return nil, util.HumaError(err)
		}

		resp := &UpdateResourceResponse{}
		resp.Body = dto.Resource{
			Id: int(role.Id),
			Name: role.Name,
		}

		return resp, nil
	})
}