package views

type UserItem struct {
	Username       string `json:"username"`
	Nickname       string `json:"nickname"`
	FollowersCount int    `json:"followers_count"`
	FollowingCount int    `json:"following_count"`
	AvatarURL      string `json:"avatar_url"`
}
