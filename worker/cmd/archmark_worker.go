package main

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	worker "github.com/contribsys/faktory_worker_go"
	"github.com/google/uuid"
	"github.com/sparkymat/archmark/archive"
	"github.com/sparkymat/archmark/database"
	"github.com/sparkymat/archmark/dbx"
	"github.com/sparkymat/archmark/internal/config"
)

const (
	base10        = 10
	sixtyFourBits = 64
)

type DatabaseService interface {
	FetchBookmarkByID(ctx context.Context, id int64) (dbx.Bookmark, error)
	MarkBookmarkFetched(ctx context.Context, arg dbx.MarkBookmarkFetchedParams) error
	UpdateBookmarkDetails(ctx context.Context, arg dbx.UpdateBookmarkDetailsParams) error
}

type ConfigService interface {
	DownloadPath() string
	MonolithPath() string
}

//nolint:funlen,revive
func saveWebPage(cfg ConfigService, db DatabaseService) func(ctx context.Context, args ...interface{}) error {
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

		bookmarkID, err := strconv.ParseInt(bookmarkIDString, base10, sixtyFourBits)
		if err != nil {
			log.Printf("Non-ID param found for job %s with err: %v\n", help.Jid(), err)

			return nil
		}

		bookmark, err := db.FetchBookmarkByID(ctx, bookmarkID)
		if err != nil {
			log.Printf("Unable to load bookmark for job %s\n", help.Jid())

			return fmt.Errorf("failed to find bookmark. err: %w", err)
		}

		if bookmark.Status != "pending" {
			log.Printf("No pending bookmark for job %s\n", help.Jid())

			return nil
		}

		// FETCH DETAILS
		archiver := archive.New()

		pageInfo, err := archiver.FetchDetails(ctx, bookmark.Url)
		if err != nil {
			log.Printf("Failed to fetch page details for job %s\n", help.Jid())

			return fmt.Errorf("failed to fetch page details. err: %w", err)
		}

		err = db.UpdateBookmarkDetails(ctx, dbx.UpdateBookmarkDetailsParams{
			ID:    bookmark.ID,
			Title: pageInfo.Title,
			Html:  pageInfo.HTMLContent,
		})
		if err != nil {
			log.Printf("Failed to update bookmark details for job %s\n", help.Jid())

			return fmt.Errorf("failed to update bookmark details. err: %w", err)
		}

		// DOWNLOAD
		fileName := strings.ReplaceAll(uuid.New().String(), "-", "")
		fileName = fmt.Sprintf("%s.html", fileName)
		filePath := filepath.Join(cfg.DownloadPath(), fileName)

		if err = downloadPageWithMonolith(help.Jid(), cfg.MonolithPath(), bookmark.Url, filePath); err != nil {
			log.Printf("Download failed for job %s with err: %v\n", help.Jid(), err)

			return nil
		}

		err = db.MarkBookmarkFetched(ctx, dbx.MarkBookmarkFetchedParams{
			ID:       bookmark.ID,
			FilePath: fileName,
		})
		if err != nil {
			log.Printf("Failed to mark bookmark as complete for job %s\n", help.Jid())

			return fmt.Errorf("failed to mark bookmark as completed. err: %w", err)
		}

		log.Printf("Completed job %s\n", help.Jid())

		return nil
	}
}

func main() {
	appConfig, err := config.New()
	if err != nil {
		panic(err)
	}

	dbDriver, err := database.New(appConfig.DatabaseURL())
	if err != nil {
		panic(err)
	}

	db := dbx.New(dbDriver.DB())

	mgr := worker.NewManager()
	mgr.Register("SaveWebPage", saveWebPage(appConfig, db))
	mgr.Concurrency = 5
	mgr.ProcessStrictPriorityQueues("critical", "default")
	mgr.Run() //nolint:errcheck
}

func downloadPageWithMonolith(jobID string, monolithPath, url, filePath string) error {
	err := exec.Command(monolithPath, "-esM", url, "-o", filePath).Run()
	if err != nil {
		log.Printf("monolith download for job %s failed with: %v", jobID, err)
	}

	if err != nil {
		return fmt.Errorf("download failed. err: %w", err)
	}

	return nil
}
