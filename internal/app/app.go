package app

import (
	"github.com/DevAthhh/url-shortener/internal/database"
	"github.com/DevAthhh/url-shortener/internal/handler"
	"github.com/DevAthhh/url-shortener/internal/initializers"
	"github.com/DevAthhh/url-shortener/internal/lib/logger"
	"go.uber.org/zap/zapcore"
)

func Run() {
	initializers.LoadEnv()
	cfg := initializers.LoadConfig()
	logger.LoadLogger(cfg)
	db := database.LoadDatabase()

	logger.Logger.Info("the program has started working")

	server := handler.NewServer(cfg, db)

	if err := server.Start(); err != nil {
		logger.Logger.Fatal("the server was not running", zapcore.Field{String: err.Error()})
	}
}
