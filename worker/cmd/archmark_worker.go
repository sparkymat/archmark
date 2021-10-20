package main

import (
	"context"
	"log"
	"strconv"

	worker "github.com/contribsys/faktory_worker_go"
	"github.com/sparkymat/archmark/archive"
	"github.com/sparkymat/archmark/config"
	"github.com/sparkymat/archmark/database"
)

func saveWebPage(cfg config.API, db database.API) func(ctx context.Context, args ...interface{}) error {
	return func(ctx context.Context, args ...interface{}) error {
		help := worker.HelperFor(ctx)
		log.Printf("Working on job %s\n", help.Jid())

		if len(args) == 0 {
			log.Printf("No params found for job %s\n", help.Jid())
			return nil
		}

		bookmarkIDString, isString := args[0].(string)
		if !isString {
			log.Printf("Non-string param passed for job %s\n", help.Jid())
			return nil
		}

		bookmarkID, err := strconv.ParseUint(bookmarkIDString, 10, 64)
		if err != nil {
			log.Printf("Non-ID param found for job %s\n", help.Jid())
			return nil
		}

		bookmark, err := db.FindBookmark(uint(bookmarkID))
		if err != nil {
			log.Printf("Unable to load bookmark for job %s\n", help.Jid())
			return err
		}

		if bookmark.Status != "pending" {
			log.Printf("No penidng bookmark for job %s\n", help.Jid())
			return nil
		}

		archiveAPI := archive.New(archive.Config{
			MonolithPath:   cfg.MonolithPath(),
			DownloadFolder: cfg.DownloadPath(),
		})

		return nil
	}
}

func main() {
	cfg := config.New()
	db := database.New(database.Config{
		ConnectionString: cfg.DBConnectionString(),
	})

	mgr := worker.NewManager()
	mgr.Register("SaveWebPage", saveWebPage(cfg, db))
	mgr.Concurrency = 5
	mgr.ProcessStrictPriorityQueues("critical", "default")
	mgr.Run()
}
