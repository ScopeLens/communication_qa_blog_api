package tables

import "time"

type View struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Username  string    `gorm:"not null"`       // 关联用户表的 username
	PostID    uint      `gorm:"not null"`       // 关联帖子表的 post_id
	CreatedAt time.Time `gorm:"autoCreateTime"` // 自动创建时间
	Post      Post      `gorm:"foreignKey:PostID;references:PostID"`
}
