package route

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	return cv.validator.Struct(i) //nolint:wrapcheck
}

func Setup(e *echo.Echo, cfg ConfigService, db DatabaseService) {
	e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("/health", func(c echo.Context) error {
		//nolint:wrapcheck
		return c.String(http.StatusOK, "Hello, World!")
	})

	app := e.Group("")
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] Got ${method} on ${uri} from ${remote_ip}. Responded with ${status} in ${latency_human}.\n",
	}))
	app.Use(middleware.Recover())

	app.Static("/js", "public/js")
	app.Static("/css", "public/css")
	app.Static("/images", "public/images")
	app.Static("/fonts", "public/fonts")
	app.Static("/uploads", cfg.DownloadPath())

	app.Use(session.Middleware(sessions.NewCookieStore([]byte(cfg.SessionSecret()))))

	registerAPIRoutes(app, cfg, db)
	registerWebRoutes(app, cfg, db)
}
