package schema

import "errors"

var (
	ErrMigration    = errors.New("error in migrating authz schema")
	ErrBadNamespace = errors.New("bad namespace, format should namespace:uuid")
)