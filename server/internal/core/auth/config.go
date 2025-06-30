package auth

import (
	"github.com/server/internal/core/auth/domain"
	"github.com/server/internal/core/auth/session"
	"github.com/server/internal/core/auth/strategy"
	"github.com/server/internal/core/auth/token"
)

type Config struct {
  	// CallbackURLs is external host used for redirect uri
	// host specified at 0th index will be used as default
	CallbackURLs []string `yaml:"callback_urls" default:"[http://localhost:8084/api/v1/auth/callback]"`
	
	AuthorizedRedirectURLs []string `yaml:"authorized_redirect_urls"`

	Session session.Config `yaml:"session"`

	Token token.Config  `yaml:"token"`
	
	OIDC map[domain.Strategy]strategy.OIDCConfig `yaml:"oicd"`

	OTP strategy.MailConfig `yaml:"mail_otp"`

	Link strategy.MailConfig `yaml:"mail_link"`

	Password strategy.PasswordConfig `yaml:"password"`
}
