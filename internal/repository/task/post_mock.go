package task

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-loremipsum/loremipsum"

	"github.com/J4stEu/task/internal/model"
	repoMessage "github.com/J4stEu/task/internal/repository/message"
)

// Tasl repository implementation with mock data
type TaskMock struct{}

// Get Task repository implementation with mock data
func NewMock() Task {
	return &TaskMock{}
}

// Mock data
func getMockedTasks() []model.Task {
	tasks := make([]model.Task, 0, 10)

	for i := range 10 {
		randomDays := 24 * (1 + rand.Intn(30-1))
		createdAt := time.Now().Add(-time.Hour * time.Duration(randomDays))
		loremIpsumGenerator := loremipsum.NewWithSeed(1234)

		tasks = append(tasks, model.Task{
			ID: uint32(i) + 1,

			Title:       fmt.Sprintf("Task %v", uint32(i)+1),
			Description: loremIpsumGenerator.Words(20),

			CreatedAt: createdAt,
			UpdatedAt: createdAt,
		})
	}

	return tasks
}

func (*TaskMock) New(task model.Task) (model.Task, error) {
	tasks := getMockedTasks()
	task.ID = tasks[len(tasks)-1].ID + 1

	return task, nil
}

func (*TaskMock) Update(task model.Task) (model.Task, error) {
	return task, nil
}

func (*TaskMock) GetByID(id uint32) (model.Task, error) {
	tasks := getMockedTasks()
	for i := range tasks {
		task := tasks[i]
		if id == task.ID {
			return task, nil
		}
	}

	return model.Task{}, repoMessage.ErrRecordNotFound
}

func (*TaskMock) GetAll() ([]model.Task, error) {
	return getMockedTasks(), nil
}

func (*TaskMock) Delete(id uint32) error {
	return nil
}
