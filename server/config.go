package server

import "fmt"

const (
	defaultPort = 9001
)

type Config struct {
	// The port to listen to
	Port   int
	DbURL  string
	DbName string
}

func NewConfig() *Config {
	return &Config{
		Port: defaultPort,
	}
}

func (cfg *Config) GetPort() string {
	return fmt.Sprintf(":%d", cfg.Port)
}
