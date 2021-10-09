package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sparkymat/archmark/database"
	"github.com/sparkymat/archmark/middleware"
	"github.com/sparkymat/archmark/model"
	"github.com/sparkymat/archmark/view"
)

var (
	ErrSiteNameBlank      = errors.New("site name can't be blank")
	ErrPasswordTooShort   = errors.New("admin password must be at least 8 characters")
	ErrPasswordsDontMatch = errors.New("admin passwords don't match")
)

type setupInput struct {
	SiteName                  string `form:"site_name"`
	AdminPassword             string `form:"admin_password"`
	AdminPasswordConfirmation string `form:"admin_password_confirmation"`
}

func (i *setupInput) Validate() []error {
	var errors []error
	if i.SiteName == "" {
		errors = append(errors, ErrSiteNameBlank)
	}
	if len(i.AdminPassword) < 8 {
		errors = append(errors, ErrPasswordTooShort)
	} else if i.AdminPassword != i.AdminPasswordConfirmation {
		errors = append(errors, ErrPasswordsDontMatch)
	}
	return errors
}

func ShowSetup(c *gin.Context) {
	renderSetupPage(c, []string{})
}

func DoSetup(c *gin.Context) {
	var input setupInput
	c.Bind(&input)

	errors := input.Validate()
	if len(errors) > 0 {
		errorStrings := []string{}
		for _, err := range errors {
			errorStrings = append(errorStrings, err.Error())
		}
		renderSetupPage(c, errorStrings)
		return
	}

	dbVal, valFound := c.Get(middleware.DBKey)
	if !valFound {
		c.String(http.StatusInternalServerError, "db connection not found")
		return
	}
	db := dbVal.(database.API)

	_, err := db.CreateAdminUser(input.AdminPassword)
	if err != nil {
		c.String(http.StatusInternalServerError, "admin creation failed")
		return
	}

	c.String(http.StatusOK, "done")
}

func renderSetupPage(c *gin.Context, errors []string) {
	siteConfigVal, valFound := c.Get(middleware.SiteConfigurationKey)
	if !valFound {
		c.String(http.StatusInternalServerError, "config not found")
		return
	}
	siteConfig := siteConfigVal.(model.Configuration)

	pageHTML := view.Setup(siteConfig.SiteName, errors)
	html := view.Layout(siteConfig.SiteName, pageHTML)
	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, html)
}
