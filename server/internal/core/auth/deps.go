package auth

import (
	"context"
	"time"

	auth_store "github.com/server/internal/adapters/store/postgres/auth"
	"github.com/server/internal/adapters/store/postgres/tx"
	user_store "github.com/server/internal/adapters/store/postgres/user"
	"github.com/server/internal/adapters/store/redis"

	"github.com/server/internal/core/auth/domain"
	serviceuser "github.com/server/internal/core/serviceuser/domain"

	"github.com/server/pkg/mailer"
	"github.com/server/pkg/token"
	"github.com/server/pkg/utils"
)

type IAuthStore interface {
	SelectUser(ctx context.Context, id utils.UUID) (*user_store.UserModel, error)
	Select(ctx context.Context, email domain.Email) (*auth_store.AuthModel, error)
	Update(ctx context.Context, auth domain.SanitizedAuth) error
	Insert(ctx context.Context, auth domain.SanitizedAuth) (*auth_store.AuthModel, error)
}

type IEmailService interface {
    DialAndSend(m *mailer.Message) error
    FromHeader() string
    NewMessage(settings ...mailer.MessageSetting) *mailer.Message
}

type ITXAuthService interface {
 	UserCreationTransaction(ctx context.Context, cmd tx.TXAuthCmd) error
}

type IPasswordService interface {
	CreateHash(password string) (hash string, err error)
	ComparePasswordAndHash(password, hash string) (match bool, err error)
}

type IFlowStore interface { 
	Select(ctx context.Context, key domain.FlowKey) (*redis.FlowModel, error)
	Insert(ctx context.Context, key domain.FlowKey, value domain.Flow, duration time.Duration) error
	Delete(ctx context.Context, key domain.FlowKey) (error)
	GenerateKey(id string) domain.FlowKey 
}

type ITokenService interface {
	GetPublicKeySet() token.JWkSet
	Build(subjectID string, metadata map[string]string) ([]byte, error)
	Parse(ctx context.Context, authToken []byte) (string, map[string]any, error)
}

type ISessionService interface {
	New(ctx context.Context, auth domain.SanitizedAuth) ([]byte, error)
	VerifyStrict(ctx context.Context, sessionTkn []byte) (*domain.Session, error)
	RemoveAllUserSessions(ctx context.Context, authId string) error
}	

type IServiceUserService interface {
	GetByJWT(ctx context.Context, tkn []byte) (*serviceuser.ServiceUser, error)
	GetBySecret(ctx context.Context, credId string, secret string) (*serviceuser.ServiceUser, error)	
 	Get(ctx context.Context, id string) (*serviceuser.ServiceUser, error) 
}
