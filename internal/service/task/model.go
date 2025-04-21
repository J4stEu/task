package task

type CreateTaskDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateTaskDTO CreateTaskDTO
