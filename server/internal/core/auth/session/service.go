package session

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/server/internal/adapters/store/redis"
	"github.com/server/internal/core/auth/domain"
	"github.com/server/pkg/crypt"
	"github.com/server/pkg/token"
	"github.com/server/pkg/utils"
)

var sessionDuration = (time.Hour * 24 * 30)

var ErrSessionService = errors.New("session service error") 
var ErrSigningToken = errors.New("invalid signature") 
var ErrInvalidToken = errors.New("invalid token")
var ErrSessionHasExpired = errors.New("session has expired")
var ErrCouldntCreateSession = errors.New("couldn't create session")

/* 
	High level overview of SessionService should not be directly imported.
 	Copy interface and use dependancy injection over direct import.
*/
type ISessionService interface {
	New(ctx context.Context, auth domain.SanitizedAuth) ([]byte, error)
	Verify(ctx context.Context, sessionTkn []byte) (*domain.Session, error)
	VerifyStrict(ctx context.Context, sessionTkn []byte) (*domain.Session, error)
	Refresh(ctx context.Context, sess domain.Session) ([]byte, error) 
	RemoveAllUserSessions(ctx context.Context, authId string) error
	Remove(ctx context.Context, id, authId string) error
	KillThemAll(ctx context.Context) error
}	

type SessionService struct {
	cfg Config
	sessionStore ISessionStore
	authStore IAuthStore
	Now func () time.Time
}

func NewSessionService(
	cfg Config,	
	auth IAuthStore,
	session ISessionStore,
) *SessionService {
	return &SessionService{
		cfg: cfg,
		sessionStore: session,
		authStore: auth,
		Now: func () time.Time {
			return time.Now().UTC()
		},
	}
}

func (s SessionService) New(ctx context.Context, auth domain.SanitizedAuth) ([]byte, error) {
	if (s.cfg.Validity == 0) {
		s.cfg.Validity = sessionDuration
	}

	if (s.cfg.Secret == "") {
		rdmLength, err := crypt.GererateRandomInt(20, 124)

		if (err != nil) {
			return nil, utils.NewErrorf(utils.ErrorCodeUnknown, "failed to generate random number")
		}

		s.cfg.Secret  = crypt.GenerateRandomStringFromLetters(int(rdmLength), crypt.CharList)
	}

	session := domain.Session{
		Id: utils.NewUUID(),
		AuthId: utils.UUID(auth.Id), 
		Email: domain.Email(auth.Email),
		Version: auth.Version,
		ExpiresAt: s.Now().Add(sessionDuration),
		CreatedAt: s.Now(),
	}

	claims := map[string]string{
		"version":  strconv.Itoa(session.Version),
	}

	tkn, err := token.BuildStdToken(
		s.cfg.Secret, 
		session.Id.String(), 
		s.cfg.Issuer,
		session.AuthId.String(), 
		sessionDuration, 
		claims,
	)

	if err != nil {
		return nil, utils.NewErrorf(utils.ErrorCodeUnknown, "token error: %v", ErrSigningToken)
	}

	key := s.sessionStore.GenerateKey(session.Id.String(), session.AuthId.String())
	err = s.sessionStore.Insert(ctx, key, session, sessionDuration)

	if err != nil {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, ErrSessionService.Error())
	}

	return tkn, err
}

// only validates the token portion of the session. we can still run into  time based 
// validity issues as we are not directly calling session storage. Only call if speed 
// is the primary and authorization aspect isn't too important    
func (s SessionService) Verify(ctx context.Context, sessionTkn []byte) (*domain.Session, error) {
	tkn, err := token.Parse( 
		sessionTkn,
		token.WithKey(token.HS256(), s.cfg.Secret),
		token.WithIssuer(s.cfg.Issuer),
		token.WithValidate(true),
	)

	var authId string
	err = tkn.Get("authId", &authId)

	if err != nil {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeNotAuthorized, ErrSessionHasExpired.Error())
	}

	id, ok := tkn.JwtID()

	if !ok {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeNotAuthorized, ErrSessionHasExpired.Error())
	}

	key := s.sessionStore.GenerateKey(id, authId)
	sess, err := s.sessionStore.Select(ctx, key)

	if err != nil {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeNotAuthorized, ErrSessionHasExpired.Error())
	}

	session := sess.Transform()

	return &session, nil
}

func (s SessionService) VerifyStrict(ctx context.Context, sessionTkn []byte) (*domain.Session, error) {
	tkn, err := token.Parse( 
		sessionTkn,
		token.WithKey(token.HS256(), s.cfg.Secret),
		token.WithIssuer(s.cfg.Issuer),
		token.WithValidate(true),
	)

	if err != nil {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, ErrSessionHasExpired.Error())
	}

	id, ok := tkn.JwtID()

	if !ok {
		return nil, utils.NewErrorf(utils.ErrorCodeNotAuthorized, ErrSessionHasExpired.Error()) 
	}

	var versionStr string
	err = tkn.Get("version", versionStr)

	if err != nil {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeNotAuthorized, ErrSessionHasExpired.Error())
	}

	var authId string
	err = tkn.Get("authId", authId)

	if err != nil {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeNotAuthorized, ErrSessionHasExpired.Error())
	}

	key := s.sessionStore.GenerateKey(id, authId)
	sess, err := s.sessionStore.Select(ctx, key)

	if err != nil || sess.ExpiresAt.After(time.Now()) {
		return nil, utils.WrapErrorf(err,  utils.ErrorCodeNotAuthorized, ErrSessionHasExpired.Error())
	}

	// technically, we don't need to check version. we can remove it later if preformace is an issue. it's more a peice of mind check
	version, err := strconv.Atoi(versionStr)

	if sess.Version != version || err != nil {
		return nil, utils.WrapErrorf(err,  utils.ErrorCodeNotAuthorized, ErrSessionHasExpired.Error())
	}

	ses := sess.Transform()

	return &ses, nil
}	

// prev session should be verified before calling refresh
func (s SessionService) Refresh(ctx context.Context, sess domain.Session) ([]byte, error) {

	err := s.Remove(ctx, sess.Id.String(), sess.AuthId.String())

	if (err != nil) {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, ErrSessionService.Error())
	}

	// technically, we don't need to call authstore. we can remove it later if preformace is an issue. it's more a peice of mind check
	authModel, err := s.authStore.Select(ctx, sess.Email)

	if err != nil {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, ErrSessionService.Error())
	}

	tkn, err := s.New(ctx, domain.SanitizedAuth{
		Id: authModel.Id,
		Email: authModel.Email,
		Version: authModel.Version,
	})

	if err != nil {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, ErrSessionService.Error())
	}

	return tkn, nil
}

func (s SessionService) RemoveAllUserSessions(ctx context.Context, authId string) error {
	key := s.sessionStore.GenerateKey(authId, "*")
	err := s.sessionStore.DeleteByPrefix(ctx, key)

	if (err != nil) {
		return utils.WrapErrorf(err, utils.ErrorCodeUnknown, "failed to remove session")
	}	

	return nil
}

func (s SessionService) Remove(ctx context.Context, id, authId string) error {
	key := s.sessionStore.GenerateKey(id, authId)
	err := s.sessionStore.Delete(ctx, key)

	if (err != nil) {
		return utils.WrapErrorf(err, utils.ErrorCodeUnknown, "failed to remove session")
	}	

	return nil
}

func (s SessionService) KillThemAll(ctx context.Context) error {
	key := domain.SessionKey(strings.Join([]string{redis.SessionPrefix, "*"},  utils.KeyDilimeter))
	err := s.sessionStore.DeleteByPrefix(ctx, key)

	if (err != nil) {
		return utils.WrapErrorf(err, utils.ErrorCodeUnknown, "failed to remove all sessions")
	}

	return nil
}