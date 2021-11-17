package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

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
