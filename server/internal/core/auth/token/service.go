package token

import (
	"context"
	"encoding/base64"
	"errors"
	"log"
	"time"

	"github.com/lestrrat-go/jwx/v3/jwk"

	"github.com/server/pkg/token"
	"github.com/server/pkg/utils"
)

var (
	ErrMissingRSADisableToken = errors.New("rsa key missing in config, generate and pass file path")
	ErrInvalidToken           = errors.New("failed to verify a valid token")
	ErrParsingRSAFile         = errors.New("failed to parse rsa key")
)

const (
	GeneratedClaimKey   = "gen"
	GeneratedClaimValue = "system"
	OrgIDsClaimKey      = "org_ids"
	SubTypeClaimsKey    = "sub_type"
	SubEmailClaimsKey   = "email"
)

/* 
	High level overview of TokenService should not be directly imported.
 	Copy interface and use dependancy injection over direct import.
*/
type ITokenService interface {
	GetPublicKeySet() token.JWkSet
	Build(subjectId string, metadata map[string]string) ([]byte, error)
	Parse(ctx context.Context, userToken []byte) (string, map[string]any, error)
}	

type TokenService struct {
	keySet       token.JWkSet
	publicKeySet token.JWkSet
	issuer       string
	validity     time.Duration
}

func NewTokenService(cfg Config) TokenService {
	keySet, err := tryToParseKeySet(cfg.RSAPath, cfg.RSABase64)

	if err != nil {
		log.Fatal(err)
	}

	publicKeySet := token.JWKNewSet()

	if keySet != nil {

		pub, err := token.GetPublicKeySet(keySet)

		if err != nil {
			log.Fatal(err)
		}

		publicKeySet = pub
	}

	return TokenService{
		keySet:       keySet,
		issuer:       cfg.Issuer,
		publicKeySet: publicKeySet,
		validity:     cfg.Validity,
	}
}

// GetPublicKeySet returns the public keys to verify the access token
func (s TokenService) GetPublicKeySet() token.JWkSet {
	return s.publicKeySet
}

// Build creates an access token for the given subjectId
func (s TokenService) Build(subjectId string, metadata map[string]string) ([]byte, error) {
	if s.keySet == nil {
		return nil, ErrMissingRSADisableToken
	}

	// use first key to sign token
	rsaKey, ok := s.keySet.Key(0)
	
	if !ok {
		return nil, errors.New("missing rsa key to generate token")
	}

	// generated token has an extra custom claim
	// used to identify which public key to use to verify the token
	metadata[GeneratedClaimKey] = GeneratedClaimValue
	return token.BuildToken(rsaKey, s.issuer, subjectId, s.validity, metadata)
}

func (s TokenService) Parse(ctx context.Context, userToken []byte) (string, map[string]any, error) {
	if s.keySet == nil {
		return "", nil, ErrMissingRSADisableToken
	}

	verifiedToken, err := token.Parse(userToken, token.WithKeySet(s.publicKeySet))

	if err != nil {
		return "", nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "%s", ErrInvalidToken)
	}

	tokenClaims := token.AsMap(verifiedToken)

	sub, ok := verifiedToken.Subject()

	if (!ok) {
		return "", nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "%s", ErrInvalidToken)
	}
	
	return sub, tokenClaims, nil
}














func tryToParseKeySet(rsapath string, rsabase64 string) (token.JWkSet, error) {
	var tokenKeySet token.JWkSet
	if rsapath != "" {
		ks, err := jwk.ReadFile(rsapath);

		if err != nil {
			return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "%s", ErrParsingRSAFile)
		} 

		tokenKeySet = ks
	}
	
	if len(rsabase64) > 0 {
		rawDecoded, err := base64.StdEncoding.DecodeString(rsabase64)
		
		if err != nil {
			return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "%s: %s", "failed to decode rsa key as std-base64", "base64.StdEncoding.DecodeString")
		}

		ks, err := jwk.Parse(rawDecoded);

		if err != nil {
			return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "%s: %s", ErrParsingRSAFile, "jwk.Parse")
		} 
		
		tokenKeySet = ks
	}

	return tokenKeySet, nil
}