package ports

import "github.com/server/internal/adapters/auth/argon"


type PasswordService interface {
	CreateHash(password string) (hash string, err error)
	ComparePasswordAndHash(password, hash string) (match bool, err error)
	CheckHash(password, hash string) (match bool, params *argon.Params, err error)
	DecodeHash(hash string) (params *argon.Params, salt, key []byte, err error)
}