package archive

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var (
	ErrUnableToCreateFile = errors.New("unable to create file")
	ErrURLFetchFailed     = errors.New("url fetch failed")
)

const okStatusPrefix = 2

type ArchivedPage struct {
	Title       string
	HTMLContent string
}

func New() Service {
	return Service{}
}

type Service struct {
}

func (s *Service) FetchDetails(ctx context.Context, url string) (*ArchivedPage, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create http request. err: %w", err)
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.81 Safari/537.36")

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

	// Check if it is a 2xx status code
	if resp.StatusCode/100 != okStatusPrefix {
		return nil, ErrURLFetchFailed
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("failed to load goquery doc. err: %w", err)
	}

	title := doc.Find("title").Text()

	// Try and extract body text
	doc.Find("script").Each(func(_ int, el *goquery.Selection) {
		el.Remove()
	})

	bodyText := doc.Text()
	page := &ArchivedPage{
		Title:       title,
		HTMLContent: bodyText,
	}

	return page, nil
}
