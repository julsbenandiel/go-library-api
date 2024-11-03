-- +goose Up
CREATE TABLE users (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  first_name VARCHAR(255), 
  last_name VARCHAR(255),
  email VARCHAR(255) NOT NULL,
  username VARCHAR(255) NOT NULL,
  birth_date DATE,
  address VARCHAR(255) NOT NULL,

  created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
  updpated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP 
);

-- +goose Down
DROP TABLE users;