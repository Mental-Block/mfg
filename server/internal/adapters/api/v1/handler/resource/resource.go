package resource

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/core/ports"
)

type ResourceHandler struct {
	resourceService ports.ResourceService
}

func NewResourceHandler(service ports.ResourceService) *ResourceHandler {
	return &ResourceHandler{
		resourceService: service,
	}
}

func (s *ResourceHandler) Routes(parrentGrp *huma.Group) {
	rescGrp := huma.NewGroup(parrentGrp, "/resources")

	s.createResource(rescGrp)
	s.deleteResource(rescGrp)
	s.getResource(rescGrp)
	s.getResources(rescGrp)
	s.updateResource(rescGrp)
}
