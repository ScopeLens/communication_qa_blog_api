package dao

import (
	"communication_qa_blog_api/models"
	"communication_qa_blog_api/models/tables"
	"errors"
	"gorm.io/gorm"
)

func CreateLike(username string, PostID uint) error {
	like := tables.Like{
		Username: username,
		PostID:   PostID,
	}
	err := models.DB.Create(&like).Error
	if err != nil {
		return err
	}
	return nil
}

func IsLike(username string, PostID uint) bool {
	var like tables.Like
	err := models.DB.Model(&tables.Like{}).Where("username = ? And post_id=?", username, PostID).First(&like).Error
	// 查询结果为空
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return true
	}
	return false
}

func DeleteLike(username string, PostID uint) error {
	like := tables.Like{
		Username: username,
		PostID:   PostID,
	}
	err := models.DB.Delete(&like).Error
	if err != nil {
		return err
	}
	return nil
}
