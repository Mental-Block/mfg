package resource

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/adapters/api/v1/dto"
	"github.com/server/internal/adapters/api/v1/util"
)

type ResourcesResponse struct {
	Body struct {
		Resources []dto.Resource `json:"resources" doc:"array of resources"`
	}
}

func (s *ResourceHandler) getResources(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"resource"},
		OperationID:   "get-resources",
		Summary:       "Gets resources",
		Description:   "gets resources",
		Path:          "/",
		Method:        http.MethodGet,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, req *struct{}) (*ResourcesResponse, error) {

		data, err := s.resourceService.GetResources(ctx)

		if (err != nil) {
			return nil, util.HumaError(err)
		}

		resources := make([]dto.Resource, len(data))

		for i, v := range data {
			resources[i].Id = int(v.Id)
			resources[i].Name = string(v.Name)
		}

		resp := &ResourcesResponse{}
		resp.Body.Resources = resources

		return resp, nil
	})
}