package dao

import (
	"communication_qa_blog_api/models"
	"communication_qa_blog_api/models/tables"
	"communication_qa_blog_api/models/views"
)

func CreateComment(comment tables.Comment) error {
	if err := models.DB.Model(tables.Comment{}).Create(&comment).Error; err != nil {
		return err
	}
	return nil
}

func DeleteComment(commentId uint) error {
	if err := models.DB.Where("comment_id = ? or parent_id=?", commentId, commentId).
		Delete(&tables.Comment{}).Error; err != nil {
		return err
	}
	return nil
}

func FindCommentIDsByUsername(username string) ([]uint, error) {
	var commentIDs []uint
	if err := models.DB.Table("comments").Where("username = ? AND parent_id IS NULL", username).Pluck("comment_id", &commentIDs).Error; err != nil {
		return nil, err
	}
	return commentIDs, nil
}

func FindRepliesByParentID(commentIDs []uint) ([]tables.Comment, error) {
	var replies []tables.Comment
	if err := models.DB.Where("parent_id IN ?", commentIDs).Find(&replies).Error; err != nil {
		return nil, err
	}
	return replies, nil
}

func FindCommentsByPostID(postID uint) ([]views.CommentItem, error) {
	var comments []views.CommentItem
	if err := models.DB.
		Model(&tables.Comment{}).
		Select("comments.*, users.nickname,users.avatar_url").
		Joins("JOIN users ON comments.username = users.username").
		Where("post_id = ?", postID).
		Order("parent_id ASC").
		Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func FindCommentsByUsername(username string) ([]tables.Comment, error) {
	var comments []tables.Comment
	if err := models.DB.Table("comments").Where("username = ?", username).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
