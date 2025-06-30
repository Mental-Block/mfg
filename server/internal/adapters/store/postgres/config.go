package postgres

type Config struct {
	Username  string	`yaml:"username" default:"username"`
	Password  string	`yaml:"password" default:"password"`
	URL       string	`yaml:"url" default:"postgres://username:password@postgres:5432/db?sslmode=disable"`
	Port      string	`yaml:"port" default:"5432"`
	Host      string	`yaml:"host" default:"localhost"`
	DefaultDB string	`yaml:"db" default:"db"`
}
