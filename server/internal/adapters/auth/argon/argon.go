package argon

import (
	"encoding/base64"

	"github.com/server/internal/core/domain"
	"github.com/server/internal/core/ports"
	"golang.org/x/crypto/argon2"
)

type Password struct {
	salt []byte
}

func New(salt string) ports.PasswordService {
	return &Password{
		salt: []byte(salt),
	}
}

func (p *Password) HashPassword(password domain.Password) domain.Password {
	return domain.Password(base64.StdEncoding.EncodeToString(argon2.Key([]byte(password), p.salt, 1, 32*1024, 4, 32)))
}

func (p *Password) VerifyPassword(password domain.Password, hash domain.Password) bool {
	newHash := p.HashPassword(password)
	return newHash == hash
}
