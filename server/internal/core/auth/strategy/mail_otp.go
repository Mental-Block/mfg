package strategy

import (
	"bytes"
	"fmt"
	"html/template"
	"time"

	"github.com/server/pkg/crypt"
)

// MailOTP sends a mail with a one time password to user's email id
// and verifies the OTP. On successful verification, it creates a session
type MailOTP struct {
	dialer  IEmailService
	subject string
	body    string
	Now     func() time.Time
}

func NewMailOTP(d IEmailService, subject, body  string) *MailOTP {
	return &MailOTP{
    	dialer:  d,
		subject: subject,
		body: body,
		Now: func() time.Time {
			return time.Now().UTC()
		},
	}
}

// SendMail sends a mail with a one time password embedded link to user's email id
func (m MailOTP) SendMail(id, to string) (string, error) {
	var otp string

	otp = crypt.GenerateRandomStringFromLetters(otpLen, otpLetterRunes)

	tpl := template.New("body")
	
	t, err := tpl.Parse(m.body)
	
	if err != nil {
		return "", fmt.Errorf("failed to parse email template: %w", err)
	}
	
	var tplBuffer bytes.Buffer
	
	if err = t.Execute(&tplBuffer, map[string]string{
		"Otp": otp,
	}); err != nil {
		return "", fmt.Errorf("failed to parse email template: %w", err)
	}
	
	tplBody := tplBuffer.String()

	tpl = template.New("sub")
	
	t, err = tpl.Parse(m.subject)
	
	if err != nil {
		return "", fmt.Errorf("failed to parse email template: %w", err)
	}
	
	tplBuffer.Reset()
	
	if err = t.Execute(&tplBuffer, map[string]string{
		"Otp":   otp,
		"Email": to,
	}); err != nil {
		return "", fmt.Errorf("failed to parse email template: %w", err)
	}

	tplSub := tplBuffer.String()

	msg := m.dialer.NewMessage()
	msg.SetHeader("From", m.dialer.FromHeader())
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", tplSub)
	msg.SetBody("text/html", tplBody)
	msg.SetDateHeader("Date", m.Now())

	return otp, m.dialer.DialAndSend(msg)
}

func DefaultOTPTemplate(reason string, s MailTemplateConfig) MailTemplateConfig {
	subject :=  fmt.Sprintf("Subject - %s", reason)
	body := fmt.Sprintf(`
			Please copy/paste the OneTimePassword in OTP form.
			<h2>{{.Otp}}</h2>
			This code will expire in %v minutes.
	`, DefaultOLTPDuration) 
	
	return MailTemplateConfig{
		Body: body,
		Subject: subject,
		Validity: DefaultOLTPDuration,
	}
}
