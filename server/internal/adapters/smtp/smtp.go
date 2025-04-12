package smtp

type SMTPServer struct {
	Email    string
	Password string
	Host     string
	Port     string
}

type SMTP struct {
	smtp SMTPServer
}

func NewSMTP(
	email string,
	password string,
	host string,
	port string,
) *SMTP {
	return &SMTP{
		smtp: SMTPServer{
			Email:    email,
			Password: password,
			Host:     host,
			Port:     port,
		},
	}
}
