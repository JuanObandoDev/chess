package config

import (
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/fx"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Production bool          `envconfig:"PRODUCTION" default:"false"`
	LogLevel   zapcore.Level `envconfig:"LOG_LEVEL"  default:"info"`

	DatabaseURL      string `envconfig:"DATABASE_URL"       required:"true"`
	ListenAddr       string `envconfig:"LISTEN_ADDR"        default:"0.0.0.0:8000"`
	PublicWebAddress string `envconfig:"PUBLIC_WEB_ADDRESS" default:"http://localhost:3000"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := envconfig.Process("", cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

var Module = fx.Options(
	fx.Provide(New),
)
