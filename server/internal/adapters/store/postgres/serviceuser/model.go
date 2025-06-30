package serviceuser

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/server/internal/core/serviceuser/domain"
	"github.com/server/pkg/metadata"
	"github.com/server/pkg/token"
	"github.com/server/pkg/utils"
)

type ServiceUserModel struct {
	Id       		utils.UUID         	`db:"serviceuser_id"`
	OrganizationId  utils.UUID          `db:"organization_id"`
	Title     		sql.NullString 		`db:"title"`
	State     		domain.State 		`db:"state"`
	Metadata  		[]byte         		`db:"metadata"`
	UpdateBy 		sql.NullString 	 	`db:"updated_by"`
	UpdatedDT 		sql.NullTime 		`db:"updated_dt"`
	CreatedBy 	  	string 			    `db:"created_by"`
	CreatedDt 	  	time.Time			`db:"created_dt"`
}

func (s ServiceUserModel) Transform() (domain.SanitizedServiceUser, error) {
	
	unMarshallMetadataData, err := metadata.UnMarshallMetadata(s.Metadata)
	
	if (err != nil) {
		return domain.SanitizedServiceUser{}, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "metadata.UnMarshallMetadata")
	}

	return domain.SanitizedServiceUser{
		Id: s.Id,
		OrgId: s.OrganizationId,
		Title: s.Title.String,
		State: s.State,
		Metadata: unMarshallMetadataData,
	}, nil
}

type ServiceUserCredentialModel struct {
	Id            string         	`db:"serviceuser_credential_id"`
	ServiceUserId string         	`db:"serviceuser_id"`
	Type          sql.NullString 	`db:"type"`
	SecretHash    sql.RawBytes 		`db:"secret_hash"`
	PublicKey     sql.RawBytes    	`db:"public_key"`
	Title         sql.NullString 	`db:"title"`
	Metadata      sql.RawBytes      `db:"metadata"`
	DeletedBy 	  sql.NullString 	`db:"deleted_by"`
	DeletedDT 	  sql.NullTime 		`db:"deleted_dt"`
	UpdateBy 	  sql.NullString 	`db:"updated_by"`
	UpdatedDT 	  sql.NullTime 		`db:"updated_dt"`
	CreatedBy 	  string 			`db:"created_by"`
	CreatedDt 	  time.Time			`db:"created_dt"`
}

func (s ServiceUserCredentialModel) Transform() (domain.SanitizedServiceUserCredential, error) {
	
	unMarshallMetadataData, err := metadata.UnMarshallMetadata(s.Metadata)
	
	if (err != nil) {
		return domain.SanitizedServiceUserCredential{}, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "metadata.UnMarshallMetadata")
	}

	var keySet token.JWkSet
	//if a secret hash is created, public key would be nil
	if len(string(s.SecretHash)) == 0 {
		set, err := token.JWKParse(s.PublicKey)

		if err != nil {
			return domain.SanitizedServiceUserCredential{}, fmt.Errorf("failed to parse public key: %w", err)
		}
		
		keySet = set
	}

	return domain.SanitizedServiceUserCredential{
		Id:            utils.UUID(s.Id),
		ServiceUserId: utils.UUID(s.ServiceUserId),
		Type:          domain.CredentialType(s.Type.String),
		SecretHash:    s.SecretHash,
		PublicKey:     keySet,
		Title:         s.Title.String,
		Metadata:      unMarshallMetadataData,
	}, nil
}
