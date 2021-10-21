package archive_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sparkymat/archmark/archive"
	"github.com/stretchr/testify/assert"
)

func TestFetch(t *testing.T) {
	archiveConfig := archive.Config{
		DownloadFolder: "/tmp",
	}
	s := archive.New(archiveConfig)

	htmlBody := "<html><head><title>THE TITLE</title></head><body>THIS IS THE BODY</body></html>"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, htmlBody)
	}))
	defer ts.Close()

	page, err := s.Save(ts.URL, "foobar")

	assert.NoError(t, err)
	assert.NotNil(t, page)
	assert.Equal(t, "THE TITLE", page.Title)
	assert.Equal(t, "THE TITLETHIS IS THE BODY\n", page.HTMLContent)
}
