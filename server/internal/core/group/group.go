package group

import (
	"github.com/server/pkg/metadata"
	"github.com/server/pkg/utils"
)

type Group struct {
	GroupId        	string  
	Name      		string         		
	Title     		string
	OrgId     		string         		
	State     		string
	Metadata  		metadata.Metadata
}


func DBModelToGroup(model GroupModel) (*Group, error) {
	data, err := metadata.UnMarshallMetadata(model.Metadata)
	
	if (err != nil) {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "unMarshalledMetadata")
	}

	return &Group{
		GroupId:	model.GroupId,  
		Name:       model.Name,     		
		Title:      model.Title.String,
		OrgId:    	model.OrgId,       		
		State:      model.State.String,
		Metadata:	data,
	}, nil
}

// func (from Group) transformToGroup() (group.Group, error) {
	
// 	return group.Group{
// 		ID:             from.ID,
// 		Name:           from.Name,
// 		Title:          from.Title.String,
// 		OrganizationID: from.OrgID,
// 		Metadata:       unmarshalledMetadata,
// 		State:          group.State(from.State.String),
// 		CreatedAt:      from.CreatedAt,
// 		UpdatedAt:      from.UpdatedAt,
// 	}, nil
// }