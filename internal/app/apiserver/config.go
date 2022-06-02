package apiserver

type Config struct {
	BindAddr    string `toml:"bind_addr`
	LogLevel    string `toml:"log_level`
	Database    string `toml:database`
	DatabaseURL string `toml:"database_url`
	SessionKey  string `toml:session_key`
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8080",
		LogLevel:    "debug",
		DatabaseURL: "host=localhost port=5432 dbname=fencing sslmode=disable user=postgres password=postgres",
	}
}
