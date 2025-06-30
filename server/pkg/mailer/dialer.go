package mailer

import (
	"crypto/tls"
	"strings"

	"gopkg.in/mail.v2"
)

type Config struct {
	SMTPHost     string `yaml:"host" default:"localhost"`                 
	SMTPPort     int    `yaml:"port" default:"578"`     
	SMTPUsername string `yaml:"username" default:"example@example.com"`               
	SMTPPassword string `yaml:"password" default:"password"`           
	SMTPInsecure bool   `yaml:"insecure" default:"true"`
	Headers      map[string]string `yaml:"headers" default:"{}"`

	// SMTP TLS policy to use when establishing a connection.
	// Defaults to MandatoryStartTLS.
	// Possible values are:
	// opportunistic: Use STARTTLS if the server supports it, otherwise connect without encryption.
	// mandatory: Always use STARTTLS.
	// none: Never use STARTTLS.
	SMTPTLSPolicy string `yaml:"tls_policy" default:"mandatory"`
}

func (c *Config) TLSPolicy() mail.StartTLSPolicy {
	switch strings.ToLower(c.SMTPTLSPolicy) {
	case "opportunistic":
		return mail.OpportunisticStartTLS
	case "mandatory":
		return mail.MandatoryStartTLS
	case "none":
		return mail.NoStartTLS
	}
	return mail.MandatoryStartTLS
}

type Message = mail.Message
type MessageSetting =  mail.MessageSetting
type Dialer = *mail.Dialer

type dialerImpl struct {
	dialer  Dialer
	headers map[string]string
}

func NewDialer(cfg Config) *dialerImpl {
	
	d := mail.NewDialer(cfg.SMTPHost, cfg.SMTPPort, cfg.SMTPUsername, cfg.SMTPPassword)
	
	d.TLSConfig = &tls.Config{
		InsecureSkipVerify: cfg.SMTPInsecure,
		ServerName:         cfg.SMTPHost,
	}

	d.StartTLSPolicy = cfg.TLSPolicy()
	
	return &dialerImpl{
		dialer:  d,
		headers: cfg.Headers,
	}
}

func (m dialerImpl) FromHeader() string {
	
	if _, ok := m.headers["from"]; !ok {
		return m.dialer.Username
	}

	return m.headers["from"]
}

func (m dialerImpl) DialAndSend(msg *Message) error {
	return m.dialer.DialAndSend(msg)
}

func (m dialerImpl) NewMessage(settings ...MessageSetting) *Message {
  return mail.NewMessage(settings...)
}