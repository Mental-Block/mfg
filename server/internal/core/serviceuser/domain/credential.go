package domain

import (
	"encoding/json"

	"github.com/server/pkg/metadata"
	"github.com/server/pkg/token"
	"github.com/server/pkg/utils"
)

type ServiceUserCredential struct {	
	// Id is the unique identifier of the credential.
	// This is also used as kid in JWT, the spec doesn't
	// state how the kid should be generated as anyway this token
	// is owned by mfg, and we are in control of key generation
	// any arbitrary string can be used as kid as long as its unique
	Id     						string         
	ServiceUserId 				string         
	Type          				string
	// SecretHash used for basic auth
	SecretHash    				[]byte
	// PublicKey used for JWT verification using RSA
	PublicKey     				[]byte 

	// never save PrivateKey to database always send to caller
	PrivateKey					[]byte
	Title         				string 
	Metadata      				metadata.Metadata 
}

func (s ServiceUserCredential) Transform() (*SanitizedServiceUserCredential, error)  {

	UUID, err := utils.ConvertStringToUUID(s.Id) 
	
	if err != nil {
		return nil, utils.NewErrorf(utils.ErrorCodeInvalidArgument, utils.ErrInvalidUUIDFormat, err)
	}

	serviceUserUUID, err := utils.ConvertStringToUUID(s.ServiceUserId) 
	
	if err != nil {
		return nil, utils.NewErrorf(utils.ErrorCodeInvalidArgument, utils.ErrInvalidUUIDFormat, err)
	}

	switch s.Type {
		case OpaqueTokenCredentialType.String():
			break;
		case JWTCredentialType.String():
			break;
		case ClientSecretCredentialType.String():
			break;
		case "":
			break; //valid for secrets
		default:
			return nil, utils.NewErrorf(utils.ErrorCodeInvalidArgument, "invalid credential type")
	}

	// try to convert to key set. if its empty let it pass
	var keySet token.JWkSet = nil 
	if (len(s.PublicKey) != 0 || s.PublicKey != nil) {
		set, err := token.JWKParse(s.PublicKey)

		if err != nil {
			return nil, utils.NewErrorf(utils.ErrorCodeInvalidArgument, "invalid key provided")
		} else {
			keySet = set
		}
	}

	// if we let the public pass and its empty we must have a secret hash
	if (len(s.PublicKey) != 0 || s.PublicKey != nil || len(s.SecretHash) != 0) {
		return &SanitizedServiceUserCredential{
			Id: UUID,
			ServiceUserId: serviceUserUUID, 				         
			Type: CredentialType(s.Type),          				
			SecretHash: []byte(s.SecretHash),    				
			PublicKey:  keySet,      				      
			Title: s.Title,         				 
			Metadata: s.Metadata,      		
		}, nil
	} else {
		return nil, utils.NewErrorf(utils.ErrorCodeInvalidArgument, "no credential was presented. token, secret, key pair")
	}
}	

type SanitizedServiceUserCredential struct {
	Id     						utils.UUID         
	ServiceUserId 				utils.UUID         
	Type          				CredentialType
	SecretHash    				[]byte
	PublicKey     				token.JWkSet      
	Title         				string 
	Metadata      				metadata.Metadata 
}

func (s SanitizedServiceUserCredential) Transform() (*ServiceUserCredential, error) {
	
	publicKey, err := json.Marshal(s.PublicKey.Keys())
	
	if (err != nil) {
		return nil, err
	}

	return &ServiceUserCredential{
		Id: s.Id.String(),
		ServiceUserId: s.ServiceUserId.String(), 				         
		Type: s.Type.String(),          				
		SecretHash: s.SecretHash,    				
		PublicKey: publicKey,      				      
		Title: s.Title,         				 
		Metadata: s.Metadata, 					 
	}, nil
}

type CredentialType string

func (c CredentialType) String() string {
	return string(c)
}

const (
	ClientSecretCredentialType CredentialType = "client_credential"
	JWTCredentialType          CredentialType = "jwt_bearer"
	OpaqueTokenCredentialType  CredentialType = "opaque_token"
)
