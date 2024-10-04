package tables

type Tag struct {
	TagID uint   `gorm:"primaryKey"`
	Name  string `gorm:"unique;not null"` // 标签名称
}
