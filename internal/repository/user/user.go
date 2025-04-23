package user

import "github.com/J4stEu/task/internal/model"

type User interface {
	// New - create new user
	New(model.User) (model.User, error)

	// GetByID - get user by id
	GetByID(uint32) (model.User, error)
	// GetByLogin - get user by login
	GetByLogin(string) (model.User, error)
	// GetAll - get all users
	GetAll() ([]model.User, error)
}
