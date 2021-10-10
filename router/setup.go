package router

import (
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/sparkymat/archmark/config"
	"github.com/sparkymat/archmark/database"
	"github.com/sparkymat/archmark/internal/handler"
	"github.com/sparkymat/archmark/middleware"
	"github.com/sparkymat/archmark/model"
)

type loginInput struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func Setup(r *gin.Engine, cfg config.API, db database.API, siteConfig model.Configuration) {
	app := r.Group("")

	app.Use(middleware.ConfigInjector(cfg, db, siteConfig))
	app.Use(middleware.SetupRedirect(cfg, db))

	app.GET("/setup", handler.ShowSetup)
	app.POST("/setup", handler.DoSetup)
	app.POST("/add", handler.Create)

	r.Static("/css", "public/css")
	r.Static("/javascript", "public/javascript")

	authApp := r.Group("")
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:      "archmark",
		Key:        []byte(cfg.JWTSecret()),
		Timeout:    time.Hour * 4,
		MaxRefresh: time.Hour * 24 * 365,
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals loginInput
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.Username
			password := loginVals.Password

			user, err := db.FindUser(userID)
			if err != nil {
				return "", jwt.ErrFailedAuthentication
			}

			if user.ValidatePassword(password) == nil {
				return &user, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Unauthorized: func(c *gin.Context, statusCode int, message string) {
			c.Redirect(http.StatusSeeOther, "/login")
		},
		LoginResponse: func(c *gin.Context, statusCode int, message string, loginTime time.Time) {
			c.Redirect(http.StatusSeeOther, "/")
		},
	})
	if err != nil {
		panic(err)
	}
	app.GET("/login", handler.ShowLogin)
	app.POST("/login", authMiddleware.LoginHandler)
	app.GET("/refresh_token", authMiddleware.RefreshHandler)
	authApp.Use(authMiddleware.MiddlewareFunc())
	authApp.GET("/", handler.Home)
}
