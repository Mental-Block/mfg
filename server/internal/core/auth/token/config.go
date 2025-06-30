package token

import "time"

type Config struct {
	// Path to rsa key file, it can contain more than one key as a json array
	// jwt will be signed by first key, but will be tried to be decoded by all matching key ids, this helps in key rotation.
	// If not provided, access token will not be generated
	RSAPath string `yaml:"rsa_path"`
	
	// RSABase64 is base64 encoded rsa key, it can contain more than one key as a json array
	RSABase64 string `yaml:"rsa_base64"`

	// Issuer uniquely identifies the service that issued the token
	// a good example could be fully qualified domain name
	Issuer string `yaml:"iss" default:"mfg"`

	// Validity is the duration for which the token is valid
	Validity time.Duration `yaml:"validity" mapstructure:"validity" default:"1h"`

	Claims ClaimConfig `yaml:"claims" mapstructure:"claims"`
}

type ClaimConfig struct {
	AddOrgIDsClaim    bool `yaml:"add_org_ids" mapstructure:"add_org_ids" default:"true"`
	AddUserEmailClaim bool `yaml:"add_user_email" mapstructure:"add_user_email" default:"true"`
}