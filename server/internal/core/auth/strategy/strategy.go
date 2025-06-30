package strategy

import (
	"time"

	"github.com/server/internal/core/auth/domain"
	"github.com/server/pkg/mailer"
)

type IEmailService interface {
    DialAndSend(m *mailer.Message) error
    FromHeader() string
    NewMessage(settings ...mailer.MessageSetting) *mailer.Message
}

var (
	otpLetterRunes = []rune("ABCDEFGHJKMNPQRSTWXYZ23456789")
	otpLen         = 6
	MaxOTPAttempt  = 3
	OtpAttemptKey  = "attempt"
)

const (
	PasswordStrategy domain.Strategy = "password"
	LinkStrategy 	 domain.Strategy = "link"
	OTPStrategy  	 domain.Strategy = "otp"
	OIDCStrategy     domain.Strategy = "oicd"
)

var (
	ForgotPasswordReason domain.Reason = "forgotpassword"
	LoginReason domain.Reason = "login"
	RegisterReason domain.Reason = "register"
)

var (
	AccountForgotPasswordEmailDuration = time.Minute * 10
	AccountVerificationEmailDuration = time.Hour * 2
	AccountLoginEmailDuration = time.Minute * 15
	DefaultOLTPDuration = time.Minute * 15
	DefaultMajicLinkDuration = time.Minute * 15
)