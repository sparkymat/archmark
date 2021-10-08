package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sparkymat/archmark/middleware"
	"github.com/sparkymat/archmark/model"
	"github.com/sparkymat/archmark/view"
)

func Setup(c *gin.Context) {
	siteConfigVal, valFound := c.Get(middleware.SiteConfigurationKey)
	if !valFound {
		c.String(http.StatusInternalServerError, "config not found")
		return
	}
	siteConfig := siteConfigVal.(model.Configuration)

	pageHTML := view.Setup(siteConfig.SiteName, []string{"error1", "error2"})
	html := view.Layout(siteConfig.SiteName, pageHTML)
	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, html)
}
