package postgres

import (
	"errors"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx"
	"github.com/server/pkg/utils"
)

var (
	ErrDuplicateKey              = errors.New("duplicate key")
	ErrCheckViolation            = errors.New("check constraint violation")
	ErrForeignKeyViolation       = errors.New("foreign key violation")
	ErrInvalidTextRepresentation = errors.New("invalid input syntax type")
	ErrInvalidID                 = errors.New("invalid id")
)

func CheckPostgresError(err error) error {
	var pgErr *pgx.PgError

	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case pgerrcode.UniqueViolation:
			return utils.WrapErrorf(err, utils.ErrorCodeInvalidArgument, "%s", ErrDuplicateKey)
		case pgerrcode.CheckViolation:
			return utils.WrapErrorf(err, utils.ErrorCodeInvalidArgument, "%s", ErrCheckViolation)
		case pgerrcode.ForeignKeyViolation:
			return utils.WrapErrorf(err, utils.ErrorCodeInvalidArgument, "%s", ErrForeignKeyViolation)
		case pgerrcode.InvalidTextRepresentation:
			return utils.WrapErrorf(err, utils.ErrorCodeInvalidArgument, "fix your query: %s", ErrInvalidTextRepresentation)
		}
	}

	return utils.WrapErrorf(err, utils.ErrorCodeUnknown, "ohh crap investigate: %s", err)
}