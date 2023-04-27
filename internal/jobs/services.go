package jobs

import "context"

type ConfigService interface {
	DeleteTimerHours() int32
}

type DatabaseService interface {
	DeleteBookmarks(ctx context.Context, agehours int32) error
}
