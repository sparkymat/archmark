package bookmark

import "github.com/sparkymat/archmark/internal/service"

func New(storageFolder string, dbx service.DatabaseProvider, llm service.LLMProvider) *Service {
	return &Service{
		storageFolder: storageFolder,
		dbx:           dbx,
		llm:           llm,
	}
}

type Service struct {
	storageFolder string
	dbx           service.DatabaseProvider
	llm           service.LLMProvider
}
