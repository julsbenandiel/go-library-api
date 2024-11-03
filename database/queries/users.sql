-- name: GetUsers :many
SELECT * FROM users ORDER BY created_at DESC;

-- name: CreateUser :one
INSERT INTO users (
  first_name, last_name, email, username, birth_date, address
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;