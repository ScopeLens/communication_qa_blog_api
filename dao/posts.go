package dao

import (
	"communication_qa_blog_api/models"
	"communication_qa_blog_api/models/tables"
	"gorm.io/gorm"
)

// 创建新的帖子
func CreatePost(post tables.Post) error {
	if err := models.DB.Create(&post).Error; err != nil {
		return err
	}
	return nil
}

// 删除帖子
func DeletePost(postId uint) error {
	if err := models.DB.Where("post_id = ?", postId).
		Delete(&tables.Post{}).Error; err != nil {
		return err
	}
	return nil
}

// 查看帖子
func FirstPostByID(postId uint) (*tables.Post, error) {
	var post tables.Post
	if err := models.DB.
		Where("post_id = ?", postId).
		First(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

// 更新收藏数
func IncrementFavoritesCount(postID uint) error {
	return models.DB.Model(&tables.Post{}).Where("post_id = ?", postID).Update("favorites_count", gorm.Expr("favorites_count + 1")).Error
}
func DecrementFavoritesCount(postID uint) error {
	return models.DB.Model(&tables.Post{}).Where("post_id = ?", postID).Update("favorites_count", gorm.Expr("favorites_count - 1")).Error
}

// 更新点赞数
func IncrementLikesCount(postID uint) error {
	return models.DB.Model(&tables.Post{}).Where("post_id = ?", postID).Update("likes_count", gorm.Expr("likes_count + 1")).Error
}
func DecrementLikesCount(postID uint) error {
	return models.DB.Model(&tables.Post{}).Where("post_id = ?", postID).Update("likes_count", gorm.Expr("likes_count - 1")).Error
}

// 更新浏览量
func IncrementViewsCount(postID uint) error {
	return models.DB.Model(&tables.Post{}).Where("post_id = ?", postID).Update("views_count", gorm.Expr("views_count + 1")).Error
}
func DecrementViewsCount(postID uint) error {
	return models.DB.Model(&tables.Post{}).Where("post_id = ?", postID).Update("views_count", gorm.Expr("views_count - 1")).Error
}
