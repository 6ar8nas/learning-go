package tasks

import (
	"6ar8nas/test-app/database"
	"6ar8nas/test-app/types"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type Repository struct {
	*database.ConnectionPool
}

func NewRepository(database *database.ConnectionPool) *Repository {
	return &Repository{ConnectionPool: database}
}

func (s *Repository) GetTaskById(id uuid.UUID) (*types.Task, error) {
	row := s.DB.QueryRow("SELECT * FROM tasks WHERE id = $1", id)
	return scanRow(row)
}

func (s *Repository) CreateTask(req types.TaskCreateRequest) (*types.Task, error) {
	row := s.DB.QueryRow("INSERT INTO tasks (type, status, result) VALUES ($1, $2, $3) RETURNING id, type, status, result", req.Type, types.Scheduled, "")
	return scanRow(row)
}

func scanRow(row *sql.Row) (*types.Task, error) {
	task := new(types.Task)
	switch err := row.Scan(
		&task.Id,
		&task.Type,
		&task.Status,
		&task.Result,
	); err {
	case nil:
		return task, nil
	case sql.ErrNoRows:
		return nil, fmt.Errorf("requested task does not exist")
	default:
		return nil, err
	}
}
