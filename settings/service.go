package settings

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/localize"
)

type API interface {
	Language(c echo.Context) localize.Language
	SetLanguage(c echo.Context, lang localize.Language)
}

func New() API {
	return &service{}
}

type service struct{}

func (*service) Language(c echo.Context) localize.Language {
	langString, err := readValueFromCookie(c, "language")
	if err != nil {
		return localize.LanguageFromString("") // Default language
	}

	return localize.LanguageFromString(langString)
}

func (*service) SetLanguage(c echo.Context, lang localize.Language) {
	writeValueToCookie(c, "language", string(lang))
}

func writeValueToCookie(c echo.Context, key, value string) {
	cookie := new(http.Cookie)
	cookie.Name = key
	cookie.Value = value
	c.SetCookie(cookie)
}

func readValueFromCookie(c echo.Context, key string) (string, error) {
	cookie, err := c.Cookie(key)
	if err != nil {
		return "", fmt.Errorf("failed to get cookie. err: %w", err)
	}

	return cookie.Value, nil
}
