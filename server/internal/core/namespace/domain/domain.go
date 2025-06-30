package domain

import (
	"github.com/server/pkg/metadata"
	"github.com/server/pkg/utils"
)

type Namespace struct {
	Id 			string
	Name      	string
	Metadata  	metadata.Metadata
}

func (n Namespace)Transform()(*SanitizedNamespace, error) {

	if (n.Id == "") {
		n.Id = utils.NewUUID().String()
	}

	UUID, err := utils.ConvertStringToUUID(n.Id) 
	
	if err != nil {
		return &SanitizedNamespace{}, utils.NewErrorf(utils.ErrorCodeInvalidArgument, utils.ErrInvalidUUIDFormat, err)
	}

	name := NameSpaceName(n.Name)

	if err := name.IsValid(); err != nil {
		return &SanitizedNamespace{}, utils.NewErrorf(utils.ErrorCodeInvalidArgument, err.Error()) 
	}

	return &SanitizedNamespace {
		Id: UUID,
		Name: name,
		Metadata:  n.Metadata,
	}, nil
}

type SanitizedNamespace struct {
	Id 			utils.UUID
	Name      	NameSpaceName
	Metadata  	metadata.Metadata
}

func (n SanitizedNamespace) Transform()(*Namespace) {
	return &Namespace {
		Id: n.Id.String(),
		Name: n.Name.String(),
		Metadata: n.Metadata,
	}
}
