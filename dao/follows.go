package dao

import (
	"communication_qa_blog_api/models"
	"communication_qa_blog_api/models/tables"
	"fmt"
)

func CreateFollow(follow tables.Follow) error {
	if err := models.DB.Create(&follow).Error; err != nil {
		fmt.Printf("Follow.Create err:%v\n", err)
		return err
	}
	return nil
}

func DeleteFollow(follow tables.Follow) error {
	if err := models.DB.Where("Follower = ? AND Followee = ?", follow.Follower, follow.Followee).Delete(&tables.Follow{}).Error; err != nil {
		fmt.Printf("Follow.delete err:%v\n", err)
		return err
	}
	return nil
}

// 查找关注的人的用户名列表
func PluckFolloweeByUsername(username string) ([]string, error) {
	var followees []string
	if err := models.DB.Model(&tables.Follow{}).Where("follower = ?", username).Pluck("followee", &followees).Error; err != nil {
		return nil, err
	}
	return followees, nil
}

// 查找粉丝的用户名列表
func PluckFollowerByUsername(username string) ([]string, error) {
	var followers []string
	if err := models.DB.Model(&tables.Follow{}).Where("followee = ?", username).Pluck("follower", &followers).Error; err != nil {
		return nil, err
	}
	return followers, nil
}
