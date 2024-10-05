package services

import (
	"communication_qa_blog_api/dao"
	"communication_qa_blog_api/models/views"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取关注列表
type FollowingReq struct {
	Username string `json:"username"`
}
type FollowingResp struct {
	UserItem []views.UserItem
}

func GetFollowing(ctx *gin.Context) {
	var req FollowingReq
	var userItems []views.UserItem

	followees, err := dao.PluckFolloweeByUsername(req.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(followees) != 0 {
		userItems, err = dao.SelectUserItem(followees)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	ctx.JSON(http.StatusOK, userItems)
}

// 获取粉丝列表
type FollowerReq struct {
	Username string `json:"username"`
}
type FollowerResp struct {
	UserItem []views.UserItem
}

func GetFollowers(ctx *gin.Context) {
	var req FollowerReq
	var userItems []views.UserItem

	followers, err := dao.PluckFollowerByUsername(req.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(followers) != 0 {
		userItems, err = dao.SelectUserItem(followers)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	ctx.JSON(http.StatusOK, userItems)
}
