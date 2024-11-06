https://go.dev/doc/modules/layout

***Foldering
- internal: shared internal packages, utilities, helpers, structs, models
- database: database, sql, queries, sqlc gen type defs

***Tools
- goose:
  - installation: go install github.com/pressly/goose/v3/cmd/goose@latest
  - description: for migration
- sqlc:
  - installation: go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
  - description: for code gen sql queries

***Notes
- to use uuid -> CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
- run migratations
  - goose postgres postgres://dev:admin123@localhost:5432/perks up
