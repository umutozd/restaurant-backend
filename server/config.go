package server

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/umutozd/restaurant-backend/types"
)

const (
	defaultPort = 9001
)

type Config struct {
	Debug bool
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

// Validate validates the config object for any missing values
func (cfg *Config) Validate() error {
	if cfg.DbURL == "" {
		return types.Errf(types.ERR_CFG_DB_URL_NOT_SPECIFIED, "DbURL must be specified")
	}
	if cfg.DbName == "" {
		return types.Errf(types.ERR_CFG_DB_NAME_NOT_SPECIFIED, "DbName must be specified")
	}
	return nil
}

func (cfg *Config) GetPort() string {
	return fmt.Sprintf(":%d", cfg.Port)
}

func (cfg *Config) ToString() string {
	b, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		logrus.WithError(err).Error("error marshaling config")
		return fmt.Sprint(*cfg)
	}
	return string(b)
}
