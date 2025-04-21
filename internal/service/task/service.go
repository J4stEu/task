package task

import "github.com/J4stEu/task/internal/model"

// Task is responsible for tasks
type Task interface {
	// New - create new task
	New(CreateTaskDTO) (model.Task, error)

	// Update - update task
	Update(model.Task, UpdateTaskDTO) (model.Task, error)

	// GetByID - get task by id
	GetByID(uint32) (model.Task, error)
	// GetAll - get all tasks
	GetAll() ([]model.Task, error)

	// Delete - delete task
	Delete(uint32) error
}
