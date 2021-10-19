package presenter

import (
	"fmt"

	"github.com/sparkymat/archmark/model"
)

type ApiToken struct {
	ID    string
	Token string
}

func PresentApiTokens(tokens []model.ApiToken) []ApiToken {
	presentedTokens := []ApiToken{}
	for _, token := range tokens {
		presentedTokens = append(presentedTokens, ApiToken{
			ID:    fmt.Sprintf("%d", token.ID),
			Token: token.Token,
		})
	}
	return presentedTokens
}
