package posts

import "github.com/enzoenrico/go_backend/app/users"

type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	User  users.User `json:"user"`
	Timestamp int64 `json:"timestamp"`
}
