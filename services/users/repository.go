package users

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

func (s *Repository) GetUsers() ([]*types.User, error) {
	rows, err := s.DB.Query("SELECT id, username, password, is_admin FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRows(rows)
}

func (s *Repository) GetUserById(id uuid.UUID) (*types.User, error) {
	row := s.DB.QueryRow("SELECT id, username, password, is_admin FROM users WHERE id = $1", id)
	return scanRow(row)
}

func (s *Repository) GetUserByUsername(username string) (*types.User, error) {
	row := s.DB.QueryRow("SELECT id, username, password, is_admin FROM users WHERE username = $1", username)
	return scanRow(row)
}

func (s *Repository) CreateUser(req types.UserHashedAuthRequest) (*types.User, error) {
	row := s.DB.QueryRow("INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id, username, password, is_admin", req.Username, req.Password)
	return scanRow(row)
}

func scanRow(row *sql.Row) (*types.User, error) {
	user := new(types.User)
	switch err := row.Scan(
		&user.Id,
		&user.Username,
		&user.Password,
		&user.Admin,
	); err {
	case nil:
		return user, nil
	case sql.ErrNoRows:
		return nil, types.ErrorNotFound
	default:
		return nil, err
	}
}

func scanRows(rows *sql.Rows) ([]*types.User, error) {
	users := make([]*types.User, 0)
	for rows.Next() {
		user := new(types.User)
		if err := rows.Scan(
			&user.Id,
			&user.Username,
			&user.Password,
			&user.Admin,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
