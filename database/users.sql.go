// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  id, first_name, last_name, email, username, birth_date, address, created_at, updpated_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING id, first_name, last_name, email, username, birth_date, address, created_at, updpated_at
`

type CreateUserParams struct {
	ID         uuid.UUID `json:"id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email"`
	Username   string    `json:"username"`
	BirthDate  time.Time `json:"birth_date"`
	Address    string    `json:"address"`
	CreatedAt  time.Time `json:"created_at"`
	UpdpatedAt time.Time `json:"updpated_at"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.ID,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Username,
		arg.BirthDate,
		arg.Address,
		arg.CreatedAt,
		arg.UpdpatedAt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Username,
		&i.BirthDate,
		&i.Address,
		&i.CreatedAt,
		&i.UpdpatedAt,
	)
	return i, err
}

const getUserStubs = `-- name: GetUserStubs :many
SELECT 
  u.id, u.first_name, u.last_name, u.email, u.username, u.birth_date, u.address, u.created_at, u.updpated_at, 
  array_agg(stubs) as stubs
FROM users u
JOIN stubs ON stubs.created_by = u.id
WHERE u.id = $1
GROUP BY u.id
ORDER BY u.created_at DESC
`

type GetUserStubsRow struct {
	ID         uuid.UUID   `json:"id"`
	FirstName  string      `json:"first_name"`
	LastName   string      `json:"last_name"`
	Email      string      `json:"email"`
	Username   string      `json:"username"`
	BirthDate  time.Time   `json:"birth_date"`
	Address    string      `json:"address"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdpatedAt time.Time   `json:"updpated_at"`
	Stubs      interface{} `json:"stubs"`
}

func (q *Queries) GetUserStubs(ctx context.Context, id uuid.UUID) ([]GetUserStubsRow, error) {
	rows, err := q.db.Query(ctx, getUserStubs, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserStubsRow
	for rows.Next() {
		var i GetUserStubsRow
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Username,
			&i.BirthDate,
			&i.Address,
			&i.CreatedAt,
			&i.UpdpatedAt,
			&i.Stubs,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUsers = `-- name: GetUsers :many
SELECT id, first_name, last_name, email, username, birth_date, address, created_at, updpated_at FROM users ORDER BY created_at DESC
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Username,
			&i.BirthDate,
			&i.Address,
			&i.CreatedAt,
			&i.UpdpatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
