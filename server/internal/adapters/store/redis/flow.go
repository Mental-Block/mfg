package redis

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/server/internal/core/auth/domain"
	"github.com/server/pkg/metadata"
	"github.com/server/pkg/utils"
)

type FlowModel struct {
	Id utils.UUID
	Reason domain.Reason
	Strategy domain.Strategy
	Email domain.Email
	StartURL string
	FinishURL string
	Nonce string
	Metadata metadata.Metadata
	CreatedAt time.Time
	ExpiresAt time.Time
}

func (f FlowModel) Transform() domain.Flow {
	return domain.Flow{
		Id: f.Id,
		Reason: f.Reason,
		Strategy: f.Strategy,
		Email: f.Email,
		StartURL: f.StartURL,
		FinishURL: f.FinishURL,
		Nonce: f.Nonce,
		Metadata: f.Metadata,
		CreatedAt: f.CreatedAt,
		ExpiresAt: f.ExpiresAt,
	} 
}

var (
	ErrFlowNotFound = errors.New("Flow not found")
	FlowPrefix = "flow"
)

/*
	High level overview of FlowStore should not be directly imported.
	Copy interface and use dependancy injection over direct import.
*/
type IFlowStore interface { 
	Select(ctx context.Context, key domain.FlowKey) (*FlowModel, error)
	Insert(ctx context.Context, key domain.FlowKey, value domain.Flow, duration time.Duration) error
	Delete(ctx context.Context, key domain.FlowKey) (error)
	GenerateKey(id string) domain.FlowKey 
}

type FlowStore struct {
	cache *Store
}

func NewFlowStore(cache *Store) *FlowStore {
	return &FlowStore{
		cache: cache,
	}
}

func (rd *FlowStore) GenerateKey(id string) domain.FlowKey {
	return domain.FlowKey(strings.Join([]string{ FlowPrefix, id }, utils.KeyDilimeter))
}

func (rd *FlowStore) Insert(ctx context.Context, key domain.FlowKey, value domain.Flow, duration time.Duration) error {

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

func (rd *FlowStore) Select(ctx context.Context, key domain.FlowKey) (*FlowModel, error) {

	result, err := rd.cache.Get(ctx, key.String())

	if (err != nil) {
		return nil, CheckError(err)
	}

	flow := &FlowModel{}
	
	err = json.Unmarshal(result, flow)

	if (err != nil) {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "json.Unmarshal")
	}

	return flow, nil
}

func (rd *FlowStore) Delete(ctx context.Context, key domain.FlowKey) (error) {

	err := rd.cache.Delete(ctx, key.String())

	if (err != nil) {
		return CheckError(err)
	}

	return nil
}
 