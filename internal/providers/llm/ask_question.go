package llm

import (
	"context"
	"errors"
)

var ErrInvalidConversation = errors.New("invalid conversation")

func (s *Service) AskQuestion(ctx context.Context, questionContext string, currentConversation []string, question string) (string, error) {
	return s.engine.AskQuestion(ctx, questionContext, currentConversation, question)
}
