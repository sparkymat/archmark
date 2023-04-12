package api

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/sparkymat/archmark/auth"
	"github.com/sparkymat/archmark/dbx"
)

func CategoriesList(_ ConfigService, db DatabaseService) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, isUser := c.Get(auth.UserKey).(dbx.User)
		if !isUser {
			return renderError(c, http.StatusInternalServerError, "failed to load user", nil)
		}

		categories, err := db.FetchCategories(c.Request().Context(), user.ID)
		if err != nil {
			return renderError(c, http.StatusInternalServerError, "failed to load categories", err)
		}

		validCategories := lo.Filter(categories, func(ps pgtype.Text, i int) bool { return ps.Valid })
		categoryStrings := lo.Map(validCategories, func(ps pgtype.Text, i int) string { return ps.String })

		return c.JSON(http.StatusOK, categoryStrings)
	}
}
