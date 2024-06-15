package gpt3

import (
	"context"
	"errors"
	"fmt"
	httplib "net/http"

	"github.com/samber/lo"
	"github.com/sparkymat/archmark/internal/http"
)

const endpoint = "https://api.openai.com/v1/chat/completions"

var ErrInvalidConversation = errors.New("invalid conversation")

type OpenAIRequest struct {
	Model    string                 `json:"model"`
	Messages []OpenAIRequestMessage `json:"messages"`
}

type RequestRole string

const (
	RequestRoleSystem    RequestRole = "system"
	RequestRoleUser      RequestRole = "user"
	RequestRoleAssistant RequestRole = "assistant"
)

type OpenAIRequestMessage struct {
	Role    RequestRole `json:"role"`
	Content string      `json:"content"`
}

type OpenAIResponse struct {
	ID        string                 `json:"id"`
	Object    string                 `json:"object"`
	CreatedAt uint64                 `json:"created"`
	Model     string                 `json:"model"`
	Choices   []OpenAIResponseChoice `json:"choices"`
}

type OpenAIResponseChoice struct {
	Index        int                  `json:"index"`
	Message      OpenAIRequestMessage `json:"message"`
	FinishReason string               `json:"finish_reason"`
}

func (s *Service) AskQuestion(ctx context.Context, questionContext string, conversation []string, question string) (string, error) {
	if len(conversation)%2 != 0 {
		return "", ErrInvalidConversation
	}

	messages := []OpenAIRequestMessage{
		{
			Role:    RequestRoleSystem,
			Content: "You are a helpful assistant.",
		},
	}

	existingMessages := lo.Map(conversation, func(msg string, idx int) OpenAIRequestMessage {
		var role RequestRole
		if idx == 0 {
			role = RequestRoleUser
		} else {
			role = RequestRoleAssistant
		}

		return OpenAIRequestMessage{
			Role:    role,
			Content: msg,
		}
	})

	messages = append(messages, existingMessages...)

	// Add new question

	newMessageContent := question

	if len(conversation) == 0 {
		newMessageContent = "Answer questions about the following text: " + questionContext + "\n\n" + newMessageContent
	}

	messages = append(messages, OpenAIRequestMessage{
		Role:    RequestRoleUser,
		Content: newMessageContent,
	})

	request := OpenAIRequest{
		Model:    "gpt-3.5-turbo",
		Messages: messages,
	}

	headers := map[string]string{
		"Authorization": "Bearer " + s.apiKey,
	}

	response, err := http.PostJSON[OpenAIResponse](ctx, httplib.DefaultClient, endpoint, headers, request)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}

	return response.Choices[0].Message.Content, nil
}
