package data

import (
	"database/sql"
	"time"
)


type Comment struct {
	ID        int64     `json:"id"`         // Unique identifier for the comment
	BlogID    int64     `json:"blogId"`     // The ID of the blog post the comment belongs to
	Author    string    `json:"author"`     // Name of the person commenting
	Content   string    `json:"content"`    // The actual comment text
	CreatedAt time.Time `json:"createdAt"`  // Timestamp for when the comment was created
	UpdatedAt time.Time `json:"updatedAt"`  // Timestamp for when the comment was last updated
}

type CommentModel struct {
	DB *sql.DB
}