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
		MonolithPath: "/usr/local/bin/monolith",
	}
	s := archive.New(archiveConfig)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "<html></html>")
	}))
	defer ts.Close()

	title, body, err := s.Fetch(ts.URL)

	assert.NoError(t, err)
	assert.Equal(t, "title", title)
	assert.Equal(t, "<html></html>", body)
}
