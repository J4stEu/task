package task

import (
	"fmt"
	"net/http"

	"github.com/J4stEu/task/internal/api/http/middleware"
	"github.com/J4stEu/task/internal/service/task"
)

type Task struct {
	path    string
	service *TaskService
}

type TaskService struct {
	Task task.Task
}

func Init(mux *http.ServeMux, authMiddleware *middleware.AuthMiddleware, path string, service *TaskService) {
	taskAPI := new(path, service)
	taskAPI.init(mux, authMiddleware)
}

func new(path string, service *TaskService) *Task {
	return &Task{
		path:    path,
		service: service,
	}
}

func (p *Task) init(mux *http.ServeMux, authMiddleware *middleware.AuthMiddleware) {
	authM := authMiddleware.AuthenticationMiddleware
	taskPath := p.path + "/tasks"

	mux.HandleFunc(fmt.Sprintf("GET %s/{id}", taskPath), authM(p.getTask))
	mux.HandleFunc(fmt.Sprintf("GET %s", taskPath), authM(p.getTasks))
	mux.HandleFunc(fmt.Sprintf("GET %s/analytics", taskPath), authM(p.getAnalytics))
}
