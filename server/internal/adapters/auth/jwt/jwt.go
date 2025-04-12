package jwt

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/server/internal/core/ports"
)

var ErrInvalidToken = errors.New("invalid token")

type JWTConfig struct {
	secret []byte
}

func New(salt string) ports.TokenService {
	return &JWTConfig{
		secret: []byte(salt),
	}
}

func (t *JWTConfig) CreateToken(claims jwt.MapClaims) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(t.secret)

	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

func (t *JWTConfig) ParseToken(token string) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return t.secret, nil
	})

	if err != nil {
		return nil, err

	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}

func (t *JWTConfig) VerifyToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		return t.secret, nil
	})

	if err != nil {
		return err
	}

	if !parsedToken.Valid {
		return ErrInvalidToken
	}

	return nil
}
