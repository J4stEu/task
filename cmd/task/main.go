package main

import (
	"context"

	api "github.com/J4stEu/task/internal/api/http"
	"github.com/J4stEu/task/internal/repository"
	"github.com/J4stEu/task/internal/service"
)

func main() {
	repo := repository.NewMock()
	service := service.NewMock(repo)
	httpServer := api.New(service)

	_ = httpServer.Run(context.Background())
}
