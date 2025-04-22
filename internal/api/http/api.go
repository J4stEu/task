package api

import (
	"net/http"

	"github.com/J4stEu/task/internal/api/http/auth"
	"github.com/J4stEu/task/internal/api/http/response"
	"github.com/J4stEu/task/internal/api/http/session"
	"github.com/J4stEu/task/internal/api/http/task"
	"github.com/J4stEu/task/internal/service"
)

func InitMux(apiSecret []byte, service *service.Service) *http.ServeMux {
	mux := http.NewServeMux()

	sess := session.New(apiSecret)

	path := "/api"

	auth.Init(
		mux,
		sess,
		path,
		&auth.AuthService{
			Auth: service.Auth,
			User: service.User,
		},
	)

	task.Init(
		mux,
		path,
		&task.TaskService{
			Task: service.Task,
		},
	)

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		response.Response(w, http.StatusOK, map[string]string{
			"apiWelcomeMessage": "Welcome to the HTTP API",
		})
	})

	return mux
}
