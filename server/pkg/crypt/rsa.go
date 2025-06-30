package crypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

// RsaKeySet holds the PEM-encoded private and public keys.
type RsaKeySet struct {
	PrivateKeyPEM []byte
	PublicKeyPEM  []byte
}

// GenerateRsaKeySet generates a new RSA private/public key pair and encodes them to PEM format.
func GenerateRsaKeySet(bits int) (*RsaKeySet, error) {
	// Generate a new RSA private key.
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	
	if err != nil {
		return nil, fmt.Errorf("error generating private key: %w", err)
	}

	// Encode the private key to PEM format.
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	// Extract and encode the public key to PEM format.
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&privateKey.PublicKey),
	})

	return &RsaKeySet{
		PrivateKeyPEM: privateKeyPEM,
		PublicKeyPEM:  publicKeyPEM,
	}, nil
}
