package main

import "github.com/gin-gonic/gin"

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello",
	})
}
func main() {
	// 默认路由引擎
	r := gin.Default()

	//r.GET("/hello", sayHello)

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get",
		})
	})
	r.POST("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "post",
		})
	})
	r.PUT("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "put",
		})
	})
	r.DELETE("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "delete",
		})
	})
	r.Run()
}
