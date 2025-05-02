package ports

type SMTPService interface {
	Send(to []string, subject string, mime string, msg string) error
	DNSLookUp(email string) error
	VerificationTemplate(apiEndpoint string) string
	RestPasswordTemplate(apiEndpoint string) string
}
