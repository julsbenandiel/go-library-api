-- +goose Up
CREATE TABLE stubs (
  id uuid PRIMARY KEY,
  name VARCHAR NOT NULL,
  description TEXT NOT NULL,

  created_at TIMESTAMPTZ NOT NULL,
  updpated_at TIMESTAMPTZ NOT NULL,
  
  created_by uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE stubs;