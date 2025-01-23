package data

import (
	"database/sql"
	"time"
)

type Blog struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
	Author string `json:"author"`
	Tags []string `json:"tags"`
	Views int64 `json:"views"`
	Likes int64 `json:"likes"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Comments    []Comment  `json:"comments,omitempty"` 
	CommentsCount int64    `json:"commentsCount"`
}

type BlogModel struct {
	DB *sql.DB
}

