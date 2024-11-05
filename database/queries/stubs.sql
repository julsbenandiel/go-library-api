-- name: GetStubs :many
SELECT s.*, sqlc.embed(users)
FROM stubs s
JOIN users ON users.id = s.created_by
ORDER BY s.created_at DESC;


-- name: CreateStub :one
INSERT INTO stubs (
  id, name, description, created_by, created_at, updpated_at
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;