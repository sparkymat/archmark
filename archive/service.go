package archive

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var ErrUnableToCreateFile = errors.New("unable to create file")

type Config struct {
	DownloadFolder string
}

type ArchivedPage struct {
	Title       string
	HTMLContent string
	FilePath    string
}

type API interface {
	Save(ctx context.Context, url string, filename string) (*ArchivedPage, error)
}

func New(cfg Config) API {
	return &service{
		config: cfg,
	}
}

type service struct {
	config Config
}

func (s *service) Save(ctx context.Context, url string, fileName string) (*ArchivedPage, error) {
	// Check if file already exists
	filePath := filepath.Join(s.config.DownloadFolder, fmt.Sprintf("%s.html", fileName))
	if _, err := os.Stat(filePath); err == nil || !os.IsNotExist(err) {
		return nil, ErrUnableToCreateFile
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create http request. err: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch url. err: %w", err)
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body. err: %w", err)
	}

	body := string(bodyBytes)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("failed to load goquery doc. err: %w", err)
	}

	title := doc.Find("title").Text()

	// Try and extract body text
	doc.Find("script").Each(func(i int, el *goquery.Selection) {
		el.Remove()
	})

	bodyText := doc.Text()
	page := &ArchivedPage{
		Title:       title,
		HTMLContent: bodyText,
		FilePath:    filePath,
	}

	return page, nil
}
