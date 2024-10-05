package services

import (
	"communication_qa_blog_api/dao"
	"communication_qa_blog_api/models/tables"
	"communication_qa_blog_api/models/views"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SearchUsersByNickname(ctx *gin.Context) {
	var users []views.UserItem
	nickname := ctx.Query("nickname") // 获取昵称的查询参数

	if nickname == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "nickname query parameter is required"})
		return
	}

	// 查询昵称中包含关键字的用户
	users, err := dao.FindUsersByNickname(nickname)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func SearchUsersByUsername(ctx *gin.Context) {
	var user *tables.User
	username := ctx.GetString("username") // 获取昵称的查询参数

	if username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username query parameter is required"})
		return
	}

	// 查询昵称中包含关键字的用户
	user, err := dao.FirstByUsername(username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func SearchPostsByTitle(ctx *gin.Context) {
	var posts []views.PostDetail
	title := ctx.Query("title") // 获取标题的查询参数
	if title == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "title query parameter is required"})
		return
	}
	posts = dao.FindPostDetailByTitle(title)

	ctx.JSON(http.StatusOK, posts)
}

func SearchPostsByTag(ctx *gin.Context) {
	var posts []views.PostDetail
	tagName := ctx.Query("tag_name")
	if tagName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "tag query parameter is required"})
		return
	}
	posts = dao.FindPostDetailByTag(tagName)
	ctx.JSON(http.StatusOK, posts)
}

func SearchPostsByUsername(ctx *gin.Context) {
	var posts []views.PostDetail
	username := ctx.Query("username")
	if username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "tag query parameter is required"})
		return
	}
	posts = dao.FindPostDetailByUsername(username)
	ctx.JSON(http.StatusOK, posts)
}
