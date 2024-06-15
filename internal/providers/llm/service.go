package llm

import (
	"errors"

	"github.com/sparkymat/archmark/internal/providers/llm/gpt3"
)

type EngineType string

const (
	EngineTypeGPT3   EngineType = "gpt3"
	EngineTypeUknown EngineType = "unknown"
)

var ErrUnknownEngineType = errors.New("unknown engine type")

func New(engineType EngineType, token string) (*Service, error) {
	var engine Engine

	switch engineType {
	case EngineTypeGPT3:
		engine = gpt3.New(token)
	case EngineTypeUknown:
		return nil, ErrUnknownEngineType
	default:
		return nil, ErrUnknownEngineType
	}

	return &Service{
		engine: engine,
	}, nil
}

type Service struct {
	engine Engine
}
