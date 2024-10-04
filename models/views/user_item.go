package views

type UserItem struct {
	Username       string
	Nickname       string
	FollowersCount int `gorm:"default:0"` // 粉丝数
	FollowingCount int `gorm:"default:0"` // 关注数
	AvatarURL      string
}
