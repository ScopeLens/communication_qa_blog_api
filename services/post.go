package services

import (
	"communication_qa_blog_api/dao"
	"communication_qa_blog_api/models/tables"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// 新建帖子
type CreatePostReq struct {
	Title   string   `json:"title" binding:"required"`   // 帖子标题
	Content string   `json:"content" binding:"required"` // 帖子内容
	Tags    []string `json:"tags"`                       // 标签数组
}

// 创建帖子
func CreatePost(ctx *gin.Context) {
	var req CreatePostReq
	req.Title = ctx.PostForm("title")
	req.Content = ctx.PostForm("content")
	req.Tags = ctx.PostFormArray("tags")
	// 从 token 中提取用户名
	username := ctx.GetString("username") // 假设使用了中间件提取用户名

	var imageUrls []string
	form, err := ctx.MultipartForm()
	fmt.Println(req, form, "!!!!!!!!!!")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "发布成功",
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		files := form.File["images"] // 获取所有上传的图片
		postImgPath := fmt.Sprintf("%s_%d", username, time.Now().Unix())
		for _, file := range files {
			filePath := fmt.Sprintf("uploadFiles/postImg/%s/%s", postImgPath, file.Filename) // 图片存储路径
			if err := ctx.SaveUploadedFile(file, filePath); err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
				return
			}
			imageUrls = append(imageUrls, filePath) // 添加保存后的图片 URL
		}
	}

	imagesJSON, err := json.Marshal(imageUrls)
	if err != nil {
		log.Fatalf("Error converting images to JSON: %v", err)
	}

	tagsJSON, err := json.Marshal(req.Tags)
	if err != nil {
		log.Fatalf("Error converting tags to JSON: %v", err)
	}

	post := tables.Post{
		Title:     req.Title,
		Content:   req.Content,
		Images:    string(imagesJSON), // 保存图片 URL 列表
		Tags:      string(tagsJSON),   // 标签数组
		Username:  username,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = dao.CreatePost(post)
	if err != nil {
		fmt.Println("创建帖子失败 err:", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "发布成功",
	})
}

// 删除帖子
type DeletePostReq struct {
	PostId uint `json:"post_id"`
}

func DeletePost(ctx *gin.Context) {
	var req DeletePostReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = dao.DeletePost(req.PostId)
	fmt.Println("这里是传输进来的postid", req.PostId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
}

type PostIDReq struct {
	PostID uint `json:"post_id" binding:"required"`
}

func CheckPostDetail(ctx *gin.Context) {
	var req PostIDReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post, err := dao.FirstPostByID(req.PostID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, post)
}

func AddFavoritesCount(ctx *gin.Context) {
	var req PostIDReq
	user, _ := ctx.Get("username")
	username := user.(string)
	postID := req.PostID
	if dao.IsFavorite(username, postID) {
		err := dao.IncrementFavoritesCount(postID)
		if err != nil {
			return
		}
		err = dao.CreateFavorite(username, postID)
		if err != nil {
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"massage": "收藏成功",
	})
}
func DelFavoritesCount(ctx *gin.Context) {
	var req PostIDReq
	user, _ := ctx.Get("username")
	username := user.(string)
	postID := req.PostID
	if !dao.IsFavorite(username, postID) {
		err := dao.DecrementFavoritesCount(postID)
		if err != nil {
			return
		}
		err = dao.DeleteFavorite(username, postID)
		if err != nil {
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"massage": "取消收藏成功",
	})
}
func AddLikesCount(ctx *gin.Context) {
	var req PostIDReq
	user, _ := ctx.Get("username")
	username := user.(string)
	postID := req.PostID
	if dao.IsLike(username, postID) {
		err := dao.IncrementLikesCount(postID)
		if err != nil {
			return
		}
		err = dao.CreateLike(username, postID)
		if err != nil {
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"massage": "点赞成功",
	})
}
func DelLikesCount(ctx *gin.Context) {
	var req PostIDReq
	user, _ := ctx.Get("username")
	username := user.(string)
	postID := req.PostID
	if !dao.IsLike(username, postID) {
		err := dao.DecrementLikesCount(postID)
		if err != nil {
			return
		}
		err = dao.DeleteLike(username, postID)
		if err != nil {
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"massage": "取消点赞成功",
	})
}
func AddViewsCount(ctx *gin.Context) {
	var req PostIDReq
	user, _ := ctx.Get("username")
	username := user.(string)
	postID := req.PostID
	if dao.IsView(username, postID) {
		err := dao.IncrementViewsCount(postID)
		if err != nil {
			return
		}
		err = dao.CreateView(username, postID)
		if err != nil {
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"massage": "浏览成功",
	})
}
func DelViewsCount(ctx *gin.Context) {
	var req PostIDReq
	user, _ := ctx.Get("username")
	username := user.(string)
	postID := req.PostID
	if dao.IsView(username, postID) {
		err := dao.DecrementViewsCount(postID)
		if err != nil {
			return
		}
		err = dao.DeleteView(username, postID)
		if err != nil {
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"massage": "取消浏览成功",
	})
}
