package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, "cc")
	})
	r.Run()
}
