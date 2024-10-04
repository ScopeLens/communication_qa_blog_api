package tables

import "time"

type Favorite struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Username  string    `gorm:"size:255;not null"` // 关联用户表的 username
	PostID    uint      `gorm:"not null"`          // 关联帖子表的 post_id
	CreatedAt time.Time `gorm:"autoCreateTime"`    // 自动创建时间
}
