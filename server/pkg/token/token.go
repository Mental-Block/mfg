package token

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/lestrrat-go/jwx/v3/jwa"
	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jwt"
)

// https://github.com/lestrrat-go/jwx/blob/develop/v3/Changes-v3.md
type JWkSet jwk.Set

const KeyIDKey = jwk.KeyIDKey

func WithKeySet(set jwk.Set, options ...interface{}) jwt.ParseOption {
	return jwt.WithKeySet(set, options)
} 

func Parse(s []byte, options ...jwt.ParseOption) (jwt.Token, error) {
	return jwt.Parse(s, options...)
}

func WithKey(alg jwa.KeyAlgorithm, key interface{}, suboptions ...jwt.Option) jwt.SignEncryptParseOption {
	return jwt.WithKey(alg, key, suboptions...)
}

func WithJwtId(s string) jwt.ValidateOption {
	return jwt.WithJwtID(s)
}

func WithValidate(v bool) jwt.ParseOption {
	return jwt.WithValidate(v)
}

func WithIssuer(s string)  jwt.ValidateOption {
	return jwt.WithIssuer(s)
}

func WithSubject(s string) jwt.ValidateOption {
	return jwt.WithSubject(s)
}

func HS256() jwa.SignatureAlgorithm {
	return jwa.HS256()
}

func ParseInsecure(s []byte, options ...jwt.ParseOption) (jwt.Token, error) {
	return jwt.ParseInsecure(s, options...)
}

func JWKParse(src []byte, options ...jwk.ParseOption) (jwk.Set, error) {
	return jwk.Parse(src, options...)
}

func JWKPem(v interface{}) ([]byte, error) {
	return jwk.Pem(v)
}

func JWKNewSet() jwk.Set {
	return jwk.NewSet()
}



// Deprecated: In v3... Not a standard func in library. Please avoid using
func AsMap(token jwt.Token) (map[string]any) {
	claimsMap := make(map[string]any)
	
	for _, key := range token.Keys() {
		var val interface{}
		if err := token.Get(key, &val); err != nil {
			continue
		}
		claimsMap[key] = val
	}

	return claimsMap
}

func ReadFile(path string, options ...jwk.ReadFileOption) (jwk.Set, error){
	return jwk.ReadFile(path, options...) 
}

const (
	RSAKeySize = 2048
)

func CreateJWKs(numOfKeys int) (jwk.Set, error) {
	keySet := JWKNewSet()
	for ; numOfKeys > 0; numOfKeys-- {
		
		keyRaw, err := rsa.GenerateKey(rand.Reader, RSAKeySize)

		if err != nil {
			return nil, err
		}

		key, err := jwk.Import(keyRaw)


		if err != nil {
			return nil, err
		}

		rsaKey, ok := key.(jwk.RSAPrivateKey)

		if !ok {
			return nil, fmt.Errorf("failed to convert jwk.Key into jwk.RSAPrivateKey")
		}

		pubKey, err := rsaKey.PublicKey()

		if err != nil {
			return nil, err
		}

		thumb, err := pubKey.Thumbprint(crypto.SHA256)

		if err != nil {
			return nil, err
		}

		if err := rsaKey.Set(jwk.AlgorithmKey, "RS256"); err != nil {
			return nil, err
		}

		if err := rsaKey.Set(jwk.KeyUsageKey, "sig"); err != nil {
			return nil, err
		}

		if err := rsaKey.Set(jwk.KeyIDKey, base64.RawURLEncoding.EncodeToString(thumb)); err != nil {
			return nil, err
		}

		if err := keySet.AddKey(rsaKey); err != nil {
			return nil, err
		}
	}
	
	return keySet, nil
}

// generate key
func CreateJWKWithKID(id string) (jwk.Key, error) {

	keyRaw, err := rsa.GenerateKey(rand.Reader, RSAKeySize)
	
	if err != nil {
		return nil, err
	}

	rsaKey, err := jwk.Import(keyRaw)
	
	if err != nil {
		return nil, err
	}

	if err := rsaKey.Set(jwk.AlgorithmKey, "RS256"); err != nil {
		return nil, err
	}

	if err := rsaKey.Set(jwk.KeyUsageKey, "sig"); err != nil {
		return nil, err
	}

	if err := rsaKey.Set(jwk.KeyIDKey, id); err != nil {
		return nil, err
	}

	return rsaKey, nil
}

// GetPublicKeySet convert private to public
func GetPublicKeySet(privateKeySet jwk.Set) (jwk.Set, error) {
	
	publicKeySet, err := jwk.PublicSetOf(privateKeySet)

	if err != nil {
		return nil, fmt.Errorf("failed to generate public key from private rsa: %w", err)
	}

	return publicKeySet, nil
}

// BuildToken creates a signed jwt using provided private key
// RS256/ES256 (Asymmetric)
// Ensure the key contains kid else the operation fails
func BuildToken(rsaKey jwk.Key, issuer, sub string, validity time.Duration, customClaims map[string]string) ([]byte, error) {
		
	key, ok := rsaKey.KeyID() 

	if key == "" || !ok {
		return nil, fmt.Errorf("key id is empty")
	}

	body := jwt.NewBuilder().
		Issuer(issuer).
		IssuedAt(time.Now().UTC()).
		NotBefore(time.Now().UTC()).
		Expiration(time.Now().UTC().Add(validity)).
		JwtID(uuid.New().String()).
		Subject(sub)
	
	body.Claim(jwk.KeyIDKey, key)

	for claimKey, claimVal := range customClaims {
		body = body.Claim(claimKey, claimVal)
	}

	tok, err := body.Build()

	if err != nil {
		return nil, err
	}

	return jwt.Sign(tok, jwt.WithKey(jwa.RS256(), rsaKey))
}

// Build Std token without needing a public, private encryption key pair set.  
// This is a symmetric (HS256) 
// Only use this with internal applications and never with a 3rd party sevices. 
func BuildStdToken(secret, id, issuer, sub string, validity time.Duration, customClaims map[string]string) ([]byte, error) {
	body := jwt.NewBuilder().
		Issuer(issuer).
		IssuedAt(time.Now().UTC()).
		NotBefore(time.Now().UTC()).
		Expiration(time.Now().UTC().Add(validity)).
		JwtID(id).
		Subject(sub)

	for claimKey, claimVal := range customClaims {
		body = body.Claim(claimKey, claimVal)
	}

	tok, err := body.Build()

	if (err != nil) {
		return nil, err
	}

	return jwt.Sign(tok, jwt.WithKey(jwa.HS256(), secret))
}
