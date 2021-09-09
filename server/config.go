package server

const (
	defaultPort = 9001
)

type Config struct {
	// The port to listen to
	Port int
}

func NewConfig() *Config {
	return &Config{
		Port: defaultPort,
	}
}
