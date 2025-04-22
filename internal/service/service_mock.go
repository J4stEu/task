package service

import (
	"github.com/J4stEu/task/internal/repository"
	"github.com/J4stEu/task/internal/service/auth"
	"github.com/J4stEu/task/internal/service/task"
	"github.com/J4stEu/task/internal/service/user"
)

func NewMock(apiSecret []byte, repo *repository.Repository) *Service {
	return &Service{
		Task: task.NewMock(
			&task.TaskMockRepo{
				Task: repo.Task,
			},
		),
		Auth: auth.NewMock(apiSecret, &auth.AuthMockRepo{
			Token: repo.Token,
			User:  repo.User,
		}),
		User: user.NewMock(&user.UserMockRepo{
			User: repo.User,
		}),
	}
}
