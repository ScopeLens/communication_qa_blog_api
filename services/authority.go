package services

import (
	"communication_qa_blog_api/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPower(ctx *gin.Context) {
	username := ctx.GetString("username")
	power, err := dao.IsPower(username)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, power)
}
