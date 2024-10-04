package tables

import "time"

type User struct {
	Username       string `gorm:"primaryKey;not null"` // 主键是 username
	Nickname       string `gorm:"not null"`
	Email          string `gorm:"unique;not null"` // Email 保持唯一性
	Password       string `gorm:"not null"`
	FollowersCount int    `gorm:"default:0"` // 粉丝数
	FollowingCount int    `gorm:"default:0"` // 关注数
	AvatarURL      string // 头像 URL
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
