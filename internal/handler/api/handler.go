package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func renderError(c echo.Context, statusCode int, message string, err error) error {
	if err != nil {
		log.Errorf("err: %v", err)
	}

	//nolint:wrapcheck
	return c.JSON(statusCode, map[string]string{
		"error": message,
	})
}
