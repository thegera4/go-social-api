package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

// (Model) Post struct that represents a post in the database.
type Post struct {
	ID        int64    `json:"id"`
	Content   string   `json:"content"`
	Title     string   `json:"title"`
	UserID    int64    `json:"user_id"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

// Custom type to hold the database connection.
type PostStore struct {
	db *sql.DB
}

// Method that creates a new post in the database.
func (s *PostStore) Create(ctx context.Context, post *Post) error {
	query := `INSERT INTO posts (content, title, user_id, tags) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`

	err := s.db.QueryRowContext(ctx, query, post.Content, post.Title, post.UserID, pq.Array(post.Tags)).
							Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)
	if err != nil { return err }

	return nil
}
