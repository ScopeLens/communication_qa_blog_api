package dao

import (
	"communication_qa_blog_api/models"
	"communication_qa_blog_api/models/tables"
)

func FindAllTag() ([]tables.Tag, error) {
	var tags []tables.Tag
	if err := models.DB.Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}
