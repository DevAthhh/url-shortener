package initializers

import (
	"log"
	"os"

	"github.com/DevAthhh/url-shortener/internal/config"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

func LoadConfig() *config.Config {
	path := os.Getenv("PATH_TO_CONFIG")
	if path == "" {
		log.Fatal("the path to the configuration has not been declared")
	}

	if _, err := os.Stat(path); err != nil {
		log.Fatal("the configuration file was not found")
	}

	viper.AddConfigPath(os.Getenv("PATH_TO_CONFIG"))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	cfg := config.Config{
		Enviroment: viper.GetString("env"),
		Server: config.HttpServer{
			Port:        viper.GetString("http_server.port"),
			Host:        viper.GetString("http_server.host"),
			Timeout:     viper.GetDuration("http_server.timeout"),
			IdleTimeout: viper.GetDuration("http_server.idle_timeout"),
		},
	}

	return &cfg
}
