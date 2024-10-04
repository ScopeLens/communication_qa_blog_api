package services

import (
	"communication_qa_blog_api/dao"
	"communication_qa_blog_api/models/tables"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type RegisterReq struct {
	Username string `json:"username" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Register(ctx *gin.Context) {
	var req RegisterReq
	//绑定结构体
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "缺少必要参数"})
		return
	}
	//账号校验
	exist, err := dao.IsExist(req.Username)
	if err != nil {
		fmt.Println("账号校验 系统错误")
		return
	}
	if exist {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "账号已存在"})
		return
	}

	//密码加密
	hashPassword, err := encryptedPassword(req.Password)
	fmt.Println(hashPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 账号信息持久化
	user := tables.User{
		Username: req.Username,
		Nickname: req.Nickname,
		Email:    req.Email,
		Password: hashPassword,
	}
	err = dao.CreateUser(user)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "注册成功",
	})
}

// 账号加密
func encryptedPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("encryptedPassword err:", err)
		return "", err
	}
	return string(hashPassword), err
}
