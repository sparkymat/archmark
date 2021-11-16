package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	mw "github.com/sparkymat/archmark/middleware"
)

func appServices(c echo.Context) *mw.AppServices {
	appServicesVal := c.Get(mw.AppServicesKey)
	if appServicesVal == nil {
		return nil
	}

	appServices, ok := appServicesVal.(*mw.AppServices)
	if !ok {
		return nil
	}

	return appServices
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
