package controllers

import (
	"net/http"

	"github.com/DevAthhh/url-shortener/internal/lib/logger"
	"github.com/DevAthhh/url-shortener/internal/lib/transport"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UrlSaver interface {
	SaveURL(url string, size int) (string, error)
}

type UrlGetter interface {
	GetUrl(alias string) (string, error)
}

func CreateController(saver UrlSaver) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req transport.RequestToSave
		if err := c.BindJSON(&req); err != nil {
			logger.Logger.Debug("error with getting root url from req", zap.Error(err))

			c.JSON(http.StatusBadRequest, gin.H{
				"error": "err with getting root url from request",
			})
			return
		}

		if err := req.Validate(); err != nil {
			logger.Logger.Debug("size or url invalid", zap.Error(err))

			c.JSON(http.StatusBadRequest, gin.H{
				"error": "size or url invalid",
			})
			return
		}

		alias, err := saver.SaveURL(req.Root, req.Size)
		if err != nil {
			logger.Logger.Debug("error with saving to db", zap.Error(err))

			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error with saving to db",
			})
			return
		}

		response := transport.ResponseFromSave{
			Status: "url has been saved to DB",
			Alias:  alias,
		}

		c.JSON(http.StatusOK, response)
	}
}

func GetController(urlGetter UrlGetter) func(c *gin.Context) {
	return func(c *gin.Context) {
		alias := c.Param("alias")
		url, err := urlGetter.GetUrl(alias)
		if err != nil {
			logger.Logger.Debug("error with getting alias from request", zap.Error(err))

			c.JSON(http.StatusBadRequest, gin.H{
				"error": "err with getting alias from request",
			})
			return
		}

		c.Redirect(http.StatusFound, url)
	}
}
