package tasks

import (
	"database/sql"

	"github.com/6ar8nas/learning-go/database"
	"github.com/6ar8nas/learning-go/types"
	"github.com/google/uuid"
)

type Repository struct {
	*database.ConnectionPool
}

func NewRepository(database *database.ConnectionPool) *Repository {
	return &Repository{ConnectionPool: database}
}

func (s *Repository) GetTasks(userId uuid.UUID, isAdmin bool) ([]*types.Task, error) {
	rows, err := s.DB.Query("SELECT id, type, status, result, user_id FROM tasks WHERE user_id = $1 OR $2", userId, isAdmin)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRows(rows)
}

func (s *Repository) GetTaskById(id uuid.UUID, userId uuid.UUID, isAdmin bool) (*types.Task, error) {
	row := s.DB.QueryRow("SELECT id, type, status, result, user_id FROM tasks WHERE id = $1 and (user_id = $2 OR $3)", id, userId, isAdmin)
	return scanRow(row)
}

func (s *Repository) CreateTask(userId uuid.UUID, req types.TaskCreateRequest) (*types.Task, error) {
	row := s.DB.QueryRow("INSERT INTO tasks (type, status, result, user_id) VALUES ($1, $2, $3, $4) RETURNING id, type, status, result, user_id", req.Type, types.Scheduled, nil, userId)
	return scanRow(row)
}

func (s *Repository) UpdateTask(id uuid.UUID, userId uuid.UUID, isAdmin bool, req types.TaskUpdateRequest) (*types.Task, error) {
	row := s.DB.QueryRow("UPDATE tasks SET status = COALESCE($4, status), result = COALESCE($5, result) WHERE id = $1 and (user_id = $2 OR $3) RETURNING id, type, status, result, user_id", id, userId, isAdmin, req.Status, req.Result)
	return scanRow(row)
}

func scanRow(row *sql.Row) (*types.Task, error) {
	task := new(types.Task)
	switch err := row.Scan(
		&task.Id,
		&task.Type,
		&task.Status,
		&task.Result,
		&task.UserId,
	); err {
	case nil:
		return task, nil
	case sql.ErrNoRows:
		return nil, types.ErrorNotFound
	default:
		return nil, err
	}
}

func scanRows(rows *sql.Rows) ([]*types.Task, error) {
	tasks := make([]*types.Task, 0)
	for rows.Next() {
		task := new(types.Task)
		if err := rows.Scan(
			&task.Id,
			&task.Type,
			&task.Status,
			&task.Result,
			&task.UserId,
		); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}
