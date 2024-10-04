package dao

import (
	"communication_qa_blog_api/models"
	"communication_qa_blog_api/models/tables"
)

func FindCommentIDsByUsername(username string) ([]uint, error) {
	var commentIDs []uint
	if err := models.DB.Table("comments").Where("username = ? AND parent_comment_id IS NULL", username).Pluck("comment_id", &commentIDs).Error; err != nil {
		return nil, err
	}
	return commentIDs, nil
}

func FindRepliesByParentID(commentIDs []uint) ([]tables.Comment, error) {
	var replies []tables.Comment
	if err := models.DB.Where("parent_comment_id IN ?", commentIDs).Find(&replies).Error; err != nil {
		return nil, err
	}
	return replies, nil
}
