package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	var myMiddleWare = func(ctx *gin.Context) {
		fmt.Println("start")
	}
	engine.Use(myMiddleWare)
	engine.POST("/ping", func(ctx *gin.Context) {
		ctx.Writer.Write([]byte("pong"))
	})
	if err := engine.Run(":8080"); err != nil {
		panic(err)
	}

}
