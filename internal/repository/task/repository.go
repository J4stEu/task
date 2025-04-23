package task

import "github.com/J4stEu/task/internal/model"

// Task repository is responsible for tasks
type Task interface {
	// New - create new task
	New(model.Task) (model.Task, error)

	// Update - update task
	Update(model.Task) (model.Task, error)

	// GetByUserIDAndID - get task by user id and id
	GetByUserIDAndID(uint32, uint32) (model.Task, error)
	// GetAllByUserID - get all user tasks
	GetAllByUserID(uint32) ([]model.Task, error)

	Delete(uint32) error
}
