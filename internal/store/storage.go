package store

import (
	"context"
	"database/sql"
)

// Custom type to hold all the interfaces for interacting with the database.
type Storage struct {
	// Contains methods for interacting with the posts table in the database.
	Posts interface {
		Create(context.Context, *Post) error
	}

	// Contains methods for interacting with the users table in the database.
	Users interface {
		Create(context.Context, *User) error
	}
}

// Function that creates a new Storage struct and initializes the Posts and Users fields.
func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostStore{db},
		Users: &UsersStore{db},
	}
}
