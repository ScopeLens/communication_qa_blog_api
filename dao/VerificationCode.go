package dao

import (
	"communication_qa_blog_api/models"
	"communication_qa_blog_api/models/tables"
	"fmt"
)

func CreateVerificationCode(code tables.VerificationCode) error {
	if err := models.DB.Create(&code).Error; err != nil {
		fmt.Println("验证码存储失败", err)
		return err
	}
	return nil
}

func FirstCodeByEmail(email string, code string) (*tables.VerificationCode, error) {
	var verificationCode tables.VerificationCode
	err := models.DB.Where("email = ? AND code = ? AND used = ?", email, code, false).
		First(&verificationCode).Error
	if err != nil {
		return nil, err
	}
	return &verificationCode, nil
}

func SaveCode(verificationCode *tables.VerificationCode) error {
	if err := models.DB.Save(verificationCode).Error; err != nil {
		return err
	}
	return nil
}
