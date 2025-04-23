package task

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-loremipsum/loremipsum"

	"github.com/J4stEu/task/internal/model"
	repoMessage "github.com/J4stEu/task/internal/repository/message"
	userRepo "github.com/J4stEu/task/internal/repository/user"
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
	statuses := []model.TaskStatus{
		model.TaskStatusPending,
		model.TaskStatusInProgress,
		model.TaskStatusDone,
	}

	for i := range 10 {
		randomStatus := statuses[rand.Intn(3)]
		randomDays := 24 * (1 + rand.Intn(30-1))
		createdAt := time.Now().Add(-time.Hour * time.Duration(randomDays))
		loremIpsumGenerator := loremipsum.NewWithSeed(1234)

		tasks = append(tasks, model.Task{
			ID: uint32(i) + 1,

			Status: randomStatus,

			Title:       fmt.Sprintf("Task %v", uint32(i)+1),
			Description: loremIpsumGenerator.Words(20),

			DueDate:   createdAt,
			CreatedAt: createdAt,
			UpdatedAt: createdAt,
		})
	}

	users := userRepo.GetMockedUsers()
	// WARN: Slices should be equal
	for i := range tasks {
		tasks[i].UserID = users[i].ID
	}

	admin := users[10]
	for i := range 10 {
		randomDays := 24 * (1 + rand.Intn(30-1))
		createdAt := time.Now().Add(-time.Hour * time.Duration(randomDays))
		loremIpsumGenerator := loremipsum.NewWithSeed(1234)

		tasks = append(tasks, model.Task{
			ID:     uint32(i) + 1,
			UserID: admin.ID,

			Title:       fmt.Sprintf("Admin task %v", uint32(i)+1),
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

func (*TaskMock) GetByUserIDAndID(userID, id uint32) (model.Task, error) {
	tasks := getMockedTasks()
	for i := range tasks {
		task := tasks[i]
		if userID == task.UserID && id == task.ID {
			return task, nil
		}
	}

	return model.Task{}, repoMessage.ErrRecordNotFound
}

func (*TaskMock) GetAllByUserID(id uint32) ([]model.Task, error) {
	tasks := getMockedTasks()

	userTasks := make([]model.Task, 0, len(tasks))
	for i := range tasks {
		if tasks[i].UserID == id {
			userTasks = append(userTasks, tasks[i])
		}
	}

	return userTasks, nil
}

func (*TaskMock) Delete(id uint32) error {
	return nil
}
