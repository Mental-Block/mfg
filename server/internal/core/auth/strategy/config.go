package strategy

import (
	"time"

	"github.com/server/internal/core/auth/domain"
	"github.com/server/pkg/crypt"
)

type MailConfig struct {
	MailTemplates map[domain.Reason]MailTemplateConfig `yaml:"mail_templates"`
	Enabled bool `yaml:"enabled" default:"false"`
}

type PasswordConfig struct {
	Params crypt.Params `yaml:"params"`
	MailTemplates map[domain.Reason]MailTemplateConfig `yaml:"mail_templates"` 
	Enabled bool `yaml:"enabled" default:"false"`
}

type OIDCConfig struct {
	ClientId     string  `yaml:"client_id" default:"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"`
	ClientSecret string  `yaml:"client_secret" default:"xxxxxxxxxxxxxxxxxxxxxxxxx.apps.googleusercontent.com"`
	IssuerUrl    string  `yaml:"issuer_url" default:"https://accounts.google.com"`
	Enabled bool `yaml:"enabled" default:"false"`
}

type MailTemplateConfig struct {
	Subject  string        `yaml:"subject" default:"Example - Subject"`
	Body     string        `yaml:"body" default:"Super secret Email body"`
	Validity time.Duration `yaml:"validity" default:"15m"`
	Enabled  bool 		   `yaml:"enabled" default:"true"`
}
