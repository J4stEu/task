package task

import (
	"net/http"
	"strconv"

	"github.com/J4stEu/task/internal/api/http/middleware"
	"github.com/J4stEu/task/internal/api/http/response"
)

func (p *Task) getTask(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseUint(r.PathValue("id"), 10, 32) //nolint:errcheck

	rData, ok := r.Context().Value(middleware.ReqAuthDataKey).(middleware.AuthData)
	if !ok {
		response.Response(w, http.StatusUnauthorized, nil)
		return
	}

	task, err := p.service.Task.GetByUserIDAndID(rData.User.ID, uint32(id))
	if err != nil {
		response.Response(w, http.StatusBadRequest, err)
		return
	}

	response.Response(w, http.StatusOK, task)
}

func (p *Task) getTasks(w http.ResponseWriter, r *http.Request) {
	rData, ok := r.Context().Value(middleware.ReqAuthDataKey).(middleware.AuthData)
	if !ok {
		response.Response(w, http.StatusUnauthorized, nil)
		return
	}

	tasks, err := p.service.Task.GetAllByUserID(rData.User.ID)
	if err != nil {
		response.Response(w, http.StatusBadRequest, err)
		return
	}

	response.Response(w, http.StatusOK, tasks)
}

func (p *Task) getAnalytics(w http.ResponseWriter, r *http.Request) {
	rData, ok := r.Context().Value(middleware.ReqAuthDataKey).(middleware.AuthData)
	if !ok {
		response.Response(w, http.StatusUnauthorized, nil)
		return
	}

	tasks, err := p.service.Task.GetAnalyticsByUserID(rData.User.ID)
	if err != nil {
		response.Response(w, http.StatusBadRequest, err)
		return
	}

	response.Response(w, http.StatusOK, tasks)
}
