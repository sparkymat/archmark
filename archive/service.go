package archive

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	ErrUnableToCreateFile = errors.New("unable to create file")
)

type Config struct {
	MonolithPath   string
	DownloadFolder string
}

type ArchivedPage struct {
	Title       string
	HTMLContent string
	FilePath    string
}

type API interface {
	Save(url string, filename string) (*ArchivedPage, error)
}

func New(cfg Config) API {
	return &service{
		config: cfg,
	}
}

type service struct {
	config Config
}

func (s *service) Save(url string, fileName string) (*ArchivedPage, error) {
	// Check if file already exists
	filePath := filepath.Join(s.config.DownloadFolder, fmt.Sprintf("%s.html", fileName))
	if _, err := os.Stat(filePath); err == nil || !os.IsNotExist(err) {
		return nil, ErrUnableToCreateFile
	}

	// Download page using monolith
	err := exec.Command(s.config.MonolithPath, "-esM", url, "-o", filePath).Run()
	if err != nil {
		return nil, err
	}

	// Fetch page
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse for title

	return "", string(out), nil
}
