package bookmark

import "github.com/sparkymat/archmark/internal/service"

func New(storageFolder string, dbx service.DatabaseProvider) *Service {
	return &Service{
		storageFolder: storageFolder,
		dbx:           dbx,
	}
}

type Service struct {
	storageFolder string
	dbx           service.DatabaseProvider
}
