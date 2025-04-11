package logger

import (
	"github.com/DevAthhh/url-shortener/internal/config"
	"go.uber.org/zap"
)

var (
	Logger *zap.Logger
)

func LoadLogger(cfg *config.Config) error {
	env := cfg.Enviroment

	var logger *zap.Logger
	var err error

	switch env {
	case config.Development:
		logger, err = zap.NewDevelopment()
	case config.Production:
		logger, err = zap.NewProduction()
	}

	if err != nil {
		return err
	}

	Logger = logger
	return nil
}
