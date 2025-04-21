package service

import (
	"github.com/J4stEu/task/internal/repository"
	"github.com/J4stEu/task/internal/service/task"
)

func NewMock(repo *repository.Repository) *Service {
	return &Service{
		Task: task.NewMock(
			&task.TaskMockRepo{
				Task: repo.Task,
			},
		),
	}
}
