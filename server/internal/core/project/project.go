package project

import (
	"github.com/server/pkg/metadata"
	"github.com/server/pkg/utils"
)


type Project struct {
	ProjectId	  string
	OrgId    	  string
	Name          string
	Title         string
	Metadata      metadata.Metadata
	State         string
}  

func  DBModelToProject(model ProjectModel) (*Project, error) {
	data, err := metadata.UnMarshallMetadata(model.Metadata)
	
	if (err != nil) {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "unMarshalledMetadata")
	}

	return &Project{
		ProjectId:    model.ProjectId,
		OrgId:        model.OrgId,
		Name:         model.Name,
		Title:        model.Title.String,
		Metadata:     data,
		State:        model.State.String,
	}, nil
}



