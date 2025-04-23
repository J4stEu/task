package task

import (
	"github.com/J4stEu/task/internal/model"
	"github.com/J4stEu/task/internal/repository/task"
)

// Task service implementation with mock repository
type TaskMock struct {
	repository *TaskMockRepo
}

// Task service repository
type TaskMockRepo struct {
	Task task.Task
}

func NewMock(taskMockRepo *TaskMockRepo) Task {
	return &TaskMock{
		repository: taskMockRepo,
	}
}

func (pm *TaskMock) New(createData CreateTaskDTO) (model.Task, error) {
	return pm.repository.Task.New(model.Task{
		Title:       createData.Title,
		Description: createData.Description,
	})
}

func (pm *TaskMock) Update(task model.Task, updateData UpdateTaskDTO) (model.Task, error) {
	if err := updateTask(&task, &updateData); err != nil {
		return model.Task{}, err
	}

	return pm.repository.Task.Update(task)
}

func updateTask(task *model.Task, updateData *UpdateTaskDTO) error {
	shouldUpdate := false

	if task.Title != updateData.Title {
		task.Title = updateData.Title
		shouldUpdate = true
	}
	if task.Description != updateData.Description {
		task.Description = updateData.Description
		shouldUpdate = true
	}

	if !shouldUpdate {
		return ErrNothingToUpdate
	}

	return nil
}

func (pm *TaskMock) GetByUserIDAndID(userID, id uint32) (model.Task, error) {
	return pm.repository.Task.GetByUserIDAndID(userID, id)
}

func (pm *TaskMock) GetAllByUserID(id uint32) ([]model.Task, error) {
	return pm.repository.Task.GetAllByUserID(id)
}

func (pm *TaskMock) Delete(id uint32) error {
	return pm.repository.Task.Delete(id)
}
