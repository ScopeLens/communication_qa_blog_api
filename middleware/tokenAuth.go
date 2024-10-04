package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNzI4NTYyMzAzfQ.2ZI98NWo7Lxiecueiu-msPpGyhHpaeh31WxCOaPTScg

var JwtKey = []byte("ScopeLens") // 你可以换成更安全的密钥

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func JWTAuthMiddleware(ctx *gin.Context) {
	// 从请求头中提取Authorization字段
	authHeader := ctx.Request.Header.Get("Authorization")
	if authHeader == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		ctx.Abort()
	}

	// 分离 "Bearer" 和 Token 部分
	tokenString := strings.Split(authHeader, "Bearer ")[1]

	// 验证和解析 token
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		ctx.Abort()
	}

	// 将解析到的用户名存入 context
	ctx.Set("username", claims.Username)

	ctx.Next()
}
