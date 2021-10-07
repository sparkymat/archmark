package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sparkymat/archmark/config"
	"gorm.io/gorm"
)

const (
	ConfigKey = "config"
	DBKey     = "db"
)

func ConfigInjector(cfg config.API, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(ConfigKey, cfg)
		c.Set(DBKey, db)
		c.Next()
	}
}
