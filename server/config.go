package server

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/umutozd/restaurant-backend/types"
	"github.com/umutozd/restaurant-backend/utils"
)

const (
	defaultPort       = 9001
	standartFormatter = "standard"
	jsonFormatter     = "json"
)

type Config struct {
	Debug           bool
	Port            int
	DbURL           string
	DbName          string
	LogrusFormatter string
}

func NewConfig() *Config {
	return &Config{
		Port:            defaultPort,
		LogrusFormatter: standartFormatter,
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
	validFormatters := []string{standartFormatter, jsonFormatter}
	if !utils.StringSliceContains(validFormatters, cfg.LogrusFormatter) {
		return types.Errf(types.ERR_CFG_LOGRUS_FORMATTER_INVALID, fmt.Sprintf("LogrusFormatter must be one of %s", strings.Join(validFormatters, ",")))
	}
	return nil
}

// SetFormatter sets logrus formatter. Validate should be called before this function
// because an invalid value to LogrusFormatter field will make this function no-op.
func (cfg *Config) SetFormatter() {
	switch cfg.LogrusFormatter {
	case standartFormatter:
		logrus.SetFormatter(&logrus.TextFormatter{})
	case jsonFormatter:
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
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
