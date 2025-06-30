package postgres

import (
	"context"
	"database/sql"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/server/pkg/utils"
)

type Store struct {
	*pgxpool.Pool
}

const PublicSchema = "public"

var (
	TABLE_AUTH 					  = "auth"
	TABLE_USER_AUTH				  = "user_auth"
	TABLE_USER					  = "user"
	TABLE_NAMESPACE               = "namespace"
	TABLE_RELATION                = "relation"
	TABLE_PERMISSION              = "permission"
	TABLE_ROLE                    = "role"
	TABLE_RESOURCE                = "resource"
	TABLE_SERVICEUSER             = "serviceuser"
	TABLE_SERVICEUSER_CREDENTIAL  = "serviceuser_credential"



	TABLE_GROUP                   = "group"
	TABLE_ORGANIZATION            = "organization"
	TABLE_POLICY               	  = "policy"
	TABLE_PROJECT                 = "project"
	TABLE_METASCHEMA              = "metaschema"
	TABLE_FLOWS                   = "flow"
	TABLE_INVITATION              = "invitation"
	TABLE_AUDITLOG                = "auditlog"
	TABLE_DOMAIN                  = "domain"
	TABLE_PREFERENCE              = "preferences"
	TABLE_WEBHOOK_ENDPOINT        = "webhook_endpoint"
)

var (
	pgInstance *Store
	pgOnce     sync.Once
)

func NewStore(ctx context.Context, cfg Config) (*Store, error) {
	var err error

	pgOnce.Do(func() {
		db, dbErr := pgxpool.New(ctx, cfg.URL)

		if dbErr != nil {
			err = utils.WrapErrorf(dbErr, utils.ErrorCodeUnknown, "pgxpool.New")
			return
		}

		pgInstance = &Store{db}
	})

	return pgInstance, err
}

func (pg *Store) Ping(ctx context.Context) error {
	return pg.Pool.Ping(ctx)
}

func (pg *Store) Close() {
	pg.Pool.Close()
}

func (pg *Store) PgxToStdLib() *sql.DB {
	return stdlib.OpenDBFromPool(pg.Pool)
}

