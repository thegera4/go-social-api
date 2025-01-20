package store

import (
	"context"
	"database/sql"
)

// Custom type to hold the database connection.
type UsersStore struct {
	db *sql.DB
}

// Method that creates a new user in the database.
func (s *UsersStore) Create(ctx context.Context) error {
	return nil
}