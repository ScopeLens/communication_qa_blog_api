package services

import (
	"communication_qa_blog_api/dao"
	"communication_qa_blog_api/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type LoginRsp struct {
	Nickname string `json:"nickname"`
	Username string `json:"username"`
	Token    string `json:"token"`
	Message  string `json:"message"`
}

func Login(ctx *gin.Context) {
	var req LoginReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 查找用户
	user, err := dao.FirstByUsername(req.Username)
	if err != nil {
		fmt.Println("账号不存在")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "账号不存在"})
		return
	}

	//密码校验 明文存储
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "账号密码错误"})
		return
	}

	//生成token
	token, err := GenerateToken(user.Username)
	if err != nil {
		return
	}

	//登录结果返回
	ctx.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"data": &LoginRsp{
			Nickname: user.Nickname,
			Username: user.Username,
			Token:    token,
			Message:  "success",
		},
	})
}

func GenerateToken(username string) (string, error) {
	//设置jwt的声明信息
	expirationTime := time.Now().Add(7 * 24 * time.Hour) //过期时间一星期
	Claims := middleware.Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	//使用指定的算法和密钥生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)
	tokenString, err := token.SignedString(middleware.JwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
