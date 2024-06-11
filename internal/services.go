package internal

type Services struct {
	User     UserService
	Features FeaturesService

	Bookmark BookmarkService
}
