package models

import (
	"time"
)

// Post is a model
type Post struct {
	ID      int64     `json:"id"`
	Text    string    `json:"text"`
	Created time.Time `json:"created"`
}

var posts []Post
