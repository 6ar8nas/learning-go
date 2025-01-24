package types

import "github.com/google/uuid"

type TaskRepository interface {
	GetTaskById(id uuid.UUID) (*Task, error)
	CreateTask(task TaskCreateRequest) (*Task, error)
}
