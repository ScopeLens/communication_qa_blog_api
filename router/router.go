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
	r.GET("/uploadFiles/postImg/:folder/:filename", func(c *gin.Context) {
		folder := c.Param("folder")
		filename := c.Param("filename")
		c.File("/uploadFiles/postImg/" + folder + "/" + filename) // 动态返回 ./images 文件夹下的图片
	})

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
			post.GET("/post-sort", services.GetPostListByType)   //排序并获得帖子
			post.POST("/post-status", services.GetPostStatus)    //查看帖子是否被收藏点赞浏览
		}

		postInfo := root.Group("/post-detail")
		{
			postInfo.POST("/add-favorite", services.AddFavoritesCount) //收藏帖子
			postInfo.POST("/del-favorite", services.DelFavoritesCount) //取消收藏
			postInfo.POST("/add-view", services.AddViewsCount)         //新增浏览
			postInfo.POST("/del-view", services.DelViewsCount)         //删除浏览
			postInfo.POST("/add-like", services.AddLikesCount)         //点赞帖子
			postInfo.POST("/del-like", services.DelLikesCount)         //取消点赞
			postInfo.POST("/add-reply", services.AddRepliesCount)      //添加评论
			postInfo.POST("/del-reply", services.DelRepliesCount)      //删除评论
		}

		search := root.Group("/search")
		{
			search.GET("/users", services.SearchUsersByNickname)         //搜索用户 使用昵称
			search.GET("/user-self", services.SearchUsersByUsername)     //搜索用户 使用用户名
			search.GET("/post-title", services.SearchPostsByTitle)       //搜索帖子 ByTitle
			search.GET("/post-tag", services.SearchPostsByTag)           //搜索帖子 ByTag
			search.GET("/post-username", services.SearchPostsByUsername) //搜索帖子 ByTag
		}
	}
}
