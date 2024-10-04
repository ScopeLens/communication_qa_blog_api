package services

import (
	"communication_qa_blog_api/dao"
	"communication_qa_blog_api/models/tables"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 关注
type FUReq struct {
	Follower string `json:"Follower" binding:"required"`
	Followee string `json:"Followee" binding:"required"`
}
type FURsp struct {
	Message string `json:"message"`
}

func FollowUser(ctx *gin.Context) {
	var req FUReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	follow := tables.Follow{
		Follower: req.Follower,
		Followee: req.Followee,
	}

	//数据持久化
	err := dao.CreateFollow(follow)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"info": &FURsp{
			Message: "关注成功",
		},
	})
}

// 取关
type UFUReq struct {
	Follower string `json:"Follower" binding:"required"`
	Followee string `json:"Followee" binding:"required"`
}
type UFURsp struct {
	Message string `json:"message"`
}

func UnfollowUser(ctx *gin.Context) {
	var req FUReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	follow := tables.Follow{
		Follower: req.Follower,
		Followee: req.Followee,
	}

	//取关 删除数据
	err := dao.DeleteFollow(follow)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"info": &UFURsp{
			Message: "取关成功",
		},
	})
}
