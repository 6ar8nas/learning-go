package types

import (
	"github.com/google/uuid"
)

type TaskRepository interface {
	GetTasks(userId uuid.UUID, isAdmin bool) ([]*Task, error)
	GetTaskById(id uuid.UUID, userId uuid.UUID, isAdmin bool) (*Task, error)
	CreateTask(userId uuid.UUID, task TaskCreateRequest) (*Task, error)
	UpdateTask(id uuid.UUID, userId uuid.UUID, isAdmin bool, task TaskUpdateRequest) (*Task, error)
}

type UserRepository interface {
	GetUsers() ([]*User, error)
	GetUserByUsername(username string) (*User, error)
	GetUserById(id uuid.UUID) (*User, error)
	CreateUser(req UserHashedAuthRequest) (*User, error)
}
