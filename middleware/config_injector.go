package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sparkymat/archmark/config"
	"github.com/sparkymat/archmark/model"
	"gorm.io/gorm"
)

const (
	ConfigKey            = "config"
	DBKey                = "db"
	SiteConfigurationKey = "site_configuration"
)

func ConfigInjector(cfg config.API, db *gorm.DB, siteConfig model.Configuration) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(ConfigKey, cfg)
		c.Set(DBKey, db)
		c.Set(SiteConfigurationKey, siteConfig)
		c.Next()
	}
}
