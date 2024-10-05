package dao

import (
	"communication_qa_blog_api/models"
	"communication_qa_blog_api/models/tables"
	"communication_qa_blog_api/models/views"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func IsExist(username string) (bool, error) {
	var user tables.User
	err := models.DB.Where("username = ?", username).First(&user).Error

	// 查询结果为空
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		fmt.Printf("Account.IsExist err:%v\n", err)
		return false, err
	}
	return true, nil
}

func CreateUser(user tables.User) error {
	if err := models.DB.Model(tables.User{}).Create(&user).Error; err != nil {
		fmt.Printf("AccountDao.Create err:%v\n", err)
		return err
	}
	return nil
}

func FirstByUsername(username string) (*tables.User, error) {
	var user tables.User
	err := models.DB.Where("Username=?", username).First(&user).Error
	if err != nil {
		fmt.Printf("AccountDao.FirstByUserID err:%v\n", err)
		return nil, err
	}
	return &user, nil
}

func SelectUserItem(usernames []string) ([]views.UserItem, error) {
	var userItems []views.UserItem
	if err := models.DB.Table("users").Where("username IN (?)", usernames).Select("username, nickname, avatar_url,followers_Count,following_Count").Scan(&userItems).Error; err != nil {
		return nil, err
	}
	return userItems, nil
}

func UpdateAvatarUrl(username string, avatarURL string) error {
	if err := models.DB.Model(&tables.User{}).
		Where("username = ?", username).
		Update("avatar_url", avatarURL).Error; err != nil {
		return err
	}
	return nil
}

func FindUsersByNickname(keyword string) ([]views.UserItem, error) {
	var userItems []views.UserItem
	if err := models.DB.Model(&tables.User{}).Where("nickname LIKE ?", "%"+keyword+"%").Find(&userItems).Error; err != nil {
		return nil, err
	}
	return userItems, nil
}
