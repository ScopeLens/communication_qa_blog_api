package services

import (
	"communication_qa_blog_api/models/tables"
	"github.com/gin-gonic/gin"
)

func CreateComp(ctx *gin.Context) {
	var comp tables.Component
	err := ctx.ShouldBind(&comp)
	if err != nil {
		return
	}
	err = tables.CreateComponent(comp)
	if err != nil {
		return
	}
	component, err := tables.FirstComponent(1)
	if err != nil {
		return
	}
	ctx.JSON(200, component)
}
