// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"time"

	"github.com/google/uuid"
)

type Stub struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdpatedAt  time.Time `json:"updpated_at"`
	CreatedBy   uuid.UUID `json:"created_by"`
}

type User struct {
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
