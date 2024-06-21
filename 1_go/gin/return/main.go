package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"mes": "hhh",
		})
	})
	r.GET("sjson", func(c *gin.Context) {
		var msg struct {
			Name string `json:"name"`
			Age  int
		}
		msg.Age = 11
		msg.Name = "aaa"
		c.JSON(http.StatusOK, msg)
	})
	r.GET("mjson", func(c *gin.Context) {
		//data := map[string]interface{}{
		//	"name": "sss",
		//	"age":  11,
		//}
		data := gin.H{"name": "sss", "age": 11}
		c.JSON(http.StatusOK, data)
	})
	r.Run()
}
