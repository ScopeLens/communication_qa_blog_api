package services

import (
	"communication_qa_blog_api/dao"
	"communication_qa_blog_api/models/tables"
	"communication_qa_blog_api/models/views"
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
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		files := form.File["images"] // 获取所有上传的图片
		postImgPath := fmt.Sprintf("%s_%d", username, time.Now().Unix())
		for _, file := range files {
			filePath := fmt.Sprintf("/uploadFiles/postImg/%s/%s", postImgPath, file.Filename) // 图片存储路径
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
		Images:    imagesJSON, // 保存图片 URL 列表
		Tags:      tagsJSON,   // 标签数组
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

type PostCommentResp struct {
	Post     views.PostDetail    `json:"post"`
	Comments []views.CommentTree `json:"comments"`
}

func CheckPostDetail(ctx *gin.Context) {
	var req PostIDReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post, err := dao.FirstPostDetailByID(req.PostID)
	// 使用一个 map 来建立父子关系
	var comments []views.CommentItem
	var commentTrees []views.CommentTree
	comments, err = dao.FindCommentsByPostID(req.PostID)

	// 将评论转换为 CommentTree 结构体
	commentMap := make(map[uint][]views.CommentTree)
	for _, comment := range comments {
		commentTree := views.CommentTree{
			CommentID: comment.CommentID,
			PostID:    comment.PostID,
			Username:  comment.Username,
			Nickname:  comment.Nickname,
			AvatarURL: comment.AvatarURL,
			Content:   comment.Content,
			ParentID:  comment.ParentID,
			CreatedAt: comment.CreatedAt,
		}

		if comment.ParentID == nil {
			// 如果 ParentID 为空，说明是顶级评论
			commentTrees = append(commentTrees, commentTree)
		} else {
			// 否则，将其加入到父评论的子评论中
			commentMap[*comment.ParentID] = append(commentMap[*comment.ParentID], commentTree)
		}
	}

	// 递归填充每个顶级评论的子评论
	for i := range commentTrees {
		populateReplies(&commentTrees[i], commentMap)
	}

	resp := &PostCommentResp{
		Post:     post,
		Comments: commentTrees,
	}

	ctx.JSON(http.StatusOK, resp)
}

// 递归填充子评论
func populateReplies(commentTree *views.CommentTree, commentMap map[uint][]views.CommentTree) {
	replies, exists := commentMap[commentTree.CommentID]
	if exists {
		commentTree.Replies = replies
		for i := range commentTree.Replies {
			populateReplies(&commentTree.Replies[i], commentMap)
		}
	}
}

func GetPostListByType(ctx *gin.Context) {
	SortType := ctx.Query("sort_type")
	var postList []views.PostDetail
	switch SortType {
	case "1":
		postList = dao.FindPostDetailSortByTime()
	case "2":
		postList = dao.FindPostDetailSortByLike()
	case "3":
		postList = dao.FindPostDetailSortByFavorite()
	case "4":
		postList = dao.FindPostDetailSortByReply()
	case "5":
		postList = dao.FindPostDetailSortByView()
	}
	ctx.JSON(http.StatusOK, postList)
}

func AddFavoritesCount(ctx *gin.Context) {
	var req PostIDReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		return
	}
	username := ctx.GetString("username")
	postID := req.PostID
	err = dao.IncrementFavoritesCount(postID)
	if err != nil {
		return
	}
	err = dao.CreateFavorite(username, postID)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"massage": "收藏成功",
	})
}
func DelFavoritesCount(ctx *gin.Context) {
	var req PostIDReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		return
	}
	username := ctx.GetString("username")
	postID := req.PostID
	err = dao.DecrementFavoritesCount(postID)
	if err != nil {
		return
	}
	err = dao.DeleteFavorite(username, postID)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"massage": "取消收藏成功",
	})
}
func AddLikesCount(ctx *gin.Context) {
	var req PostIDReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		return
	}
	username := ctx.GetString("username")
	postID := req.PostID
	err = dao.IncrementLikesCount(postID)
	if err != nil {
		return
	}
	err = dao.CreateLike(username, postID)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"massage": "点赞成功",
	})
}
func DelLikesCount(ctx *gin.Context) {
	var req PostIDReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		return
	}
	username := ctx.GetString("username")
	postID := req.PostID
	err = dao.DecrementLikesCount(postID)
	if err != nil {
		return
	}
	err = dao.DeleteLike(username, postID)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"massage": "取消点赞成功",
	})
}
func AddViewsCount(ctx *gin.Context) {
	var req PostIDReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		return
	}
	username := ctx.GetString("username")
	view := tables.View{
		Username: username,
		PostID:   req.PostID,
	}
	err = dao.CreateView(view)
	if err != nil {
		return
	}
	err = dao.IncrementViewsCount(req.PostID)
	if err != nil {
		return
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
func GetPostStatus(ctx *gin.Context) {
	var req PostIDReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		return
	}
	username := ctx.GetString("username")
	postStatus := views.PostStatus{
		IsViewed:   false,
		IsFavorite: false,
		IsLiked:    false,
	}
	if !dao.IsView(username, req.PostID) {
		postStatus.IsViewed = true
	}
	if !dao.IsFavorite(username, req.PostID) {
		postStatus.IsFavorite = true
	}
	if !dao.IsLike(username, req.PostID) {
		postStatus.IsLiked = true
	}

	ctx.JSON(http.StatusOK, postStatus)
}

type CommentReq struct {
	PostID   uint   `json:"post_id" binding:"required"`
	Content  string `json:"content" binding:"required"`
	ParentID *uint  `json:"parent_id"`
}

func AddRepliesCount(ctx *gin.Context) {
	var req CommentReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	comment := tables.Comment{
		Username: ctx.GetString("username"),
		PostID:   req.PostID,
		Content:  req.Content,
		ParentID: req.ParentID,
	}
	err = dao.IncrementRepliesCount(comment.PostID)
	if err != nil {
		return
	}
	err = dao.CreateComment(comment)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"massage": "评论成功",
	})
}

type DelCommentReq struct {
	PostID    uint `json:"post_id" binding:"required"`
	CommentID uint `json:"comment_id" binding:"required"`
}

func DelRepliesCount(ctx *gin.Context) {
	var req DelCommentReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		return
	}
	err = dao.DecrementRepliesCount(req.PostID)
	if err != nil {
		return
	}
	err = dao.DeleteComment(req.CommentID)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"massage": "删除评论成功",
	})
}
