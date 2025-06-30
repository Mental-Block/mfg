package serviceuser

import (
	"context"
	"crypto/sha3"
	"crypto/subtle"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/server/internal/adapters/bootstrap/schema"

	permission "github.com/server/internal/core/permission/domain"
	relation "github.com/server/internal/core/relation/domain"
	"github.com/server/internal/core/serviceuser/domain"

	"github.com/server/pkg/crypt"
	"github.com/server/pkg/token"
	"github.com/server/pkg/utils"
)

/*
 High level overview of ServiceUserService should not be directly imported.
 Copy interface and use dependancy injection over direct import.
*/
type IServiceUserService interface {
	List(ctx context.Context, filter domain.ServiceUserFilter) ([]domain.ServiceUser, error)
	ListByOrg(ctx context.Context, orgId string) ([]domain.ServiceUser, error)
	ListCreds(ctx context.Context, filter domain.CredentialFilter) ([]domain.ServiceUserCredential, error)
	New(ctx context.Context, serviceUser domain.ServiceUser) (*domain.ServiceUser, error)
	NewCredKeyPair(ctx context.Context, credential domain.ServiceUserCredential) (*domain.ServiceUserCredential, error) 
	NewCredSecret(ctx context.Context, credential domain.ServiceUserCredential) (*domain.Secret, error)
	Delete(ctx context.Context, id string) error
	DeleteCred(ctx context.Context, credId string) error
	GetByJWT(ctx context.Context, tkn []byte) (*domain.ServiceUser, error)
	GetBySecret(ctx context.Context, credId string, secret string) (*domain.ServiceUser, error)
	GetByIds(ctx context.Context, ids []string) ([]domain.ServiceUser, error)	
	GetCred(ctx context.Context, credId string) (*domain.ServiceUserCredential, error)
 	Get(ctx context.Context, id string) (*domain.ServiceUser, error) 
	IsSuper(ctx context.Context, id string, permissionName string) (bool, error)	
	FilterSupers(ctx context.Context, ids []string) ([]string, error) 
	NewSuper(ctx context.Context, id string, rel string) error
	UnSuper(ctx context.Context, id string) error
}

type ServiceUserService struct {
	cfg 							crypt.Params
	serviceUserStore				IServiceUserStore
	serviceUserCredentialStore 		IServiceUserCredentialStore     
	relationService 				IRelationService
	Now 							func() time.Time
}

func NewServiceUserService(
	cfg				crypt.Params,
	store			IServiceUserStore,
	credentialStore IServiceUserCredentialStore,     
	relationService IRelationService,
) *ServiceUserService {
	return &ServiceUserService{
		cfg: cfg,
		serviceUserStore: store,
		serviceUserCredentialStore: credentialStore,
		relationService: relationService,
		Now: func () time.Time {
			return time.Now().UTC();
		},
	}
}

func (s ServiceUserService) List(ctx context.Context, filter domain.ServiceUserFilter) ([]domain.ServiceUser, error) {

	serviceUserModels, err := s.serviceUserStore.Selects(ctx, filter)

	if err != nil {
		return nil, err
	}

	serviceUsers := make([]domain.ServiceUser, len(serviceUserModels))

	for i := range serviceUsers {
		 santizedUser, err := serviceUsers[i].Transform()
		 
		 if err != nil {
			return nil, err
		 }
		 
		serviceUsers[i] = *santizedUser.Transform()
	}

	return serviceUsers, nil
}

func (s ServiceUserService) New(ctx context.Context, serviceUser domain.ServiceUser) (*domain.ServiceUser, error) {
	
	SanitizedServiceUser, err := serviceUser.Transform()

	if err != nil {
		return nil, err
	}

	serviceUserModel, err := s.serviceUserStore.Insert(ctx, *SanitizedServiceUser)
	
	if err != nil {
		return nil, err
	}

	backToSantized, err := serviceUserModel.Transform()

	if err != nil {
		return nil, err
	}

	service := backToSantized.Transform()

	// attach service user to organization
	_, err = s.relationService.New(ctx, relation.Relation{
		Object: relation.Object{
			Id:        service.OrgId,
			Namespace: schema.OrganizationNamespace.String(),
		},
		Subject: relation.Subject{
			Id:        service.Id,
			Namespace: schema.ServiceUserPrincipal.String(),
		},
		RelationName: schema.MemberRelationName.String(),
	})

	if err != nil {
		return nil, err
	}

	_, err = s.relationService.New(ctx, relation.Relation{
		Object: relation.Object{
			Id:        service.Id,
			Namespace: schema.ServiceUserPrincipal.String(),
		},
		Subject: relation.Subject{
			Id:        service.OrgId,
			Namespace: schema.OrganizationNamespace.String(),
		},
		RelationName: schema.OrganizationRelationName.String(),
	})

	if err != nil {
		return nil, err
	}

	if len(service.CreatedByUser) > 0 {
		// TODO: write authz tests that checks if the user who created the service user
		// has the permission to interact with the service user
		// attach user to service user who created it
		_, err = s.relationService.New(ctx, relation.Relation{
			Object: relation.Object{
				Id:        service.Id,
				Namespace: schema.ServiceUserPrincipal.String(),
			},
			Subject: relation.Subject{
				Id:        service.CreatedByUser,
				Namespace: schema.UserPrincipal.String(),
			},
			RelationName: schema.UserRelationName.String(),
		})
		
		if err != nil {
			return nil, err
		}
	}

	return service, nil
}

func (s ServiceUserService) Get(ctx context.Context, id string) (*domain.ServiceUser, error) {

	uuid, err := utils.ConvertStringToUUID(id)

	if err != nil {
		return nil, err
	}

	serviceUserModel, err := s.serviceUserStore.Select(ctx, uuid)

	santizedServiceUser, err := serviceUserModel.Transform()

	if err != nil {
		return nil, err
	}

	serviceUser := santizedServiceUser.Transform()

	return serviceUser, nil
}

func (s ServiceUserService) GetByIds(ctx context.Context, ids []string) ([]domain.ServiceUser, error) {

	uids := make([]utils.UUID ,len(ids))

	for i, v := range ids {
		
		UUID, err := utils.ConvertStringToUUID(v) 
		
		if err != nil {
			return nil, utils.NewErrorf(utils.ErrorCodeInvalidArgument, utils.ErrInvalidUUIDFormat, err)
		}

		uids[i] = UUID
	}

	serviceUserModels, err := s.serviceUserStore.SelectByIds(ctx, uids)

	if err != nil {
		return nil, err
	}

	serviceUsers := make([]domain.ServiceUser, len(serviceUserModels))

	for i := range serviceUsers {
		 santizedUser, err := serviceUsers[i].Transform()
		 
		 if err != nil {
			return nil, err
		 }
		 
		serviceUsers[i] = *santizedUser.Transform()
	}

	return serviceUsers, nil
}

func (s ServiceUserService) ListByOrg(ctx context.Context, orgId string) ([]domain.ServiceUser, error) {
	
	userIds, err := s.relationService.LookupSubjects(ctx, relation.Relation{
		Object: relation.Object{
			Id:        orgId,
			Namespace: schema.OrganizationNamespace.String(),
		},
		Subject: relation.Subject{
			Namespace: schema.ServiceUserPrincipal.String(),
		},
		RelationName: schema.MembershipPermission.String(),
	})
	
	if err != nil {
		return nil, err
	}

	if len(userIds) == 0 {
		return nil, nil
	}

	servuceUsers, err := s.GetByIds(ctx, userIds)

	if err != nil {
		return nil, err
	}
	
	return servuceUsers, nil
}

// GetBySecret matches the secret with the secret hash stored in the database of the service user
// and if the secret matches, returns the service user
func (s ServiceUserService) GetBySecret(ctx context.Context, credId string, secret string) (*domain.ServiceUser, error) {
	if len(secret) <= 0 {
		return nil, ErrInvalidCred
	}

	uuid, err := utils.ConvertStringToUUID(credId)

	if err != nil {
		return nil, err
	}

	credModel, err := s.serviceUserCredentialStore.Select(ctx, uuid)
	
	if err != nil {
		return nil, err
	}

	cred, err := credModel.Transform()

	if err != nil {
		return nil, err
	}

	if len(cred.Type) == 0 || cred.Type == domain.ClientSecretCredentialType {
		if ok, err :=  crypt.ComparePasswordAndHash(secret, string(cred.SecretHash)); !ok || err != nil {
			return nil, ErrInvalidCred
		}
	}

	// decode the hex encoded secret
	if cred.Type == domain.OpaqueTokenCredentialType {
	
		decodedReqSecret := make([]byte, 32)
		
		if _, err := hex.Decode(decodedReqSecret, []byte(secret)); err != nil {
			return nil, err
		}

		reqDigest := sha3.Sum256(decodedReqSecret)

		decodedCredSecret := make([]byte, 32)

		if _, err := hex.Decode(decodedCredSecret, cred.SecretHash); err != nil {
			return nil, err
		}

		if subtle.ConstantTimeCompare(reqDigest[:], decodedCredSecret) == 0 {
			return nil, ErrInvalidCred
		}
	}

	serviceUserModel, err := s.serviceUserStore.Select(ctx, utils.UUID(credId))
	
	if err != nil {
		return nil, err
	}

	sanitizedServiceUser, err := serviceUserModel.Transform()

	if (err != nil) {
		return nil, err
	}

	serviceUser := sanitizedServiceUser.Transform()

	return serviceUser, nil
}
 
// GetByJWT returns the service user by verifying the token in authservice
func (s ServiceUserService) GetByJWT(ctx context.Context, tkn []byte) (*domain.ServiceUser, error) {
	
	insecureTkn, err := token.ParseInsecure(tkn)

	if err != nil {
		return nil, fmt.Errorf("invalid serviceuser token: %w", err)
	}

	tknKID, ok := insecureTkn.JwtID()

	if !ok {
		return nil, fmt.Errorf("invalid key id from token")
	}

	credModel, err := s.serviceUserCredentialStore.Select(ctx, utils.UUID(tknKID))

	if err != nil {
		return nil, fmt.Errorf("credential invalid of kid %s: %w", tknKID, err)
	}
	
	santizedCred, err := credModel.Transform()

	if (err != nil) {
		return nil, err
	}

	_, err = token.Parse(tkn, token.WithKeySet(santizedCred.PublicKey))

	if err != nil {
		return nil, fmt.Errorf("invalid serviceuser token: %w", err)
	}

	serviceUserModel, err := s.serviceUserStore.Select(ctx, santizedCred.ServiceUserId)

	if err != nil {
		return nil, err
	}

	santizedService, err := serviceUserModel.Transform()

	if err != nil {
		return nil, err
	}

	serviceUser := santizedService.Transform()

	return serviceUser, nil 
}

func (s ServiceUserService) Delete(ctx context.Context, id string) error {

	UUID, err := utils.ConvertStringToUUID(id) 
	
	if err != nil {
		return utils.NewErrorf(utils.ErrorCodeInvalidArgument, utils.ErrInvalidUUIDFormat, err)
	}

	if err := s.relationService.Remove(ctx, relation.Relation{
		Subject: relation.Subject{
			Id:        UUID.String(),
			Namespace: schema.ServiceUserPrincipal.String(),
		},
	}); 
	
	err != nil {
		return err
	}

	err = s.serviceUserStore.DeleteUserAndCredentials(ctx, UUID)
	
	if err != nil {
		return err
	}

	return nil
}

func (s ServiceUserService) NewCredKeyPair(ctx context.Context, credential domain.ServiceUserCredential) (*domain.ServiceUserCredential, error) {
	
	credential.Id = utils.NewUUID().String()
	credential.Type = domain.JWTCredentialType.String()

	// generate public/private key pair
	newJWK, err := token.CreateJWKWithKID(credential.Id)

	if err != nil {
		return nil, fmt.Errorf("failed to create key pair: %w", err)
	}

	jwkPEM, err := token.JWKPem(newJWK)

	if err != nil {
		return nil, fmt.Errorf("failed to convert jwk to pem: %w", err)
	}

	pubKey, err := newJWK.PublicKey()

	if err != nil {
		return nil, err
	}

	publicKeySet := token.JWKNewSet()

	if err := publicKeySet.AddKey(pubKey); err != nil {
		return nil, err
	}

	// slap it back to bytes so we can parse it using the credential transform func and validate the reset of the input 
	publicKeyBytes, err := json.Marshal(publicKeySet)

	if err != nil {
		return nil, err
	}

	credential.PublicKey = publicKeyBytes

	sanitizedCredential, err := credential.Transform()

	if err != nil {
		return nil, err
	}

	credModel, err := s.serviceUserCredentialStore.Insert(ctx, *sanitizedCredential)

	if err != nil {
		return nil, err
	}

	sanitizedCred, err := credModel.Transform()

	if err != nil {
		return nil, err
	}

	cred, err := sanitizedCred.Transform()

	if (err != nil) {
		return nil, err
	}
	
	cred.PrivateKey = jwkPEM

	return cred, nil
}

// the most insecure way to create a credential prefer tokens or keypair methods when posible or secret
func (s ServiceUserService) NewCredSecret(ctx context.Context, credential domain.ServiceUserCredential) (*domain.Secret, error) {

	credential.Id = utils.NewUUID().String()
	credential.Type = domain.ClientSecretCredentialType.String()

	rdmLength, err := crypt.GererateRandomInt(10, 40)

	if (err != nil) {
		return nil, utils.NewErrorf(utils.ErrorCodeUnknown, "failed to generate random number")
	}

	passphrase := crypt.GenerateRandomStringFromLetters(int(rdmLength), crypt.CharList)

	secret, err := crypt.CreateHash(s.cfg, passphrase)
	
	if err != nil {
		return nil, err
	} 
	
	credential.SecretHash = []byte(secret)

	sanitizedCredential, err := credential.Transform()

	if err != nil {
		return nil, err
	}
	
	createdCred, err := s.serviceUserCredentialStore.Insert(ctx, *sanitizedCredential)
	
	if err != nil {
		return nil, err
	}
	
	// we can skip transform here as we are sending back a secert and not serviceusercred

	return &domain.Secret{
		Id:        createdCred.Id,
		Title:     createdCred.Title.String,
		Value:     secret,
		CreatedAt: s.Now(),
	}, nil
}

func (s ServiceUserService) NewCredToken(ctx context.Context, credential domain.ServiceUserCredential) (*domain.Token, error) {
	
	credential.Id = utils.NewUUID().String()
	rdmbytes, err := crypt.GenerateRandomBytes(32)

	if err != nil {
		return nil, err
	}

	hash := sha3.Sum256(rdmbytes)

	credToken := hex.EncodeToString(hash[:])

	credential.SecretHash = []byte(credToken)
	
	credential.Type = domain.OpaqueTokenCredentialType.String()

	sanitizedCred, err := credential.Transform()

	if err != nil {
		return nil, err
	}

	_, err = s.serviceUserCredentialStore.Insert(ctx, *sanitizedCred)
	
	if err != nil {
		return nil, err
	}

	// we can skip transform here as we are sending back a secert and not serviceusercred
	return &domain.Token{
		Id:        credential.Id,
		Title:     credential.Title,
		Value:     credToken,
		CreatedAt: s.Now(),
	}, nil
}

func (s ServiceUserService) GetCred(ctx context.Context, credId string) (*domain.ServiceUserCredential, error) {

	UUID, err := utils.ConvertStringToUUID(credId) 
	
	if err != nil {
		return nil, utils.NewErrorf(utils.ErrorCodeInvalidArgument, utils.ErrInvalidUUIDFormat, err)
	}

	credModel, err := s.serviceUserCredentialStore.Select(ctx, UUID)

	if err != nil {
		return nil, err
	}

	sanitizedCred, err := credModel.Transform()

	if err != nil {
		return nil, err
	}

	cred, err := sanitizedCred.Transform()

	if err != nil {
		return nil, err
	}

	return cred, nil
}

func (s ServiceUserService) DeleteCred(ctx context.Context, credId string) error {

	UUID, err := utils.ConvertStringToUUID(credId) 
	
	if err != nil {
		return utils.NewErrorf(utils.ErrorCodeInvalidArgument, utils.ErrInvalidUUIDFormat, err)
	}

	_, err = s.serviceUserCredentialStore.Delete(ctx, UUID)

	if err != nil {
		return err
	}

	return nil
}

func (s ServiceUserService) ListCreds(ctx context.Context, filter domain.CredentialFilter) ([]domain.ServiceUserCredential, error) {
	if filter.Id != "" {	
		if !utils.IsValidUUID(filter.Id) {
			return nil,  utils.NewErrorf(utils.ErrorCodeInvalidArgument, utils.ErrInvalidUUIDFormat, filter.Id)
		}
	}

	serviceCredModels, err := s.serviceUserCredentialStore.Selects(ctx, filter)

	if err != nil {
		return nil, err
	}

	credentials := make([]domain.ServiceUserCredential, len(serviceCredModels))

	for i := range credentials {
		sanitizedUserCreds, err := serviceCredModels[i].Transform()

		if err != nil {
			return nil, err
		}
		
		userCreds, err := sanitizedUserCreds.Transform()

		if err != nil {
			return nil, err
		}

		credentials[i] = *userCreds
	}

	return credentials, nil
}

func (s ServiceUserService) IsSuper(ctx context.Context, id string, permissionName string) (bool, error) {
	return s.relationService.CheckPermission(ctx, relation.Relation{
		Subject: relation.Subject{
			Id:        id,
			Namespace: schema.ServiceUserPrincipal.String(),
		},
		Object: relation.Object{
			Id:        schema.PlatformId,
			Namespace: schema.PlatformNamespace.String(),
		},
		RelationName: permissionName,
	})
}

func (s ServiceUserService) FilterSupers(ctx context.Context, ids []string) ([]string, error) {
	
	relations := make([]relation.Relation, 0, len(ids))
	
	for _, id := range ids {
		rel := relation.Relation{
			Subject: relation.Subject{
				Id:        id,
				Namespace: schema.ServiceUserPrincipal.String(),
			},
			Object: relation.Object{
				Id:        schema.PlatformId,
				Namespace: schema.PlatformNamespace.String(),
			},
			RelationName: schema.PlatformSuperPermission.String(),
		}
		relations = append(relations, rel)
	}

	checkPairs, err := s.relationService.BatchCheckPermission(ctx, relations)
	
	if err != nil {
		return nil, err
	}
	
	sudoIDs := make([]string, 0, len(checkPairs))
	
	for i, checkPair := range checkPairs {
		if !checkPair.Status {
			sudoIDs = append(sudoIDs, ids[i])
		}
	}
	
	return sudoIDs, nil
}

func (s ServiceUserService) NewSuper(ctx context.Context, id string, rel string) error {
	
	UUID, err := utils.ConvertStringToUUID(id) 
	
	if err != nil {
		return utils.NewErrorf(utils.ErrorCodeInvalidArgument, utils.ErrInvalidUUIDFormat, err)
	}

	user, err := s.serviceUserStore.Select(ctx, UUID)
	
	if err != nil {
		return utils.WrapErrorf(err, utils.ErrorCodeUnknown, "s")
	}

	var name relation.RelationName
	var permission permission.PermissionName

	switch rel {
		case schema.MemberRelationName.String():
			name = schema.MemberRelationName
			permission = schema.PlatformCheckPermission
		case schema.AdminRelationName.String():
			name = schema.AdminRelationName
			permission = schema.PlatformSuperPermission
		default:
			return utils.WrapErrorf(
			errors.New("not a platform relation"), 
			utils.ErrorCodeInvalidArgument, 
			"possible options include: %s, %s",
			schema.AdminRelationName,
			schema.MemberRelationName,
		)
	}

	if _, err := s.IsSuper(ctx, user.Id.String(), permission.String()); err != nil {
		return err
	}

	_, err = s.relationService.New(ctx, relation.Relation{
		Object: relation.Object{
			Id:        schema.PlatformId,
			Namespace: schema.PlatformNamespace.String(),
		},
		Subject: relation.Subject{
			Id:        user.Id.String(),
			Namespace: schema.UserPrincipal.String(),
		},
		RelationName: name.String(),
	})
	
	return err
}

func (s ServiceUserService) UnSuper(ctx context.Context, id string) error {
	currentUser, err := s.Get(ctx, id)
	
	if err != nil {
		return err
	}

	if _, err := s.IsSuper(ctx, currentUser.Id, schema.PlatformCheckPermission.String()); err != nil {
		return err
	} 

	relationName := schema.MemberRelationName
	
	err = s.relationService.Remove(ctx, relation.Relation{
		Object: relation.Object{
			Id:        schema.PlatformId,
			Namespace: schema.PlatformNamespace.String(),
		},
		Subject: relation.Subject{
			Id:        currentUser.Id,
			Namespace: schema.ServiceUserPrincipal.String(),
		},
		RelationName: relationName.String(),
	})
	
	return err
}
