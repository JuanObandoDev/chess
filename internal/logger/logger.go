package logger

import (
	"github.com/sanpezlo/chess/internal/config"
	"github.com/sanpezlo/chess/internal/version"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(cfg *config.Config) (*zap.Logger, error) {
	var config zap.Config
	if cfg.Production {
		config = zap.NewProductionConfig()
		config.InitialFields = map[string]interface{}{"v": version.Version}
	} else {
		config = zap.NewDevelopmentConfig()
	}

	config.Level.SetLevel(cfg.LogLevel)
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}

var Module = fx.Options(
	fx.Provide(New),
	fx.Invoke(func(l *zap.Logger) {
		zap.ReplaceGlobals(l)
	}),
)
