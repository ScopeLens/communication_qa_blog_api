package services

import (
	"communication_qa_blog_api/dao"
	"communication_qa_blog_api/models/tables"
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
	"time"
)

type ECodeReq struct {
	Email string `json:"email" binding:"required,email"`
}

// 生成验证码
func GenerateVerificationCode(length int) (string, error) {
	const charset = "0123456789"
	code := make([]byte, length)
	for i := range code {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		code[i] = charset[num.Int64()]
	}
	return string(code), nil
}

// 存储并发送验证码
func SendVerificationEmail(ctx *gin.Context) {
	var req ECodeReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	code, err := GenerateVerificationCode(6)
	if err != nil {
		fmt.Println("验证码生成失败", err)
	}
	verificationCode := tables.VerificationCode{
		Email:     req.Email,
		Code:      code,
		ExpiresAt: time.Now().Add(time.Minute * 5),
	}
	err = dao.CreateVerificationCode(verificationCode)
	if err != nil {
		return
	}

	//TODO 发送邮件
}

// 验证验证码
func VerifyCode(email, code string) (bool, error) {
	verificationCode, err := dao.FirstCodeByEmail(email, code)
	if err != nil {
		return false, err
	}
	if time.Now().After(verificationCode.ExpiresAt) {
		return false, errors.New("验证码已过期")
	}

	// 验证通过后将验证码标记为已使用
	verificationCode.IsUsed = true
	err = dao.SaveCode(verificationCode)
	if err != nil {
		return false, err
	}
	return true, nil
}
