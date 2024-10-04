package tables

import "time"

type Follow struct {
	ID           uint   `gorm:"primaryKey"`
	Follower     string `gorm:"not null"` // 关注者的 username
	Followee     string `gorm:"not null"` // 被关注者的 username
	CreatedAt    time.Time
	FollowerUser User `gorm:"foreignKey:Follower;references:Username"`
	FolloweeUser User `gorm:"foreignKey:Followee;references:Username"`
}
