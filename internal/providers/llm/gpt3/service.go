package gpt3

func New(token string) *Service {
	return &Service{
		apiKey: token,
	}
}

type Service struct {
	apiKey string
}
