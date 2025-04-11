package handler

import (
	"net/http"

	"github.com/DevAthhh/url-shortener/internal/config"
	"github.com/DevAthhh/url-shortener/internal/database"
)

type SetGetURL interface {
	SaveURL(url string, size int) (string, error)
	GetUrl(alias string) (string, error)
}

type Server struct {
	server *http.Server
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

func NewServer(cfg *config.Config, db *database.Database) *Server {
	return &Server{
		server: &http.Server{
			Addr:         cfg.Server.Host + ":" + cfg.Server.Port,
			Handler:      Route(cfg, db),
			ReadTimeout:  cfg.Server.Timeout,
			WriteTimeout: cfg.Server.Timeout,
			IdleTimeout:  cfg.Server.IdleTimeout,
		},
	}
}
