-- +goose Up
CREATE TABLE users (
  id uuid PRIMARY KEY,
  first_name VARCHAR(255) NOT NULL, 
  last_name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  username VARCHAR(255) NOT NULL,
  birth_date DATE NOT NULL,
  address VARCHAR(255) NOT NULL,

  created_at TIMESTAMPTZ NOT NULL,
  updpated_at TIMESTAMPTZ NOT NULL 
);

-- +goose Down
DROP TABLE users;