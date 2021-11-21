package localize

type Language string

const (
	English   Language = "en"
	Malayalam Language = "ml"
	German    Language = "de"
)

var SupportedLanguages = []Language{ //nolint:gochecknoglobals
	English,
	Malayalam,
	German,
}

func LanguageFromString(langString string) Language {
	switch langString {
	case string(English):
		return English
	case string(Malayalam):
		return Malayalam
	case string(German):
		return German
	default:
		return English
	}
}

type StringIdentifier string

const (
	ActiveTokens          StringIdentifier = "ActiveTokens"
	APITokens             StringIdentifier = "APITokens"
	Add                   StringIdentifier = "Add"
	AddNew                StringIdentifier = "AddNew"
	AddWarning            StringIdentifier = "AddWarning"
	AddedTimeStamp        StringIdentifier = "AddedTimeStamp"
	Bookmarks             StringIdentifier = "Bookmarks"
	Cancel                StringIdentifier = "Cancel"
	CreateNewToken        StringIdentifier = "CreateNewToken"
	Delete                StringIdentifier = "Delete"
	DeleteBookmarkTitle   StringIdentifier = "DeleteBookmarkTitle"
	DeleteBookmarkWarning StringIdentifier = "DeleteBookmarkWarning"
	InternalServerError   StringIdentifier = "InternalServerError"
	LanguageLabel         StringIdentifier = "LanguageLabel"
	NoBookmarksFound      StringIdentifier = "NoBookmarksFound"
	OpenOriginalLink      StringIdentifier = "OpenOriginalLink"
	PasteURLHere          StringIdentifier = "PasteURLHere"
	Pending               StringIdentifier = "Pending"
	SaveSettings          StringIdentifier = "SaveSettings"
	SearchPlaceholder     StringIdentifier = "SearchPlaceholder"
	SelectLanguage        StringIdentifier = "SelectLanguage"
	Settings              StringIdentifier = "Settings"
)
