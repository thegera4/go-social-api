package store

import (
	"context"
	"database/sql"
)

// Custom type to hold the database connection.
type PostsStore struct {
	db *sql.DB
}

// Method that creates a new post in the database.
func (s *PostsStore) Create(ctx context.Context) error {
	return nil
}