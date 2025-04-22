package auth

import "github.com/J4stEu/task/internal/model"

type Auth interface {
	Authenticate(LoginUserDTO) (AuthData, error)

	// ActualizeToken - actualize token data
	ActualizeToken(model.Token) error

	// GetTokenByValue - get token by value
	GetTokenByValue(string) (model.Token, error)

	// Revoke - revoke token
	Revoke(model.Token) (model.Token, error)
}
