package strategy

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
	"time"

	"github.com/server/pkg/crypt"
)

type MailLink struct {
	dialer  IEmailService
	subject string
	body    string
	host    string
	returnhost string
	Now     func() time.Time
}

// MailLink sends a mail with a one time password link to user's email id.
// On successful verification, it creates a session
func NewMailLink(d IEmailService, host, returnHost, subject, body  string) *MailLink {
	return &MailLink{
    	dialer:  d,
		subject: subject,
		body: body,
		host: host,
		returnhost: returnHost,
		Now: func() time.Time {
			return time.Now().UTC()
		},
	}
}

// SendMail sends a mail with a one time password embedded link to user's email id
func (m MailLink) SendMail(id, to string) (string, error) {
	
	otp := crypt.GenerateRandomStringFromLetters(20, otpLetterRunes)

	t, err := template.New("body").Parse(m.body)
  
	if err != nil {
		return "", fmt.Errorf("failed to parse email template: %w", err)
	}

	var tpl bytes.Buffer

	link := fmt.Sprintf("%s?strategy_name=%s&code=%s&state=%s", strings.TrimRight(m.host, "/"), LinkStrategy, otp, id)
	
  	err = t.Execute(&tpl, map[string]string{
		"Link": link,
		"RLink":  m.returnhost,
	})

	if err != nil {
		return "", fmt.Errorf("failed to parse email template: %w", err)
	}
  
	msg := m.dialer.NewMessage()
	msg.SetHeader("From", m.dialer.FromHeader())
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", m.subject)
	msg.SetBody("text/html", tpl.String())
	msg.SetDateHeader("Date", m.Now())

	return otp, m.dialer.DialAndSend(msg)
}

func DefaultMajicLinkTemplate(reason string, s MailTemplateConfig) MailTemplateConfig {
	subject :=  fmt.Sprintf("Subject - %s", "")
	body := fmt.Sprintf(`
		Click on the following link or copy/paste the url in browser.
		<br>
		<h2>
			<a href='{{.Link}}' target='_blank'>%v</a>
		</h2>
		<br>
		Address: {{.Link}} 
		<br>This link will expire in %v minutes.
	`, reason, DefaultMajicLinkDuration) 
	
	return MailTemplateConfig{
		Body: body,
		Subject: subject,
		Validity: DefaultMajicLinkDuration,
	}
}
