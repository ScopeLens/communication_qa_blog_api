package tables

import "time"

type Post struct {
	PostID         uint      `gorm:"primaryKey"`                              // 帖子ID
	Title          string    `gorm:"not null"`                                // 帖子标题
	Content        string    `gorm:"type:text"`                               // 帖子内容
	Username       string    `gorm:"not null"`                                // 用户名 (外键，关联 User 表的 username)
	Images         string    `gorm:"type:json"`                               // 图片URL列表，存储为JSON
	Tags           string    `gorm:"type:json"`                               // 标签列表，存储为JSON
	FavoritesCount int       `gorm:"default:0"`                               // 收藏数
	LikesCount     int       `gorm:"default:0"`                               // 点赞数
	ViewsCount     int       `gorm:"default:0"`                               // 浏览量
	CreatedAt      time.Time `gorm:"autoCreateTime"`                          // 创建时间
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`                          // 更新时间
	User           User      `gorm:"foreignKey:Username;references:Username"` // 关联用户 (User表)
}
