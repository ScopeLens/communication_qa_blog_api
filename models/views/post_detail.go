package views

import (
	"encoding/json"
	"time"
)

type PostDetail struct {
	PostID         uint            `json:"post_id"`
	Title          string          `json:"title"`
	Content        string          `json:"content"`
	Username       string          `json:"username"`
	Nickname       string          `json:"nickname"`
	AvatarURL      string          `json:"avatar_url"`
	Images         json.RawMessage `json:"images"`
	Tags           json.RawMessage `json:"tags"`
	FavoritesCount int             `json:"favorites_count"`
	LikesCount     int             `json:"likes_count"`
	ViewsCount     int             `json:"views_count"`
	ReplyCount     int             `json:"reply_count"`
	UpdatedAt      time.Time       `json:"updated_at"`
}

type PostStatus struct {
	IsViewed   bool `json:"is_viewed"`
	IsLiked    bool `json:"is_liked"`
	IsFavorite bool `json:"is_favorite"`
}
