package services

import (
	"communication_qa_blog_api/dao"
	"communication_qa_blog_api/models"
	"communication_qa_blog_api/models/tables"
	"communication_qa_blog_api/models/views"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SearchUsersByNickname(c *gin.Context) {
	var users *[]views.UserItem
	nickname := c.Query("nickname") // 获取昵称的查询参数

	if nickname == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nickname query parameter is required"})
		return
	}

	// 查询昵称中包含关键字的用户
	users, err := dao.FindUsersByNickname(nickname)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func SearchPostsByTitleOrTag(ctx *gin.Context) {
	var posts []tables.Post
	title := ctx.Query("title") // 获取标题的查询参数
	tag := ctx.Query("tag")     // 获取标签的查询参数

	query := models.DB.Model(&tables.Post{})

	// 如果有 title 查询参数，根据 title 搜索
	if title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}

	// 如果有 tag 查询参数，根据 tag 搜索
	if tag != "" {
		query = query.Where("tags LIKE ?", "%"+tag+"%")
	}

	if err := query.Find(&posts).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No posts found"})
		return
	}

	ctx.JSON(http.StatusOK, posts)
}
