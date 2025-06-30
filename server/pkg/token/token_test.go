package token

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuildToken(t *testing.T) {
	issuer := "test"
	sub := uuid.New().String()
	validity := time.Minute * 10
	kid := uuid.New().String()
	newKey, err := CreateJWKWithKID(kid)
	assert.NoError(t, err)
	t.Run("create a valid token", func(t *testing.T) {
		got, err := BuildToken(newKey, issuer, sub, validity, nil)
		assert.NoError(t, err)
		parsedToken, err := jwt.ParseInsecure(got)
		assert.NoError(t, err)
		issuerVal, ok :=  parsedToken.Issuer(); 
		assert.True(t, ok)
		assert.Equal(t, issuer, issuerVal)
		subjectVal, ok := parsedToken.Subject()
		assert.True(t, ok)
		assert.Equal(t, sub, subjectVal)
		var gotKid string
		require.NoError(t, parsedToken.Get(jwk.KeyIDKey, &gotKid))
		assert.Equal(t, kid, gotKid)
	})
}
