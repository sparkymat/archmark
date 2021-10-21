package presenter

import (
	"fmt"

	"github.com/sparkymat/archmark/model"
)

type APIToken struct {
	ID    string
	Token string
}

func PresentAPITokens(tokens []model.APIToken) []APIToken {
	presentedTokens := []APIToken{}

	for _, token := range tokens {
		presentedTokens = append(presentedTokens, APIToken{
			ID:    fmt.Sprintf("%d", token.ID),
			Token: token.Token,
		})
	}

	return presentedTokens
}
