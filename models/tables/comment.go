package tables

import "time"

type Comment struct {
	CommentID uint   `gorm:"primaryKey"`
	PostID    uint   `gorm:"not null"`
	Username  string `gorm:"not null"` // 使用 username 关联用户
	Content   string `gorm:"not null"`
	ParentID  *uint  // 父评论ID，用于评论的回复
	CreatedAt time.Time
	User      User `gorm:"foreignKey:Username;references:Username"`
	Post      Post `gorm:"foreignKey:PostID;references:PostID"`
}
