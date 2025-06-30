package spicedb

import (
	"context"
	"errors"
	"fmt"

	authzedpb "github.com/authzed/authzed-go/proto/authzed/api/v1"
)

type SchemaStore struct {
	spiceDB *SpiceDB
}

var (
	ErrWritingSchema = errors.New("error in writing schema to spicedb")
)

func NewSchemaStore(spiceDB *SpiceDB) *SchemaStore {
	return &SchemaStore{
		spiceDB: spiceDB,
	}
}

func (r SchemaStore) WriteSchema(ctx context.Context, schema string) error {

	if _, err := r.spiceDB.Client.WriteSchema(ctx, &authzedpb.WriteSchemaRequest{Schema: schema}); err != nil {
		return fmt.Errorf("%w: %s", ErrWritingSchema, err.Error())
	}

	return nil
}
