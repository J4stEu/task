package token

import "github.com/J4stEu/task/internal/model"

type Token interface {
	// New - create new token
	New(model.Token) (model.Token, error)

	// Update - update token
	Update(model.Token) (model.Token, error)

	// GetByValue - get token by value
	GetByValue(string) (model.Token, error)
}
