package strategy

import (
	"github.com/server/internal/core/auth/domain"
	"github.com/server/pkg/crypt"
	"github.com/server/pkg/utils"
)

type Password struct {
	params crypt.Params
} 

func NewPassword(param crypt.Params) *Password{
	return &Password{
		params: crypt.Params(param),
	}
}

func IsPasswordValid(password string, hash domain.Password) (bool, error) {
	passwrd, err := domain.Password(password).NewPassword()

	if (err != nil) {
		return false, utils.WrapErrorf(err, utils.ErrorCodeInvalidArgument, "invalid password format")
	}
	
	ok, err := crypt.ComparePasswordAndHash(passwrd.String(), hash.String());

	if !ok || err != nil {
		return ok, utils.WrapErrorf(err, utils.ErrorCodeInvalidArgument, "invalid password")
	}

	return ok, nil
}

func (p Password) CreateHash(password string) (domain.Password, error) {
	pas, err := crypt.CreateHash(p.params, password)
	return domain.Password(pas), err
}
