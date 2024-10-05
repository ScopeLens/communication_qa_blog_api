package tables

import "time"

type User struct {
	Username       string    `gorm:"primaryKey;not null" json:"username"` // 主键是 username
	Nickname       string    `gorm:"not null" json:"nickname"`
	Email          string    `gorm:"unique;not null" json:"email"` // Email 保持唯一性
	Password       string    `gorm:"not null" json:"password"`
	FollowersCount int       `gorm:"default:0" json:"followers_count"` // 粉丝数
	FollowingCount int       `gorm:"default:0" json:"following_count"` // 关注数
	AvatarURL      string    `json:"avatar_url"`                       // 头像 URL
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
