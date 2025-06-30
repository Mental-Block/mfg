package bootstrap

type Config struct {
	// Useful for when we first bootstrap the app... first setting up when there is no production data or within testing  
	// Once you are done user should be delivated to a normal user
	SuperUsers []string `yaml:"super_users"`
}
