package repository

import (
	"github.com/J4stEu/task/internal/repository/task"
	"github.com/J4stEu/task/internal/repository/token"
	"github.com/J4stEu/task/internal/repository/user"
)

type Repository struct {
	Task  task.Task
	Token token.Token
	User  user.User
}
