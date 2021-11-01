package localize

type Language string

const (
	English   Language = "en"
	Malayalam Language = "ml"
)

type StringIdentifier string

const (
	SearchPlaceholder StringIdentifier = "SearchPlaceholder"
	NoBookmarksFound  StringIdentifier = "NoBookmarksFound"
	Bookmarks         StringIdentifier = "Bookmarks"
	AddNew            StringIdentifier = "AddNew"
	APITokens         StringIdentifier = "APITokens"
)
