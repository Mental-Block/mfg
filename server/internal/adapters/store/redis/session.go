package redis

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/server/internal/core/auth/domain"
	"github.com/server/pkg/utils"
)

type SessionModel struct {
	Id utils.UUID
	AuthId utils.UUID
	Email domain.Email
	Version int
	CreatedAt time.Time
	ExpiresAt time.Time
}

func (s SessionModel) Transform() domain.Session {
	return domain.Session{
		Id: s.Id,
		AuthId: s.AuthId,
		Email: s.Email,
		Version: s.Version,
		CreatedAt: s.CreatedAt,
		ExpiresAt: s.ExpiresAt,
	}
}


var (
	ErrSessionNotFound = errors.New("Session not found")
	SessionPrefix = "session"
)

/*
	High level overview of SessionStore should not be directly imported.
	Copy interface and use dependancy injection over direct import.
*/
type ISessionStore interface { 
	Select(ctx context.Context, key domain.SessionKey) (*SessionModel, error) 
	Insert(ctx context.Context, key domain.SessionKey, value domain.Session, duration time.Duration) error
	Delete(ctx context.Context, key domain.SessionKey) error
	DeleteByPrefix(ctx context.Context, key domain.SessionKey) error
	GenerateKey(id, authId string) domain.SessionKey
}

type SessionStore struct {
	cache *Store
}

func NewSessionStore(cache *Store) *SessionStore {
	return &SessionStore{
		cache: cache,
	}
}

func (rd *SessionStore) GenerateKey(id, authId string) domain.SessionKey {
	return domain.SessionKey(strings.Join([]string{SessionPrefix, authId, id}, utils.KeyDilimeter)) 
}

func (rd *SessionStore) Insert(ctx context.Context, key domain.SessionKey, value domain.Session, duration time.Duration) error {

	b, err := json.Marshal(value)

	if (err != nil) {
		return utils.NewErrorf(utils.ErrorCodeInvalidArgument, "json.Marshal")
	}

	err = rd.cache.Set(ctx, key.String(), b, duration) 

	if (err != nil) {
		return utils.WrapErrorf(err, utils.ErrorCodeUnknown, err.Error()) 
	}

	return nil
}

func (rd *SessionStore) Select(ctx context.Context, key domain.SessionKey) (*SessionModel, error) {

	result, err := rd.cache.Get(ctx, key.String())

	if (err != nil) {
		return nil, CheckError(err)
	}

	session := &SessionModel{}
	
	err = json.Unmarshal(result, session)

	if (err != nil) {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "json.Unmarshal")
	}

	return session, nil
}

func (rd *SessionStore) Delete(ctx context.Context, key domain.SessionKey) error {

	err := rd.cache.Delete(ctx, key.String())

	if (err != nil) {
		return CheckError(err)
	}

	return nil
}

func (rd *SessionStore) DeleteByPrefix(ctx context.Context, key domain.SessionKey) error {

	err := rd.cache.DeleteByPrefix(ctx, key.String())

	if (err != nil) {
		return CheckError(err)
	}

	return nil
}