package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("F:\\tryDemo\\1_go\\gin\\html\\post.html", "F:\\tryDemo\\1_go\\gin\\html\\user.html")
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "post.html", gin.H{
			"title": "posts/index",
		})
	})

	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user.html", gin.H{
			"title": "users/index",
		})
	})

	r.Run(":8080")
}
