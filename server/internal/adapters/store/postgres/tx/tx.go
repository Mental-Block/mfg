package tx

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/server/internal/adapters/store/postgres"
	"github.com/server/internal/adapters/store/postgres/auth"
	"github.com/server/internal/adapters/store/postgres/user"
	"github.com/server/pkg/utils"
)

type txAdapters struct {
	UserStore  user.UserStore
	AuthStore  auth.AuthStore
}

type TXProvider struct {
	db *postgres.Store
}

func NewTXProvider(db *postgres.Store) *TXProvider {
	return &TXProvider{
		db,
	}
}

func (p *TXProvider) Transact(ctx context.Context, txFunc func(adapters txAdapters) error) error {
	return RunInTx(ctx, p.db.Pool, func(tx pgx.Tx) error {

		adapters := txAdapters{
			UserStore:  *user.NewUserStore(p.db),
			AuthStore: 	*auth.NewAuthStore(p.db),	
		}
		
		return txFunc(adapters)
	})
}

func RunInTx(ctx context.Context, pool *pgxpool.Pool, fn func(tx pgx.Tx) error) error {
	tx, err := pool.Begin(ctx)

	if err != nil {
		return utils.WrapErrorf(err, utils.ErrorCodeUnknown, "transaction begin failed %v", postgres.CheckPostgresError(err))
	}

	if err := fn(tx); err != nil {
		_ = tx.Rollback(ctx)
		return utils.WrapErrorf(err, utils.ErrorCodeUnknown, "transaction query failed %v", postgres.CheckPostgresError(err))
	}

	if err := tx.Commit(ctx); err != nil {
		return utils.WrapErrorf(err, utils.ErrorCodeUnknown, "transaction query commit %v", postgres.CheckPostgresError(err))
	}

	return nil
}