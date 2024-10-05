package tables

import "time"

type VerificationCode struct {
	ID        uint      `gorm:"primaryKey"`       // 自增ID
	Email     string    `gorm:"not null"`         // 接收验证码的邮箱
	Code      string    `gorm:"size:10;not null"` // 验证码，长度限制为10个字符
	CreatedAt time.Time `gorm:"autoCreateTime"`   // 创建时间
	ExpiredAt time.Time `gorm:"not null"`         // 过期时间
	IsUsed    bool      `gorm:"default:false"`    // 是否已使用
}
