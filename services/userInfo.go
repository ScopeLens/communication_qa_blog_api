package services

import (
	"communication_qa_blog_api/dao"
	"communication_qa_blog_api/models/tables"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

//TODO 修改个人信息
//TODO 搜索用户，搜索帖子

// UpdateAvatar 上传头像
func UpdateAvatar(ctx *gin.Context) {
	// 从中间件中获取用户名
	username, exists := ctx.Get("username")
	userID := username.(string)
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "认证失败"})
		return
	}

	// 获取上传的文件
	file, err := ctx.FormFile("avatar")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No file is received"})
		return
	}

	// 验证文件类型（仅允许图片）  前端可以实现
	//allowedExtensions := []string{".jpg", ".jpeg", ".png", ".gif"}
	ext := strings.ToLower(filepath.Ext(file.Filename))
	//isAllowed := false
	//for _, a := range allowedExtensions {
	//	if ext == a {
	//		isAllowed = true
	//		break
	//	}
	//}
	//if !isAllowed {
	//	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported file type"})
	//	return
	//}

	// 生成唯一的文件名，避免冲突
	filename := fmt.Sprintf("%s_%d%s", userID, time.Now().Unix(), ext)
	avatarFilepath := filepath.Join("uploadFiles/avatars", filename)

	// 保存文件到服务器
	if err := ctx.SaveUploadedFile(file, avatarFilepath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
		return
	}

	// 生成文件的访问 URL
	avatarURL := fmt.Sprintf("/uploadFiles/avatars/%s", filename)

	//删除旧的头像
	user, _ := dao.FirstByUsername(userID)
	oldFilePath := user.AvatarURL // 根据您的 URL 结构调整路径
	if _, err := os.Stat(oldFilePath); err == nil {
		err := os.Remove(oldFilePath)
		if err != nil {
			return
		}
	}
	// 更新用户的 avatar_url
	err = dao.UpdateAvatarUrl(userID, avatarURL)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update avatar URL"})
		return
	}

	// 返回成功响应
	ctx.JSON(http.StatusOK, gin.H{
		"message":    "头像更新成功",
		"avatar_url": avatarURL,
	})
}

// 查看自己的收藏
func ShowFavorite(ctx *gin.Context) {
	username := ctx.GetString("username")
	postList, err := dao.GetFavoritePostList(username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get post list"})
		return
	}
	ctx.JSON(http.StatusOK, postList)

}

// 查看自己的浏览历史
func ShowView(ctx *gin.Context) {
	username := ctx.GetString("username")
	postList, err := dao.GetViewPostList(username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get post list"})
		return
	}
	ctx.JSON(http.StatusOK, postList)
}

// 获得自己的消息
func GetRepliesToMyComments(ctx *gin.Context) {
	var replies []tables.Comment
	username := ctx.GetString("username") // 假设通过Token获取登录用户的username
	// 获取用户的所有评论ID
	commentIDs, err := dao.FindCommentIDsByUsername(username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get comment ids"})
		return
	}

	if len(commentIDs) == 0 {
		ctx.JSON(http.StatusOK, gin.H{"message": "No comments found"})
		return
	}

	// 获取所有回复这些评论的回复
	replies, err = dao.FindRepliesByParentID(commentIDs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get replies"})
		return
	}

	ctx.JSON(http.StatusOK, replies)
}
