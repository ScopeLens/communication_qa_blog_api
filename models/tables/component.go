package tables

import (
	"communication_qa_blog_api/models"
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type Component struct {
	ID          uint            `gorm:"primary_key;AUTO_INCREMENT"`
	Name        string          `gorm:"type:varchar(255);not null"`
	Description string          `gorm:"type:text"`      // 组件描述
	Dimensions  json.RawMessage `gorm:"type:json"`      // 组件尺寸 (长宽存储为JSON)
	CreatedAt   time.Time       `gorm:"autoCreateTime"` // 创建时间
	UpdatedAt   time.Time       `gorm:"autoUpdateTime"` // 更新时间
	DeletedAt   gorm.DeletedAt  `gorm:"index"`          // 删除时间，支持软删除
}

func CreateComponent(component Component) error {
	if err := models.DB.Create(&component).Error; err != nil {
		return err
	}
	return nil
}
func FirstComponent(id uint) (*Component, error) {
	var component Component
	if err := models.DB.First(&component, id).Error; err != nil {
		return nil, err
	}
	return &component, nil
}
