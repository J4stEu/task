package service

import (
	"github.com/J4stEu/task/internal/service/auth"
	"github.com/J4stEu/task/internal/service/task"
	"github.com/J4stEu/task/internal/service/user"
)

type Service struct {
	Task task.Task
	Auth auth.Auth
	User user.User
}
