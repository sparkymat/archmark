package archive_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/sparkymat/archmark/archive"
	"github.com/stretchr/testify/assert"
)

func TestFetch(t *testing.T) {
	defer func() {
		if _, err := os.Stat("/tmp/foobar.html"); err == nil {
			os.Remove("/tmp/foobar.html")
		}
	}()
	archiveConfig := archive.Config{
		MonolithPath:   "/usr/local/bin/monolith",
		DownloadFolder: "/tmp",
	}
	s := archive.New(archiveConfig)

	htmlBody := "<html><head><title>THE TITLE</title></head><body></body></html>"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, htmlBody)
	}))
	defer ts.Close()

	page, err := s.Save(ts.URL, "foobar")

	assert.NoError(t, err)
	assert.NotNil(t, page)

	_, err = os.Stat("/tmp/foobar.html")
	assert.NoError(t, err)
	assert.Equal(t, "THE TITLE", page.Title)
	assert.Equal(t, htmlBody+"\n", page.HTMLContent)
}
