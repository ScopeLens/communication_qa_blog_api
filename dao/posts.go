package dao

import (
	"communication_qa_blog_api/models"
	"communication_qa_blog_api/models/tables"
	"communication_qa_blog_api/models/views"
	"gorm.io/gorm"
)

// 创建新的帖子
func CreatePost(post tables.Post) error {
	if err := models.DB.Create(&post).Error; err != nil {
		return err
	}
	return nil
}

// 删除帖子
func DeletePost(postId uint) error {
	if err := models.DB.Where("post_id = ?", postId).
		Delete(&tables.Post{}).Error; err != nil {
		return err
	}
	return nil
}

// 查看帖子
func FirstPostByID(postId uint) (*tables.Post, error) {
	var post tables.Post
	if err := models.DB.
		Where("post_id = ?", postId).
		First(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func FirstPostDetailByID(postId uint) (views.PostDetail, error) {
	var post views.PostDetail
	if err := models.DB.Model(&tables.Post{}).
		Select("posts.*, users.nickname,users.avatar_url").
		Joins("JOIN users ON posts.username = users.username").
		Where("post_id = ?", postId).
		First(&post).Error; err != nil {
		return views.PostDetail{}, err
	}
	return post, nil
}

// 查看帖子细节
func FindPostDetailSortByTime() []views.PostDetail {
	var posts []views.PostDetail
	models.DB.Model(&tables.Post{}).
		Select("posts.*, users.nickname").
		Joins("JOIN users ON posts.username = users.username").
		Order("posts.updated_at DESC").
		Scan(&posts)
	return posts
}
func FindPostDetailSortByReply() []views.PostDetail {
	var posts []views.PostDetail
	models.DB.Model(&tables.Post{}).
		Select("posts.*, users.nickname,users.avatar_url").
		Joins("JOIN users ON posts.username = users.username").
		Order("posts.reply_count DESC").
		Scan(&posts)
	return posts
}
func FindPostDetailSortByView() []views.PostDetail {
	var posts []views.PostDetail
	models.DB.Model(&tables.Post{}).
		Select("posts.*, users.nickname,users.avatar_url").
		Joins("JOIN users ON posts.username = users.username").
		Order("posts.views_count DESC").
		Scan(&posts)
	return posts
}
func FindPostDetailSortByFavorite() []views.PostDetail {
	var posts []views.PostDetail
	models.DB.Model(&tables.Post{}).
		Select("posts.*, users.nickname,users.avatar_url").
		Joins("JOIN users ON posts.username = users.username").
		Order("posts.favorites_count DESC").
		Scan(&posts)
	return posts
}
func FindPostDetailSortByLike() []views.PostDetail {
	var posts []views.PostDetail
	models.DB.Model(&tables.Post{}).
		Select("posts.*, users.nickname,users.avatar_url").
		Joins("JOIN users ON posts.username = users.username").
		Order("posts.likes_count DESC").
		Scan(&posts)
	return posts
}

func FindPostDetailByTitle(title string) []views.PostDetail {
	var posts []views.PostDetail
	models.DB.Model(&tables.Post{}).
		Select("posts.*, users.nickname,users.avatar_url").
		Joins("JOIN users ON posts.username = users.username").
		Where("posts.title like ?", "%"+title+"%").
		Order("posts.likes_count DESC").
		Scan(&posts)
	return posts
}

func FindPostDetailByTag(tagName string) []views.PostDetail {
	var posts []views.PostDetail
	models.DB.Model(&tables.Post{}).
		Select("posts.*, users.nickname,users.avatar_url").
		Joins("JOIN users ON posts.username = users.username").
		Joins("JOIN tags ON posts.post_id = tags.post_id").
		Where("tags.tag_name = ?", tagName).
		Order("posts.likes_count DESC").
		Scan(&posts)
	return posts
}

func FindPostDetailByUsername(username string) []views.PostDetail {
	var posts []views.PostDetail
	models.DB.Model(&tables.Post{}).
		Select("posts.*, users.nickname,users.avatar_url").
		Joins("JOIN users ON posts.username = users.username").
		Where("users.username = ?", username).
		Order("posts.likes_count DESC").
		Scan(&posts)
	return posts
}

// 更新收藏数
func IncrementFavoritesCount(postID uint) error {
	return models.DB.Model(&tables.Post{}).Where("post_id = ?", postID).Update("favorites_count", gorm.Expr("favorites_count + 1")).Error
}
func DecrementFavoritesCount(postID uint) error {
	return models.DB.Model(&tables.Post{}).Where("post_id = ?", postID).Update("favorites_count", gorm.Expr("favorites_count - 1")).Error
}

// 更新点赞数
func IncrementLikesCount(postID uint) error {
	return models.DB.Model(&tables.Post{}).Where("post_id = ?", postID).Update("likes_count", gorm.Expr("likes_count + 1")).Error
}
func DecrementLikesCount(postID uint) error {
	return models.DB.Model(&tables.Post{}).Where("post_id = ?", postID).Update("likes_count", gorm.Expr("likes_count - 1")).Error
}

// 更新浏览量
func IncrementViewsCount(postID uint) error {
	return models.DB.Model(&tables.Post{}).Where("post_id = ?", postID).Update("views_count", gorm.Expr("views_count + 1")).Error
}
func DecrementViewsCount(postID uint) error {
	return models.DB.Model(&tables.Post{}).Where("post_id = ?", postID).Update("views_count", gorm.Expr("views_count - 1")).Error
}

// 更新评论数
func IncrementRepliesCount(postID uint) error {
	return models.DB.Model(&tables.Post{}).Where("post_id = ?", postID).Update("reply_count", gorm.Expr("reply_count + 1")).Error
}
func DecrementRepliesCount(postID uint) error {
	return models.DB.Model(&tables.Post{}).Where("post_id = ?", postID).Update("reply_count", gorm.Expr("reply_count - 1")).Error
}
