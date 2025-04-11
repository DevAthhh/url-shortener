package handler

import (
	"github.com/DevAthhh/url-shortener/internal/config"
	"github.com/DevAthhh/url-shortener/internal/controllers"
	"github.com/DevAthhh/url-shortener/internal/database"
	"github.com/gin-gonic/gin"
)

func Route(cfg *config.Config, db *database.Database) *gin.Engine {
	switch cfg.Enviroment {
	case config.Production:
		gin.SetMode(gin.ReleaseMode)
	case config.Development:
		gin.SetMode(gin.DebugMode)
	}

	router := gin.New()

	router.GET("/:alias", controllers.GetController(db))
	router.POST("/", controllers.CreateController(db))

	return router
}
