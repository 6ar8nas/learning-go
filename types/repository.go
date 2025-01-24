package types

import (
	"errors"

	"github.com/google/uuid"
)

type TaskRepository interface {
	GetTasks() ([]*Task, error)
	GetTaskById(id uuid.UUID) (*Task, error)
	CreateTask(userId uuid.UUID, task TaskCreateRequest) (*Task, error)
	UpdateTask(id uuid.UUID, task TaskUpdateRequest) (*Task, error)
}

type UserRepository interface {
	GetUsers() ([]*User, error)
	GetUserByUsername(username string) (*User, error)
	GetUserById(id uuid.UUID) (*User, error)
	CreateUser(req UserHashedAuthRequest) (*User, error)
}

var ErrorNotFound = errors.New("requested resource doesn't exist")
