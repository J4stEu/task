package task

import (
	"net/http"
	"strconv"

	"github.com/J4stEu/task/internal/api/http/response"
)

func (p *Task) getTask(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseUint(r.PathValue("id"), 10, 32)

	task, err := p.service.Task.GetByID(uint32(id))
	if err != nil {
		response.Response(w, http.StatusBadRequest, err)
		return
	}

	response.Response(w, http.StatusOK, task)
}

func (p *Task) getTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := p.service.Task.GetAll()
	if err != nil {
		response.Response(w, http.StatusBadRequest, err)
		return
	}

	response.Response(w, http.StatusOK, tasks)
}
