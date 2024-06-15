package route

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sparkymat/archmark/internal"
	"github.com/sparkymat/archmark/internal/auth"
	"github.com/sparkymat/archmark/internal/handler/api"
)

func registerAPIRoutes(app *echo.Group, cfg ConfigService, services internal.Services) {
	apiGroup := app.Group("api")

	if cfg.ReverseProxyAuthentication() {
		apiGroup.Use(auth.ProxyAuthMiddleware(cfg, services.User))
	} else {
		apiGroup.Use(auth.APIMiddleware(cfg, services.User))
	}

	apiGroup.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:X-CSRF-Token",
	}))

	apiGroup.POST("/bookmarks", api.BookmarksCreate(services))
	apiGroup.GET("/bookmarks/search", api.BookmarksSearch(services))
	apiGroup.GET("/bookmarks/recent", api.BookmarksFetchRecent(services))
	apiGroup.GET("/bookmarks/:id", api.BookmarksShow(services))
	apiGroup.DELETE("/bookmarks/:id", api.BookmarksDestroy(services))
	apiGroup.PATCH("/bookmarks/:id/update_title", api.BookmarksUpdateTitle(services))
	apiGroup.PATCH("/bookmarks/:id/update_url", api.BookmarksUpdateUrl(services))
	apiGroup.PATCH("/bookmarks/:id/update_description", api.BookmarksUpdateDescription(services))
	apiGroup.PATCH("/bookmarks/:id/update_transcript", api.BookmarksUpdateTranscript(services))
	apiGroup.PATCH("/bookmarks/:id/upload_image", api.BookmarksUploadImage(services))

	apiGroup.POST("/analysis_notes", api.AnalysisNotesCreate(services))
	apiGroup.GET("/bookmarks/:parent_id/analysis_notes/recent", api.AnalysisNotesFetchRecent(services))
	apiGroup.GET("/analysis_notes/:id", api.AnalysisNotesShow(services))
	apiGroup.DELETE("/analysis_notes/:id", api.AnalysisNotesDestroy(services))
	apiGroup.PATCH("/analysis_notes/:id/update_content", api.AnalysisNotesUpdateContent(services))
}
