package redis

type Config struct {
	Username  string	`yaml:"username" default:"username"`
	Password  string	`yaml:"password" default:"password"`
	URL       string	`yaml:"url" default:"redis://username:password@localhost:6379/0"`
	Port      string	`yaml:"port" default:"6379"`
	Host      string	`yaml:"host" default:"localhost"`
	DefaultDB string	`yaml:"db" default:"0"`
}
