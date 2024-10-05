package views

import (
	"time"
)

type CommentTree struct {
	CommentID uint          `json:"comment_id"`
	PostID    uint          `json:"post_id"`
	Username  string        `json:"username"`
	Nickname  string        `json:"nickname"`
	AvatarURL string        `json:"avatar_url"`
	Content   string        `json:"content"`
	ParentID  *uint         `json:"parent_id"`
	Replies   []CommentTree `json:"replies"`
	CreatedAt time.Time     `json:"created_at"`
}

type CommentItem struct {
	CommentID uint      `json:"comment_id"`
	PostID    uint      `json:"post_id"`
	Nickname  string    `json:"nickname"`
	AvatarURL string    `json:"avatar_url"`
	Username  string    `json:"username"`
	Content   string    `json:"content"`
	ParentID  *uint     `json:"parent_id"`
	CreatedAt time.Time `json:"created_at"`
}
