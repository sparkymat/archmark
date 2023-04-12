package route

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sparkymat/archmark/auth"
	"github.com/sparkymat/archmark/internal/handler/api"
)

func registerAPIRoutes(app *echo.Group, cfg ConfigService, db DatabaseService) {
	apiGroup := app.Group("api")

	if cfg.ReverseProxyAuthentication() {
		apiGroup.Use(auth.ProxyAuthMiddleware(cfg, db))
	} else {
		apiGroup.Use(auth.APIMiddleware(cfg, db))
	}

	apiGroup.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:X-CSRF-Token",
	}))

	apiGroup.GET("/bookmarks", api.BookmarksList(cfg, db))
	apiGroup.POST("/bookmarks", api.BookmarksCreate(cfg, db))
	apiGroup.GET("/bookmarks/search", api.BookmarksSearch(cfg, db))
	apiGroup.GET("/categories", api.CategoriesList(cfg, db))
	apiGroup.POST("/bookmarks/:id/update_category", api.BookmarksUpdateCategory(cfg, db))
}
