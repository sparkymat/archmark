package localize

type Language string

const (
	English   Language = "en"
	Malayalam Language = "ml"
)

type StringIdentifier string

const (
	SearchPlaceholder     StringIdentifier = "SearchPlaceholder"
	NoBookmarksFound      StringIdentifier = "NoBookmarksFound"
	Bookmarks             StringIdentifier = "Bookmarks"
	AddNew                StringIdentifier = "AddNew"
	APITokens             StringIdentifier = "APITokens"
	PasteURLHere          StringIdentifier = "PasteURLHere"
	Add                   StringIdentifier = "Add"
	AddWarning            StringIdentifier = "AddWarning"
	CreateNewToken        StringIdentifier = "CreateNewToken"
	Delete                StringIdentifier = "Delete"
	Pending               StringIdentifier = "Pending"
	OpenOriginalLink      StringIdentifier = "OpenOriginalLink"
	AddedTimeStamp        StringIdentifier = "AddedTimeStamp"
	DeleteBookmarkTitle   StringIdentifier = "DeleteBookmarkTitle"
	DeleteBookmarkWarning StringIdentifier = "DeleteBookmarkWarning"
	Cancel                StringIdentifier = "Cancel"
	InternalServerError   StringIdentifier = "InternalServerError"
)
