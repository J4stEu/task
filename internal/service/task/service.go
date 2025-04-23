package task

import (
	"context"

	"github.com/J4stEu/task/internal/model"
)

// Task is responsible for tasks
type Task interface {
	// New - create new task
	New(CreateTaskDTO) (model.Task, error)

	// Update - update task
	Update(model.Task, UpdateTaskDTO) (model.Task, error)

	// GetByUserIDAndID - get task by user id and id
	GetByUserIDAndID(uint32, uint32) (model.Task, error)
	// GetAllByUserID - get all tasks by user id
	GetAllByUserID(uint32) ([]model.Task, error)
	// GetAnalyticsByUserID - get analytics by user id
	GetAnalyticsByUserID(uint32) (TasksUnalytics, error)

	WatchAnalytics(ctx context.Context)

	// Delete - delete task
	Delete(uint32) error
}
