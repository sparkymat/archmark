package jobs

import (
	"context"
	"log"
)

func DeleteExpiredBookmarks(cfg ConfigService, db DatabaseService) JobFunc {
	return func() {
		err := db.DeleteBookmarks(context.Background(), cfg.DeleteTimerHours())
		if err != nil {
			log.Printf("DeleteExpiredBookmarks failed at db.DeleteBookmarks. err: %v", err)
		}
	}
}
