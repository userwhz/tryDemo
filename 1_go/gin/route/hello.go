package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 默认路由引擎
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get",
		})
	})
	r.Any("/test", func(c *gin.Context) {
		fmt.Println("any")
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "no",
		})
	})
	// 路由组
	//userGroup := r.Group("/user")
	//{
	//	userGroup.GET("/index", func(c *gin.Context) {...})
	//	userGroup.GET("/login", func(c *gin.Context) {...})
	//	userGroup.POST("/login", func(c *gin.Context) {...})
	//
	//}

	r.Run()
}
