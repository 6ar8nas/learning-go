package types

import (
	sharedTypes "github.com/6ar8nas/learning-go/shared/types"
	"github.com/google/uuid"
)

type TaskRepository interface {
	GetTasks(userId uuid.UUID, isAdmin bool) ([]*sharedTypes.Task, error)
	GetTaskById(id uuid.UUID, userId uuid.UUID, isAdmin bool) (*sharedTypes.Task, error)
	CreateTask(userId uuid.UUID, task sharedTypes.TaskCreateRequest) (*sharedTypes.Task, error)
	UpdateTask(id uuid.UUID, userId uuid.UUID, isAdmin bool, task sharedTypes.TaskUpdateRequest) (*sharedTypes.Task, error)
}

type UserRepository interface {
	GetUsers() ([]*sharedTypes.User, error)
	GetUserByUsername(username string) (*sharedTypes.User, error)
	GetUserById(id uuid.UUID) (*sharedTypes.User, error)
	CreateUser(req UserHashedAuthRequest) (*sharedTypes.User, error)
}
