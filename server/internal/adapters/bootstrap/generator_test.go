package bootstrap_test

import (
	"context"
	_ "embed"
	"fmt"
	"sort"
	"testing"

	"github.com/authzed/spicedb/pkg/caveats/types"
	"github.com/authzed/spicedb/pkg/schemadsl/compiler"
	"github.com/server/internal/adapters/bootstrap"
	"github.com/server/internal/adapters/bootstrap/schema"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

//go:embed testdata/compute_service.yml
var computeSchemaYaml []byte

//go:embed testdata/compiled_schema.zed
var compiledSchemaZed string

func TestCompileSchema(t *testing.T) {
	tenantName := "acme"
	compiledSchema, err := compiler.Compile(compiler.InputSchema{
		Source:       "base_schema.zed",
		SchemaString: schema.BaseSchemaZed,
	}, compiler.ObjectTypePrefix(tenantName))
	assert.NoError(t, err)

	appService, err := bootstrap.BuildServiceDefinitionFromAZSchema(compiledSchema.ObjectDefinitions, "app")
	assert.NoError(t, err)
	assert.Len(t, appService.Permissions, 25)
}

func TestAddServiceToSchema(t *testing.T) {
	tenantName := "mfg"
	existingSchema, err := compiler.Compile(compiler.InputSchema{
		Source:       "base_schema.zed",
		SchemaString: schema.BaseSchemaZed,
	}, compiler.ObjectTypePrefix(tenantName))
	assert.NoError(t, err)

	computeServiceDefinition := schema.ServiceDefinition{}
	err = yaml.Unmarshal(computeSchemaYaml, &computeServiceDefinition)
	assert.NoError(t, err)

	spiceDBDefinitions := existingSchema.ObjectDefinitions
	spiceDBDefinitions, err = bootstrap.ApplyServiceDefinitionOverAZSchema(&computeServiceDefinition, spiceDBDefinitions)
	assert.NoError(t, err)

	// sort definitions, useful to keep it consistent
	for idx := range spiceDBDefinitions {
		sort.Slice(spiceDBDefinitions[idx].GetRelation(), func(i, j int) bool {
			return spiceDBDefinitions[idx].GetRelation()[i].GetName() < spiceDBDefinitions[idx].GetRelation()[j].GetName()
		})
	}
	sort.Slice(spiceDBDefinitions, func(i, j int) bool {
		return spiceDBDefinitions[i].GetName() < spiceDBDefinitions[j].GetName()
	})

	caveatTypeSet := &types.TypeSet{} 

	authzedSchemaSource, err := bootstrap.PrepareSchemaAsAZSource(spiceDBDefinitions, caveatTypeSet)
	assert.NoError(t, err)

	// compile and validate generated schema
	err = bootstrap.ValidatePreparedAZSchema(context.Background(), authzedSchemaSource)
	assert.NoError(t, err)

	if compiledSchemaZed != authzedSchemaSource {
		fmt.Println(authzedSchemaSource)
	}
	assert.Equal(t, compiledSchemaZed, authzedSchemaSource)
}
