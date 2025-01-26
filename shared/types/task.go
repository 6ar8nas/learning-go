package types

import "github.com/google/uuid"

type TaskType string

const (
	HardWork TaskType = "HardWork"
)

type TaskStatus string

const (
	Scheduled TaskStatus = "Scheduled"
	Active    TaskStatus = "Active"
	Complete  TaskStatus = "Complete"
)

type TaskCreateRequest struct {
	Type TaskType `json:"type"`
}

type TaskUpdateRequest struct {
	Status *TaskStatus `json:"status,omitempty"`
	Result *string     `json:"result,omitempty"`
}

type Task struct {
	Id     uuid.UUID  `json:"id"`
	Type   TaskType   `json:"type"`
	Status TaskStatus `json:"status"`
	Result *string    `json:"result,omitempty"`
	UserId string     `json:"-"`
}
