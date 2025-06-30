package audit

import (
	"fmt"
	"time"

	"github.com/server/pkg/metadata"
	"github.com/server/pkg/utils"
)

type Audit struct {
	AuditId     string 		
	OrgId  		string 		
	Source 		string 		
	Action 		string 		
	Actor 	    metadata.Metadata 
	Target	    metadata.Metadata 
	Metadata 	metadata.Metadata 
}

type AuditFilter struct {
	OrgID  string
	Source string
	Action string
	StartTime time.Time
	EndTime   time.Time
	IgnoreSystem bool
}

type AuditLog struct {
	AuditId     string
	OrgId  		string
	Source 		string
	Action 		string
	Actor    	Actor
	Target   	Target
	Metadata 	metadata.Metadata
	CreatedAt 	time.Time
}

type Actor struct {
	Id   string
	Type string
	Name string
}

type Target struct {
	Id   string
	Type string
	Name string
}

var (
	ErrInvalidDetail = fmt.Errorf("invalid audit details")
	ErrInvalidID     = fmt.Errorf("group id is invalid")
)

func NewAction() {
	
}

func NewSource() {

}

func DBModelToAudit(model AuditModel) (*Audit, error) {
	
	data, err := metadata.UnMarshallMetadata(model.Metadata)
	
	if (err != nil) {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "unMarshalledMetadata")
	}
	
	target, err := metadata.UnMarshallMetadata(model.Metadata)
	
	if (err != nil) {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "unMarshalledMetadata")
	}

	actor, err := metadata.UnMarshallMetadata(model.Metadata)
	
	if (err != nil) {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "unMarshalledMetadata")
	}

	return &Audit{
		AuditId: model.AuditId,
		OrgId: model.OrgId,
		Source: model.Source,
		Action: model.Action,
		Actor: actor,
		Target: target,
		Metadata: data,
	}, nil
}
