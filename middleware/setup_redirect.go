package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sparkymat/archmark/config"
	"github.com/sparkymat/archmark/database"
)

func SetupRedirect(cfg config.API, db database.API) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/setup" {
			c.Next()
			return
		}

		_, err := db.LoadAdminUser()
		if err != nil {
			c.Redirect(http.StatusSeeOther, "/setup")
			return
		}

		c.Next()
	}
}
