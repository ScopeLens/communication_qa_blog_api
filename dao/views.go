package dao

import (
	"communication_qa_blog_api/models"
	"communication_qa_blog_api/models/tables"
	"errors"
	"gorm.io/gorm"
)

func CreateView(username string, PostID uint) error {
	view := tables.View{
		Username: username,
		PostID:   PostID,
	}
	err := models.DB.Create(&view).Error
	if err != nil {
		return err
	}
	return nil
}

func IsView(username string, PostID uint) bool {
	var view tables.View
	err := models.DB.Model(&tables.View{}).Where("username = ? And post_id=?", username, PostID).First(&view).Error
	// 查询结果为空
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return true
	}
	return false
}

func DeleteView(username string, PostID uint) error {
	view := tables.View{
		Username: username,
		PostID:   PostID,
	}
	err := models.DB.Delete(&view).Error
	if err != nil {
		return err
	}
	return nil
}

// 获得PostID切片
func GetViewPostList(username string) ([]tables.Post, error) {
	var posts []tables.Post
	err := models.DB.Joins("JOIN views ON views.post_id = posts.id").
		Where("views.username = ?", username).
		Find(&posts).Error
	return posts, err
}
