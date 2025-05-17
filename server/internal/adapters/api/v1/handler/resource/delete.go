package resource

import (
	"context"
	"net/http"
	"strconv"

	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/adapters/api/v1/util"
)

type DeleteResourceRequest struct{
	Id string `json:"id" path:"resource-id" example:"1" doc:"unique identifier"`
}

type DeleteResourceResponse struct {
	Body struct {
		Id int `json:"id" example:"1" doc:"unique identifier"`
	}
}

func (s *ResourceHandler) deleteResource(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"resource"},
		OperationID:   "delete-resources",
		Summary:       "Delete resource",
		Description:   "deletes resource by id",
		Path:          "/{resource-id}/",
		Method:        http.MethodDelete,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, req *DeleteResourceRequest) (*DeleteResourceResponse, error) {
		
		if req.Id == "" {
			return nil, huma.Error400BadRequest("resource-id can't be empty")
		}

		resourceId, err := strconv.Atoi(req.Id)

		if err != nil {
			return nil, huma.Error400BadRequest("resource-id has to be a number")
		}

		id, err := s.resourceService.Remove(ctx, resourceId)

		if (err != nil) {
			return nil, util.HumaError(err)
		}

		resp := &DeleteResourceResponse{}
		resp.Body.Id = int(*id)

		return resp, nil
	})
}