package service

import "context"

type LLMProvider interface {
	AskQuestion(ctx context.Context, questionContext string, currentConversation []string, question string) (string, error)
}
