package configiface

import "github.com/sparkymat/archmark/localize"

type ConfigAPI interface {
	DBConnectionString() string
	AdminPassword() string
	MonolithPath() string
	DownloadPath() string
	DefaultLanguage() localize.Language
	DefaultTheme() string
}
