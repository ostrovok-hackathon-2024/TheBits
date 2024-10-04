package ostrovok

type Config struct {
	Auth     AuthConfig
	APIKey   string
	RedisURL string
	APIURL   string
}

type AuthConfig struct {
	Username string
	Password string
}
