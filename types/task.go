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

type Task struct {
	Id     uuid.UUID  `json:"id"`
	Type   TaskType   `json:"type"`
	Status TaskStatus `json:"status"`
	Result *string    `json:"result,omitempty"`
	UserId string     `json:"-"`
}

type TaskCreateRequest struct {
	UserId uuid.UUID `json:"userId"` // TODO: acquire userId from token
	Type   TaskType  `json:"type"`
}

type TaskUpdateRequest struct {
	Status *TaskStatus `json:"status,omitempty"`
	Result *string     `json:"result,omitempty"`
}
