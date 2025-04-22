package user

import "github.com/J4stEu/task/internal/model"

// User is responsible for users
type User interface {
	// New - create new user
	New(CreateUserDTO) (model.User, error)

	// GetByID - get user by id
	GetByID(uint32) (model.User, error)
}
