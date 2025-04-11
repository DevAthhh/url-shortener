package config

import (
	"time"
)

const (
	Production  = "prod"
	Development = "dev"
)

type Config struct {
	Enviroment string     `json:"env"`
	Server     HttpServer `json:"http_server"`
}

type HttpServer struct {
	Port        string        `json:"port"`
	Host        string        `json:"host"`
	Timeout     time.Duration `json:"timeout"`
	IdleTimeout time.Duration `json:"idle_timeout"`
}
