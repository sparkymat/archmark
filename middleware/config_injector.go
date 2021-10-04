package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sparkymat/archmark/config"
)

const ConfigKey = "config"

func ConfigInjector(cfg config.API) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(ConfigKey, cfg)
		c.Next()
	}
}
