package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sparkymat/archmark/app"
	mw "github.com/sparkymat/archmark/middleware"
)

func appServices(c echo.Context) *app.Service {
	appServiceVal := c.Get(mw.AppServicesKey)
	if appServiceVal == nil {
		return nil
	}

	appService, ok := appServiceVal.(*app.Service)
	if !ok {
		return nil
	}

	return appService
}

func getCSRFToken(c echo.Context) string {
	csrfTokenVal := c.Get(middleware.DefaultCSRFConfig.ContextKey)
	if csrfTokenVal == nil {
		return ""
	}

	csrfToken, ok := csrfTokenVal.(string)
	if !ok {
		return ""
	}

	return csrfToken
}
