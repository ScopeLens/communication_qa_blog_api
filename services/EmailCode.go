package services

import (
	"communication_qa_blog_api/dao"
	"communication_qa_blog_api/models/tables"
	"crypto/rand"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
	"math/big"
	"net/http"
	"time"
)

type ECodeReq struct {
	Email string `json:"email" binding:"required"`
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
		ExpiredAt: time.Now().Add(time.Minute * 5),
	}
	err = dao.CreateVerificationCode(verificationCode)
	if err != nil {
		return
	}
	CodeStr := fmt.Sprintf("您的验证码为：%s,有效时间为3分钟，请尽快输入", code)
	err = sendEmail(req.Email, "LenPark", CodeStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "邮件发送成功")
}

func sendEmail(toEmail, subject, body string) error {

	CodeEmail := gomail.NewMessage()
	// QQ 邮箱的 SMTP 服务器地址和端口号
	CodeEmail.SetHeaders(map[string][]string{
		"From":    {CodeEmail.FormatAddress("911263610@qq.com", "LenPark")},
		"To":      {toEmail},
		"Subject": {subject},
	})
	CodeEmail.SetBody("text/html", body)

	host := "smtp.qq.com"
	port := 587
	userName := "911263610@qq.com"
	password := "yfyrkkcawhgibchi" // qq邮箱填授权码
	d := gomail.NewDialer(
		host,
		port,
		userName,
		password,
	)
	// 发送邮件
	err := d.DialAndSend(CodeEmail)
	if err != nil {
		fmt.Println("邮件发送失败 err:", err)
		return err
	}

	return nil
}

type CodeReq struct {
	Email string `json:"email" binding:"required,email"`
	Code  string `json:"code" binding:"required"`
}

// 验证验证码
func VerifyCode(ctx *gin.Context) {
	var req CodeReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	verificationCode, err := dao.FirstCodeByEmail(req.Email, req.Code)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if time.Now().After(verificationCode.ExpiredAt) {
		ctx.JSON(http.StatusOK, false)
		return
	}

	// 验证通过后将验证码标记为已使用
	verificationCode.IsUsed = true
	err = dao.SaveCode(verificationCode)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, true)
}
