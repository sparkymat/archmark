package llm

import "context"

type Engine interface {
	AskQuestion(ctx context.Context, questionContext string, currentConversation []string, question string) (string, error)
}
