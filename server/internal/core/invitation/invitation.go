package inventation

import (
	"time"

	"github.com/server/pkg/metadata"
	"github.com/server/pkg/utils"
)

type Invitation struct {
	Id          string
	UserEmailId string
	OrgId       string
	GroupIds     []string
	RoleIds     []string
	Metadata    metadata.Metadata
	CreatedDT   time.Time
	ExpiresDT   time.Time
}

func DBModelToInvitation(model InvitationModel) (*Invitation, error) {
	data, err := metadata.UnMarshallMetadata(model.Metadata)
	
	if (err != nil) {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "unMarshalledMetadata")
	}

	groupIDs := metadata.MapToArray(data, "group_ids")

	roleIDs := metadata.MapToArray(data, "role_ids")
	
	return &Invitation{
		Id:          model.InvitationId,
		UserEmailId: model.UserId,
		OrgId:       model.OrgId,
		GroupIds:    groupIDs,
		RoleIds:     roleIDs,
		Metadata:    data,
		CreatedDT:   model.CreatedDT,
		ExpiresDT:   model.ExpiresDT,
	}, nil
}


