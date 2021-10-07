package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sparkymat/archmark/config"
	"github.com/sparkymat/archmark/model"
	"gorm.io/gorm"
)

func SetupRedirect(cfg config.API, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/setup" {
			c.Next()
			return
		}

		var user model.User

		result := db.Where("username = 'admin'").First(&user)

		if result.RowsAffected == 0 { // no results
			c.Redirect(http.StatusSeeOther, "/setup")
			return
		}

		c.Next()
	}
}
