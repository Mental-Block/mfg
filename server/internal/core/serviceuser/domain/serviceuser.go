package domain

import (
	"github.com/server/pkg/metadata"
	"github.com/server/pkg/utils"
)

type ServiceUser struct {
	Id				string
	OrgId     		string       
	Title     		string    
	State     		string 
	Metadata  		metadata.Metadata 
	CreatedByUser 				string `json:"-"`
}

func (s ServiceUser) Transform()(*SanitizedServiceUser, error)  {
	uuid, err := utils.ConvertStringToUUID(s.Id) 
	
	if err != nil {
		return nil, utils.NewErrorf(utils.ErrorCodeInvalidArgument, utils.ErrInvalidUUIDFormat, err)
	}

	orguuid, err := utils.ConvertStringToUUID(s.OrgId) 
	
	if err != nil {
		return nil, utils.NewErrorf(utils.ErrorCodeInvalidArgument, utils.ErrInvalidUUIDFormat, err)
	}

	return &SanitizedServiceUser{
		Id: uuid,
		OrgId: orguuid,
		Title: s.Title,
		State: State(s.State),
		Metadata: s.Metadata,
	}, nil
}

type SanitizedServiceUser struct {
	Id				utils.UUID
	OrgId     		utils.UUID       
	Title     		string    
	State     		State 
	Metadata  		metadata.Metadata 

	CreatedByUser 				string `json:"-"`
}


func (s SanitizedServiceUser) Transform()(*ServiceUser) {
	return &ServiceUser{
		Id: s.Id.String(),
		OrgId: s.OrgId.String(),
		Title: s.Title,
		State: s.State.String(),
		Metadata: s.Metadata,
	}
}