package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sparkymat/archmark/middleware"
	"github.com/sparkymat/archmark/model"
	"github.com/sparkymat/archmark/view"
)

func ShowLogin(c *gin.Context) {
	siteConfigVal, valFound := c.Get(middleware.SiteConfigurationKey)
	if !valFound {
		c.String(http.StatusInternalServerError, "config not found")
		return
	}

	siteConfig := siteConfigVal.(model.Configuration)
	pageHTML := view.Login(siteConfig.SiteName, []string{})
	html := view.Layout(siteConfig.SiteName, pageHTML)
	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, html)
}
