package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf("%d", time.Now().Unix()))
}
