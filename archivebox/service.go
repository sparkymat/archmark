package archivebox

type Config struct {
	Path     string
	Username string
	Password string
}

type API interface {
	ArchiveLink(url string) (string, string, error)
}

func New(cfg Config) API {
	return &service{cfg: cfg}
}

type service struct {
	cfg Config
}

func (s *service) ArchiveLink(url string) (string, string, error) {
	panic("not implemented")
}
