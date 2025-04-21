package repository

import (
	"github.com/J4stEu/task/internal/repository/task"
)

// Get repository implementation with mock data
func NewMock() *Repository {
	return &Repository{
		Task: task.NewMock(),
	}
}
