package dao

import (
	"communication_qa_blog_api/models"
	"communication_qa_blog_api/models/tables"
	"communication_qa_blog_api/models/views"
	"errors"
	"gorm.io/gorm"
)

func CreateView(view tables.View) error {
	err := models.DB.Create(&view).Error
	if err != nil {
		return err
	}
	return nil
}

func IsView(username string, PostID uint) bool {
	var view tables.View
	err := models.DB.Model(&tables.View{}).Where("username = ? And post_id = ?", username, PostID).First(&view).Error
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

// 获得浏览过的post
func GetViewPostList(username string) ([]views.PostDetail, error) {
	var posts []views.PostDetail
	models.DB.Model(&tables.Post{}).
		Select("posts.*, users.nickname").
		Joins("JOIN users ON posts.username = users.username").
		Joins("JOIN views ON posts.post_id = views.post_id").
		Where("views.username = ?", username).
		Order("views.created_at DESC").
		Scan(&posts)
	return posts, nil
}
