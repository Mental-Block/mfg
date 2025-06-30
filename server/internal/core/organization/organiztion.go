package organization

import (
	"github.com/server/pkg/metadata"
	"github.com/server/pkg/utils"
)

type Organization struct {
	OrgId     string     
	Name      string        
	Title     string
	Avatar    string 
	Metadata  metadata.Metadata       
	State     string
}

func DBModelToOrg(model OrganizationModel) (*Organization, error) {
	data, err := metadata.UnMarshallMetadata(model.Metadata)
	
	if (err != nil) {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "unMarshalledMetadata")
	}

	return &Organization{
		OrgId:     model.OrgId,
		Name:      model.Name,
		Avatar:    model.Avatar.String,
		Metadata:  data,
		State: 	   model.State.String,
	}, nil
}

// return svc.AggregatedProject{
	// 	ID:             p.ProjectID.String,
	// 	Name:           p.ProjectName.String,
	// 	Title:          p.ProjectTitle.String,
	// 	State:          project.State(p.ProjectState.String),
	// 	MemberCount:    p.MemberCount.Int64,
	// 	UserIDs:        p.UserIDs,
	// 	CreatedAt:      p.CreatedAt.Time,
	// 	OrganizationID: p.OrganizationID.String,
	// }
// func DBModelToAggregatedUser()  {
// 	reut
	
	
	// return svc.AggregatedUser{
	// 	ID:          u.UserID.String,
	// 	Name:        u.UserName.String,
	// 	Title:       u.UserTitle.String,
	// 	Avatar:      u.UserAvatar.String,
	// 	Email:       u.UserEmail.String,
	// 	State:       user.State(u.UserState.String),
	// 	RoleNames:   u.RoleNames,
	// 	RoleTitles:  u.RoleTitles,
	// 	RoleIDs:     u.RoleIDs,
	// 	OrgID:       u.OrgID.String,
	// 	OrgJoinedAt: u.OrgJoinedAt.Time,
	// }


