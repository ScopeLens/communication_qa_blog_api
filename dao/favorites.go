package dao

import (
	"communication_qa_blog_api/models"
	"communication_qa_blog_api/models/tables"
	"errors"
	"gorm.io/gorm"
)

// 新建收藏
func CreateFavorite(username string, PostID uint) error {
	favorite := tables.Favorite{
		Username: username,
		PostID:   PostID,
	}
	err := models.DB.Create(&favorite).Error
	if err != nil {
		return err
	}
	return nil
}

// 是否已经收藏
func IsFavorite(username string, PostID uint) bool {
	var favorite tables.Favorite
	err := models.DB.Model(&tables.Favorite{}).Where("username = ? And post_id=?", username, PostID).First(&favorite).Error
	// 查询结果为空
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return true
	}
	return false
}

// 取消收藏
func DeleteFavorite(username string, PostID uint) error {
	favorite := tables.Favorite{
		Username: username,
		PostID:   PostID,
	}
	err := models.DB.Delete(&favorite).Error
	if err != nil {
		return err
	}
	return nil
}

// 获得PostID切片
func GetFavoritePostList(username string) ([]tables.Post, error) {
	var posts []tables.Post
	err := models.DB.Joins("JOIN favorites ON favorites.post_id = posts.id").
		Where("favorites.username = ?", username).
		Find(&posts).Error
	return posts, err
}
