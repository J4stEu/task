package task

type CreateTaskDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateTaskDTO CreateTaskDTO

type TasksUnalytics struct {
	Pending    uint32 `json:"pending"`
	InProgress uint32 `json:"inProgress"`
	Done       uint32 `json:"completed"`
	Overdue    uint32 `json:"overdue"`
}
