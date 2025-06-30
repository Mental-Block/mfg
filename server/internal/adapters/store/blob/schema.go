package blob

import (
	"context"
	"fmt"
	"io"
	"strings"

	"gocloud.dev/blob"
	"gopkg.in/yaml.v3"

	"github.com/pkg/errors"
	"github.com/server/internal/adapters/bootstrap/schema"
)

type SchemaStore struct {
	store IBucketStore
}

func NewSchemaStore(b IBucketStore) *SchemaStore {
	return &SchemaStore{store: b}
}

// GetDefinition returns the service definition from the bucket
func (s *SchemaStore) GetDefinition(ctx context.Context) (*schema.ServiceDefinition, error) {
	var definitions []schema.ServiceDefinition

	// iterate over bucket files, only read .yml & .yaml files
	it := s.store.List(&blob.ListOptions{})

	for {
		obj, err := it.Next(ctx)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		if obj.IsDir {
			continue
		}
		
		if !(strings.HasSuffix(obj.Key, ".yaml") || strings.HasSuffix(obj.Key, ".yml")) {
			continue
		}

		fileBytes, err := s.store.ReadAll(ctx, obj.Key)

		if err != nil {
			return nil, fmt.Errorf("%s: %s", "error in reading bucket object", err.Error())
		}

		var def schema.ServiceDefinition

		if err := yaml.Unmarshal(fileBytes, &def); err != nil {
			return nil, errors.Wrap(err, "GetDefinitions: yaml.Unmarshal: "+obj.Key)
		}

		definitions = append(definitions, def)
	}

	return schema.MergeServiceDefinitions(definitions...), nil
}
