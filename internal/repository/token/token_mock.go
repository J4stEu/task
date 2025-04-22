package token

import (
	"time"

	"github.com/J4stEu/task/internal/model"
	repoMessage "github.com/J4stEu/task/internal/repository/message"
)

// Token repository implementation with mock data
type TokenMock struct{}

// Get Token repository implementation with mock data
func NewMock() Token {
	return &TokenMock{}
}

// Mock data
func getMockedTokens() []model.Token {
	tokens := make([]model.Token, 0, 1)

	tokens = append(tokens, model.Token{
		ID:       1,
		UserID:   10, // admin
		Value:    "14f65a66-455d-44f1-a2bc-f2cd89293d84",
		Lifetime: time.Duration(2419200000000000),
	})

	return tokens
}

func (*TokenMock) New(token model.Token) (model.Token, error) {
	tokens := getMockedTokens()
	token.ID = tokens[len(tokens)-1].ID + 1

	return token, nil
}

func (*TokenMock) Update(token model.Token) (model.Token, error) {
	return token, nil
}

func (*TokenMock) GetByValue(value string) (model.Token, error) {
	tokens := getMockedTokens()
	for i := range tokens {
		token := tokens[i]
		if token.Value == value {
			return token, nil
		}
	}

	return model.Token{}, repoMessage.ErrRecordNotFound
}
