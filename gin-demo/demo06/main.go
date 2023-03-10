package main

import (
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 单文件上传
	r.POST("/singleFile", func(ctx *gin.Context) {
		file, err := ctx.FormFile("fileName")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		log.Println(file.Filename)
		// dst := fmt.Sprintf("./%s", file.Filename)
		dst := path.Join("./", file.Filename)
		ctx.SaveUploadedFile(file, dst)
		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("'%s' uploaded", file.Filename),
		})
	})

	r.POST("/multFile", func(ctx *gin.Context) {
		form, _ := ctx.MultipartForm()
		files := form.File["file"]

		for index, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("./temp/%s_%d", file.Filename, index)
			err := ctx.SaveUploadedFile(file, dst)
			if err != nil {
				log.Println("err is: ", err)
			}
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d files uploaded", len(files)),
		})
	})

	// 重定向相关
	r.GET("/redirect", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")

	})

	r.GET("/redirect2", func(ctx *gin.Context) {
		ctx.Request.URL.Path = "/test"
		r.HandleContext(ctx)
	})

	r.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "hello world",
		})
	})
	r.Run(":9000")
}
