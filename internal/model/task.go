package model

import "time"

type Task struct {
	ID     uint32 `json:"id"`
	UserID uint32 `json:"userID"`

	Status TaskStatus `json:"status"`

	Title       string `json:"title"`
	Description string `json:"description"`

	DueDate time.Time `json:"dueDate"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (t *Task) IsOverdue() bool {
	now := time.Now()

	return now.After(t.DueDate) && t.Status != TaskStatusDone
}

type TaskStatus string

const (
	TaskStatusPending    TaskStatus = "pending"
	TaskStatusInProgress TaskStatus = "inProgress"
	TaskStatusDone       TaskStatus = "done"
)
