package main

import (
	"context"
	"log"
	"os/exec"
	"path/filepath"
	"strconv"

	worker "github.com/contribsys/faktory_worker_go"
	"github.com/joho/godotenv"
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
			log.Printf("No pending bookmark for job %s\n", help.Jid())
			return nil
		}

		filePath := filepath.Join(cfg.DownloadPath(), bookmark.FileName)
		err = downloadPageWithMonolith(help.Jid(), cfg.MonolithPath(), bookmark.Url, filePath)
		if err != nil {
			log.Printf("Download failed for job %s\n", help.Jid())
			return nil
		}

		err = db.MarkBookmarkCompleted(bookmark.ID)
		if err != nil {
			log.Printf("Failed to mark bookmark as complete for job %s\n", help.Jid())
			return err
		}

		log.Printf("Completed job %s\n", help.Jid())
		return nil
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file. Expecting ENV to be set")
	}

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

func downloadPageWithMonolith(jobID string, monolithPath, url, filePath string) error {
	err := exec.Command(monolithPath, "-esM", url, "-o", filePath).Run()
	if err != nil {
		log.Printf("monolith download for job %s failed with: %v", jobID, err)
	}

	return err
}
