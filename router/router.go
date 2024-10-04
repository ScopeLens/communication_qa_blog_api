package api

import (
	"communication_qa_blog_api/middleware"
	"communication_qa_blog_api/services"
	"github.com/gin-gonic/gin"
)

const (
	rootPath   = "/api"
	noAuthPath = "/out/api"
)

// 鉴定权限
func BasicRouter(r *gin.Engine) {
	r.Static("/uploadFiles/avatars", "./uploadFiles/avatars")
	r.Static("/uploadFiles/postImg", "./uploadFiles/postImg")
	noAuth := r.Group(noAuthPath)
	{
		noAuth.POST("/register", services.Register)                   //注册
		noAuth.POST("/login", services.Login)                         //登录
		noAuth.POST("/sendEmailCode", services.SendVerificationEmail) //发送验证码
	}

	root := r.Group(rootPath)
	root.Use(middleware.JWTAuthMiddleware)
	{
		personalInfo := root.Group("/personal-detail")
		{
			personalInfo.POST("/unfollow-user", services.UnfollowUser)          //取关用户
			personalInfo.POST("/follow-user", services.FollowUser)              //关注用户
			personalInfo.POST("/upload-avatar", services.UpdateAvatar)          //更新头像
			personalInfo.POST("/followers", services.GetFollowers)              //查看粉丝
			personalInfo.POST("/following", services.GetFollowing)              //查看关注
			personalInfo.POST("/show-favorites", services.ShowFavorite)         //查看收藏
			personalInfo.POST("/show-views", services.ShowView)                 //查看浏览历史
			personalInfo.POST("/show-message", services.GetRepliesToMyComments) //查看消息
		}

		post := root.Group("/post")
		{
			post.POST("/add-post", services.CreatePost)          //发布帖子
			post.POST("/del-post", services.DeletePost)          //删除帖子
			post.POST("/check-detail", services.CheckPostDetail) //查看帖子内容
		}

		postInfo := root.Group("/post-detail")
		{
			postInfo.POST("/add-favorite", services.AddFavoritesCount) //收藏帖子
			postInfo.POST("/del-favorite", services.DelFavoritesCount) //取消收藏
			postInfo.POST("/add-view", services.AddViewsCount)         //新增浏览
			postInfo.POST("/del-view", services.DelViewsCount)         //删除浏览
			postInfo.POST("/add-like", services.AddLikesCount)         //点赞帖子
			postInfo.POST("/del-like", services.DelLikesCount)         //取消点赞
		}

		search := root.Group("/search")
		{
			search.GET("/users", services.SearchUsersByNickname)   //搜索用户
			search.GET("/posts", services.SearchPostsByTitleOrTag) //搜索帖子
		}
	}
}
