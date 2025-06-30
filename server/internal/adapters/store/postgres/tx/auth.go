package tx

import (
	"context"

	auth_domain "github.com/server/internal/core/auth/domain"
	user_domain "github.com/server/internal/core/user/domain"
	"github.com/server/pkg/utils"
)

type TXAuthCmd struct {
	User user_domain.SanitizedUser
	Auth auth_domain.SanitizedAuth
}

type txAuthService struct {
	txProvider txProvider
}

type txProvider interface {
	Transact(ctx context.Context, txFunc func(adapters txAdapters) error) error
}

func NewAuthTxConsumer(
	txProvider txProvider,
) *txAuthService {
	return &txAuthService{
		txProvider: txProvider,
	}
}

func (h *txAuthService) UserCreationTransaction(ctx context.Context, cmd TXAuthCmd) error {
	return h.txProvider.Transact(ctx, func(adapters txAdapters) error {

		u, err := adapters.UserStore.Insert(ctx, cmd.User)

		if err != nil {
			return utils.WrapErrorf(err, utils.ErrorCodeUnknown, "failed to create user in transaction")
		}

		a, err := adapters.AuthStore.Insert(ctx, cmd.Auth)

		if err != nil {
			return utils.WrapErrorf(err, utils.ErrorCodeUnknown, "failed to create authentication in transaction")
		}

		err = adapters.AuthStore.AssignUser(ctx, a.Id, u.Id)

		if err != nil {
			return utils.WrapErrorf(err, utils.ErrorCodeUnknown, "failed to assign authentication to user in transaction")
		}

		return nil
	})
}
