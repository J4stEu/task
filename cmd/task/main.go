package main

import (
	"context"

	api "github.com/J4stEu/task/internal/api/http"
	"github.com/J4stEu/task/internal/repository"
	"github.com/J4stEu/task/internal/service"
)

func main() {
	apiSecret := []byte("test")

	repo := repository.NewMock()
	service := service.NewMock(apiSecret, repo)
	httpServer := api.New(apiSecret, service)

	_ = httpServer.Run(context.Background())
}
