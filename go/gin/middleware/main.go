package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func indexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})

}

func login(c *gin.Context) {
	//if 登录用户
	//c.Next()
	//else
	//c.Abort()
}

func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("name", "小王子") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		// 调用该请求的剩余处理程序
		c.Next()
		// 不调用该请求的剩余处理程序
		// c.Abort()
		// 计算耗时
		cost := time.Since(start)
		log.Println(cost)
	}
}

func m1(c *gin.Context) {
	fmt.Println("M1")
	start := time.Now()
	c.Next()
	cost := time.Since(start)
	fmt.Printf("%v\n", cost)
	//c.Abort() 阻止调用后续函数
}

func main() {
	r := gin.Default()
	//r.Use(m1,login)
	r.GET("/index", m1, indexHandler)
	r.Run(":8080")
}
