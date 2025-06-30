package session

import (
	"time"
)

type Config struct {
  // If left empty it will auto generate a secret hash key
  Secret string `yaml:"secret" default:""`
	
  // Issuer uniquely identifies the service that issued the token
	// a good example could be fully qualified domain name
	Issuer string `yaml:"iss" default:"mfg"`
	
  // Validity is the duration for which the session is valid
  Validity time.Duration `yaml:"validity" default:"720h"`  
  
  // SameSite can be set to "default", "lax", "strict" or "none". use "strict" for production
  SameSite string `yaml:"same_site" default:"lax"`
	
  // use true for production
  Secure bool `yaml:"secure" default:"false"`
}


