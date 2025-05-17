package resource

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/adapters/api/v1/dto"
	"github.com/server/internal/adapters/api/v1/util"
)


type CreateResourceRequest struct{
	Body struct {
		Name string `json:"name" example:"document" doc:"resource"`
	}
}

type CreateResourceResponse struct {
	Body dto.Resource
}

func (s *ResourceHandler) createResource(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"resource"},
		OperationID:   "create-resource",
		Summary:       "Create resource",
		Description:   "Create a new resource",
		Path:          "/",
		Method:        http.MethodPost,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, req *CreateResourceRequest) (*CreateResourceResponse, error) {
		
		resource, err := s.resourceService.New(ctx, req.Body.Name)

		if (err != nil) {
			return nil, util.HumaError(err)
		}

		resp := &CreateResourceResponse{
			Body: dto.Resource{
				Id: int(resource.Id),
				Name: resource.Name,
			},
		}

		return resp, nil
	})
}