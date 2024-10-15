package main

import (
	"communication_qa_blog_api/models"
	"communication_qa_blog_api/models/tables"
	api "communication_qa_blog_api/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	r := gin.Default()
	api.BasicRouter(r)
	DBInit()

	err := models.DB.AutoMigrate(&tables.Component{})
	if err != nil {
		return
	}

	//component := tables.Component{
	//	Name:        "组件1",
	//	Description: "组件一的描述",
	//	Dimensions:  json.RawMessage("[1,2]"),
	//}
	//err = tables.CreateComponent(component)
	//if err != nil {
	//	return
	//}
	//
	//firstComponent, err := tables.FirstComponent(1)
	//if err != nil {
	//	return
	//}
	//fmt.Println(firstComponent)

	err = r.Run(":105")
	if err != nil {
		fmt.Println(err)
		return
	}
}

//gorm

func DBInit() {
	username := "root"
	password := "911263610"
	host := "localhost"
	port := "3306"
	Dbname := "communication_qa_blog"
	timeout := "10s"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s",
		username, password, host, port, Dbname, timeout)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//SingularTable: true, //单数表名
			//NoLowerCase:   true, //不要自动转换成小写
		},
	})
	if err != nil {
		panic("连接数据库失败，err=" + err.Error())
	}
	models.DB = db.Debug()
}
