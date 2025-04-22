package repository

import (
	"github.com/J4stEu/task/internal/repository/task"
	"github.com/J4stEu/task/internal/repository/token"
	"github.com/J4stEu/task/internal/repository/user"
)

// Get repository implementation with mock data
func NewMock() *Repository {
	return &Repository{
		Task:  task.NewMock(),
		Token: token.NewMock(),
		User:  user.NewMock(),
	}
}
