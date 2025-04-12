package ports

type SMTPService interface {
	Send(to []string, subject string, mime string, msg string) error
	VerificationTemplate(apiEndpoint string) string
	RestPasswordTemplate(apiEndpoint string) string
}
