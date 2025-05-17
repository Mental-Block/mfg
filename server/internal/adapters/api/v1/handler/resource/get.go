package resource

import (
	"context"
	"net/http"
	"strconv"

	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/adapters/api/v1/dto"
	"github.com/server/internal/adapters/api/v1/util"
)

type ResourceRequest struct{
	Id string `json:"id" path:"resource-id" example:"1" doc:"unique identifier"`
}

type ResourceResponse struct {
	Body *dto.Resource
}

func (s *ResourceHandler) getResource(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"resource"},
		OperationID:   "get-resource",
		Summary:       "Get resource",
		Description:   "Gets resource by id",
		Path:          "/{resource-id}/",
		Method:        http.MethodGet,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, req *ResourceRequest) (*ResourceResponse, error) {
		
		if req.Id == "" {
			return nil, huma.Error400BadRequest("resource-id can't be empty")
		}

		id, err := strconv.Atoi(req.Id)

		if err != nil {
			return nil, huma.Error400BadRequest("resource-id has to be a number")
		}

		resource, err := s.resourceService.GetResource(ctx, id)

		if (err != nil) {
			return nil, util.HumaError(err)
		}

		resp := &ResourceResponse{}
	 	resp.Body = &dto.Resource{
			Id: int(resource.Id),
			Name: resource.Name,
		}

		return resp, nil
	})
}
