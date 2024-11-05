-- name: GetUsers :many
SELECT * FROM users ORDER BY created_at DESC;

-- name: GetUserStubs :many
SELECT 
  u.*, 
  array_agg(stubs) as stubs
FROM users u
JOIN stubs ON stubs.created_by = u.id
WHERE u.id = $1
GROUP BY u.id
ORDER BY u.created_at DESC;

-- name: CreateUser :one
INSERT INTO users (
  id, first_name, last_name, email, username, birth_date, address, created_at, updpated_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;

